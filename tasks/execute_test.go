package tasks

import (
	"testing"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/handlers"
	"github.com/bianjieai/irita-sync/models"
)

func TestParse(t *testing.T) {
	block := int64(11921)
	c := pool.GetClient()
	defer func() {
		c.Release()
	}()

	if blockDoc, txDocs, ops, err := handlers.ParseBlockAndTxs(block, c); err != nil {
		t.Fatal(err)
	} else {
		err := saveDocsWithTxn(blockDoc,txDocs,models.SyncTask{},ops)
		if err != nil {
			t.Fatal(err)
		}

		//b, _ := hex.DecodeString("736572766963652063616c6c20726573706f6e7365")
		//t.Log(string(b))
	}
}
