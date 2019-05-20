// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"fmt"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/pb"
)

var MeteringTaskRunner = fmt.Sprintf("%s:%d", constants.MeteringManagerHost, constants.MeteringManagerPort)

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
