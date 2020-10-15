package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNameNft = "nft"

	NftIdTag        = "id"
	NftNameTag      = "name"
	NftUriTag       = "uri"
	NftDataTag      = "data"
	NftOwnerTag     = "owner"
	NftOfDenomIdTag = "denom_id"
)

type (
	Nft struct {
		ID      bson.ObjectId `bson:"_id"`
		Id      string        `bson:"id"`
		NftName string        `bson:"name"`
		Owner   string        `bson:"owner"`
		Uri     string        `bson:"uri"`
		Data    string        `bson:"data"`
		DenomId string        `bson:"denom_id"`
	}
)

func (m Nft) Name() string {
	return CollectionNameNft
}

func (m Nft) EnsureIndexes() {
	var indexes []mgo.Index
	indexes = append(indexes, mgo.Index{
		Key:        []string{NftIdTag, NftOfDenomIdTag},
		Unique:     true,
		Background: true,
	})
	ensureIndexes(m.Name(), indexes)
}

func (m Nft) PkKvPair() map[string]interface{} {
	return bson.M{NftIdTag: m.Id, NftOfDenomIdTag: m.DenomId}
}

func (m Nft) Save(nft Nft) error {
	return Save(nft)
}

func (m Nft) Update(nft Nft) error {
	return Update(nft)
}

func (m Nft) Delete(nft Nft) error {
	return Delete(nft)
}

func (m Nft) AllNftsMaps() (map[string]bson.ObjectId, error) {
	cond := bson.M{}
	var nfts []Nft
	fn := func(c *mgo.Collection) error {
		return c.Find(cond).All(&nfts)
	}
	if err := ExecCollection(m.Name(), fn); err != nil {
		return nil, err
	}
	mapData := make(map[string]bson.ObjectId, len(nfts))
	for _, val := range nfts {
		mapData[val.Id+val.DenomId] = val.ID
	}
	return mapData, nil
}
