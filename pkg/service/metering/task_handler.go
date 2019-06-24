// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"

	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/yamlutil"
)

const (
	MeterHandler            = "metering"
	ActionInitMetering      = "InitMetering"
	ActionUpdateMetering    = "UpdateMetering"
	ActionStopMetering      = "StopMetering"
	ActionStartMetering     = "StartMetering"
	ActionTerminateMetering = "TerminateMetering"
)

var constWrapperMap sync.Map

func getConstantWrapper(constant string) *wrappers.StringValue {
	v, _ := constWrapperMap.Load(constant)
	if v == nil {
		cw := pbutil.ToProtoString(constant)
		constWrapperMap.Store(constant, cw)
		return cw
	} else {
		return v.(*wrappers.StringValue)
	}
}

func newCreateTaskRequest(action, conf string) *pb.CreateScheduleTaskRequest {
	return &pb.CreateScheduleTaskRequest{
		Handler: getConstantWrapper(MeterHandler),
		Action:  getConstantWrapper(action),
		Conf:    pbutil.ToProtoString(conf),
	}
}

func parseConf(conf string, obj interface{}) error {
	return yamlutil.Decode([]byte(conf), obj)
}

func handleInitMeteringTask(ctx context.Context, req *pb.HandleTaskRequest) (*pb.HandleTaskResponse, error) {
	initReq := &pb.InitMeteringRequest{}
	err := yamlutil.Decode([]byte(req.GetConf().GetValue()), initReq)
	if err != nil {
		logger.Error(ctx, "Fail to parse configration to struct:%+v", err)
		return nil, err
	}

	err = initMetering(ctx, initReq)
	if err != nil {
		logger.Error(ctx, "Fail to init metering:%+v", err)
		return nil, err
	}
	return &pb.HandleTaskResponse{Ok: pbutil.ToProtoBool(true)}, nil
}

func initMetering(ctx context.Context, req *pb.InitMeteringRequest) error {
	var leasings []*models.Leasing
	now := time.Now()

	for _, skuM := range req.GetSkuMeterings() {
		skuId := skuM.GetSkuId().GetValue()
		renewalTime, _ := renewalTimeFromSku(ctx, skuId, now)
		leasings = append(leasings, models.NewLeasing(skuM.GetValue(), req.GetUserId().GetValue(), req.GetResourceId().GetValue(), skuId, now, *renewalTime))
	}

	//insert leasings
	err := insertLeasings(ctx, leasings)
	if err != nil {
		return err
	}

	//TODO: How to guarantee consistency operations.
	for _, l := range leasings {
		err = leasingToRedis(*l)
		//TODO: check if BillingService exist, then add billing task for pre-charging by curl TaskService
	}
	return nil
}

func updateMeteringHandler(ctx context.Context, req *pb.HandleTaskRequest) (*pb.HandleTaskResponse, error) {

	return &pb.HandleTaskResponse{}, nil
}

func (s *Server) HandleTask(ctx context.Context, req *pb.HandleTaskRequest) (*pb.HandleTaskResponse, error) {
	switch req.GetAction().GetValue() {
	case ActionInitMetering:
		return handleInitMeteringTask(ctx, req)
	case ActionUpdateMetering:
		return updateMeteringHandler(ctx, req)
	default:
		return nil, errors.New("The action not exist!")
	}

	return &pb.HandleTaskResponse{Ok: pbutil.ToProtoBool(true)}, nil
}
