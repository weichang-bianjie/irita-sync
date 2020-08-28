package ibc

import (
	"encoding/json"
	. "github.com/bianjieai/irita-sync/msgs"
)

// MsgTimeout receives timed-out packet
type DocTxMsgIBCTimeout struct {
	Packet                  `bson:"packet"`
	NextSequenceRecv uint64 `bson:"next_sequence_recv"`
	Proof            string `bson:"proof"`
	ProofHeight      uint64 `bson:"proof_height"`
	Signer           string `bson:"signer"`
}

func (m *DocTxMsgIBCTimeout) GetType() string {
	return MsgTypeIBCMsgTimeout
}

func (m *DocTxMsgIBCTimeout) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgIBCTimeout)

	var sendpacket SendPacket
	json.Unmarshal(msg.Packet.GetData(), &sendpacket)

	packet := Packet{
		Sequence:           msg.Packet.GetSequence(),
		TimeoutHeight:      msg.Packet.GetTimeoutHeight(),
		SourcePort:         msg.Packet.GetSourcePort(),
		SourceChannel:      msg.Packet.GetSourceChannel(),
		DestinationPort:    msg.Packet.GetDestPort(),
		DestinationChannel: msg.Packet.GetDestChannel(),
		Data:               sendpacket,
	}
	m.Packet = packet
	proofdata, _ := json.Marshal(msg.Proof)
	m.Proof = string(proofdata)
	m.NextSequenceRecv = msg.NextSequenceRecv
	m.ProofHeight = msg.ProofHeight
	m.Signer = msg.Signer.String()
}

func (m *DocTxMsgIBCTimeout) HandleTxMsg(msg MsgIBCTimeout) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
