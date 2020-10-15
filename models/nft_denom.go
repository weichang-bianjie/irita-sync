package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameNftDenom = "denom"

	NftDenomIdTag = "id"
)

type (
	NftDenom struct {
		Id          string `bson:"id"`     //denom id
		DenomName   string `bson:"name"`   //name
		Schema      string `bson:"schema"` //json schema'
		Creator     string `bson:"creator"`
		TxHash      string `bson:"tx_hash"`
		BlockHeight int64  `bson:"block_height"`
	}
)

func (m NftDenom) Name() string {
	return CollectionNameNftDenom
}

func (m NftDenom) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{NftDenomIdTag},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(m.Name(), indexes)
}

func (m NftDenom) PkKvPair() map[string]interface{} {
	return bson.M{NftDenomIdTag: m.Id}
}

func (m NftDenom) Save(denom NftDenom) error {
	return Save(denom)
}

func (m NftDenom) Update(denom NftDenom) error {
	return Update(denom)
}

func (m NftDenom) Delete(denom NftDenom) error {
	return Delete(denom)
}

func (m NftDenom) List(denom string) ([]NftDenom, error) {
	cond := bson.M{}
	if denom != "" {
		cond[NftDenomIdTag] = denom
	}
	var denoms []NftDenom
	fn := func(c *mgo.Collection) error {
		return c.Find(cond).All(&denoms)
	}
	if err := ExecCollection(m.Name(), fn); err != nil {
		return denoms, err
	}
	return denoms, nil
}
