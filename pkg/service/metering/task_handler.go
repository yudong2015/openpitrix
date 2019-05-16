// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"

	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/pb"
)

func (ts *TaskServer) InitMeteringTask(ctx context.Context, req *pb.InitMeteringRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *TaskServer) UpdateMeteringTask(ctx context.Context, req *pb.UpdateMeteringRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *TaskServer) StopMeteringsTask(ctx context.Context, req *pb.StopMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *TaskServer) StartMeteringsTask(ctx context.Context, req *pb.StartMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

func (ts *TaskServer) TerminateMeteringsTask(ctx context.Context, req *pb.TerminateMeteringsRequest) (*pb.MeteringTaskResponse, error) {
	//TODO: impl
	return &pb.MeteringTaskResponse{}, nil
}

type MultiQueue struct {
	Queues []*etcd.Queue
	Topic string
	Size int
}

func (mq *MultiQueue) MutilEnqueue(val string, ) error {
	return nil
}

func (mq *MultiQueue) MutilDequeue(val string, ) error {
	return nil
}
