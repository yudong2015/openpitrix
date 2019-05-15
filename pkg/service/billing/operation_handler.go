// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"context"

	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func (s *Server) CreatePrice(ctx context.Context, req *pb.CreatePriceRequest) (*pb.CreatePriceResponse, error) {
	price := models.PbToPrice(req)

	//TODO: how to check skuId and attributeId

	//insert price
	err := insertPrice(ctx, price)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreatePriceResponse{PriceId: pbutil.ToProtoString(price.PriceId)}, nil
}

func (s *Server) DescribePrices(ctx context.Context, req *pb.DescribePricesRequest) (*pb.DescribePricesResponse, error) {
	//TODO: impl DescribePrices
	return &pb.DescribePricesResponse{}, nil
}

func (s *Server) ModifyPrice(ctx context.Context, req *pb.ModifyPriceRequest) (*pb.ModifyPriceResponse, error) {
	//TODO: impl ModifyPrice
	return &pb.ModifyPriceResponse{}, nil
}

func (s *Server) DeletePrices(ctx context.Context, req *pb.DeletePricesRequest) (*pb.DeletePricesResponse, error) {
	//TODO: impl DeletePrices
	return &pb.DeletePricesResponse{}, nil
}

func (s *Server) DescribeLeasingContracts(ctx context.Context, req *pb.DescribeLeasingContractsRequest) (*pb.DescribeLeasingContractsResponse, error){
	//TODO: impl
	return &pb.DescribeLeasingContractsResponse{}, nil
}

func (s *Server) DescribeLeasedContracts(ctx context.Context, req *pb.DescribeLeasedContractsRequest) (*pb.DescribeLeasedContractsResponse, error){
	//TODO: impl
	return &pb.DescribeLeasedContractsResponse{}, nil
}

