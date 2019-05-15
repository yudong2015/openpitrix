// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package mbing

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pi"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/pb"
)

type TaskServer struct{}

func NewTaskServer() (*TaskServer, error) {
	server := &TaskServer{}
	return server, nil
}

func TaskServe(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s, _ := NewTaskServer()
	manager.NewGrpcServer(constants.BillingManagerHost, constants.BillingManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterMeteringTaskManagerHandler(server, &s)
		})
}
