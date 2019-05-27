// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"google.golang.org/grpc"

	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pi"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/pb"
)

type Server struct{}

func Serve(cfg *config.Config) {
	pi.SetGlobal(cfg)
	s := &Server{}
	manager.NewGrpcServer(constants.MeteringManagerHost, constants.MeteringManagerPort).
		ShowErrorCause(cfg.Grpc.ShowErrorCause).
		Serve(func(server *grpc.Server) {
			pb.RegisterSkuManagerServer(server, s)
			pb.RegisterMeteringManagerServer(server, s)
			pb.RegisterPromotionSkuManagerServer(server, s)
			pb.RegisterTaskHandlerManagerServer(server, s)
		})
}
