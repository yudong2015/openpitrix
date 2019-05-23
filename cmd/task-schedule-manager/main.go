// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

//openpitrix mbing task manager
package main

import (
	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/service/task_schedule"
)

func main() {
	cfg := config.LoadConf()
	task_schedule.ExecutorServe(cfg)
}
