// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"context"

	"openpitrix.io/openpitrix/pkg/pb"
)

//billing task impl
func (es *ExecutorServer) InitBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (es *ExecutorServer) UpdateBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (es *ExecutorServer) StopBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (es *ExecutorServer) StartBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (es *ExecutorServer) TerminateBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

//charging task impl
func (es *ExecutorServer) CreateChargingTask(ctx context.Context, req *pb.LeasingContract) (*pb.ChargingTaskResponse, error) {
	//TODO: impl
	return &pb.ChargingTaskResponse{}, nil
}

func (es *ExecutorServer) CreateRefundTask(ctx context.Context, req *pb.LeasingContract) (*pb.ChargingTaskResponse, error) {
	//TODO: impl
	return &pb.ChargingTaskResponse{}, nil
}
