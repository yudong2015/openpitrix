// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import "time"

type Charge struct {
	ChargeId   string
	ContractId string
	UserId     string
	Fee        float32
	Currency   string
	CreateTime time.Time
}

type ReCharge struct {
	ReChargeId string
	ContractId string
	UserId     string
	Fee        float32
	Currency   string
	Operator   string
	CreateTime time.Time
	remark     string
}

type Income struct {
	IncomeId string
}
