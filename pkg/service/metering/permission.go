// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/models"

	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

//TODO: manage relation ship of resource and user only in mbing-system
func UserId(ctx context.Context) string {
	return ctxutil.GetSender(ctx).UserId
}

func CheckAttributeTermPermission(ctx context.Context, attributeTermId string) error {
	return CheckAttributeTermsPermission(ctx, []string{attributeTermId})
}

func CheckAttributeTermsPermission(ctx context.Context, attributeTermIds []string) error {
	attTerms, err := getAttributeTermsByIds(ctx, attributeTermIds)
	if err != nil {
		return internalError(ctx, err)
	}

	if attTerms != nil || len(attTerms) == 0 {
		return gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "AttributeTerm", "AttributeTerm", "属性项")
	}

	for _, attTerm := range attTerms {
		if attTerm.Provider != UserId(ctx) {
			return gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, attTerm.AttributeTermId)
		}
	}

	return nil
}

func CheckAttributeUnitsPermission(ctx context.Context, attributeUnitIds []string) error {
	attUnits, err := getAttributeUnitsByIds(ctx, attributeUnitIds)
	if err != nil {
		return internalError(ctx, err)
	}

	if attUnits != nil {
		return gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "AttributeUnit", "AttributeUnit", "属性单位")
	}

	for _, attUnit := range attUnits {
		if attUnit.Provider != UserId(ctx) {
			return gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, attUnit.AttributeUnitId)
		}
	}

	return nil
}

func CheckAttributePermission(ctx context.Context, attributeId string) (*models.Attribute, error) {
	atts, err := CheckAttributesPermission(ctx, []string{attributeId})
	if err != nil {
		return nil, err
	}
	return atts[0], nil
}

func CheckAttributesPermission(ctx context.Context, attributeIds []string) ([]*models.Attribute, error) {
	atts, err := getAttributesByIds(ctx, attributeIds)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	if atts != nil {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "Attribute", "Attribute", "属性")
	}

	for _, att := range atts {
		if att.Provider != UserId(ctx) {
			return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, att.AttributeUnitId)
		}
	}

	return atts, nil
}

func CheckSpuPermission(ctx context.Context, spuId string) error {
	return CheckSpusPermission(ctx, []string{spuId})
}

func CheckSpusPermission(ctx context.Context, spuIds []string) error {
	spus, err := getSpusByIds(ctx, spuIds)
	if err != nil {
		return internalError(ctx, err)
	}

	if spus != nil {
		return gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "Spu", "Spu", "Spu")
	}

	for _, spu := range spus {
		if spu.Owner != UserId(ctx) {
			return gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, spu.SpuId)
		}
	}

	return nil
}

func checkSkuPermission(ctx context.Context, skuId string) error {
	return checkSkusPermission(ctx, []string{skuId})
}

func checkSkusPermission(ctx context.Context, skuIds []string) error {
	skus, err := getSkusByIds(ctx, skuIds, []string{constants.ColumnSpuId})
	if err != nil {
		return internalError(ctx, err)
	}
	var spuIds []string
	for _, sku := range skus {
		spuIds = append(spuIds, sku.SpuId)
	}
	return CheckSpusPermission(ctx, spuIds)
}
