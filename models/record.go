package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const (
	CollectionNameRecord = "record"
	RecordIdTag          = "id"
)

type (
	Record struct {
		ID          bson.ObjectId `bson:"_id"`
		Id          string        `bson:"id"` //record id
		TimeStamp   int64         `bson:"timestamp"`
		TxHash      string        `bson:"tx_hash"`
		BlockHeight int64         `bson:"block_height"`
		Contents    []Content     `bson:"contents"`
		Source      Source        `bson:"source"`
		Creator     string        `bson:"creator"` //record creator in global chain (ibc record signer)
	}
	Content struct {
		DigestAlgo string `bson:"digest_algo"`
		Digest     string `bson:"digest"`
		Uri        string `bson:"uri"` //offchain object id
		Meta       string `bson:"meta"`
	}
	Source struct {
		Locality    string `bson:"locality" json:"locality"`
		BlockHeight int64  `bson:"block_height" json:"block_height"` //local chain block height
		TxHash      string `bson:"tx_hash" json:"tx_hash"`           //local chain tx hash
		Creator     string `bson:"creator" json:"creator"`           //record creator
	}
)

func (r Record) Name() string {
	return CollectionNameRecord
}

func (r Record) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{RecordIdTag},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(r.Name(), indexes)
}

func (r Record) PkKvPair() map[string]interface{} {
	return bson.M{RecordIdTag: r.Id}
}
