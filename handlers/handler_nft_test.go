package handlers

import (
	"testing"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"encoding/json"
)

func Test_parseNftDataForRecordId(t *testing.T) {
	data := "{\"data\":{\"header\":{\"model\":{\"protocol\":\"跨链监管数据模型\",\"version\":\"1.0\"},\"content\":{\"type\":\"主体\",\"action\":\"Update\",\"timestamp\":1596697200},\"source\":{\"locality\":\"jiangsu\",\"block_height\":837770000,\"tx_hash\":\"TCUAXHXKQFDAFPLSJFBC\",\"creator\":\"COFFSPQKXSLEFZAPAJZL\"},\"body\":{\"对象标识\":\"745B5F8995D66D65462E69A2E0C6FE033C86E1ADF7484B1B1F3638A817D04796\",\"主体标识\":\"A8825A938728D0DF85ECBC4FB9286AE3A683E2714B7626367F8DAF11D6C00AC3\"}}}}"
	recordid := parseNftDataForRecordId(data)
	t.Log(recordid)

	var record models.Record
	//packet
	record.Source.BlockHeight = int64(837770000)
	record.Source.TxHash = "TCUAXHXKQFDAFPLSJFBC"
	record.Source.Creator = "COFFSPQKXSLEFZAPAJZL"
	record.Source.Locality = "jiangsu"

	data1,_ := json.Marshal(record.Source)

	record.Id = utils.Md5(string(data1))
	t.Log(record.Id)
}

