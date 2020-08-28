package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"time"
)

// MsgCreateClient defines a message to create an IBC client
type DocMsgCreateClient struct {
	ClientID        string        `bson:"client_id" yaml:"client_id"`
	Header          Header        `bson:"header" yaml:"header"`
	TrustingPeriod  time.Duration `bson:"trusting_period" yaml:"trusting_period"`
	UnbondingPeriod time.Duration `bson:"unbonding_period" yaml:"unbonding_period"`
	MaxClockDrift   time.Duration `bson:"max_clock_drift" yaml:"max_clock_drift"`
	Signer          string        `bson:"signer" yaml:"signer"`
}

func (m *DocMsgCreateClient) GetType() string {
	return MsgTypeIBCCreateClient
}

func (m *DocMsgCreateClient) BuildMsg(v interface{}) {
	msg := v.(MsgCreateClient)

	m.ClientID = msg.ClientID
	m.Signer = msg.Signer.String()
	m.MaxClockDrift = msg.MaxClockDrift
	m.TrustingPeriod = msg.TrustingPeriod
	m.UnbondingPeriod = msg.UnbondingPeriod
	m.Header = Header{}.Build(msg.Header)
}

func (m *DocMsgCreateClient) HandleTxMsg(msg MsgCreateClient) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
