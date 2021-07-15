package handlers

import abci "github.com/tendermint/tendermint/abci/types"

type (
	DeliverTx struct {
		Tx       string                 `json:"tx"`
		TxResult abci.ResponseDeliverTx `json:"tx_result"`
	}
	Txs struct {
		Txs []DeliverTx `json:"txs"`
	}
)
