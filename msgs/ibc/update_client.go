package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
)

// MsgUpdateClient defines a message to update an IBC client
type DocMsgUpdateClient struct {
	ClientID string `bson:"client_id" yaml:"client_id"`
	Header   Header `bson:"header" yaml:"header"`
	Signer   string `bson:"signer" yaml:"signer"`
}

func (m *DocMsgUpdateClient) GetType() string {
	return MsgTypeIBCUpdateClient
}

func (m *DocMsgUpdateClient) BuildMsg(v interface{}) {
	msg := v.(MsgUpdateClient)

	m.ClientID = msg.ClientID
	m.Signer = msg.Signer.String()
	m.Header = Header{}.Build(msg.Header)
}

func (m *DocMsgUpdateClient) HandleTxMsg(msg MsgUpdateClient) MsgDocInfo {
	var (
		addrs []string
	)
	addrs = append(addrs, m.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
