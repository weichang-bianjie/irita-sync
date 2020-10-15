package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameTx = "tx"
	TxHashTag        = "hash"
)

type (
	Tx struct {
		//Time      int64       `bson:"time"`
		//Height    int64       `bson:"height"`
		//TxHash    string      `bson:"tx_hash"`
		//Type      string      `bson:"type"` // parse from first msg
		//Memo      string      `bson:"memo"`
		//Status    uint32      `bson:"status"`
		//Log       string      `bson:"log"`
		//Fee       *Fee        `bson:"fee"`
		//Types     []string    `bson:"types"`
		//Events    []Event     `bson:"events"`
		//Signers   []string    `bson:"signers"`
		//DocTxMsgs []DocTxMsg  `bson:"msgs"`
		//Addrs     []string    `bson:"addrs"`
		//Ext       interface{} `bson:"ext"`
		TimeStamp   int64      `bson:"timestamp"`
		BlockHeight int64      `bson:"block_height"`
		Hash        string     `bson:"hash"`
		Memo        string     `bson:"memo"`
		Status      uint32     `bson:"status"`
		Type        string     `bson:"type"` // parse from first msg
		Events      []Event    `bson:"events"`
		DocTxMsgs   []DocTxMsg `bson:"msgs"`
		Signers     []string   `bson:"signers"`
		Addrs       []string   `bson:"addrs"`
	}

	Event struct {
		Type       string   `bson:"type"`
		Attributes []KvPair `bson:"attributes"`
	}

	KvPair struct {
		Key   string `bson:"key"`
		Value string `bson:"value"`
	}

	DocTxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Fee struct {
		Amount []Coin `bson:"amount" json:"amount"`
		Gas    int64  `bson:"gas" json:"gas"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}
)

func (d Tx) Name() string {
	return CollectionNameTx
}

func (d Tx) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-tx_hash"},
		Unique:     true,
		Background: true,
	})
	indexes = append(indexes, mgo.Index{
		Key:        []string{"-height"},
		Unique:     true,
		Background: true,
	})

	ensureIndexes(d.Name(), indexes)
}

func (d Tx) PkKvPair() map[string]interface{} {
	return bson.M{TxHashTag: d.Hash}
}