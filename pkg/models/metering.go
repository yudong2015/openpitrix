// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"time"

	"openpitrix.io/openpitrix/pkg/pb"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"

	"openpitrix.io/openpitrix/pkg/util/idutil"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func newAttributeTermId() string {
	return idutil.GetUuid("term-")
}

func newAttributeUnitId() string {
	return idutil.GetUuid("unit-")
}

func newAttributeId() string {
	return idutil.GetUuid("att-")
}

func newSpuId() string {
	return idutil.GetUuid("spu-")
}

func newSkuId() string {
	return idutil.GetUuid("sku-")
}

func newLeasingId() string {
	return idutil.GetUuid("leasing-")
}

func newMeteringId() string {
	return idutil.GetUuid("metering-")
}

type AttributeTerm struct {
	AttributeTermId string
	Name            string
	Description     string
	Type            string
	Provider        string
	Status          string
	CreateTime      time.Time
	StatusTime      time.Time
}

func NewAttributeTerm(name, termType, description, provider string) *AttributeTerm {
	return &AttributeTerm{
		AttributeTermId: newAttributeTermId(),
		Name:            name,
		Description:     description,
		Type:            termType,
		Provider:        provider,
		Status:          constants.StatusActive,
	}
}

func PbToAttributeTerm(pbAttTerm *pb.CreateAttributeTermRequest, provider string) *AttributeTerm {
	return NewAttributeTerm(
		pbAttTerm.GetName().GetValue(),
		pbAttTerm.GetDescription().GetValue(),
		pbAttTerm.GetType().String(),
		provider,
	)
}

func AttributeTermToPb(attName *AttributeTerm) *pb.AttributeTerm {
	return &pb.AttributeTerm{
		AttributeTermId: pbutil.ToProtoString(attName.AttributeTermId),
		Name:            pbutil.ToProtoString(attName.Name),
		Description:     pbutil.ToProtoString(attName.Description),
		Type:            pb.AttributeTermType(pb.AttributeTermType_value[attName.Type]),
		CreateTime:      pbutil.ToProtoTimestamp(attName.CreateTime),
		StatusTime:      pbutil.ToProtoTimestamp(attName.StatusTime),
	}
}

type AttributeUnit struct {
	AttributeUnitId string
	Name            string
	Provider        string
	Status          string
	CreateTime      time.Time
	StatusTime      time.Time
}

func NewAttributeUnit(name, provider string) *AttributeUnit {
	return &AttributeUnit{
		AttributeUnitId: newAttributeUnitId(),
		Name:            name,
		Provider:        provider,
		Status:          constants.StatusActive,
	}
}

func PbToAttributeUnit(pbAttUnit *pb.CreateAttributeUnitRequest, providor string) *AttributeUnit {
	return NewAttributeUnit(
		pbAttUnit.GetName().GetValue(),
		providor,
	)
}

func AttributeUnitToPb(attUnit *AttributeUnit) *pb.AttributeUnit {
	return &pb.AttributeUnit{
		AttributeUnitId: pbutil.ToProtoString(attUnit.AttributeUnitId),
		Name:            pbutil.ToProtoString(attUnit.Name),
		CreateTime:      pbutil.ToProtoTimestamp(attUnit.CreateTime),
		StatusTime:      pbutil.ToProtoTimestamp(attUnit.StatusTime),
	}
}

type Attribute struct {
	AttributeId     string
	AttributeTermId string
	AttributeUnitId string
	Value           string
	ValueRange      string
	Provider        string
	Status          string
	CreateTime      time.Time
	StatusTime      time.Time
}

var AttributeColumns = db.GetColumnsFromStruct(&Attribute{})

func NewAttribute(attNameId, attUnitId, value, valueRange, provider string) *Attribute {
	return &Attribute{
		AttributeId:     newAttributeId(),
		AttributeTermId: attNameId,
		AttributeUnitId: attUnitId,
		Value:           value,
		ValueRange:      valueRange,
		Provider:        provider,
		Status:          constants.StatusActive,
	}
}

func PbToAttribute(pbAttribute *pb.CreateAttributeRequest, provider string) *Attribute {
	return NewAttribute(
		pbAttribute.GetAttributeTermId().GetValue(),
		pbAttribute.GetAttributeUnitId().GetValue(),
		pbAttribute.GetValue().GetValue(),
		pbAttribute.GetValueRange().GetValue(),
		provider,
	)
}

func AttributeToPb(att *Attribute) *pb.Attribute {
	return &pb.Attribute{
		AttributeId:     pbutil.ToProtoString(att.AttributeId),
		AttributeTermId: pbutil.ToProtoString(att.AttributeTermId),
		AttributeUnitId: pbutil.ToProtoString(att.AttributeUnitId),
		Value:           pbutil.ToProtoString(att.Value),
		ValueRange:      pbutil.ToProtoString(att.ValueRange),
		CreateTime:      pbutil.ToProtoTimestamp(att.CreateTime),
		StatusTime:      pbutil.ToProtoTimestamp(att.StatusTime),
	}
}

//SPU: standard product unit
type Spu struct {
	SpuId      string
	ProductId  string
	Owner      string
	Status     string
	CreateTime time.Time
	StatusTime time.Time
}

func NewSpu(productId, owner string) *Spu {
	return &Spu{
		SpuId:     newSpuId(),
		ProductId: productId,
		Owner:     owner,
		Status:    constants.StatusActive,
	}
}

func PbToSpu(pbSpu *pb.CreateSpuRequest, owner string) *Spu {
	return NewSpu(pbSpu.GetProductId().GetValue(), owner)
}

func SpuToPb(spu *Spu) *pb.Spu {
	return &pb.Spu{
		SpuId:      pbutil.ToProtoString(spu.SpuId),
		ProductId:  pbutil.ToProtoString(spu.ProductId),
		Owner:      pbutil.ToProtoString(spu.Owner),
		CreateTime: pbutil.ToProtoTimestamp(spu.CreateTime),
		StatusTime: pbutil.ToProtoTimestamp(spu.StatusTime),
	}
}

//SKU: stock keeping unit
type Sku struct {
	SkuId                string
	SpuId                string
	AttributeIds         []string
	MeteringAttributeIds []string
	Status               string
	FeePolicy            string
	CreateTime           time.Time
	StatusTime           time.Time
}

func NewSku(spuId, feePolicy string, attributeIds, meteringAttributeIds []string) *Sku {
	return &Sku{
		SkuId:                newSkuId(),
		SpuId:                spuId,
		AttributeIds:         attributeIds,
		MeteringAttributeIds: meteringAttributeIds,
		Status:               constants.StatusActive,
		FeePolicy:            feePolicy,
	}
}

func PbToSku(pbSku *pb.CreateSkuRequest) *Sku {
	return NewSku(
		pbSku.GetSpuId().GetValue(),
		pbSku.GetFeePolicy().GetValue(),
		pbSku.GetAttributeIds(),
		pbSku.GetMeteringAttributeIds(),
	)
}

func SkuToPb(sku *Sku) *pb.Sku {
	return &pb.Sku{
		SkuId:                pbutil.ToProtoString(sku.SkuId),
		SpuId:                pbutil.ToProtoString(sku.SpuId),
		FeePolicy:            pbutil.ToProtoString(sku.FeePolicy),
		AttributeIds:         sku.AttributeIds,
		MeteringAttributeIds: sku.MeteringAttributeIds,
		CreateTime:           pbutil.ToProtoTimestamp(sku.CreateTime),
		StatusTime:           pbutil.ToProtoTimestamp(sku.StatusTime),
	}
}

type Leasing struct {
	LeasingId  string
	UserId     string
	ResourceId string
	CreateTime time.Time
	StatusTime time.Time
	Status     string
}

var LeasingColumns = db.GetColumnsFromStruct(&Leasing{})

func NewLeasing(userId, resourceId string) *Leasing {
	return &Leasing{
		LeasingId:  newLeasingId(),
		UserId:     userId,
		ResourceId: resourceId,
		Status:     constants.StatusActive,
	}
}

func PbToLeasing(req *pb.InitMeteringRequest) *Leasing {
	return NewLeasing(req.GetUserId().GetValue(), req.GetResourceId().GetValue())
}

type Leased struct {
	LeasedId          string
	UserId            string
	ResourceId        string
	LeasingCreateTime time.Time
	CreateTime        time.Time
}

func (l *Leasing) ToLeased() *Leased {
	return &Leased{
		LeasedId:          l.LeasingId,
		UserId:            l.UserId,
		ResourceId:        l.ResourceId,
		LeasingCreateTime: l.CreateTime,
	}
}

//SkuMetering
type Metering struct {
	MeteringId         string
	LeasingId          string
	SkuId              string
	Values             map[string]string //{attributeId: value, ..}
	StartTime          time.Time         //action_time
	UpdateDurationTime time.Time         //update time for duration
	RenewalTime        time.Time         //next update time
	CreateTime         time.Time
	StatusTime         time.Time               //update time by other services(cluster_manager)
	StopTimes          map[time.Time]time.Time //{closeTime: restartTime, ..}
	Status             string
}

func NewMetering(leasingId, skuId string, startTime time.Time, values map[string]string) *Metering {
	return &Metering{
		MeteringId:  newMeteringId(),
		LeasingId:   leasingId,
		SkuId:       skuId,
		Values:      values,
		StartTime:   startTime,
		Status:      constants.StatusActive,
	}
}

func PbToMetering(leasingId string, pbMetering *pb.SkuMetering) *Metering {
	return NewMetering(leasingId,
		pbMetering.GetSkuId().GetValue(),
		pbutil.FromProtoTimestamp(pbMetering.GetActionTime()),
		pbMetering.GetValue(),
	)
}
