// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package task_schedule

import (
	"google.golang.org/grpc"

	"openpitrix.io/logger"
	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
)

type ExecutorServer struct {
	*TaskInfoClient
	*TaskQueue
	*TaskRunnerManager
}

func NewExecutorServer() *ExecutorServer {
	return &ExecutorServer{
		TaskInfoClient: NewTaskInfoClient(),
		TaskQueue: GetTaskQueue(),
		TaskRunnerManager: NewTaskRunnerManager(),
	}
}

func ExecutorServe(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s := NewExecutorServer()

	//** start task runner **
	logger.Infof(nil, "[%s]", "** start TaskRunnerManager **")
	go s.TaskRunnerManager.Serve()

	logger.Infof(nil, "[%s]", "** start TaskScheduleService **")
	manager.NewGrpcServer(constants.MbingExecutorManagerHost, constants.MbingExecutorManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterTaskScheduleManagerServer(server, s)
		})
}
