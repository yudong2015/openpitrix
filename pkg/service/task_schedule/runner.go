// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task_schedule

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

type TaskRunnerManager struct {
	runners []*TaskRunner
}

type TaskRunner struct {
	index int //the index of runner
	*RegisterClient
	*TaskInfoClient
	*TaskQueue
}

func NewTaskRunnerManager() *TaskRunnerManager {
	runners := make([]*TaskRunner, TaskRunnerNum)
	for i := 0; i < TaskRunnerNum; i++ {
		runners = append(runners, &TaskRunner{
			index:          i,
			RegisterClient: NewRegisterClient(),
			TaskInfoClient: NewTaskInfoClient(),
			TaskQueue:      GetTaskQueue(),
		})
	}
	return &TaskRunnerManager{runners}
}

func (tr *TaskRunner) GetTaskInfo(taskId string) (*models.ScheduleTask, error) {
	task := &models.ScheduleTask{}
	//get task info
	key := fmt.Sprintf(TaskInfoFmt, taskId)
	err := tr.TaskInfoClient.DlockWithTimeout(key, 60*time.Second, func() error {
		get, e := tr.TaskInfoClient.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debug(nil, "[TaskRunner-%d]The task info of %s is null!", tr.index, taskId)
		} else {
			e = yamlutil.Decode(get.Kvs[0].Value, task)
			if e != nil {
				return e
			}
			//update task info to running
			taskStr, e := task.UpdateToRun(Executor, strconv.Itoa(tr.index))
			if e != nil {
				return e
			}
			_, e = tr.TaskInfoClient.Put(nil, key, string(taskStr))
			if e != nil {
				return e
			}
		}
		return e
	})
	if err != nil {
		logger.Error(nil, "[TaskRunner-%d]Fail to get task(%s) info: %+v", tr.index, taskId, err)
	}
	return task, err
}

func (tr *TaskRunner) ConsumeTask() (*models.ScheduleTask, error) {
	var taskId string
	//consume task id
	for {
		taskId, err := tr.TaskQueue.DequeueTask()
		if err != nil {
			logger.Error(nil, "[TaskRunner-%d]Fail to dequeue task id from etcd queue: %+v", tr.index, err)
			time.Sleep(3 * time.Second)
			continue
		}
		logger.Debug(nil, "[TaskRunner-%d]Dequeue task id [%s] from etcd queue succeed", tr.index, taskId)
		break
	}

	//get task info
	return tr.GetTaskInfo(taskId)
}

func (tr *TaskRunner) run(ctx context.Context, task models.ScheduleTask) error {
	handlerClient, err := GetHandlerClient(task.Handler)
	if err != nil {
		return err
	}

	in := &pb.HandleTaskRequest{
		TaskId: pbutil.ToProtoString(task.Id),
		Action: pbutil.ToProtoString(task.Action),
		Conf:   pbutil.ToProtoString(task.Conf),
	}

	_, err = handlerClient.HandleTask(ctx, in)
	return err
}

func (tr *TaskRunner) Fail(task models.ScheduleTask) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.TaskInfoClient.DlockWithTimeout(key, 60*time.Second, func() error {
		get, e := tr.TaskInfoClient.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debug(nil, "The task info of %s is null!", task.Id)
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
			_, e = tr.TaskInfoClient.Put(nil, key, string(taskStr))
			if e != nil {
				return e
			}
		}
		return e
	})
}

func (tr *TaskRunner) Done(task models.ScheduleTask) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.TaskInfoClient.DlockWithTimeout(key, 60*time.Second, func() error {
		_, err := tr.TaskInfoClient.Delete(nil, key)
		if err != nil {
			logger.Error(nil, "[Runner-%d]Fail to update task[%s] to done: %+v", tr.index, task.Id, err)
		}
		return err
	})
}

func (tr *TaskRunner) Start(ctx context.Context) error {
	cancelCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	//register to etcd
	err := tr.RegisterClient.Register(cancelCtx, tr.index)
	if err != nil {
		logger.Error(ctx, "[Runner-%d]Fail to register to etcd: %+v", tr.index, err)
		return err
	}

	//loop to consume and handle task
	go func() {
		for {
			logger.Info(cancelCtx, "[Runner-%d]Consume task..", tr.index)
			task, err := tr.ConsumeTask()
			if err != nil {
				logger.Error(nil, "[Runner-%d]Fail to consume task: %+v", tr.index, err)
				//TODO: send alert email
				continue
			}

			logger.Info(cancelCtx, "[Runner-%d]Run task %s", tr.index, task.Id)
			err = tr.run(cancelCtx, *task)
			if err != nil {
				logger.Error(nil, "[Runner-%d]Fail to run task: %+v", tr.index, err)
				tr.Fail(*task)
				//TODO: send alert email
				continue
			}

			tr.Done(*task)

			//check if need to exit
			select {
			case <-ctx.Done():
				logger.Error(ctx, "[Runner-%d]Exit.", tr.index)
				break
			}
		}
	}()
	return nil
}

func (manager *TaskRunnerManager) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < TaskRunnerNum; i++ {
		err := manager.runners[i].Start(ctx)
		if err != nil {
			logger.Error(ctx, "Runner-%d failed to start: %+v", i, err)
			//TODO: handle err
		}
	}
}
