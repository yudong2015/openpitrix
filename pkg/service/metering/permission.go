// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"

	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/util/ctxutil"
)

//TODO: manage relation ship of resource and user only in mbing-system
func UserId(ctx context.Context) string {
	return ctxutil.GetSender(ctx).UserId
}

func CheckAttributeTermPermission(ctx context.Context, attributeTermId string) (*models.AttributeTerm, error) {
	attTerms, err := getAttributeTermsByIds(ctx, []string{attributeTermId})
	if err != nil {
		return nil, err
	}
	return attTerms[0], nil
}

func CheckAttributeTermsPermission(ctx context.Context, attributeTermIds []string) ([]*models.AttributeTerm, error) {
	attTerms, err := getAttributeTermsByIds(ctx, attributeTermIds)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	if attTerms != nil || len(attTerms) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "AttributeTerm", "AttributeTerm", "属性项")
	}

	for _, attTerm := range attTerms {
		if attTerm.Provider != UserId(ctx) {
			return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, attTerm.AttributeTermId)
		}
	}

	return attTerms, nil
}

func CheckAttributeUnitsPermission(ctx context.Context, attributeUnitIds []string) ([]*models.AttributeUnit, error) {
	attUnits, err := getAttributeUnitsByIds(ctx, attributeUnitIds)
	if err != nil {
		return nil, internalError(ctx, err)
	}

	if attUnits != nil {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "AttributeUnit", "AttributeUnit", "属性单位")
	}

	for _, attUnit := range attUnits {
		if attUnit.Provider != UserId(ctx) {
			return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, attUnit.AttributeUnitId)
		}
	}

	return attUnits, nil
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
