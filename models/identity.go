package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

const (
	CollectionNameIdentity = "identity"

	IdentityIdTag           = "id"
	IdentityOwnerTag        = "owner"
	IdentityPubkeysag       = "pubkeys"
	IdentityCertificatesTag = "certificates"
	IdentityCredentialsTag  = "credentials"
)

type (
	Identity struct {
		ID           bson.ObjectId `bson:"_id"`
		Id           string        `bson:"id"`
		Owner        string        `bson:"owner"`
		Pubkeys      []PubKeyInfo  `bson:"pubkeys"`
		Certificates []string      `bson:"certificates"`
		Credentials  string        `bson:"credentials"`
	}
	// PubKey represents a public key along with the corresponding algorithm
	PubKeyInfo struct {
		PubKey    string `bson:"pubkey"`
		Algorithm string `bson:"algorithm"`
	}
)

func (m Identity) Name() string {
	return CollectionNameIdentity
}

func (m Identity) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{IdentityIdTag},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(m.Name(), indexes)
}

func (m Identity) PkKvPair() map[string]interface{} {
	return bson.M{IdentityIdTag: m.Id}
}

func (m Identity) AllIdentitysMaps() (map[string]Identity, error) {
	cond := bson.M{}
	var idens []Identity
	fn := func(c *mgo.Collection) error {
		return c.Find(cond).All(&idens)
	}
	if err := ExecCollection(m.Name(), fn); err != nil {
		return nil, err
	}
	mapData := make(map[string]Identity, len(idens))
	for _, val := range idens {
		mapData[val.Id] = val
	}
	return mapData, nil
}
