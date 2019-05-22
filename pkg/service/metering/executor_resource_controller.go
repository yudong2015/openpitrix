// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package metering

import (
	"context"
	"fmt"

	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/models"
)


func PublishTask(ctx context.Context, etcd *etcd.Etcd, task models.MbingTask) error {
	key := fmt.Sprintf(TaskInfoFmt, task.Id)
	return etcd.Dlock(ctx, key, func() error {
		taskStr, err := task.String()
		if err != nil {
			return err
		}
		_, err = etcd.Put(ctx, key, taskStr)
		return err
	})
}

func EnquequeTask(queue *etcd.Queue, taskId string) error {
	return queue.Enqueue(taskId)
}
