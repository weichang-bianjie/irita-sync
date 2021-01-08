package ibc

import (
	"github.com/bianjieai/irita-sync/libs/cdc"
	. "github.com/bianjieai/irita-sync/msgs"
)

// MsgUpdateClient defines a message to update an IBC client
type DocMsgUpdateClient struct {
	ClientId string      `bson:"client_id" yaml:"client_id"`
	Header   interface{} `bson:"header" yaml:"header"`
	Signer   string      `bson:"signer" yaml:"signer"`
}

func (m *DocMsgUpdateClient) GetType() string {
	return MsgTypeUpdateClient
}

func (m *DocMsgUpdateClient) BuildMsg(v interface{}) {
	msg := v.(MsgUpdateClient)

	m.ClientId = msg.ClientId
	m.Signer = msg.Signer
	m.Header = msg.Header
}

func (m *DocMsgUpdateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgUpdateClient
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
