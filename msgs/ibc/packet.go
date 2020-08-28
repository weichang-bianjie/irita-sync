package ibc

import (
	"github.com/bianjieai/irita-sync/models"
	"encoding/json"
	. "github.com/bianjieai/irita-sync/msgs"
)

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence           uint64     `bson:"sequence" `            // number corresponds to the order of sends and receives, where a Packet with an earlier sequence number must be sent and received before a Packet with a later sequence number.
	SourcePort         string     `bson:"source_port" `         // identifies the port on the sending chain.
	SourceChannel      string     `bson:"source_channel" `      // identifies the channel end on the sending chain.
	DestinationPort    string     `bson:"destination_port" `    // identifies the port on the receiving chain.
	DestinationChannel string     `bson:"destination_channel" ` // identifies the channel end on the receiving chain.
	TimeoutHeight      uint64     `bson:"timeout_height" `      // block height after which the packet times out
	Data               SendPacket `bson:"data"`                 // opaque value which can be defined by the application logic of the associated modules.
}

type SendPacket struct {
	Amount   []models.Coin `bson:"amount" json:"amount"`
	Receiver string        `bson:"receiver" json:"receiver"`
	Sender   string        `bson:"sender" json:"sender"`
	Source   bool          `bson:"source" json:"source"`
	Timeout  string        `bson:"timeout" json:"timeout"`
}

// MsgPacket receives incoming IBC packet
type DocTxMsgIBCPacket struct {
	Packet      Packet `bson:"packet"`
	Proof       string `bson:"proof"`
	ProofHeight uint64 `bson:"proof_height" `
	Signer      string `bson:"signer"`
}

func (m *DocTxMsgIBCPacket) GetType() string {
	return MsgTypeIBCMsgPacket
}

func (m *DocTxMsgIBCPacket) BuildMsg(txMsg interface{}) {
	msg := txMsg.(MsgIBCPacket)
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
	m.Signer = msg.Signer.String()
	m.ProofHeight = msg.ProofHeight
	proofdata, _ := json.Marshal(msg.Proof)
	m.Proof = string(proofdata)
}

func (m *DocTxMsgIBCPacket) HandleTxMsg(msg MsgIBCPacket) MsgDocInfo {
	var (
		addrs []string
	)

	addrs = append(addrs, m.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(msg, handler)
}
