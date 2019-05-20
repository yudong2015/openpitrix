// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"fmt"
	"time"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/jsonutil"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

const (
	TaskQueueNum      = 100
	TaskQueueTopic    = "/task/ids"
	TaskQueueTopicFmt = TaskQueueTopic + "/%d"
	TaskInfoDir       = "/task/infos"
	TaskInfoFmt       = TaskInfoDir + "/%s"
	Executor          = "node_01"
)

type TaskRunner struct {
	etcd   *etcd.Etcd
	queues []*etcd.Queue
}

type Task struct {
	Id         string
	Runner     string
	Action     string
	Conf       string
	Exector    string
	Status     string
	RetryTimes int
}

func NewTaskRunner() *TaskRunner {
	e := pi.Global().Etcd(nil)
	queques := []*etcd.Queue{}
	for i := 0; i < TaskQueueNum; i++ {
		queques = append(queques, e.NewQueue(fmt.Sprintf(TaskQueueTopicFmt, i)))
	}
	return &TaskRunner{
		etcd:   e,
		queues: queques,
	}
}

func (t *Task) updateToRun() (string, error) {
	t.Exector = Executor
	t.Status = constants.StatusRunning
	b, err := jsonutil.Encode(t)
	return string(b), err
}

func (t *Task) fail() (string, error) {
	t.Status = constants.StatusFailed
	b, err := jsonutil.Encode(t)
	return string(b), err
}

func (tr *TaskRunner) getTaskInfo(taskId string) (*Task, error) {
	task := &Task{}
	//get task info
	key := fmt.Sprintf(TaskInfoFmt, taskId)
	err := tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		get, e := tr.etcd.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debugf(nil, "The task info of %s is null!", taskId)
		} else {
			e = yamlutil.Decode(get.Kvs[0].Value, task)
			if e != nil {
				return e
			}
			//update task info to running
			taskStr, e := task.updateToRun()
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
		logger.Errorf(nil, "Failed to get task info of %s : %+v", taskId, err)
	}
	return task, err
}

func (tr *TaskRunner) consumeTask(num int) (*Task, error) {
	var taskId string
	//consume task id
	for {
		taskId, err := tr.queues[num].Dequeue()
		if err != nil {
			logger.Errorf(nil, "Failed to dequeue task id from etcd queue: %+v", err)
			time.Sleep(3 * time.Second)
			continue
		}
		logger.Debugf(nil, "Dequeue task id [%s] from etcd queue succeed", taskId)
		break
	}

	//get task info
	return tr.getTaskInfo(taskId)
}

func (tr *TaskRunner) run(task Task) error {
	return nil
}

func (tr *TaskRunner) fail(task Task) {
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
			taskStr, e := task.fail()
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

func (tr *TaskRunner) done(task Task) {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	tr.etcd.DlockWithTimeout(key, 60*time.Second, func() error {
		_, err := tr.etcd.Delete(nil, key)
		if err != nil {
			logger.Errorf(nil, "Failed to update task[%s] to done: %+v", task.Id, err)
		}
		return err
	})
}

func (tr *TaskRunner) Serve(num int) {
	for {
		task, err := tr.consumeTask(num)
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
