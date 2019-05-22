// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"fmt"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

const (
	ActionInitMetering      = "InitMetering"
	ActionUpdateMetering    = "UpdateMetering"
	ActionStopMetering      = "StopMetering"
	ActionStartMetering     = "StartMetering"
	ActionTerminateMetering = "TerminateMetering"
)

var MeteringTaskHandler = fmt.Sprintf("%s:%d", constants.MeteringManagerHost, constants.MeteringManagerPort)

func (es *ExecutorServer) InitMeteringTask(ctx context.Context, req *pb.InitMeteringRequest) (*pb.MeteringTaskResponse, error) {
	task, err := models.PbToMbingTask(ctx, req, MeteringTaskHandler, ActionInitMetering)
	if err != nil {
		logger.Errorf(ctx, "Failed to convert [%+v] to MbingTask: %+v", req.String(), err)
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}

	err = PublishTask(ctx, pi.Global().Etcd(ctx), *task)
	if err != nil {
		logger.Errorf(ctx, "Failed to publish MbingTask: %+v", err)
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}

	err = EnquequeTask(es.TaskQueue, task.Id)
	if err != nil {
		logger.Errorf(ctx, "Failed to enqueue task to queue: %+v", err)
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}

	return &pb.MeteringTaskResponse{TaskId: pbutil.ToProtoString(task.Id)}, nil
}

func (es *ExecutorServer) UpdateMeteringTask(ctx context.Context, req *pb.UpdateMeteringRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	task, err := models.PbToMbingTask(ctx, req, MeteringTaskHandler, ActionUpdateMetering)
	return &pb.MeteringTaskResponse{}, nil
}

func (es *ExecutorServer) StopMeteringsTask(ctx context.Context, req *pb.StopMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (es *ExecutorServer) StartMeteringsTask(ctx context.Context, req *pb.StartMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (es *ExecutorServer) TerminateMeteringsTask(ctx context.Context, req *pb.TerminateMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}
