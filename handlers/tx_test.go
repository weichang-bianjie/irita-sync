package handlers

import (
	"github.com/bianjieai/irita-sync/config"
	"github.com/bianjieai/irita-sync/utils"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(26941)
	conf, err := config.ReadConfig()
	if err != nil {
		t.Fatal(err.Error())
	}
	InitRouter(conf)

	if blockDoc, txDocs, _, err := ParseBlockAndTxs(block, conf); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))
	}
}
