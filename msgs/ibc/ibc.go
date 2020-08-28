package ibc

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(msg sdk.Msg) (MsgDocInfo, bool) {
	ok := true
	switch msg.(type) {
	case MsgIBCPacket:
		docMsg := DocTxMsgIBCPacket{}
		return docMsg.HandleTxMsg(msg.(MsgIBCPacket)), ok
	case MsgIBCTimeout:
		docMsg := DocTxMsgIBCTimeout{}
		return docMsg.HandleTxMsg(msg.(MsgIBCTimeout)), ok
	case MsgIBCTransfer:
		docMsg := DocTxMsgIBCTransfer{}
		return docMsg.HandleTxMsg(msg.(MsgIBCTransfer)), ok
	case MsgCreateClient:
		docMsg := DocMsgCreateClient{}
		return docMsg.HandleTxMsg(msg.(MsgCreateClient)), ok
	case MsgUpdateClient:
		docMsg := DocMsgUpdateClient{}
		return docMsg.HandleTxMsg(msg.(MsgUpdateClient)), ok
	case MsgSubmitClientMisbehaviour:
		docMsg := DocMsgSubmitClientMisbehaviour{}
		return docMsg.HandleTxMsg(msg.(MsgSubmitClientMisbehaviour)), ok
	default:
		ok = false
	}
	return MsgDocInfo{}, ok
}
