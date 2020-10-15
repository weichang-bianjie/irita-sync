package nft

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, models.Nft, models.NftDenom, bool) {
	var (
		msgDocInfo MsgDocInfo
		nftInfo    models.Nft
		nftDenom   models.NftDenom
	)
	ok := true
	switch v.Type() {
	case new(MsgNFTMint).Type():
		docMsg := DocMsgNFTMint{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		nftInfo.Id = docMsg.Id
		nftInfo.DenomId = docMsg.Denom
		nftInfo.Owner = docMsg.Sender
		nftInfo.Uri = docMsg.URI
		nftInfo.Data = docMsg.Data
		nftInfo.NftName = docMsg.Name
		break
	case new(MsgNFTEdit).Type():
		docMsg := DocMsgNFTEdit{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		nftInfo.Id = docMsg.Id
		nftInfo.DenomId = docMsg.Denom
		nftInfo.Owner = docMsg.Sender
		nftInfo.Uri = docMsg.URI
		nftInfo.Data = docMsg.Data
		nftInfo.NftName = docMsg.Name
		break
	case new(MsgNFTTransfer).Type():
		docMsg := DocMsgNFTTransfer{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		nftInfo.Id = docMsg.Id
		nftInfo.DenomId = docMsg.Denom
		nftInfo.Owner = docMsg.Recipient
		nftInfo.Uri = docMsg.URI
		nftInfo.Data = docMsg.Data
		nftInfo.NftName = docMsg.Name
		break
	case new(MsgNFTBurn).Type():
		docMsg := DocMsgNFTBurn{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		nftInfo.Id = docMsg.Id
		nftInfo.DenomId = docMsg.Denom
		nftInfo.Owner = docMsg.Sender
		break
	case new(MsgIssueDenom).Type():
		docMsg := DocMsgIssueDenom{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		nftDenom.Creator = docMsg.Sender
		nftDenom.Schema = docMsg.Schema
		nftDenom.Id = docMsg.Id
		nftDenom.DenomName = docMsg.Name
		break
	default:
		ok = false
	}
	return msgDocInfo, nftInfo, nftDenom, ok
}
