// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package billing

import (
	"context"
	"openpitrix.io/openpitrix/pkg/manager"
	"openpitrix.io/openpitrix/pkg/pb"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/logger"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
)

func updateStatusToDeleted(ctx context.Context, table string, ids []string) error {
	columnId := table + "_id"
	_, err := pi.Global().DB(ctx).
		Update(table).
		Set(constants.ColumnStatus, constants.StatusDeleted).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(db.Eq(columnId, ids)).
		Exec()

	if err != nil {
		logger.Error(ctx, "Failed to status of %s to deleted: [%+v]", table, err)
	}
	return err
}

func insertPrice(ctx context.Context, price *models.Price) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TablePrice).
		Record(price).
		Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert price, Errors: [%+v]", err)
	}
	return err
}

func getPrices(ctx context.Context, req *pb.DescribePricesRequest) (int, []*models.Price, error) {
	var prices []*models.Price
	count, err := pi.Global().DB(ctx).
		Select(constants.IndexedColumns[constants.TablePrice]...).
		From(constants.TablePrice).
		Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
		Where(manager.BuildFilterConditions(req, constants.TablePrice)).
		Load(&prices)

	if err != nil {
		logger.Error(ctx, "Failed to insert price, Errors: [%+v]", err)
	}
	return count, prices, err
}

func updatePrice(ctx context.Context, req *pb.ModifyPriceRequest) error {
	buildedAttributes := manager.BuildUpdateAttributes(req,
		constants.ColumnPriceName,
		constants.ColumnAttributeIds,
		constants.ColumnPrices,
		constants.ColumnPricePolicy)

	var err error
	if len(buildedAttributes) > 0 {
		_, err = pi.Global().DB(ctx).
			Update(constants.TablePrice).
			SetMap(buildedAttributes).
			Where(db.Eq(constants.ColumnStatus, constants.StatusActive)).
			Where(db.Eq(constants.ColumnPriceId, req.GetPriceId().GetValue())).
			Exec()

		if err != nil {
			logger.Error(ctx, "Failed to update attribute_term: [%+v]", err)
		}
	}

	return err
}

//promotion
func insertCombinationPrice(ctx context.Context, comPrice *models.CombinationPrice) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableCombinationPrice).
		Record(comPrice).Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert Combination_Price, Error: [%+v].", err)
	} else {
		logger.Info(ctx, "Insert Combination_Price successfully.")
	}
	return err
}

func insertProbation(ctx context.Context, probation *models.Probation) error {
	_, err := pi.Global().DB(ctx).
		InsertInto(constants.TableProbation).
		Record(probation).Exec()
	if err != nil {
		logger.Error(ctx, "Failed to insert probation, Error: [%+v].", err)
	} else {
		logger.Info(ctx, "Insert probation successfully.")
	}
	return err
}

func insertLeasingContract(ctx context.Context, contract *models.LeasingContract) error {
	//TODO: impl insert
	return nil
}

func getLeasingContract(ctx context.Context, contractId, leasingId string) (*models.LeasingContract, error) {
	//TODO: impl
	return &models.LeasingContract{}, nil
}

func insufficientBalanceToEtcd(resourceId, skuId, userId string) error {
	//TODO: add resourceId, skuId, userId to Etcd(insufficient queue)
	return nil
}
