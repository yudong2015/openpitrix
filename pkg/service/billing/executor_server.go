// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
)

type ExecutorServer struct{}

func NewExecutorServer() *ExecutorServer {
	return &ExecutorServer{}
}

func ExecutorServe(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s := NewExecutorServer()

	/**********************************************************
	** start task runner **
	**********************************************************/
	logger.Info(nil, "[%s]", "/**********************************************************")
	logger.Info(nil, "[%s]", "** start start task runner **")
	logger.Info(nil, "[%s]", "**********************************************************/")
	logger.Info(nil, "[%s]", "")
	manager.NewGrpcServer(constants.TaskScheduleManagerHost, constants.TaskScheduleManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterBillingTaskManagerServer(server, s)
			pb.RegisterChargingTaskManagerServer(server, s)
		})
}
