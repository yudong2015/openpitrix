// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"context"
	"openpitrix.io/openpitrix/pkg/constants"

	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func (s *Server) CreatePrice(ctx context.Context, req *pb.CreatePriceRequest) (*pb.CreatePriceResponse, error) {
	price := models.PbToPrice(req)

	//TODO: how to check skuId and attributeIds
	//TODO: check if duplicate attributeIds

	//insert price
	err := insertPrice(ctx, price)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreatePriceResponse{PriceId: pbutil.ToProtoString(price.PriceId)}, nil
}

func (s *Server) DescribePrices(ctx context.Context, req *pb.DescribePricesRequest) (*pb.DescribePricesResponse, error) {
	count, prices, err := getPrices(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	var pbPrices []*pb.Price
	for _, price := range prices {
		pbPrices = append(pbPrices, models.PriceToPb(price))
	}

	return &pb.DescribePricesResponse{TotalCount: uint32(count), PriceSet: pbPrices}, nil
}

func (s *Server) ModifyPrice(ctx context.Context, req *pb.ModifyPriceRequest) (*pb.ModifyPriceResponse, error) {
	//TODO: how to check skuId and attributeIds
	//TODO: check if duplicate attributeIds
	err := updatePrice(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.ModifyPriceResponse{PriceId: req.GetPriceId()}, nil
}

func (s *Server) DeletePrices(ctx context.Context, req *pb.DeletePricesRequest) (*pb.DeletePricesResponse, error) {
	err := updateStatusToDeleted(ctx, constants.TablePrice, req.GetPriceIds())
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.DeletePricesResponse{PriceIds: req.GetPriceIds()}, nil
}

func (s *Server) DescribeLeasingContracts(ctx context.Context, req *pb.DescribeLeasingContractsRequest) (*pb.DescribeLeasingContractsResponse, error) {
	//TODO: impl
	return &pb.DescribeLeasingContractsResponse{}, nil
}

func (s *Server) DescribeLeasedContracts(ctx context.Context, req *pb.DescribeLeasedContractsRequest) (*pb.DescribeLeasedContractsResponse, error) {
	//TODO: impl
	return &pb.DescribeLeasedContractsResponse{}, nil
}
