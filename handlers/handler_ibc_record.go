package handlers

import (
	"github.com/bianjieai/irita-sync/models"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2/bson"
)

func handleIbcRecordData(docTx *models.Tx, recordInfo *models.Record) (Ops []txn.Op) {
	recordInfo.TxHash = docTx.Hash
	recordInfo.BlockHeight = docTx.BlockHeight
	recordInfo.TimeStamp = docTx.TimeStamp
	recordInfo.ID = bson.NewObjectId()
	op := txn.Op{
		C:      models.CollectionNameRecord,
		Id:     bson.NewObjectId(),
		Insert: recordInfo,
	}
	Ops = append(Ops, op)
	return
}

