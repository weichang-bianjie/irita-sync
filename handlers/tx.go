package handlers

import (
	"github.com/bianjieai/irita-sync/libs/cdc"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/libs/pool"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/utils/constant"
	aTypes "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	"github.com/tendermint/tendermint/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)

func ParseBlockAndTxs(b int64, client *pool.Client) (*models.Block, []*models.Tx, error) {
	var (
		blockDoc models.Block
		block    *ctypes.ResultBlock
	)

	if v, err := client.Block(&b); err != nil {
		logger.Warn("parse block fail, now try again", logger.Int64("height", b),
			logger.String("err", err.Error()))
		if v2, err := client.Block(&b); err != nil {
			logger.Error("parse block fail", logger.Int64("height", b),
				logger.String("err", err.Error()))
			return &blockDoc, nil, err
		} else {
			block = v2
		}
	} else {
		block = v
	}
	blockDoc = models.Block{
		Height:   block.Block.Height,
		Time:     block.Block.Time.Unix(),
		Hash:     block.Block.Header.Hash().String(),
		Txn:      int64(len(block.Block.Data.Txs)),
		Proposer: block.Block.ProposerAddress.String(),
	}

	txDocs := make([]*models.Tx, 0, len(block.Block.Txs))
	if len(block.Block.Txs) > 0 {
		for _, v := range block.Block.Txs {
			txDoc := parseTx(client, v, block.Block)
			if txDoc.TxHash != "" && len(txDoc.Type) > 0 {
				txDocs = append(txDocs, &txDoc)
			}
		}
	}

	return &blockDoc, txDocs, nil
}

func parseTx(c *pool.Client, txBytes types.Tx, block *types.Block) models.Tx {
	var (
		docTx     models.Tx
		docTxMsgs []models.DocTxMsg
	)

	Tx, err := cdc.GetTxDecoder()(txBytes)
	if err != nil {
		logger.Error(err.Error())
		return docTx
	}
	height := block.Height
	txHash := utils.BuildHex(txBytes.Hash())

	authTx := Tx.(signing.Tx)
	fee := models.BuildFee(authTx.GetFee(), authTx.GetGas())
	memo := authTx.GetMemo()

	txResult, err := c.Tx(txBytes.Hash(), false);
	if err != nil {
		logger.Error("get tx result fail", logger.String("txHash", txBytes.String()),
			logger.String("err", err.Error()))
		return docTx
	}
	status := parseTxStatus(txResult.TxResult.Code)
	log := txResult.TxResult.Log

	msgs := authTx.GetMsgs()
	if len(msgs) == 0 {
		return docTx
	}
	docTx = models.Tx{
		Height: height,
		Time:   block.Time.Unix(),
		TxHash: txHash,
		Fee:    &fee,
		Memo:   memo,
		Status: status,
		Log:    log,
		Events: parseEvents(txResult.TxResult.Events),
	}
	for i, v := range msgs {
		msgDocInfo := HandleTxMsg(v)
		if len(msgDocInfo.Addrs) == 0 {
			continue
		}
		if i == 0 {
			docTx.Type = msgDocInfo.DocTxMsg.Type
		}
		for _, signer := range v.GetSigners() {
			docTx.Signers = append(docTx.Signers, signer.String())
		}

		docTx.Addrs = append(docTx.Addrs, removeDuplicatesFromSlice(msgDocInfo.Addrs)...)
		docTxMsgs = append(docTxMsgs, msgDocInfo.DocTxMsg)
		docTx.Types = append(docTx.Types, msgDocInfo.DocTxMsg.Type)
	}
	docTx.Signers = removeDuplicatesFromSlice(docTx.Signers)
	docTx.Types = removeDuplicatesFromSlice(docTx.Types)
	docTx.Addrs = removeDuplicatesFromSlice(docTx.Addrs)

	docTx.DocTxMsgs = docTxMsgs

	// don't save txs which have not parsed
	if docTx.Type == "" || docTx.TxHash == "" {
		return models.Tx{}
	}

	return docTx
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
