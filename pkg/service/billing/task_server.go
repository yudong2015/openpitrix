// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	m "openpitrix.io/openpitrix/pkg/service/metering"
)

type TaskServer struct {
	*m.TaskServer
}

func NewTaskServer() (*TaskServer, error){
	server := &TaskServer{&m.TaskServer{}}
	return server, nil
}

func TaskServe(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s, _ := NewTaskServer()
	manager.NewGrpcServer(constants.MbingTaskManagerHost, constants.MbingTaskManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterMeteringTaskManagerServer(server, s.TaskServer)
			pb.RegisterBillingTaskManagerServer(server, s)
		})
}
