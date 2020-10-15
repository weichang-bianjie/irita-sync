package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameNftHistory = "nft_history"

	NftHistoryTxHashTag   = "tx_hash"
	NftHistoryIdTag       = "id"
	NftHistoryDenomIdTag  = "denom_id"
	NftHistoryTxIndexTag  = "tx_index"
	NftHistoryMsgIndexTag = "msg_index"
)

type (
	NftHistory struct {
		Id          string `bson:"id"`
		NftName     string `bson:"name"`
		Owner       string `bson:"owner"`
		Uri         string `bson:"uri"`
		Data        string `bson:"data"`
		DenomId     string `bson:"denom_id"`
		TxHash      string `bson:"tx_hash"`
		TxType      string `bson:"tx_type"`
		BlockHeight int64  `bson:"block_height"`
		RecordId    string `bson:"record_id"`
		TxIndex     int    `bson:"tx_index"`
		MsgIndex    int    `bson:"msg_index"`
	}
)

func (m NftHistory) Name() string {
	return CollectionNameNftHistory
}

func (m NftHistory) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{NftHistoryTxHashTag, NftHistoryTxIndexTag, NftHistoryMsgIndexTag},
		Unique:     true,
		Background: true,
	})
	indexes = append(indexes, mgo.Index{
		Key:        []string{NftHistoryIdTag, NftHistoryDenomIdTag},
		Background: true,
	})
	ensureIndexes(m.Name(), indexes)
}

func (m NftHistory) PkKvPair() map[string]interface{} {
	return bson.M{NftHistoryTxHashTag: m.TxHash}
}

func (m NftHistory) Save(history NftHistory) error {
	return Save(history)
}
