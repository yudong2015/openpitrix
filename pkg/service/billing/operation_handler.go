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

func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error){
	//TODO: impl
	return &pb.CreateAccountResponse{UserId: pbutil.ToProtoString("init_user")}, nil
}

func (s *Server) DescribeAccounts(ctx context.Context, req *pb.DescribeAccountsRequest) (*pb.DescribeAccountsResponse, error){
	//TODO: impl
}

func (s *Server) ModifyAccount(ctx context.Context, req *pb.ModifyAccountRequest) (*pb.ModifyAccountResponse, error){
	//TODO: impl
}

func (s *Server) DeleteAccounts(ctx context.Context, req *pb.DeleteAccountsRequest) (*pb.DeleteAccountsResponse, error){
	//TODO: impl
}

func (s *Server) DescribeLeasingContracts(ctx context.Context, req *pb.DescribeLeasingContractsRequest) (*pb.DescribeLeasingContractsResponse, error){
	//TODO: impl
}

func (s *Server) DescribeLeasedContracts(ctx context.Context, req *pb.DescribeLeasedContractsRequest) (*pb.DescribeLeasedContractsResponse, error){
	//TODO: impl
}

func (s *Server) CreateRecharge(ctx context.Context, req *pb.CreateRechargeRequest) (*pb.CreateRechargeResponse, error){
	//TODO: impl
}

func (s *Server) DescribeRecharges(ctx context.Context, req *pb.DescribeRechargesRequest) (*pb.DescribeRechargesResponse, error){
	//TODO: impl
}

func (s *Server) DescribeCharges(ctx context.Context, req *pb.DescribeChargesRequest) (*pb.DescribeChargesResponse, error){
	//TODO: impl
}

func (s *Server) DescribeRefunds(ctx context.Context, req *pb.DescribeRefundsRequest) (*pb.DescribeRefundsResponse, error){
	//TODO: impl
}

func (s *Server) DescribeIncomes(ctx context.Context, req *pb.DescribeIncomesRequest) (*pb.DescribeIncomesResponse, error){
	//TODO: impl
}

func (s *Server) CreateWithdraw(ctx context.Context, req *pb.CreateWithdrawRequest) (*pb.CreateWithdrawResponse, error){
	//TODO: impl
}

func (s *Server) DescribeWithdraws(ctx context.Context, req *pb.DescribeWithdrawsRequest) (*pb.DescribeWithdrawsResponse, error){
	//TODO: impl
}
