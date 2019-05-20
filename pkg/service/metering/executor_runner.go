// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"fmt"
	"time"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

const (
	TaskRunnerNum  = 10
	TaskQueueTopic = "/task/ids"
	TaskInfoDir    = "/task/infos"
	TaskInfoFmt    = TaskInfoDir + "/%s"
	Executor       = "node_01"
)

type TaskRunner struct {
	etcd  *etcd.Etcd
	queue *etcd.Queue
}

func NewTaskRunner() *TaskRunner {
	e := pi.Global().Etcd(nil)
	//queques := []*etcd.Queue{}
	//for i := 0; i < TaskQueueNum; i++ {
	//	queques = append(queques, e.NewQueue(fmt.Sprintf(TaskQueueTopicFmt, i)))
	//}
	return &TaskRunner{
		etcd:  e,
		queue: e.NewQueue(TaskQueueTopic),
	}
}

func (tr *TaskRunner) getTaskInfo(taskId string, index int) (*models.MbingTask, error) {
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
			logger.Debugf(nil, "[TaskRunner-%d]The task info of %s is null!", index, taskId)
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
		logger.Errorf(nil, "[TaskRunner-%d]Failed to get task(%s) info: %+v", index, taskId, err)
	}
	return task, err
}

func (tr *TaskRunner) consumeTask(index int) (*models.MbingTask, error) {
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
	return tr.getTaskInfo(taskId, index)
}

func (tr *TaskRunner) run(task models.MbingTask) error {
	return nil
}

func (tr *TaskRunner) fail(task models.MbingTask) {
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

func (tr *TaskRunner) done(task models.MbingTask) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		_, err := tr.etcd.Delete(nil, key)
		if err != nil {
			logger.Errorf(nil, "Failed to update task[%s] to done: %+v", task.Id, err)
		}
		return err
	})
}

func (tr *TaskRunner) registToEtcd(index int) error {
	//TODO: regist
	return nil
}

func (tr *TaskRunner) Start(index int) {
	tr.registToEtcd(index)
	for {
		task, err := tr.consumeTask(index)
		if err != nil {
			logger.Errorf(nil, "Failed to consume task: %+v", err)
			//TODO: send alert email
		}
		err = tr.run(*task)
		if err != nil {
			logger.Errorf(nil, "Failed to run task: %+v", err)
			tr.fail(*task)
			//TODO: send alert email
		}
		tr.done(*task)
	}
}


func (tr *TaskRunner) Serve() {
	for i:=0;i<TaskRunnerNum;i++{
		go tr.Start(i)
	}
}
