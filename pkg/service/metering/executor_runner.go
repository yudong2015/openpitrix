// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"go.etcd.io/etcd/clientv3"
	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/service/app"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

const (
	TaskRunnerNum  = 10
	TaskQueueTopic = "/task/ids"
	TaskInfoDir    = "/task/infos"
	TaskInfoFmt    = TaskInfoDir + "/%s"
	Executor       = "node_01"
	LeaseTimeout = 60
	AliveReportInterval = 30*time.Second
)

type TaskRunnerManager struct {
	runners []*TaskRunner
}

type TaskRunner struct {
	index int //the index of runner
	etcd  *etcd.Etcd
	queue *etcd.Queue
}

func NewTaskRunnerManager() *TaskRunnerManager {
	e := pi.Global().Etcd(nil)
	//queques := []*etcd.Queue{}
	//for i := 0; i < TaskQueueNum; i++ {
	//	queques = append(queques, e.NewQueue(fmt.Sprintf(TaskQueueTopicFmt, i)))
	//}

	queue := e.NewQueue(TaskQueueTopic)
	runners := make([]*TaskRunner, TaskRunnerNum)
	for i:=0;i<TaskRunnerNum;i++{
		runners = append(runners, &TaskRunner{
			index: i,
			etcd:  e,
			queue: queue,
		})
	}
	return &TaskRunnerManager{runners}
}

func (tr *TaskRunner) GetTaskInfo(taskId string) (*models.MbingTask, error) {
	task := &models.MbingTask{}
	//get task info
	key := fmt.Sprintf(TaskInfoFmt, taskId)
	err := tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		get, e := tr.etcd.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debugf(nil, "[TaskRunner-%d]The task info of %s is null!", tr.index, taskId)
		} else {
			e = yamlutil.Decode(get.Kvs[0].Value, task)
			if e != nil {
				return e
			}
			//update task info to running
			taskStr, e := task.UpdateToRun(Executor)
			if e != nil {
				return e
			}
			_, e = tr.etcd.Put(nil, key, string(taskStr))
			if e != nil {
				return e
			}
		}
		return e
	})
	if err != nil {
		logger.Errorf(nil, "[TaskRunner-%d]Failed to get task(%s) info: %+v", tr.index, taskId, err)
	}
	return task, err
}

func (tr *TaskRunner) ConsumeTask() (*models.MbingTask, error) {
	var taskId string
	//consume task id
	for {
		taskId, err := tr.queue.Dequeue()
		if err != nil {
			logger.Errorf(nil, "[TaskRunner-%d]Failed to dequeue task id from etcd queue: %+v", index, err)
			time.Sleep(3 * time.Second)
			continue
		}
		logger.Debugf(nil, "[TaskRunner-%d]Dequeue task id [%s] from etcd queue succeed", index, taskId)
		break
	}

	//get task info
	return tr.GetTaskInfo(taskId)
}

func (tr *TaskRunner) run(task models.MbingTask) error {
	return nil
}

func (tr *TaskRunner) Fail(task models.MbingTask) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		get, e := tr.etcd.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debugf(nil, "The task info of %s is null!", task.Id)
		} else {
			e = yamlutil.Decode(get.Kvs[0].Value, task)
			if e != nil {
				return e
			}
			//update task info to running
			taskStr, e := task.Fail()
			if e != nil {
				return e
			}
			_, e = tr.etcd.Put(nil, key, string(taskStr))
			if e != nil {
				return e
			}
		}
		return e
	})
}

func (tr *TaskRunner) Done(task models.MbingTask) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		_, err := tr.etcd.Delete(nil, key)
		if err != nil {
			logger.Errorf(nil, "Failed to update task[%s] to done: %+v", task.Id, err)
		}
		return err
	})
}

func (tr *TaskRunner) RegistToEtcd(ctx context.Context) error {
	grantRes, err := tr.etcd.Lease.Grant(ctx, LeaseTimeout)
	if err != nil {
		logger.Errorf(ctx, "Failed to grant a lease: %+v", err)
		return err
	}
	_, err = tr.etcd.Put(ctx, Executor, strconv.Itoa(tr.index), clientv3.WithLease(grantRes.ID))
	if err != nil {
		logger.Errorf(ctx, "Failed to put: %+v", err)
		return err
	}
	//heartBeat
	cancelCtx, cancel := context.WithCancel(ctx)
	go tr.alive(cancelCtx, grantRes.ID)
	return nil
}

func (tr *TaskRunner) alive(ctx context.Context, ID clientv3.LeaseID){
	for {
		_, err := tr.etcd.Lease.KeepAliveOnce(ctx, ID)
		if err != nil {
			logger.Errorf(ctx, "Failed to KeepAliveOnce: %+v", err)
		}
		time.Sleep(AliveReportInterval)
	}
}

func (tr *TaskRunner) Start() {

	for {
		task, err := tr.ConsumeTask()
		if err != nil {
			logger.Errorf(nil, "Failed to consume task: %+v", err)
			//TODO: send alert email
		}
		err = tr.run(*task)
		if err != nil {
			logger.Errorf(nil, "Failed to run task: %+v", err)
			tr.Fail(*task)
			//TODO: send alert email
		}
		tr.Done(*task)
	}
}

func (manager *TaskRunnerManager) Serve() {
	ctx := context.Background()
	for i := 0; i < TaskRunnerNum; i++ {
		//TODO: goroutine report themself to TaskRunner
		go manager.runners[i].Start()
	}
}
