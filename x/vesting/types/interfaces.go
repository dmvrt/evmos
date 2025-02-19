package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected interface contract the vesting module
// requires for storing accounts.
type AccountKeeper interface {
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI
	SetAccount(sdk.Context, authtypes.AccountI)
	NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

// BankKeeper defines the expected interface contract the vesting module requires
// for creating vesting accounts with funds.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	IsSendEnabledCoins(ctx sdk.Context, coins ...sdk.Coin) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	BlockedAddr(addr sdk.AccAddress) bool
}

// StakingKeeper defines the expected interface contract the vesting module
// requires for finding and changing the delegated tokens, used in clawback.
type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
	GetDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) []stakingtypes.Delegation
	GetUnbondingDelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) []stakingtypes.UnbondingDelegation
	GetValidator(ctx sdk.Context, valAddr sdk.ValAddress) (stakingtypes.Validator, bool)

	// Created custom Agoric stakingkeeper functions on vestingkeeper (commented out below)
	// GetDelegatorUnbonding(ctx sdk.Context, delegator sdk.AccAddress) sdk.Int
	// GetDelegatorBonded(ctx sdk.Context, delegator sdk.AccAddress) sdk.Int
	// TransferUnbonding(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, valAddr sdk.ValAddress, wantAmt sdk.Int) sdk.Int
	// TransferDelegation(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, valAddr sdk.ValAddress, wantShares sdk.Dec) sdk.Dec
	GetUnbondingDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (stakingtypes.UnbondingDelegation, bool)
	HasMaxUnbondingDelegationEntries(ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress) bool
	SetUnbondingDelegationEntry(ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress, creationHeight int64, minTime time.Time, balance sdk.Int) stakingtypes.UnbondingDelegation
	InsertUBDQueue(ctx sdk.Context, ubd stakingtypes.UnbondingDelegation, completionTime time.Time)
	RemoveUnbondingDelegation(ctx sdk.Context, ubd stakingtypes.UnbondingDelegation)
	SetUnbondingDelegation(ctx sdk.Context, ubd stakingtypes.UnbondingDelegation)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (stakingtypes.Delegation, bool)
	GetRedelegation(ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) (stakingtypes.Redelegation, bool)
	MaxEntries(ctx sdk.Context) uint32
	SetDelegation(ctx sdk.Context, delegation stakingtypes.Delegation)
	RemoveDelegation(ctx sdk.Context, delegation stakingtypes.Delegation)
	GetRedelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) []stakingtypes.Redelegation
	SetRedelegationEntry(ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorSrcAddr, validatorDstAddr sdk.ValAddress, creationHeight int64, minTime time.Time, balance sdk.Int, sharesSrc, sharesDst sdk.Dec) stakingtypes.Redelegation
	InsertRedelegationQueue(ctx sdk.Context, red stakingtypes.Redelegation, completionTime time.Time)
	SetRedelegation(ctx sdk.Context, red stakingtypes.Redelegation)
	RemoveRedelegation(ctx sdk.Context, red stakingtypes.Redelegation)
}

// DistributionHooks is the expected interface for distribution module hooks.
type DistributionHooks interface {
	// AllowWithdrawAddr tells whether to honor the delegation withdraw
	// address associated with the address (if any). The distribution
	// keeper will call this before each reward withdrawal.
	// If multiple distribution hooks are set, then any of them may
	// disallow the withdraw address.
	AllowWithdrawAddr(ctx sdk.Context, delAddr sdk.AccAddress) bool

	// AfterDelegationReward is called after the reward has been transferred the address.
	AfterDelegationReward(ctx sdk.Context, delAddr, withdrawAddr sdk.AccAddress, reward sdk.Coins)
}

// type StakingHooks interface {
// 	BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)        // Must be called when a delegation is created
// 	BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) // Must be called when a delegation's shares are modified
// 	BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)        // Must be called when a delegation is removed
// 	AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress)
// }
