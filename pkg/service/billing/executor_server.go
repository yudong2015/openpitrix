// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"google.golang.org/grpc"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	m "openpitrix.io/openpitrix/pkg/service/metering"
)

type ExecutorServer struct {
	*m.ExecutorServer
}

func NewExecutorServer() *ExecutorServer {
	return &ExecutorServer{
		m.NewExecutorServer(),
	}
}

func ExecutorServe(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s := NewExecutorServer()

	/**********************************************************
	** start task runner **
	**********************************************************/
	logger.Infof(nil, "[%s]", "/**********************************************************")
	logger.Infof(nil, "[%s]", "** start start task runner **")
	logger.Infof(nil, "[%s]", "**********************************************************/")
	logger.Infof(nil, "[%s]", "")
	go s.Runner.Serve()
	manager.NewGrpcServer(constants.MbingExecutorManagerHost, constants.MbingExecutorManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterMeteringTaskManagerServer(server, s.ExecutorServer)
			pb.RegisterBillingTaskManagerServer(server, s)
			pb.RegisterChargingTaskManagerServer(server, s)
		})
}
