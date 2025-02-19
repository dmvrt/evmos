package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	"github.com/tharsis/evmos/x/vesting/types"
)

var _ distrtypes.StakingHooks = Hooks{}

// Hooks wrapper struct for vesting keeper
type Hooks struct {
	k Keeper
}

// Return the wrapper struct
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

func (k Keeper) AllowWithdrawAddr(ctx sdk.Context, delAddr sdk.AccAddress) bool {
	acc := k.accountKeeper.GetAccount(ctx, delAddr)
	_, isClawback := acc.(*types.ClawbackVestingAccount)
	return !isClawback
}

func (k Keeper) AfterDelegationReward(ctx sdk.Context, delAddr, withdrawAddr sdk.AccAddress, reward sdk.Coins) {
	acc := k.accountKeeper.GetAccount(ctx, delAddr)
	cva, isClawback := acc.(*types.ClawbackVestingAccount)
	if isClawback {
		k.PostReward(ctx, *cva, reward)
	}
}

// ___________________________________________________________________________//

// staking hooks
func (h Hooks) AllowWithdrawAddr(ctx sdk.Context, delAddr sdk.AccAddress) bool {
	return h.k.AllowWithdrawAddr(ctx, delAddr)
}

func (h Hooks) AfterDelegationReward(ctx sdk.Context, delAddr, withdrawAddr sdk.AccAddress, reward sdk.Coins) {
	h.k.AfterDelegationReward(ctx, delAddr, withdrawAddr, reward)
}

func (h Hooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress)   {}
func (h Hooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {}
func (h Hooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}

func (h Hooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
}
func (h Hooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {}
