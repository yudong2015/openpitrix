// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package mbing

import (
	"openpitrix.io/openpitrix/pkg/models"
)

func charge(contract *models.LeasingContract) (string, error) {
	//TODO:
	return "Charge.id", nil
}

func reCharge(contract *models.LeasingContract) (string, error) {
	return "ReCharge.id", nil
}

func addCredit(userId string, currency string, fee float32) (string, error) {
	return "ReCharge.id", nil
}
