// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task_schedule

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/service/billing"
	"openpitrix.io/openpitrix/pkg/service/metering"
)

const (
	TaskRunnerNum      = 10
	TaskQueueTopic     = "/task/ids"
	TaskInfoDir        = "/task/infos"
	TaskInfoFmt        = TaskInfoDir + "/%s"
	Executor           = "node_01"
	RunnerLeaseTimeout = 60
	HeartBeatInterval  = 30 * time.Second
	RunnerRegisterDir  = "/runners"
	RunnerRegisterFmt  = RunnerRegisterDir + "/%s/%d"
)

func NewRegisterClient() *RegisterClient {
	return &RegisterClient{
		pi.Global().Etcd(nil),
	}
}

func NewTaskInfoClient() *TaskInfoClient {
	return &TaskInfoClient{
		pi.Global().Etcd(nil),
	}
}

type RegisterClient struct {
	*etcd.Etcd
}

type TaskInfoClient struct {
	*etcd.Etcd
}

type TaskQueue struct {
	*etcd.Queue
}

//register client operations
func (rc *RegisterClient) Register(ctx context.Context, runnerIndex int) error {
	grantRes, err := rc.Lease.Grant(ctx, RunnerLeaseTimeout)
	if err != nil {
		return err
	}
	key := fmt.Sprintf(RunnerRegisterFmt, Executor, runnerIndex)
	_, err = rc.Put(ctx, key, strconv.Itoa(runnerIndex), clientv3.WithLease(grantRes.ID))
	if err != nil {
		return err
	}

	go rc.HeartBeat(ctx, runnerIndex, grantRes.ID)

	return nil
}

func (rc *RegisterClient) HeartBeat(ctx context.Context, runnerIndex int, ID clientv3.LeaseID) {
	for {
		_, err := rc.Lease.KeepAliveOnce(ctx, ID)
		if err != nil {
			logger.Errorf(ctx, "Runner-%d failed to heart beat: %+v", runnerIndex, err)
		}
		time.Sleep(HeartBeatInterval)
		select {
		case <-ctx.Done():
			break
		}
	}
}

//task info client operation
func (cli *TaskInfoClient) PublishTask(ctx context.Context, task models.ScheduleTask) error {
	taskStr, err := task.String()
	if err != nil {
		return err
	}

	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	_, err = cli.Put(ctx, key, taskStr)
	return err
}

//task queue operations
var taskQueue *TaskQueue
var once sync.Once

func GetTaskQueue() *TaskQueue {
	once.Do(func() {
		taskQueue = &TaskQueue{
			pi.Global().Etcd(nil).NewQueue(TaskQueueTopic),
		}
	})
	return taskQueue
}

func (tq *TaskQueue) EnqueueTask(taskId string) error {
	return tq.Enqueue(taskId)
}

func (tq *TaskQueue) DequeueTask() (string, error) {
	return tq.Dequeue()
}

//get handle task client
//return host and port by handler type
func getHandlerEndpoint(handler string) (string, int) {
	switch handler {
	case metering.MeterHandler:
		return constants.MeteringManagerHost, constants.MeteringManagerPort
	case billing.BillingHandler:
		return constants.BillingManagerHost, constants.BillingManagerPort
	default:
		return "", -1
	}
}

func GetHandlerClient(handler string) (pb.TaskHandlerManagerClient, error) {
	host, port := getHandlerEndpoint(handler)
	if host == "" || port == -1 {
		return nil, errors.New("The handler client endpoint not exist.")
	}
	conn, err := manager.NewClient(host, port)
	if err != nil {
		return nil, err
	}
	return pb.NewTaskHandlerManagerClient(conn), err
}
