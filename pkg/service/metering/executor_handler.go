// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"fmt"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
)

const (
	ActionInitMetering      = "InitMetering"
	ActionUpdateMetering    = "UpdateMetering"
	ActionStopMetering      = "StopMetering"
	ActionStartMetering     = "StartMetering"
	ActionTerminateMetering = "TerminateMetering"
)

var MeteringTaskRunner = fmt.Sprintf("%s:%d", constants.MeteringManagerHost, constants.MeteringManagerPort)

func publishTask(ctx context.Context, etcd *etcd.Etcd, task models.MbingTask) error {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	return etcd.Dlock(ctx, key, func() error {
		taskStr, err := task.String()
		if err != nil {
			return err
		}
		_, err = etcd.Put(ctx, key, taskStr)
		return err
	})
}

func EnquequeMbingTask(ctx context.Context, etcd *etcd.Etcd, task models.MbingTask) error {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	return etcd.Dlock(ctx, key, func() error {
		taskStr, err := task.String()
		if err != nil {
			return err
		}
		_, err = etcd.Put(ctx, key, taskStr)
		return err
	})

	return etcd.NewQueue()
}

func (ts *ExecutorServer) InitMeteringTask(ctx context.Context, req *pb.InitMeteringRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	task, err := models.PbToMbingTask(ctx, req, MeteringTaskRunner, ActionInitMetering, Executor)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}
	err = publishTask(ctx, pi.Global().Etcd(ctx), *task)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorCreateMbingTask)
	}

	return &pb.MeteringTaskResponse{}, nil
}

func (ts *ExecutorServer) UpdateMeteringTask(ctx context.Context, req *pb.UpdateMeteringRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *ExecutorServer) StopMeteringsTask(ctx context.Context, req *pb.StopMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *ExecutorServer) StartMeteringsTask(ctx context.Context, req *pb.StartMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *ExecutorServer) TerminateMeteringsTask(ctx context.Context, req *pb.TerminateMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}
