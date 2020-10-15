package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameBlock = "block"
)

type (
	Block struct {
		Height    int64  `bson:"height"`
		Hash      string `bson:"hash"`
		Txn       int64  `bson:"txn"`
		TimeStamp int64  `bson:"timestamp"`
	}
)

func (d Block) Name() string {
	return CollectionNameBlock
}

func (d Block) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-height"},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(d.Name(), indexes)
}

func (d Block) PkKvPair() map[string]interface{} {
	return bson.M{"height": d.Height}
}
