// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task_schedule

import (
	"context"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func (es *ExecutorServer) CreateTask(ctx context.Context, req *pb.CreateScheduleTaskRequest) (*pb.CreateScheduleTaskResponse, error) {
	task := models.PbToScheduleTask(req)

	err := PublishTask(ctx, pi.Global().Etcd(ctx), *task)
	if err != nil {
		logger.Errorf(ctx, "Failed to publish task: %+v", err)
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}

	//if error, it can be enqueue by HeathChecker
	EnquequeTask(es.TaskQueue, task.Id)
	return &pb.CreateScheduleTaskResponse{TaskId: pbutil.ToProtoString(task.Id)}, nil
}
