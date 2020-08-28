package ibc

import (
	"time"
	"github.com/cosmos/cosmos-sdk/x/ibc/07-tendermint/types"
)

// Header defines the Tendermint consensus Header
type Header struct {
	SignedHeader `bson:"signed_header" yaml:"signed_header"` // contains the commitment root
	//ValidatorSet *tmtypes.ValidatorSet `bson:"validator_set" yaml:"validator_set"`
}

func (h Header) Build(msgHeader types.Header) Header {
	header := Header{
		SignedHeader: SignedHeader{
			BlockHeader: BlockHeader{
				ChainID:            msgHeader.ChainID,
				Height:             msgHeader.Height,
				Time:               msgHeader.Time,
				LastCommitHash:     msgHeader.LastCommitHash.String(),
				LastResultsHash:    msgHeader.LastResultsHash.String(),
				DataHash:           msgHeader.DataHash.String(),
				NextValidatorsHash: msgHeader.NextValidatorsHash.String(),
				ValidatorsHash:     msgHeader.ValidatorsHash.String(),
				ConsensusHash:      msgHeader.ConsensusHash.String(),
				EvidenceHash:       msgHeader.EvidenceHash.String(),
				AppHash:            msgHeader.AppHash.String(),
				ProposerAddress:    msgHeader.ProposerAddress.String(),
				LastBlockID: BlockID{
					Hash: msgHeader.LastBlockID.Hash.String(),
					PartsHeader: PartSetHeader{
						Total: msgHeader.LastBlockID.PartsHeader.Total,
						Hash:  msgHeader.LastBlockID.PartsHeader.Hash.String(),
					},
				},
				Version: Consensus{
					Block: msgHeader.Version.Block.Uint64(),
					App:   msgHeader.Version.App.Uint64(),
				},
			},
		},
	}
	return header
}

type SignedHeader struct {
	BlockHeader `bson:"header"`
	//Commit  *Commit `bson:"commit"`
}

// Consensus captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
type Consensus struct {
	Block uint64 `bson:"block"`
	App   uint64 `bson:"app"`
}

type BlockHeader struct {
	// basic block info
	Version Consensus `bson:"version"`
	ChainID string    `bson:"chain_id"`
	Height  int64     `bson:"height"`
	Time    time.Time `bson:"time"`

	// prev block info
	LastBlockID BlockID `bson:"last_block_id"`

	// hashes of block data
	LastCommitHash string `bson:"last_commit_hash"` // commit from validators from the last block
	DataHash       string `bson:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     string `bson:"validators_hash"`      // validators for the current block
	NextValidatorsHash string `bson:"next_validators_hash"` // validators for the next block
	ConsensusHash      string `bson:"consensus_hash"`       // consensus params for current block
	AppHash            string `bson:"app_hash"`             // state after txs from the previous block
	// root hash of all results from the txs from the previous block
	LastResultsHash string `bson:"last_results_hash"`

	// consensus info
	EvidenceHash    string `bson:"evidence_hash"`    // evidence included in the block
	ProposerAddress string `bson:"proposer_address"` // original proposer of the block
}

// BlockID defines the unique ID of a block as its Hash and its PartSetHeader
type BlockID struct {
	Hash        string        `bson:"hash"`
	PartsHeader PartSetHeader `bson:"parts"`
}

type PartSetHeader struct {
	Total int    `bson:"total"`
	Hash  string `bson:"hash"`
}
