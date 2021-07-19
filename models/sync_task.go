package models

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNameSyncTask = "sync_task"

	// value of status
	SyncTaskStatusCatchUping = "catchuping"
	SyncTaskStatusUnderway   = "underway"
)

type (
	SyncTask struct {
		EndHeight int64  `bson:"end_height"` // task end height
		Status    string `bson:"status"`     // task status
		CreateAt  int64  `bson:"create_at"`
		UpdateAt  int64  `bson:"update_at"` // unix timestamp
	}
)

func (d SyncTask) Name() string {
	if GetSrvConf().ChainId == "" {
		return CollectionNameSyncTask
	}
	return fmt.Sprintf("sync_%v_task", GetSrvConf().ChainId)
}

func (d SyncTask) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-end_height"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d SyncTask) PkKvPair() map[string]interface{} {
	return bson.M{"end_height": d.EndHeight}
}

// query valid follow way
func (d SyncTask) QueryValidFollow() (bool, error) {
	var syncTasks []SyncTask
	q := bson.M{}

	q["status"] = SyncTaskStatusUnderway

	q["end_height"] = bson.M{
		"$eq": 0,
	}

	fn := func(c *mgo.Collection) error {
		return c.Find(q).All(&syncTasks)
	}

	err := ExecCollection(d.Name(), fn)

	if err != nil {
		return false, err
	}

	if len(syncTasks) == 1 {
		return true, nil
	}

	return false, nil
}

func (d SyncTask) Create() error {
	timenow := time.Now().Unix()
	task := &SyncTask{
		EndHeight: 0,
		Status:    SyncTaskStatusCatchUping,
		CreateAt:  timenow,
		UpdateAt:  timenow,
	}
	fn := func(c *mgo.Collection) error {
		pk := task.PkKvPair()
		n, _ := c.Find(pk).Count()
		if n >= 1 {
			return fmt.Errorf("record exists while save record")
		}
		return c.Insert(task)
	}
	return ExecCollection(d.Name(), fn)
}

func (d SyncTask) Update(oldStatus, newStatus string) error {
	fn := func(c *mgo.Collection) error {
		err := c.Update(
			bson.M{
				"end_height": 0,
				"status":     oldStatus,
			},
			bson.M{
				"$set": bson.M{
					"status":    newStatus,
					"update_at": time.Now().Unix(),
				},
			})
		return err
	}
	return ExecCollection(d.Name(), fn)
}
