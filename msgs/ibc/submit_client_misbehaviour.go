package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
)
// MsgSubmitClientMisbehaviour defines an sdk.Msg type that supports submitting
// Evidence for client misbehaviour.
type DocMsgSubmitClientMisbehaviour struct {
	Evidence  string `bson:"evidence" yaml:"evidence"`
	Submitter string `bson:"submitter" yaml:"submitter"`
}

func (m *DocMsgSubmitClientMisbehaviour) GetType() string {
	return MsgTypeIBCSubmitClientMisbehaviour
}

func (m *DocMsgSubmitClientMisbehaviour) BuildMsg(v interface{}) {
	msg := v.(MsgSubmitClientMisbehaviour)

	m.Evidence = msg.Evidence.String()
	m.Submitter = msg.Submitter.String()
}

func (m *DocMsgSubmitClientMisbehaviour) HandleTxMsg(msg MsgSubmitClientMisbehaviour) MsgDocInfo {
	var (
		addrs []string
	)
	addrs = append(addrs, m.Submitter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}