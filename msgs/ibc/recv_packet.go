package ibc

import (
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/utils"
)

type DocMsgRecvPacket struct {
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	// number corresponds to the order of sends and receives, where a Packet
	// with an earlier sequence number must be sent and received before a Packet
	// with a later sequence number.
	Sequence uint64 `bson:"sequence" json:"sequence"`
	// identifies the port on the sending chain.
	SourcePort string `bson:"source_port" json:"source_port"`
	// identifies the channel end on the sending chain.
	SourceChannel string `bson:"source_channel" json:"source_channel"`
	// identifies the port on the receiving chain.
	DestinationPort string `bson:"destination_port" json:"destination_port"`
	// identifies the channel end on the receiving chain.
	DestinationChannel string `bson:"destination_channel" json:"destination_channel"`
	// actual opaque bytes transferred directly to the application module
	Data []byte `bson:"data" json:"data"`
	// block height after which the packet times out
	TimeoutHeight Height `bson:"timeout_height" json:"timeout_height"`
	// block timestamp (in nanoseconds) after which the packet times out
	TimeoutTimestamp uint64 `bson:"timeout_timestamp" json:"timeout_timestamp"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)
	m.ProofHeight = Height{
		RevisionNumber: msg.ProofHeight.RevisionNumber,
		RevisionHeight: msg.ProofHeight.RevisionHeight,
	}
	m.Signer = msg.Signer
	m.ProofCommitment = string(msg.ProofCommitment)
	m.Packet = DecodeToIBCRecord(msg.Packet)
}
func DecodeToIBCRecord(packet interface{}) (ret Packet) {
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(packet), &ret)
	return
}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
