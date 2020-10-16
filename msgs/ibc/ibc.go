package ibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/bianjieai/irita-sync/models"
	"github.com/bianjieai/irita-sync/utils"
	"github.com/bianjieai/irita-sync/msgs/ibc/record"
	"encoding/json"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, models.Record, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecvPacket).Type():
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		var recordData models.Record
		//packet
		ibcRecord := docMsg.Packet.Data
		recordData.Creator = docMsg.Signer
		recordData.Contents = loadContents(ibcRecord.Contents)

		recordData.Source.BlockHeight = int64(ibcRecord.Height)
		recordData.Source.TxHash = ibcRecord.TxHash
		//record.Source.Creator = ibcRecord.Value.Creator //nft data source not exist this segment
		recordData.Source.Locality = docMsg.ClientID

		md5Feed, _ := json.Marshal(recordData.Source)
		recordData.Id = utils.Md5(string(md5Feed))
		return msgDocInfo, recordData, ok
	case new(MsgCreateClient).Type():
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpdateClient).Type():
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	default:
		ok = false
	}
	return msgDocInfo, models.Record{}, ok
}

func loadContents(contents []*record.Content) []models.Content {
	sliceContent := make([]models.Content, 0, len(contents))
	for _, val := range contents {
		sliceContent = append(sliceContent, models.Content{
			Digest:     val.Digest,
			DigestAlgo: val.DigestAlgo,
			Meta:       val.Meta,
			Uri:        val.URI,
		})
	}
	return sliceContent
}