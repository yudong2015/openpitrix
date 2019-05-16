// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"fmt"
	"time"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

const (
	TaskQueueNum   = 100
	TaskQueueTopic = "/task/ids"
	TaskInfoDir    = "/task/infos"
)

type TaskRunner struct {
	etcd   *etcd.Etcd
	queues []*etcd.Queue
}

type Task struct {
	Id         string
	Type       string
	Params     map[string]interface{}
	Runner     string
	Status     string
	RetryTimes int
}

func (tr *TaskRunner) getTaskInfo(taskId string) (*Task, error) {
	task := &Task{}
	//get task info
	key := fmt.Sprintf("%s/%s", TaskInfoDir, taskId)
	err := tr.etcd.DlockWithTimeout(key, 30*time.Second, func() error {
		get, e := tr.etcd.Get(nil, key)
		if e != nil {
			return e
		}
		// parse value
		if get.Count == 0 {
			logger.Debugf(nil, "The task info of %s is null!", taskId)
		} else {
			e = yamlutil.Decode(get.Kvs[0].Value, task)
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

func (tr *TaskRunner) running(task Task) error {
	//TODO: curl
	return nil
}

func (tr *TaskRunner) Run(num int) {
	for {
		task, err := tr.consumeTask(num)
		if err != nil {
			logger.Errorf(nil, "Failed to consume task: %+v", err)
			//TODO: send alert email
		}
		err = tr.running(*task)
		if err != nil {
			logger.Errorf(nil, "Failed to run task: %+v", err)
			//TODO: send alert email
		}
	}
}

