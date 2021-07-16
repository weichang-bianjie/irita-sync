package tasks

import (
	"fmt"
	"github.com/bianjieai/irita-sync/config"
	"github.com/bianjieai/irita-sync/handlers"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/utils/constant"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"os"
	"time"
)

func Start(conf *config.Config) {

	heightChan := make(chan int64)
	go getHeightTask(heightChan, conf)
	models.SyncTaskModel.Create()

	for {
		select {
		case height, ok := <-heightChan:
			if ok {
				parseBlockAndSave(height, conf)
			}

		}

	}

}

func getHeightTask(chanHeight chan int64, conf *config.Config) {
	inProcessBlock := int64(1)
	if maxHeight, err := new(models.Block).GetMaxBlockHeight(); err != nil {
		if err != mgo.ErrNotFound {
			logger.Fatal("get max height in block table have error",
				logger.String("err", err.Error()))
		}
	} else {
		inProcessBlock = maxHeight.Height
	}
	if conf.Server.ChainBlockResetHeight > 0 {
		inProcessBlock = conf.Server.ChainBlockResetHeight
	}
	for {
		//check file if exist
		filepath := fmt.Sprintf("%v/%v_%v_block", conf.Server.WriteDir, conf.Server.FilePrefix, inProcessBlock)
		exist, err := checkFileExist(filepath)
		if err != nil {
			logger.Warn("check file exist failed", logger.String("err", err.Error()))
			continue
		}
		if !exist {
			models.SyncTaskModel.Update(models.SyncTaskStatusUnderway)
			logger.Info("wait blockChain latest height update",
				logger.Int64("curSyncedHeight", inProcessBlock-1),
				logger.Int64("blockChainLatestHeight", inProcessBlock))
			time.Sleep(1 * time.Second)
			continue
		} else {
			models.SyncTaskModel.Update(models.SyncTaskStatusCatchUping)
			chanHeight <- inProcessBlock
			inProcessBlock++
		}
	}
}

// check file whether exist
// return true if exist, otherwise return false
func checkFileExist(filePath string) (bool, error) {
	exists := true
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			// unknown err
			return false, err
		}
	}
	return exists, nil
}

func parseBlockAndSave(height int64, conf *config.Config) {

	noFinish := true
	for noFinish {
		// parse data from block
		blockDoc, txDocs, ops, err := handlers.ParseBlockAndTxs(height, conf)
		if err != nil {
			logger.Error("Parse block fail",
				logger.Int64("height", height),
				logger.String("errTag", utils.GetErrTag(err)),
				logger.String("err", err.Error()))
			time.Sleep(time.Second)
			continue

		}
		if err := saveDocsWithTxn(blockDoc, txDocs, ops); err != nil {
			logger.Error("save docs fail", logger.String("err", err.Error()))
			time.Sleep(time.Second)
			continue
		}
		noFinish = false
	}

	logger.Info("sync blockChain have ok",
		logger.Int64("curSyncedHeight", height))

	return
}

func saveDocsWithTxn(blockDoc *models.Block, txDocs []*models.Tx, opsDoc []txn.Op) error {
	var (
		ops, txsOps []txn.Op
	)

	if blockDoc.Height == 0 {
		return fmt.Errorf("invalid block, height equal 0")
	}

	blockOp := txn.Op{
		C:      models.BlockModel.Name(),
		Id:     bson.NewObjectId(),
		Insert: blockDoc,
	}

	if length := len(txDocs); length > 0 {
		for _, v := range txDocs {
			op := txn.Op{
				C:      models.TxModel.Name(),
				Id:     bson.NewObjectId(),
				Insert: v,
			}
			txsOps = append(txsOps, op)
		}
	}

	ops = append(ops, txsOps...)
	ops = append(ops, blockOp)

	if len(opsDoc) > 0 {
		ops = append(ops, opsDoc...)
	}

	return batchSave(ops)
}

func batchSave(ops []txn.Op) error {
	if len(ops) <= constant.MaxRecordNumForBatchInsert {
		err := models.Txn(ops)
		if err != nil {
			return err
		}
		return nil
	}

	if batchSize := len(ops); batchSize > 0 {
		num := 1
		if batchSize%constant.MaxRecordNumForBatchInsert > 0 {
			num = batchSize/constant.MaxRecordNumForBatchInsert + 1
		} else if batchSize%constant.MaxRecordNumForBatchInsert == 0 {
			num = batchSize / constant.MaxRecordNumForBatchInsert
		}
		for i := 0; i < num; i++ {
			start := i * constant.MaxRecordNumForBatchInsert
			end := start + constant.MaxRecordNumForBatchInsert
			if i == num-1 {
				end = start + batchSize - i*constant.MaxRecordNumForBatchInsert
			}
			//fmt.Println(start, end)
			err := models.Txn(ops[start:end])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
