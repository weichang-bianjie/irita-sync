package handlers

import abci "github.com/tendermint/tendermint/abci/types"

type (
	LastestBlock struct {
		Height   int64  `json:"height"`
		Hash     string `json:"hash"`
		Time     int64  `json:"time"`
		Proposer string `json:"proposer"`
	}
	DeliverTx struct {
		Tx       string                 `json:"tx"`
		TxResult abci.ResponseDeliverTx `json:"tx_result"`
	}
	Txs struct {
		Txs []DeliverTx `json:"txs"`
	}
)
