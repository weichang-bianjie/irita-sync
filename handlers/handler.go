package handlers

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/bianjieai/irita-sync/msgs"
	"github.com/bianjieai/irita-sync/msgs/nft"
	"github.com/bianjieai/irita-sync/msgs/identity"
	"github.com/bianjieai/irita-sync/msgs/ibc"
	"github.com/bianjieai/irita-sync/libs/logger"
	"github.com/bianjieai/irita-sync/models"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2"
)

func HandleTxMsg(v types.Msg, tx *models.Tx, txIndex, msgIndex int) (MsgDocInfo, []txn.Op) {
	//if BankDocInfo, ok := bank.HandleTxMsg(v); ok {
	//	return BankDocInfo, nil
	//}
	//if IServiceDocInfo, ok := service.HandleTxMsg(v); ok {
	//	return IServiceDocInfo, nil
	//}
	if NftDocInfo, nftDoc, nftDenomDoc, ok := nft.HandleTxMsg(v); ok {
		mapPkObjIdData, err := nftDoc.AllNftsMaps()
		if err != nil && err != mgo.ErrNotFound {
			logger.Warn("get All NftsIds Maps have error", logger.String("err", err.Error()))
		}
		opsNft := handleNftData(tx, NftDocInfo.DocTxMsg.Type, &nftDenomDoc, &nftDoc, mapPkObjIdData, txIndex, msgIndex)
		return NftDocInfo, opsNft
	}
	//if RecordDocInfo, ok := record.HandleTxMsg(v); ok {
	//	return RecordDocInfo, nil
	//}
	//if TokenDocInfo, ok := token.HandleTxMsg(v); ok {
	//	return TokenDocInfo, nil
	//}
	//if CoinswapDocInfo, ok := coinswap.HandleTxMsg(v); ok {
	//	return CoinswapDocInfo, nil
	//}
	//if CrisisDocInfo, ok := crisis.HandleTxMsg(v); ok {
	//	return CrisisDocInfo, nil
	//}
	//if DistrubutionDocInfo, ok := distribution.HandleTxMsg(v); ok {
	//	return DistrubutionDocInfo, nil
	//}
	//if SlashingDocInfo, ok := slashing.HandleTxMsg(v); ok {
	//	return SlashingDocInfo, nil
	//}
	//if EvidenceDocInfo, ok := evidence.HandleTxMsg(v); ok {
	//	return EvidenceDocInfo, nil
	//}
	//if HtlcDocInfo, ok := htlc.HandleTxMsg(v); ok {
	//	return HtlcDocInfo, nil
	//}
	//if StakingDocInfo, ok := staking.HandleTxMsg(v); ok {
	//	return StakingDocInfo, nil
	//}
	//if GovDocInfo, ok := gov.HandleTxMsg(v); ok {
	//	return GovDocInfo, nil
	//}
	if IdentityDocInfo, identityDoc, ok := identity.HandleTxMsg(v); ok {
		mapObjData, err := identityDoc.AllIdentitysMaps()
		if err != nil && err != mgo.ErrNotFound {
			logger.Warn("get All IdentityIds Maps have error", logger.String("err", err.Error()))
		}
		opsIdentity := handleIdentityData(IdentityDocInfo.DocTxMsg.Type, &identityDoc, mapObjData)
		return IdentityDocInfo, opsIdentity
	}
	//if RandomDocInfo, ok := random.HandleTxMsg(v); ok {
	//	return RandomDocInfo, nil
	//}
	//if OracleDocInfo, ok := oracle.HandleTxMsg(v); ok {
	//	return OracleDocInfo, nil
	//}
	if IbcDocinfo, recordDoc, ok := ibc.HandleTxMsg(v); ok {
		opsRecord := handleIbcRecordData(tx, &recordDoc)
		return IbcDocinfo, opsRecord
	}
	return MsgDocInfo{}, nil
}

func removeDuplicatesFromSlice(data []string) (result []string) {
	tempSet := make(map[string]string, len(data))
	for _, val := range data {
		if _, ok := tempSet[val]; ok || val == "" {
			continue
		}
		tempSet[val] = val
	}
	for one := range tempSet {
		result = append(result, one)
	}
	return
}

//
//func handlerIbcClient(msgType string, client *models.IbcClient) (Ops []txn.Op) {
//	switch msgType {
//	case MsgTypeCreateClient:
//		client.ID = bson.NewObjectId()
//		op := txn.Op{
//			C:      models.CollectionNameIbcClient,
//			Id:     bson.NewObjectId(),
//			Insert: client,
//		}
//		Ops = append(Ops, op)
//	case MsgTypeUpdateClient:
//		v := client
//		mapObjId, err := client.AllIbcClientMaps()
//		if err != nil {
//			return
//		}
//		if id, ok := mapObjId[v.ClientId]; ok {
//			v.ID = id
//		}
//		if !v.ID.Valid() {
//			return
//		}
//		updateOp := txn.Op{
//			C:      models.CollectionNameIbcClient,
//			Id:     v.ID,
//			Assert: txn.DocExists,
//			Update: bson.M{
//				"$set": bson.M{
//					models.IbcClientHeaderTag:   v.Header,
//					models.IbcClientSignerTag:   v.Signer,
//					models.IbcClientUpdateAtTag: v.UpdateAt,
//				},
//			},
//		}
//		Ops = append(Ops, updateOp)
//	}
//	return
//}
