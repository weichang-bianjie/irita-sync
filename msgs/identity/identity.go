package identity

import (
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	idtypes "gitlab.bianjie.ai/irita-pro/iritamod/modules/identity/types"
)

func HandleTxMsg(v types.Msg) (MsgDocInfo, models.Identity, bool) {
	var (
		msgDocInfo  MsgDocInfo
		identityDoc models.Identity
	)
	ok := true
	switch v.Type() {
	case new(MsgCreateIdentity).Type():
		docMsg := DocMsgCreateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		identityDoc.Owner = docMsg.Owner
		identityDoc.Credentials = docMsg.Credentials
		identityDoc.Id = docMsg.Id
		if docMsg.Certificate != "" {
			identityDoc.Certificates = append(identityDoc.Certificates, docMsg.Certificate)
		}
		if docMsg.PubKey != nil {
			identityDoc.Pubkeys = append(identityDoc.Pubkeys, models.PubKeyInfo{
				Algorithm: idtypes.PubKeyAlgorithm_name[docMsg.PubKey.Algorithm],
				PubKey:    docMsg.PubKey.PubKey,
			})
		}
		break
	case new(MsgUpdateIdentity).Type():
		docMsg := DocMsgUpdateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		identityDoc.Owner = docMsg.Owner
		identityDoc.Credentials = docMsg.Credentials
		identityDoc.Id = docMsg.Id
		if docMsg.Certificate != "" {
			identityDoc.Certificates = append(identityDoc.Certificates, docMsg.Certificate)
		}
		if docMsg.PubKey != nil {
			identityDoc.Pubkeys = append(identityDoc.Pubkeys, models.PubKeyInfo{
				Algorithm: idtypes.PubKeyAlgorithm_name[docMsg.PubKey.Algorithm],
				PubKey:    docMsg.PubKey.PubKey,
			})
		}
		break
	default:
		ok = false
	}
	return msgDocInfo, identityDoc, ok
}
