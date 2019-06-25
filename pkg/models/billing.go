// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"time"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/util/idutil"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func NewPriceId() string {
	return idutil.GetUuid("price-")
}

func NewContractId() string {
	return idutil.GetUuid("contract-")
}

type PriceItem struct {
	Price uint64
	Range map[string]string
}

func pbPricesToPriceItems(pbPrices []*pb.PriceItem) []PriceItem {
	priceItems := []PriceItem{}
	for _, item := range pbPrices {
		priceItems = append(priceItems, PriceItem{
			Price: item.GetPrice().GetValue(),
			Range: item.GetRange(),
		})
	}
	return priceItems
}

func priceItemsToPbPrices(priceItems []PriceItem) []*pb.PriceItem {
	pbPriceItems := []*pb.PriceItem{}
	for _, item := range priceItems {
		pbPriceItems = append(pbPriceItems, &pb.PriceItem{
			Price: pbutil.ToProtoUInt64(item.Price),
			Range: item.Range,
		})
	}
	return pbPriceItems
}

type Price struct {
	PriceId      string
	PriceName    string
	SkuId        string
	AttributeIds []string
	Prices       []PriceItem
	PricePolicy  string
	Currency     string
	Status       string
	StartTime    time.Time
	EndTime      time.Time
	CreateTime   time.Time
	StatusTime   time.Time
}

func NewPrice(priceName, skuId, policy, currency string, attributeIds []string, prices []PriceItem, startTime, endTime time.Time) *Price {
	now := time.Now()
	if (time.Time{}) == startTime {
		startTime = now
	}
	return &Price{
		PriceId:      NewPriceId(),
		PriceName:    priceName,
		SkuId:        skuId,
		AttributeIds: attributeIds,
		Prices:       prices,
		PricePolicy:  policy,
		Currency:     currency,
		Status:       constants.StatusActive,
		StartTime:    startTime,
		EndTime:      endTime,
		CreateTime:   now,
		StatusTime:   now,
	}
}

func PbToPrice(pbPrice *pb.CreatePriceRequest) *Price {
	return NewPrice(
		pbPrice.GetPriceName().GetValue(),
		pbPrice.GetSkuId().GetValue(),
		pbPrice.GetPricePolicy().GetValue(),
		pbPrice.GetCurrency().String(),
		pbPrice.GetAttributeIds(),
		pbPricesToPriceItems(pbPrice.GetPrices()),
		pbutil.FromProtoTimestamp(pbPrice.GetStartTime()),
		pbutil.FromProtoTimestamp(pbPrice.GetEndTime()),
	)
}

func PriceToPb(price *Price) *pb.Price {
	return &pb.Price{
		PriceId:      pbutil.ToProtoString(price.PriceId),
		PriceName:    pbutil.ToProtoString(price.PriceName),
		SkuId:        pbutil.ToProtoString(price.SkuId),
		AttributeIds: price.AttributeIds,
		Prices:       priceItemsToPbPrices(price.Prices),
		PricePolicy:  pbutil.ToProtoString(price.PricePolicy),
		Currency:     pb.Currency(pb.Currency_value[price.Currency]),
		StartTime:    pbutil.ToProtoTimestamp(price.StartTime),
		EndTime:      pbutil.ToProtoTimestamp(price.EndTime),
		CreateTime:   pbutil.ToProtoTimestamp(price.CreateTime),
		StatusTime:   pbutil.ToProtoTimestamp(price.StatusTime),
	}
}

type LeasingContract struct {
	ContractId       string
	LeasingId        string
	ResourceId       string
	SkuId            string
	UserId           string
	MeteringValues   map[string]float64
	FeeInfo          string
	Fee              float64
	DueFee           float64
	OtherContractFee float64
	CouponFee        float64
	RealFee          float64
	Currency         string
	Status           string //active/updating/deleted
	StartTime        time.Time
	StatusTime       time.Time
	CreateTime       time.Time
}

func NewLeasingContract(leasingId, resourceId, skuId, userId, currency string,
	meteringValues map[string]float64,
	startTime, updateDurationTime time.Time) *LeasingContract {

	now := time.Now()
	return &LeasingContract{
		ContractId:     NewContractId(),
		LeasingId:      leasingId,
		ResourceId:     resourceId,
		SkuId:          skuId,
		UserId:         userId,
		MeteringValues: meteringValues,
		Status:         constants.StatusActive,
		StartTime:      startTime,
		StatusTime:     now,
		CreateTime:     now,
		Currency:       currency,
	}
}

type LeasedContract struct {
	ContractId       string
	LeasingId        string
	ResourceId       string
	SkuId            string
	UserId           string
	MeteringValues   map[string]float64
	FeeInfo          string
	Fee              float64
	DueFee           float64
	OtherContractFee float64
	CouponFee        float64
	RealFee          float64
	Currency         string
	StartTime        time.Time
	EndTime          time.Time
	CreateTime       time.Time
}

func (leasingContract LeasingContract) ToLeasedContract() *LeasedContract {
	return &LeasedContract{
		ContractId:       leasingContract.ContractId,
		LeasingId:        leasingContract.LeasingId,
		ResourceId:       leasingContract.ResourceId,
		SkuId:            leasingContract.SkuId,
		UserId:           leasingContract.UserId,
		MeteringValues:   leasingContract.MeteringValues,
		FeeInfo:          leasingContract.FeeInfo,
		Fee:              leasingContract.Fee,
		DueFee:           leasingContract.DueFee,
		OtherContractFee: leasingContract.OtherContractFee,
		CouponFee:        leasingContract.CouponFee,
		RealFee:          leasingContract.RealFee,
		Currency:         leasingContract.Currency,
		StartTime:        leasingContract.StartTime,
		EndTime:          leasingContract.StatusTime,
		CreateTime:       time.Now(),
	}
}

type Account struct {
	UserId     string
	UserType   string
	Balance    float64
	Currency   string
	Income     map[string]float64
	Status     string
	CreateTime time.Time
	StatusTime time.Time
}

type Charge struct {
	ChargeId   string
	ContractId string
	Fee        float64
	Currency   string
	Status     string
	CreateTime time.Time
}

type Refund struct {
	RefundId   string
	ContractId string
	Fee        float64
	Currency   string
	Status     string
	CreateTime time.Time
}

type Recharge struct {
	ReChargeId  string
	Balance     float64
	Currency    string
	Status      string
	CreateTime  time.Time
	Description string
}

type Income struct {
	IncomeId   string
	ContractId string
	Balance    string
	Currency   string
	CreateTime time.Time
}
