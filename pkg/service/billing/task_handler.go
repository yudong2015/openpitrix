// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"context"

	"openpitrix.io/openpitrix/pkg/pb"
)

//billing task impl
func (ts TaskServer) InitBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (ts TaskServer) UpdateBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (ts TaskServer) StopBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (ts TaskServer) StartBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}

func (ts TaskServer) TerminateBillingTask(ctx context.Context, req *pb.Leasing) (*pb.BillingTaskResponse, error) {
	//TODO: impl
	return &pb.BillingTaskResponse{}, nil
}


//charging task impl
func (ts TaskServer) CreateChargingTask(ctx context.Context, req *pb.LeasingContract) (*pb.ChargingTaskResponse, error) {
	//TODO: impl
	return &pb.ChargingTaskResponse{}, nil
}

func (ts TaskServer) CreateRefundTask(ctx context.Context, req *pb.LeasingContract) (*pb.ChargingTaskResponse, error) {
	//TODO: impl
	return &pb.ChargingTaskResponse{}, nil
}
