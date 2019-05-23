// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"fmt"
	"time"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/idutil"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

func NewScheduleTaskId() string {
	return idutil.GetUuid("mbt-")
}

type ScheduleTask struct {
	Id     string
	Handler    string
	Action     string
	Conf       string
	Runner     string
	Status     string
	RetryTimes int
	CreateTime time.Time
}

func PbToScheduleTask(req *pb.CreateScheduleTaskRequest) *ScheduleTask {
	return &ScheduleTask{
		Id:     NewScheduleTaskId(),
		Handler:    req.Handler.GetValue(),
		Action:     req.Action.GetValue(),
		Conf:       req.Conf.GetValue(),
		Status:     constants.StatusReady,
		RetryTimes: 0,
		CreateTime: time.Now(),
	}
}

func (t *ScheduleTask) UpdateToRun(executor, runner string) (string, error) {
	t.Runner = fmt.Sprintf("%s-%s", executor, runner)
	t.Status = constants.StatusRunning
	return t.String()
}

func (t *ScheduleTask) Fail() (string, error) {
	t.Status = constants.StatusFailed
	return t.String()
}

func (t *ScheduleTask) String() (string, error) {
	b, err := yamlutil.Encode(t)
	if err != nil {
		logger.Errorf(nil, "Failed to encode ScheduleTask to string: %+v", err)
		return "", err
	}
	return string(b), err
}
