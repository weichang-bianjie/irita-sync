package ibc

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

type DocTxMsgIBCTransfer struct {
	SourcePort        string        `bson:"source_port"`        // the port on which the packet will be sent
	SourceChannel     string        `bson:"source_channel"`     // the channel by which the packet will be sent
	DestinationHeight uint64        `bson:"destination_height"` // the current height of the destination chain
	Amount            []models.Coin `bson:"amount"`             // the tokens to be transferred
	Sender            string        `bson:"sender"`             // the sender address
	Receiver          string        `bson:"receiver"`           // the recipient address on the destination chain
}


func (m *DocTxMsgIBCTransfer) GetType() string {
	return MsgTypeIBCTransfer
}

func (m *DocTxMsgIBCTransfer) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgIBCTransfer)

	m.SourcePort = msg.SourcePort
	m.SourceChannel = msg.SourceChannel
	m.DestinationHeight = msg.DestinationHeight
	m.Amount = models.BuildDocCoins(msg.Amount)
	m.Sender = msg.Sender.String()
	m.Receiver = msg.Receiver
}

func (m *DocTxMsgIBCTransfer) HandleTxMsg(msg MsgIBCTransfer) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Sender, m.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}