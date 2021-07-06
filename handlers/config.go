package handlers

import (
	"github.com/bianjieai/irita-sync/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// PrefixValidator is the prefix for validator keys
	PrefixValidator = "val"
	// PrefixConsensus is the prefix for consensus keys
	PrefixConsensus = "cons"
	// PrefixPublic is the prefix for public keys
	PrefixPublic = "pub"
	// PrefixOperator is the prefix for operator keys
	PrefixOperator = "oper"
)

func initBech32Prefix(conf *config.Config) {
	var (
		// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
		Bech32PrefixAccAddr = conf.Server.Bech32AccPrefix
		// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
		Bech32PrefixAccPub = Bech32PrefixAccAddr + PrefixPublic
		// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
		Bech32PrefixValAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator
		// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
		Bech32PrefixValPub = Bech32PrefixAccAddr + PrefixValidator + PrefixOperator + PrefixPublic
		// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
		Bech32PrefixConsAddr = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus
		// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
		Bech32PrefixConsPub = Bech32PrefixAccAddr + PrefixValidator + PrefixConsensus + PrefixPublic
	)
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
	cfg.Seal()
}
