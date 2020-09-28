package msgs

import (
	"github.com/bianjieai/irita-sync/models"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	service "github.com/irisnet/irismod/modules/service/types"
	nft "github.com/irisnet/irismod/modules/nft/types"
	record "github.com/irisnet/irismod/modules/record/types"
	token "github.com/irisnet/irismod/modules/token/types"
	"gitlab.bianjie.ai/irita-pro/iritamod/modules/identity"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

const (
	MsgTypeSend         = "send"
	MsgTypeMultiSend    = "multisend"
	MsgTypeNFTMint      = "mint_nft"
	MsgTypeNFTEdit      = "edit_nft"
	MsgTypeNFTTransfer  = "transfer_nft"
	MsgTypeNFTBurn      = "burn_nft"
	MsgTypeIssueDenom   = "issue_denom"
	MsgTypeRecordCreate = "create_record"

	MsgTypeMintToken          = "mint_token"
	MsgTypeEditToken          = "edit_token"
	MsgTypeIssueToken         = "issue_token"
	MsgTypeTransferTokenOwner = "transfer_token_owner"

	MsgTypeDefineService             = "define_service"               // type for MsgDefineService
	MsgTypeBindService               = "bind_service"                 // type for MsgBindService
	MsgTypeUpdateServiceBinding      = "update_service_binding"       // type for MsgUpdateServiceBinding
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address" // type for MsgSetWithdrawAddress
	MsgTypeDisableServiceBinding     = "disable_service_binding"      // type for MsgDisableServiceBinding
	MsgTypeEnableServiceBinding      = "enable_service_binding"       // type for MsgEnableServiceBinding
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"       // type for MsgRefundServiceDeposit
	MsgTypeCallService               = "call_service"                 // type for MsgCallService
	MsgTypeRespondService            = "respond_service"              // type for MsgRespondService
	MsgTypePauseRequestContext       = "pause_request_context"        // type for MsgPauseRequestContext
	MsgTypeStartRequestContext       = "start_request_context"        // type for MsgStartRequestContext
	MsgTypeKillRequestContext        = "kill_request_context"         // type for MsgKillRequestContext
	MsgTypeUpdateRequestContext      = "update_request_context"       // type for MsgUpdateRequestContext
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"         // type for MsgWithdrawEarnedFees

	MsgTypeStakeCreateValidator           = "create_validator"
	MsgTypeStakeEditValidator             = "edit_validator"
	MsgTypeStakeDelegate                  = "delegate"
	MsgTypeStakeBeginUnbonding            = "begin_unbonding"
	MsgTypeBeginRedelegate                = "begin_redelegate"
	MsgTypeUnjail                         = "unjail"
	MsgTypeSetWithdrawAddress             = "set_withdraw_address"
	MsgTypeWithdrawDelegatorReward        = "withdraw_delegator_reward"
	MsgTypeMsgFundCommunityPool           = "fund_community_pool"
	MsgTypeMsgWithdrawValidatorCommission = "withdraw_validator_commission"
	MsgTypeSubmitProposal                 = "submit_proposal"
	MsgTypeDeposit                        = "deposit"
	MsgTypeVote                           = "vote"

	MsgTypeCreateHTLC = "create_htlc"
	MsgTypeClaimHTLC  = "claim_htlc"
	MsgTypeRefundHTLC = "refund_htlc"

	MsgTypeAddLiquidity    = "add_liquidity"
	MsgTypeRemoveLiquidity = "remove_liquidity"
	MsgTypeSwapOrder       = "swap_order"

	MsgTypeSubmitEvidence  = "submit_evidence"
	MsgTypeVerifyInvariant = "verify_invariant"

	MsgTypeUpdateIdentity = "update_identity"
	MsgTypeCreateIdentity = "create_identity"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.DocTxMsg
		Addrs    []string
		Signers  []string
	}
	Msg models.Msg
	SdkMsg sdk.Msg

	Coin models.Coin

	Coins []*Coin

	MsgSend = bank.MsgSend
	MsgMultiSend = bank.MsgMultiSend

	MsgNFTMint = nft.MsgMintNFT
	MsgNFTEdit = nft.MsgEditNFT
	MsgNFTTransfer = nft.MsgTransferNFT
	MsgNFTBurn = nft.MsgBurnNFT
	MsgIssueDenom = nft.MsgIssueDenom

	MsgDefineService = service.MsgDefineService
	MsgBindService = service.MsgBindService
	MsgCallService = service.MsgCallService
	MsgRespondService = service.MsgRespondService

	MsgUpdateServiceBinding = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit = service.MsgRefundServiceDeposit
	MsgPauseRequestContext = service.MsgPauseRequestContext
	MsgStartRequestContext = service.MsgStartRequestContext
	MsgKillRequestContext = service.MsgKillRequestContext
	MsgUpdateRequestContext = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees = service.MsgWithdrawEarnedFees

	MsgRecordCreate = record.MsgCreateRecord

	MsgIssueToken = token.MsgIssueToken
	MsgEditToken = token.MsgEditToken
	MsgMintToken = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner

	//MsgStakeCreate = stake.MsgCreateValidator
	//MsgStakeEdit = stake.MsgEditValidator
	//MsgStakeDelegate = stake.MsgDelegate
	//MsgStakeBeginUnbonding = stake.MsgUndelegate
	//MsgBeginRedelegate = stake.MsgBeginRedelegate
	//MsgUnjail = slashing.MsgUnjail
	//MsgStakeSetWithdrawAddress = dtypes.MsgSetWithdrawAddress
	//MsgWithdrawDelegatorReward = distribution.MsgWithdrawDelegatorReward
	//MsgFundCommunityPool = distribution.MsgFundCommunityPool
	//MsgWithdrawValidatorCommission = distribution.MsgWithdrawValidatorCommission
	//StakeValidator = stake.Validator
	//Delegation = stake.Delegation
	//UnbondingDelegation = stake.UnbondingDelegation
	//
	//MsgDeposit = gov.MsgDeposit
	//MsgSubmitProposal = gov.MsgSubmitProposal
	//TextProposal = gov.TextProposal
	//MsgVote = gov.MsgVote
	//Proposal = gov.Proposal
	//SdkVote = gov.Vote

	//MsgSwapOrder = coinswap.MsgSwapOrder
	//MsgAddLiquidity = coinswap.MsgAddLiquidity
	//MsgRemoveLiquidity = coinswap.MsgRemoveLiquidity
	//
	//MsgClaimHTLC = htlc.MsgClaimHTLC
	//MsgCreateHTLC = htlc.MsgCreateHTLC
	//MsgRefundHTLC = htlc.MsgRefundHTLC

	//MsgSubmitEvidence = evidence.MsgSubmitEvidence
	//MsgVerifyInvariant = crisis.MsgVerifyInvariant
	//
	MsgCreateIdentity = identity.MsgCreateIdentity
	MsgUpdateIdentity = identity.MsgUpdateIdentity
)
