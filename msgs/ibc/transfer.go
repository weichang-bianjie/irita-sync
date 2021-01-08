package ibc

import (
	"github.com/bianjieai/irita-sync/libs/cdc"
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DocMsgTransfer struct {
	// the port on which the packet will be sent
	SourcePort string `bson:"source_port" json:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `bson:"source_channel" json:"source_channel"`
	// the tokens to be transferred
	Token models.Coin `bson:"token" json:"token"`
	// the sender address
	Sender string `bson:"sender" json:"sender"`
	// the recipient address on the destination chain
	Receiver string `bson:"receiver" json:"receiver"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight Height `bson:"timeout_height" json:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `bson:"timeout_timestamp" json:"timeout_timestamp"`
}

type Height struct {
	// the revision that the client is currently on
	RevisionNumber uint64 `bson:"revision_number" json:"revision_number"`
	// the height within the given revision
	RevisionHeight uint64 `bson:"revision_height" json:"revision_height"`
}

func (m *DocMsgTransfer) GetType() string {
	return MsgTypeTransfer
}

func (m *DocMsgTransfer) BuildMsg(v interface{}) {
	handle_fn := func(v sdk.Msg) {
		var msg MsgTransfer
		data, _ := cdc.GetMarshaler().MarshalJSON(v)
		cdc.GetMarshaler().UnmarshalJSON(data, &msg)
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msg), m)
	}
	handle_fn(v.(sdk.Msg))
}

func (m *DocMsgTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTransfer
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
