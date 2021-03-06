package handlers

import (
	"encoding/hex"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/utils"
	"testing"
)

func TestParseTxs(t *testing.T) {
	block := int64(18687)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, err := ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		t.Log(utils.MarshalJsonIgnoreErr(blockDoc))
		t.Log(utils.MarshalJsonIgnoreErr(txDocs))

		b, _ := hex.DecodeString("736572766963652063616c6c20726573706f6e7365")
		t.Log(string(b))
	}
}
