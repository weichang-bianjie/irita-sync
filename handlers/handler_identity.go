package handlers

import (
	"github.com/bianjieai/irita-sync/models"
	. "github.com/bianjieai/irita-sync/msgs"
	"gopkg.in/mgo.v2/txn"
	"gopkg.in/mgo.v2/bson"
)

func handleIdentityData(msgType string, identityInfo *models.Identity, mapObj map[string]models.Identity) (Ops []txn.Op) {
	switch msgType {
	case MsgTypeCreateIdentity:
		identityInfo.ID = bson.NewObjectId()
		op := txn.Op{
			C:      models.CollectionNameIdentity,
			Id:     bson.NewObjectId(),
			Insert: identityInfo,
		}
		Ops = append(Ops, op)
	case MsgTypeUpdateIdentity:
		v := identityInfo
		if obj, ok := mapObj[v.Id]; ok {
			v.ID = obj.ID
			//first one is new value
			if len(obj.Certificates) > 0 {
				v.Certificates = append(v.Certificates, obj.Certificates...)
			}
			if len(obj.Pubkeys) > 0 {
				v.Pubkeys = append(v.Pubkeys, obj.Pubkeys...)
			}
		}
		if !v.ID.Valid() {
			return
		}
		updateOp := txn.Op{
			C:      models.CollectionNameIdentity,
			Id:     v.ID,
			Assert: txn.DocExists,
			Update: bson.M{
				"$set": bson.M{
					models.IdentityIdTag:           v.Id,
					models.IdentityCertificatesTag: v.Certificates,
					models.IdentityCredentialsTag:  v.Credentials,
					models.IdentityOwnerTag:        v.Owner,
					models.IdentityPubkeysag:       v.Pubkeys,
				},
			},
		}
		Ops = append(Ops, updateOp)
	}

	return
}
