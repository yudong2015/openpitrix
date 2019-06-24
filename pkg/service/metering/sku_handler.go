// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"time"

	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)


func (s *Server) CreateAttributeTerm(ctx context.Context, req *pb.CreateAttributeTermRequest) (*pb.CreateAttributeTermResponse, error) {
	term := models.PbToAttributeTerm(req, UserId(ctx))
	err := insertAttributeTerm(ctx, term)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateAttributeTermResponse{AttributeTermId: pbutil.ToProtoString(term.AttributeTermId)}, nil
}

func (s *Server) DescribeAttributeTerms(ctx context.Context, req *pb.DescribeAttributeTermsRequest) (*pb.DescribeAttributeTermsResponse, error) {
	attributeTerms, err := getAttributeTerms(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	var pbAttributeTerms []*pb.AttributeTerm
	for _, attTerm := range attributeTerms {
			pbAttributeTerms = append(pbAttributeTerms, models.AttributeTermToPb(attTerm))
	}

	return &pb.DescribeAttributeTermsResponse{AttributeTermSet: pbAttributeTerms}, nil
}

func (s *Server) ModifyAttributeTerm(ctx context.Context, req *pb.ModifyAttributeTermRequest) (*pb.ModifyAttributeTermResponse, error) {
	attTerm, err := CheckAttributeTermPermission(ctx, req.GetAttributeTermId().GetValue())
	if err != nil {
		return nil, err
	}

	attTerm.Name = req.GetName().GetValue()
	attTerm.Description = req.GetDescription().GetValue()

	err = updateAttributeTerm(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.ModifyAttributeTermResponse{AttributeTermId: req.AttributeTermId}, nil
}

func (s *Server) DeleteAttributeTerms(ctx context.Context, req *pb.DeleteAttributeTermsRequest) (*pb.DeleteAttributeTermsResponse, error) {
	attTermIds := req.GetAttributeTermIds()
	_, err := CheckAttributeTermsPermission(ctx, attTermIds)
	if err != nil {
		return nil, err
	}

	err = deleteAttributeTerms(ctx, attTermIds)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.DeleteAttributeTermsResponse{AttributeTermIds: req.AttributeTermIds}, nil
}

func (s *Server) CreateAttributeUnit(ctx context.Context, req *pb.CreateAttributeUnitRequest) (*pb.CreateAttributeUnitResponse, error) {
	attributeUnit := models.PbToAttributeUnit(req, UserId(ctx))
	err := insertAttributeUnit(ctx, attributeUnit)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateAttributeUnitResponse{AttributeUnitId: pbutil.ToProtoString(attributeUnit.AttributeUnitId)}, nil
}

func (s *Server) DescribeAttributeUnits(ctx context.Context, req *pb.DescribeAttributeUnitsRequest) (*pb.DescribeAttributeUnitsResponse, error) {
	attributeUnits, err := getAttributeUnits(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	var pbAttributeUnits []*pb.AttributeUnit
	for _, attUnit := range attributeUnits {
		pbAttributeUnits = append(pbAttributeUnits, models.AttributeUnitToPb(attUnit))
	}

	return &pb.DescribeAttributeUnitsResponse{AttributeUnitSet: pbAttributeUnits}, nil
}

func (s *Server) DeleteAttributeUnits(ctx context.Context, req *pb.DeleteAttributeUnitsRequest) (*pb.DeleteAttributeUnitsResponse, error) {
	attUnitIds := req.GetAttributeUnitIds()
	_, err := CheckAttributeUnitsPermission(ctx, attUnitIds)
	if err != nil {
		return nil, err
	}

	err = deleteAttributeUnits(ctx, attUnitIds)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.DeleteAttributeUnitsResponse{AttributeUnitIds: req.GetAttributeUnitIds()}, nil
}

func (s *Server) CreateAttribute(ctx context.Context, req *pb.CreateAttributeRequest) (*pb.CreateAttributeResponse, error) {
	attribute := models.PbToAttribute(req, UserId(ctx))

	err := CheckAttributeRequirements(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	//insert into attribute
	err = insertAttribute(ctx, attribute)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateAttributeResponse{AttributeId: pbutil.ToProtoString(attribute.AttributeId)}, nil
}

func (s *Server) DescribeAttributes(ctx context.Context, req *pb.DescribeAttributesRequest) (*pb.DescribeAttributesResponse, error) {
	attributes, err := getAttributes(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	var pbAttributes []*pb.Attribute
	for _, att := range attributes {
		pbAttributes = append(pbAttributes, models.AttributeToPb(att))
	}

	return &pb.DescribeAttributesResponse{AttributeSet: pbAttributes}, nil
}

func (s *Server) ModifyAttribute(ctx context.Context, req *pb.ModifyAttributeRequest) (*pb.ModifyAttributeResponse, error) {
	attribute, err := CheckAttributePermission(ctx, req.GetAttributeId().GetValue())
	if err != nil {
		return nil, err
	}

	attribute.AttributeTermId = req.GetAttributeTermId().GetValue()
	attribute.AttributeUnitId = req.GetAttributeUnitId().GetValue()
	err = CheckAttributeRequirements(ctx, *attribute)
	if err != nil {
		return nil, err
	}

	err = updateAttribute(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	return &pb.ModifyAttributeResponse{AttributeId: req.GetAttributeId()}, nil
}

func (s *Server) DeleteAttributes(ctx context.Context, req *pb.DeleteAttributesRequest) (*pb.DeleteAttributesResponse, error) {
	attributeIds := req.GetAttributeIds()
	_, err := CheckAttributesPermission(ctx, attributeIds)
	if err != nil {
		return nil, err
	}

	err = deleteAttributes(ctx, attributeIds)

	return &pb.DeleteAttributesResponse{AttributeIds: req.GetAttributeIds()}, nil
}

func (s *Server) CreateSpu(ctx context.Context, req *pb.CreateSpuRequest) (*pb.CreateSpuResponse, error) {
	//TODO: get id of current user
	spu := models.PbToSpu(req, UserId(ctx))

	//insert spu
	err := insertSpu(ctx, spu)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateSpuResponse{SpuId: pbutil.ToProtoString(spu.SpuId)}, nil
}

func (s *Server) DescribeSpus(ctx context.Context, req *pb.DescribeSpusRequest) (*pb.DescribeSpusResponse, error) {
	//TODO: impl DescribeSpus
	return &pb.DescribeSpusResponse{}, nil
}

func (s *Server) ModifySpu(ctx context.Context, req *pb.ModifySpuRequest) (*pb.ModifySpuResponse, error) {
	//TODO: impl ModifySpu
	return &pb.ModifySpuResponse{}, nil
}

func (s *Server) DeleteSpus(ctx context.Context, req *pb.DeleteSpusRequest) (*pb.DeleteSpusResponse, error) {
	//TODO: impl DeleteSpus
	return &pb.DeleteSpusResponse{}, nil
}

func (s *Server) CreateSku(ctx context.Context, req *pb.CreateSkuRequest) (*pb.CreateSkuResponse, error) {
	sku := models.PbToSku(req)

	//get spu
	spu, err := getSpu(ctx, sku.SpuId)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	if spu == nil {
		return nil, notExistError(ctx, models.Spu{}, sku.SpuId)
	}

	//TODO: check attribute_ids

	//insert sku
	err = insertSku(ctx, sku)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateSkuResponse{SkuId: pbutil.ToProtoString(sku.SkuId)}, nil
}

func (s *Server) DescribeSkus(ctx context.Context, req *pb.DescribeSkusRequest) (*pb.DescribeSkusResponse, error) {
	//TODO: impl DescribeSkus
	return &pb.DescribeSkusResponse{}, nil
}

func (s *Server) ModifySku(ctx context.Context, req *pb.ModifySkuRequest) (*pb.ModifySkuResponse, error) {
	//TODO: impl ModifySku
	return &pb.ModifySkuResponse{}, nil
}

func (s *Server) DeleteSkus(ctx context.Context, req *pb.DeleteSkusRequest) (*pb.DeleteSkusResponse, error) {
	//TODO: impl DeleteSkus
	return &pb.DeleteSkusResponse{}, nil
}

func renewalTimeFromSku(ctx context.Context, skuId string, actionTime time.Time) (*time.Time, error) {
	sku, err := getSku(ctx, skuId)

	if err != nil {
		logger.Error(ctx, "Failed to convert renewal time from sku, Error: [%+v]", err)
		return nil, err
	}

	//TODO: calculate renewalTime
	//TODO: check if duration in metering_attributes
	renewalTime := sku.CreateTime

	return &renewalTime, nil
}
