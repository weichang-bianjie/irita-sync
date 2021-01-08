package ibc

import (
	"github.com/bianjieai/irita-sync/libs/cdc"
	. "github.com/bianjieai/irita-sync/msgs"
)

// MsgCreateClient defines a message to create an IBC client
type DocMsgCreateClient struct {
	ClientState    interface{} `bson:"client_state"`
	ConsensusState interface{} `bson:"consensus_state"`
	Signer         string      `bson:"signer" yaml:"signer"`
}

func (m *DocMsgCreateClient) GetType() string {
	return MsgTypeCreateClient
}

func (m *DocMsgCreateClient) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateClient)

	m.Signer = msg.Signer
	m.ClientState = msg.ClientState
	m.ConsensusState = msg.ConsensusState

}

func (m *DocMsgCreateClient) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgCreateClient
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
