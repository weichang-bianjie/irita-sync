package handlers

import (
	"encoding/hex"
	"fmt"
	"github.com/bianjieai/irita-sync/config"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/libs/msgparser"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/utils/constant"
	"github.com/kaifei-bianjie/msg-parser/codec"
	msgsdktypes "github.com/kaifei-bianjie/msg-parser/types"
	aTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/types"
	"gopkg.in/mgo.v2/txn"
	"io/ioutil"
	"strings"
)

var _parser msgparser.MsgParser

func InitRouter(conf *config.Config) {
	var router msgparser.Router
	if conf.Server.OnlySupportModule != "" {
		modules := strings.Split(conf.Server.OnlySupportModule, ",")
		msgRoute := msgparser.NewRouter()
		for _, one := range modules {
			fn, exist := msgparser.RouteHandlerMap[one]
			if !exist {
				logger.Fatal("no support module: " + one)
			}
			msgRoute = msgRoute.AddRoute(one, fn)
			if one == msgparser.IbcRouteKey {
				msgRoute = msgRoute.AddRoute(msgparser.IbcTransferRouteKey, msgparser.RouteHandlerMap[one])
			}
		}
		if msgRoute.GetRoutesLen() > 0 {
			router = msgRoute
		}
	} else {
		router = msgparser.RegisteRouter()
	}
	_parser = msgparser.NewMsgParser(router)

	Init(conf)
}

func ParseBlockAndTxs(b int64, conf *config.Config) (*models.Block, []*models.Tx, []txn.Op, error) {
	var (
		blockDoc models.Block
		txnOps   []txn.Op
		txs      []*models.Tx
	)
	blockfile := fmt.Sprintf("%v/%v_%v_block", conf.Server.WriteDir, conf.Server.FilePrefix, b)
	bytesblock, err := ioutil.ReadFile(blockfile)
	if err != nil {
		return &blockDoc, nil, nil, err
	}
	utils.UnMarshalJsonIgnoreErr(string(bytesblock), &blockDoc)

	if blockDoc.Txn > 0 {
		txs = make([]*models.Tx, 0, blockDoc.Txn)
		txfile := fmt.Sprintf("%v/%v_%v_txs", conf.Server.WriteDir, conf.Server.FilePrefix, b)
		txbytes, err := ioutil.ReadFile(txfile)
		if err != nil {
			return &blockDoc, nil, nil, err
		}
		var txsData Txs
		utils.UnMarshalJsonIgnoreErr(string(txbytes), &txsData)
		for i, v := range txsData.Txs {
			tx, _ := hex.DecodeString(v.Tx)
			txDoc, ops, err := parseTx(tx, v.TxResult, blockDoc, i)
			if err != nil {
				return &blockDoc, txs, txnOps, err
			}
			if txDoc.TxHash != "" && len(txDoc.Type) > 0 {
				txs = append(txs, &txDoc)
				if len(ops) > 0 {
					txnOps = append(txnOps, ops...)
				}
			}
		}
	}

	return &blockDoc, txs, txnOps, nil
}

func parseTx(txBytes types.Tx, txResult aTypes.ResponseDeliverTx, block models.Block, txIndex int) (models.Tx, []txn.Op, error) {
	var (
		docTx     models.Tx
		docTxMsgs []msgsdktypes.TxMsg
		txnOps    []txn.Op
	)
	txHash := utils.BuildHex(txBytes.Hash())
	docTx.Time = block.Time
	docTx.Height = block.Height
	docTx.TxHash = txHash
	docTx.Status = parseTxStatus(txResult.Code)
	if docTx.Status == constant.TxStatusFail {
		docTx.Log = txResult.Log
	}

	docTx.Events = parseEvents(txResult.Events)
	docTx.EventsNew = parseABCILogs(txResult.Log)
	docTx.TxIndex = uint32(txIndex)

	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		logger.Warn(err.Error(),
			logger.String("errTag", "TxDecoder"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return docTx, txnOps, nil
	}
	docTx.Fee = msgsdktypes.BuildFee(authTx.GetFee(), authTx.GetGas())
	docTx.Memo = authTx.GetMemo()

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx, txnOps, nil
	}

	for i, v := range msgs {
		msgDocInfo, ops := _parser.HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}

		docTx.Signers = append(docTx.Signers, removeDuplicatesFromSlice(msgDocInfo.Signers)...)
		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
		if len(ops) > 0 {
			txnOps = append(txnOps, ops...)
		}
	}

	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.DocTxMsgs = docTxMsgs

	// don't save txs which have not parsed
	if docTx.Type == "" {
		logger.Warn(constant.NoSupportMsgTypeTag,
			logger.String("errTag", "TxMsg"),
			logger.String("txhash", txHash),
			logger.Int64("height", block.Height))
		return models.Tx{}, txnOps, nil
	}

	return docTx, txnOps, nil
}

func parseTxStatus(code uint32) uint32 {
	if code == 0 {
		return constant.TxStatusSuccess
	} else {
		return constant.TxStatusFail
	}
}

func parseEvents(events []aTypes.Event) []models.Event {
	var eventDocs []models.Event
	if len(events) > 0 {
		for _, e := range events {
			var kvPairDocs []models.KvPair
			if len(e.Attributes) > 0 {
				for _, v := range e.Attributes {
					kvPairDocs = append(kvPairDocs, models.KvPair{
						Key:   string(v.Key),
						Value: string(v.Value),
					})
				}
			}
			eventDocs = append(eventDocs, models.Event{
				Type:       e.Type,
				Attributes: kvPairDocs,
			})
		}
	}

	return eventDocs
}

// parseABCILogs attempts to parse a stringified ABCI tx log into a slice of
// EventNe types. It ignore error upon JSON decoding failure.
func parseABCILogs(logs string) []models.EventNew {
	var res []models.EventNew
	utils.UnMarshalJsonIgnoreErr(logs, &res)
	return res
}

func removeDuplicatesFromSlice(data []string) (result []string) {
	tempSet := make(map[string]string, len(data))
	for _, val := range data {
		if _, ok := tempSet[val]; ok || val == "" {
			continue
		}
		tempSet[val] = val
	}
	for one := range tempSet {
		result = append(result, one)
	}
	return
}
