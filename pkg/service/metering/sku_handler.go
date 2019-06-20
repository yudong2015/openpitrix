// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/gerr"
	"time"

	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

func (s *Server) CreateAttributeTerm(ctx context.Context, req *pb.CreateAttributeTermRequest) (*pb.CreateAttributeTermResponse, error) {
	attName := models.PbToAttributeTerm(req)
	err := insertAttributeTerm(ctx, attName)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.CreateAttributeTermResponse{AttributeTermId: pbutil.ToProtoString(attName.AttributeTermId)}, nil
}

func (s *Server) DescribeAttributeTerms(ctx context.Context, req *pb.DescribeAttributeTermsRequest) (*pb.DescribeAttributeTermsResponse, error) {
	attributeTerms, err := getAttributeTerms(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	var pbAttributeTerms []*pb.AttributeTerm
	for _, attName := range attributeTerms {
		pbAttributeTerms = append(pbAttributeTerms, models.AttributeTermToPb(attName))
	}

	return &pb.DescribeAttributeTermsResponse{AttributeTermSet: pbAttributeTerms}, nil
}

func (s *Server) ModifyAttributeTerm(ctx context.Context, req *pb.ModifyAttributeTermRequest) (*pb.ModifyAttributeTermResponse, error) {
	err := updateAttributeTerm(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.ModifyAttributeTermResponse{AttributeTermId: req.AttributeTermId}, nil
}

func (s *Server) DeleteAttributeTerms(ctx context.Context, req *pb.DeleteAttributeTermsRequest) (*pb.DeleteAttributeTermsResponse, error) {
	err := deleteAttributeTerms(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.DeleteAttributeTermsResponse{AttributeTermIds: req.AttributeTermIds}, nil
}

func (s *Server) CreateAttributeUnit(ctx context.Context, req *pb.CreateAttributeUnitRequest) (*pb.CreateAttributeUnitResponse, error) {
	attributeUnit := models.PbToAttributeUnit(req)
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

func (s *Server) ModifyAttributeUnit(ctx context.Context, req *pb.ModifyAttributeUnitRequest) (*pb.ModifyAttributeUnitResponse, error) {
	err := updateAttributeUnit(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.ModifyAttributeUnitResponse{AttributeUnitId: req.GetAttributeUnitId()}, nil
}

func (s *Server) DeleteAttributeUnits(ctx context.Context, req *pb.DeleteAttributeUnitsRequest) (*pb.DeleteAttributeUnitsResponse, error) {
	err := deleteAttributeUnits(ctx, req)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	return &pb.DeleteAttributeUnitsResponse{AttributeUnitIds: req.GetAttributeUnitIds()}, nil
}

func (s *Server) CreateAttribute(ctx context.Context, req *pb.CreateAttributeRequest) (*pb.CreateAttributeResponse, error) {
	userId := ctxutil.GetSender(ctx).UserId
	attribute := models.PbToAttribute(req, userId)

	//check if attribute_term exist
	exist, err := checkExistById(ctx, constants.TableAttributeTerm, attribute.AttributeTermId)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	if !exist {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "attribute_term", "attribute_term", "属性项")
	}

	//check if attribute_unit exist
	exist, err = checkExistById(ctx, constants.TableAttributeUnit, attribute.AttributeUnitId)
	if err != nil {
		return nil, internalError(ctx, err)
	}
	if !exist {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "attribute_unit", "attribute_unit", "属性项")
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
	//TODO: impl ModifyAttribute
	return &pb.ModifyAttributeResponse{}, nil
}

func (s *Server) DeleteAttributes(ctx context.Context, req *pb.DeleteAttributesRequest) (*pb.DeleteAttributesResponse, error) {
	//TODO: impl DeleteAttributes
	return &pb.DeleteAttributesResponse{}, nil
}

func (s *Server) CreateSpu(ctx context.Context, req *pb.CreateSpuRequest) (*pb.CreateSpuResponse, error) {
	//TODO: get id of current user
	spu := models.PbToSpu(req, "")

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
