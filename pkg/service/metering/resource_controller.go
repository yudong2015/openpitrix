// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package main

import (
	"context"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/util/stringutil"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pb"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/pbutil"
)

func CheckExist(ctx context.Context, table string, id string) (bool, error) {
	idName := stringutil.CamelCaseToUnderscore(table) + "_id"
	count, err := pi.Global().DB(ctx).
		Select(idName).
		From(table).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(idName, id)).
		Limit(1).Count()

	if err != nil {
		logger.Error(ctx, "Failed to connect to database: [%+v]", err)
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	logger.Info(ctx, "The %s not exist in %s!", id, table)

	return false, nil
}

func CheckAttributeRequirements(ctx context.Context, attribute models.Attribute) error {
	//check if attribute_term exist
	exist, err := CheckExist(ctx, constants.TableAttributeTerm, attribute.AttributeTermId)
	if err != nil {
		return internalError(ctx, err)
	}
	if !exist {
		return gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "attribute_term", "attribute_term", "属性项")
	}

	//check if attribute_unit exist
	exist, err = CheckExist(ctx, constants.TableAttributeUnit, attribute.AttributeUnitId)
	if err != nil {
		return internalError(ctx, err)
	}
	if !exist {
		return gerr.New(ctx, gerr.NotFound, gerr.ErrorNotExist, "attribute_unit", "attribute_unit", "属性项")
	}
	return nil
}

func insertAttributeTerm(ctx context.Context, term *models.AttributeTerm) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableAttributeTerm).
		Record(term).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert attribute_term: [%+v]", err)
	}
	return err
}

func getAttributeTermsByIds(ctx context.Context, ids []string) ([]*models.AttributeTerm, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var attributeTerms []*models.AttributeTerm
	_, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttributeTerm]...).
		From(constants.TableAttributeTerm).
		Where(db.Eq(constants.ColumnAttributeTermId, ids)).
		Load(&attributeTerms)

	if err != nil {
		logger.Error(ctx, "Failed to get attribute_term by ids: [%+v]", err)
	}
	return attributeTerms, err
}

func getAttributeTerms(ctx context.Context, req *pb.DescribeAttributeTermsRequest) (int, []*models.AttributeTerm, error) {
	var attributeTerms []*models.AttributeTerm
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttributeTerm]...).
		From(constants.TableAttributeTerm).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TableAttributeTerm)).
		OrderDir(pbutil.GetSortKey(req), !req.GetReverse().GetValue()).
		Offset(pbutil.GetOffsetFromRequest(req)).
		Limit(pbutil.GetLimitFromRequest(req)).
		Load(&attributeTerms)

	if err != nil {
		logger.Error(ctx, "Failed to get attribute_terms: [%+v]", err)
	}
	return count, attributeTerms, err
}

func updateAttributeTerm(ctx context.Context, req *pb.ModifyAttributeTermRequest) error {
	buildedAttributes := manager.BuildUpdateAttributes(req,
		constants.ColumnName,
		constants.ColumnDescription)

	var err error
	if len(buildedAttributes) > 0 {
		_, err = pi.Global().DB(ctx).
			Update(constants.TableAttributeTerm).
			SetMap(buildedAttributes).
			Where(db.Eq(constants.ColumnAttributeTermId, req.GetAttributeTermId().GetValue())).
			Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
			Exec()

		if err != nil {
			logger.Error(ctx, "Failed to update attribute_term: [%+v]", err)
		}
	}

	return err
}

func deleteAttributeTerms(ctx context.Context, attTermIds []string) error {
	_, err := pi.Global().DB(ctx).
		Update(constants.TableAttributeTerm).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnAttributeTermId, attTermIds)).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to delete attribute_term(set status to deleted): [%+v]", err)
	}

	return err
}

func insertAttributeUnit(ctx context.Context, attUnit *models.AttributeUnit) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableAttributeUnit).
		Record(attUnit).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert attribute_unit: [%+v]", err)
	}
	return err
}

func getAttributeUnits(ctx context.Context, req *pb.DescribeAttributeUnitsRequest) (int, []*models.AttributeUnit, error) {
	var attUnits []*models.AttributeUnit
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttributeUnit]...).
		From(constants.TableAttributeUnit).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TableAttributeUnit)).
		OrderDir(pbutil.GetSortKey(req), !req.GetReverse().GetValue()).
		Offset(pbutil.GetOffsetFromRequest(req)).
		Limit(pbutil.GetLimitFromRequest(req)).
		Load(&attUnits)

	if err != nil {
		logger.Error(ctx, "Failed to get attribute_units: [%+v]", err)
	}
	return count, attUnits, err
}

func getAttributeUnitsByIds(ctx context.Context, attributeUnitIds []string) ([]*models.AttributeUnit, error) {
	var attUnits []*models.AttributeUnit
	_, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttributeUnit]...).
		From(constants.TableAttributeUnit).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnAttributeUnitId, attributeUnitIds)).
		Load(&attUnits)

	if err != nil {
		logger.Error(ctx, "Failed to get attribute_units: [%+v]", err)
	}
	return attUnits, err
}

func deleteAttributeUnits(ctx context.Context, attUnitIds []string) error {
	_, err := pi.Global().DB(ctx).
		Update(constants.TableAttributeUnit).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnAttributeUnitId, attUnitIds)).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to delete attribute_unit(set status to deleted): [%+v]", err)
	}

	return err
}

func insertAttribute(ctx context.Context, attribute *models.Attribute) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableAttribute).
		Record(attribute).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert attribute: [%+v]", err)
	}
	return err
}

func getAttributes(ctx context.Context, req *pb.DescribeAttributesRequest) (int, []*models.Attribute, error) {
	var attributes []*models.Attribute
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttribute]...).
		From(constants.TableAttribute).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TableAttribute)).
		OrderDir(pbutil.GetSortKey(req), !req.GetReverse().GetValue()).
		Offset(pbutil.GetOffsetFromRequest(req)).
		Limit(pbutil.GetLimitFromRequest(req)).
		Load(&attributes)

	if err != nil {
		logger.Error(ctx, "Failed to get attributes: [%+v]", err)
	}
	return count, attributes, err
}

func getAttributesByIds(ctx context.Context, attIds []string) ([]*models.Attribute, error) {
	var attributes []*models.Attribute
	_, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableAttribute]...).
		From(constants.TableAttribute).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnAttributeId, attIds)).
		Load(&attributes)

	if err != nil {
		logger.Error(ctx, "Failed to get attributes: [%+v]", err)
	}
	return attributes, err
}

func updateAttribute(ctx context.Context, req *pb.ModifyAttributeRequest) error {
	buildAttributes := manager.BuildUpdateAttributes(req,
		constants.ColumnAttributeTermId,
		constants.ColumnAttributeUnitId,
		constants.ColumnValue,
		constants.ColumnValueRange)

	_, err := pi.Global().DB(ctx).
		Update(constants.TableAttribute).
		SetMap(buildAttributes).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnAttributeId, req.GetAttributeId().GetValue())).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to update attribute: [%+v]", err)
	}

	return err
}

func deleteAttributes(ctx context.Context, attributeIds []string) error {
	_, err := pi.Global().DB(ctx).
		Update(constants.TableAttribute).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnAttributeId, attributeIds)).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to delete attribute(set status to deleted): [%+v]", err)
	}

	return err
}

func insertSpu(ctx context.Context, spu *models.Spu) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableSpu).
		Record(spu).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert spu: [%+v]", err)
	}
	return err
}

func getSpus(ctx context.Context, req *pb.DescribeSpusRequest) (int, []*models.Spu, error) {
	var spus []*models.Spu
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableSpu]...).
		From(constants.TableSpu).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TableSpu)).
		OrderDir(pbutil.GetSortKey(req), !req.GetReverse().GetValue()).
		Offset(pbutil.GetOffsetFromRequest(req)).
		Limit(pbutil.GetLimitFromRequest(req)).
		Load(&spus)

	if err != nil {
		logger.Error(ctx, "Failed to get spus: [%+v]", err)
	}
	return count, spus, err
}

func getSpusByIds(ctx context.Context, spuIds []string) ([]*models.Spu, error) {
	var spus []*models.Spu
	_, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableSpu]...).
		From(constants.TableSpu).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnSpuId, spuIds)).
		Load(&spus)

	if err != nil {
		logger.Error(ctx, "Failed to get spus: [%+v]", err)
	}
	return spus, err
}

func deleteSpus(ctx context.Context, spuIds []string) error {
	tx, err := pi.Global().DB(ctx).Begin()
	defer tx.RollbackUnlessCommitted()
	if err != nil {
		logger.Error(ctx, "Failed to delete spus: [%+v]", err)
		return err
	}

	tx.Update(constants.TableSpu).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnSpuId, spuIds))

	tx.Update(constants.TableSku).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnSpuId, spuIds))

	err = tx.Commit()

	if err != nil {
		logger.Error(ctx, "Failed to delete spus: [%+v]", err)
		return err
	}
	return err
}

func insertSku(ctx context.Context, sku *models.Sku) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableSku).
		Record(sku).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert sku: [%+v]", err)
	}
	return err
}

func getSkus(ctx context.Context, req *pb.DescribeSkusRequest) (int, []*models.Sku, error) {
	var skus []*models.Sku
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TableSku]...).
		From(constants.TableSku).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TableSku)).
		OrderDir(pbutil.GetSortKey(req), !req.GetReverse().GetValue()).
		Offset(pbutil.GetOffsetFromRequest(req)).
		Limit(pbutil.GetLimitFromRequest(req)).
		Load(&skus)

	if err != nil {
		logger.Error(ctx, "Failed to get skus: [%+v]", err)
	}
	return count, skus, err
}

func getSkusByIds(ctx context.Context, skuIds, columns []string) ([]*models.Sku, error) {
	var skus []*models.Sku
	if len(columns) == 0 {
		columns = constants.IndexedColumns[constants.TableSku]
	}
	_, err := pi.Global().DB(ctx).
		Select(columns...).
		From(constants.TableSku).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnSkuId, skuIds)).
		Load(&skus)

	if err != nil {
		logger.Error(ctx, "Failed to get skus: [%+v]", err)
	}
	return skus, err
}

func updateSku(ctx context.Context, req *pb.ModifySkuRequest) error {
	buildAttributes := manager.BuildUpdateAttributes(req,
		constants.ColumnAttributeIds,
		constants.ColumnMeteringAttributeIds,
		constants.ColumnFeePolicy)

	_, err := pi.Global().DB(ctx).
		Update(constants.TableSku).
		SetMap(buildAttributes).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(constants.ColumnSkuId, req.GetSkuId().GetValue())).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to update sku: [%+v]", err)
	}

	return err
}

func deleteSkus(ctx context.Context, skuIds []string) error {
	_, err := pi.Global().DB(ctx).
		Update(constants.TableSku).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnSkuId, skuIds)).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to delete skus: [%+v]", err)
	}
	return err
}

//Metering
func insertLeasings(ctx context.Context, leasings []*models.Leasing) error {
	result, err := pi.Global().DB(ctx).
		InsertInto(constants.TableLeasing).
		Record(leasings).Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert leasings, Error: [%+v].", err)
	} else {
		count, _ := result.RowsAffected()
		logger.Info(ctx, "Insert %d leasings successfully.", count)
	}
	return err
}

func getLeasing(ctx context.Context, leasingID, resourceId, skuId string) (*models.Leasing, error) {
	//ids: leasingId, resourceId, skuId
	//TODO: impl get leasing
	return &models.Leasing{}, nil
}

//promotion
func insertCombination(ctx context.Context, com *models.Combination) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableCombination).
		Record(com).Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert combination, Error: [%+v].", err)
	} else {
		logger.Info(ctx, "Insert combination successfully.")
	}
	return err
}

func insertCombinationSku(ctx context.Context, comSku *models.CombinationSku) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableCombinationSku).
		Record(comSku).Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert combination_sku: [%+v].", err)
	} else {
		logger.Info(ctx, "Insert combination_sku successfully.")
	}
	return err
}

func leasingToRedis(leasing models.Leasing) error {
	//TODO: impl add leasing to redis
	return nil
}
