package handlers

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"encoding/json"
	"github.com/bianjieai/irita-sync/utils"
)

func handleNftData(docTx *models.Tx, msgType string, denomInfo *models.NftDenom, nftInfo *models.Nft, mapPkObjIdData map[string]bson.ObjectId, txIndex, msgIndex int) (Ops []txn.Op) {

	switch msgType {
	case MsgTypeIssueDenom:
		denomInfo.TxHash = docTx.Hash
		denomInfo.BlockHeight = docTx.BlockHeight
		op := txn.Op{
			C:      models.CollectionNameNftDenom,
			Id:     bson.NewObjectId(),
			Insert: denomInfo,
		}
		Ops = append(Ops, op)
	case MsgTypeNFTMint, MsgTypeNFTEdit, MsgTypeNFTBurn, MsgTypeNFTTransfer:
		v := nftInfo
		if msgType != MsgTypeNFTMint {
			pk := v.Id + v.DenomId
			if id, ok := mapPkObjIdData[pk]; ok {
				v.ID = id
			}

			switch msgType {
			case MsgTypeNFTEdit:
				if !v.ID.Valid() {
					return
				}
				updateOp := txn.Op{
					C:      models.CollectionNameNft,
					Id:     v.ID,
					Assert: txn.DocExists,
					Update: bson.M{
						"$set": bson.M{
							models.NftUriTag:  v.Uri,
							models.NftDataTag: v.Data,
							models.NftNameTag: v.NftName,
						},
					},
				}
				Ops = append(Ops, updateOp)
			case MsgTypeNFTBurn:
				if !v.ID.Valid() {
					return
				}
				deleteOp := txn.Op{
					C:      models.CollectionNameNft,
					Id:     v.ID,
					Assert: txn.DocExists,
					Remove: true,
				}
				Ops = append(Ops, deleteOp)

			case MsgTypeNFTTransfer:
				if !v.ID.Valid() {
					return
				}
				updateOp := txn.Op{
					C:      models.CollectionNameNft,
					Id:     v.ID,
					Assert: txn.DocExists,
					Update: bson.M{
						"$set": bson.M{
							models.NftUriTag:   v.Uri,
							models.NftDataTag:  v.Data,
							models.NftNameTag:  v.NftName,
							models.NftOwnerTag: v.Owner,
						},
					},
				}
				Ops = append(Ops, updateOp)
			}

		} else { //MsgTypeNFTMint
			v.ID = bson.NewObjectId()
			op := txn.Op{
				C:      models.CollectionNameNft,
				Id:     bson.NewObjectId(),
				Insert: v,
			}
			Ops = append(Ops, op)
		}

		docNftHistory := initNftHistory(nftInfo, docTx.Hash, docTx.BlockHeight, txIndex, msgIndex)
		docNftHistory.TxType = msgType
		op := txn.Op{
			C:      models.CollectionNameNftHistory,
			Id:     bson.NewObjectId(),
			Insert: docNftHistory,
		}
		Ops = append(Ops, op)
	}
	return
}

func parseNftDataForRecordId(data string) string {
	var tmpModel struct {
		Header struct {
			Source models.Source `json:"source"`
		} `json:"header"`
		Body struct{} `json:"body"`
	}

	if err := json.Unmarshal([]byte(data), &tmpModel); err == nil {
		recordId, _ := json.Marshal(tmpModel.Header.Source)
		return utils.Md5(string(recordId))
	}
	return ""
}
func initNftHistory(nftInfo *models.Nft, txHash string, height int64, txIndex, msgIndex int) models.NftHistory {
	recordId := parseNftDataForRecordId(nftInfo.Data)
	return models.NftHistory{
		Id:          nftInfo.Id,
		NftName:     nftInfo.NftName,
		Owner:       nftInfo.Owner,
		DenomId:     nftInfo.DenomId,
		Uri:         nftInfo.Uri,
		Data:        nftInfo.Data,
		TxHash:      txHash,
		BlockHeight: height,
		RecordId:    recordId,
		TxIndex:     txIndex,
		MsgIndex:    msgIndex,
	}

}
