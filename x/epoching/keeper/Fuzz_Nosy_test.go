package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/babylonlabs-io/babylon/x/epoching/types"
	epochingtypes "github.com/babylonlabs-io/babylon/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

// skipping Fuzz_Nosy_Keeper_SetHooks__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.EpochingHooks

func Fuzz_Nosy_Hooks_AfterBlsKeyRegistered__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterBlsKeyRegistered(ctx, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterDelegationModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterDelegationModified(ctx, delAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterRawCheckpointBlsSigVerified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		h.AfterRawCheckpointBlsSigVerified(ctx, ckpt)
	})
}

func Fuzz_Nosy_Hooks_AfterRawCheckpointConfirmed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		h.AfterRawCheckpointConfirmed(ctx, epoch)
	})
}

func Fuzz_Nosy_Hooks_AfterRawCheckpointFinalized__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		h.AfterRawCheckpointFinalized(ctx, epoch)
	})
}

func Fuzz_Nosy_Hooks_AfterRawCheckpointForgotten__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		h.AfterRawCheckpointForgotten(ctx, ckpt)
	})
}

func Fuzz_Nosy_Hooks_AfterRawCheckpointSealed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		h.AfterRawCheckpointSealed(ctx, epoch)
	})
}

func Fuzz_Nosy_Hooks_AfterUnbondingInitiated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id uint64
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}

		h.AfterUnbondingInitiated(ctx, id)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorBeginUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr sdk.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorBonded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr sdk.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterValidatorBonded(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterValidatorCreated(ctx, valAddr)
	})
}

func Fuzz_Nosy_Hooks_AfterValidatorRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr sdk.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.AfterValidatorRemoved(ctx, consAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationCreated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.BeforeDelegationCreated(ctx, delAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationRemoved__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.BeforeDelegationRemoved(ctx, delAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeDelegationSharesModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeValidatorModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		h.BeforeValidatorModified(ctx, valAddr)
	})
}

func Fuzz_Nosy_Hooks_BeforeValidatorSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h Hooks
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		var fraction math.LegacyDec
		fill_err = tp.Fill(&fraction)
		if fill_err != nil {
			return
		}

		h.BeforeValidatorSlashed(ctx, valAddr, fraction)
	})
}

func Fuzz_Nosy_Keeper_AddSlashedValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.AddSlashedValidator(ctx, valAddr)
	})
}

func Fuzz_Nosy_Keeper_AfterEpochBegins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		k.AfterEpochBegins(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_AfterEpochEnds__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		k.AfterEpochEnds(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_ApplyAndReturnValidatorSetUpdates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.ApplyAndReturnValidatorSetUpdates(ctx)
	})
}

func Fuzz_Nosy_Keeper_ApplyMatureUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.ApplyMatureUnbonding(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_BeforeSlashThreshold__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valSet epochingtypes.ValidatorSet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}

		k.BeforeSlashThreshold(ctx, valSet)
	})
}

func Fuzz_Nosy_Keeper_CheckMsgCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *stakingtypes.MsgCreateValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		k.CheckMsgCreateValidator(ctx, msg)
	})
}

func Fuzz_Nosy_Keeper_ClearSlashedValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.ClearSlashedValidators(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_ClearValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.ClearValidatorSet(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_CurrentEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryCurrentEpochRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.CurrentEpoch(c, req)
	})
}

func Fuzz_Nosy_Keeper_DelegationLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryDelegationLifecycleRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.DelegationLifecycle(c, req)
	})
}

func Fuzz_Nosy_Keeper_EnqueueMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg types.QueuedMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		k.EnqueueMsg(ctx, msg)
	})
}

func Fuzz_Nosy_Keeper_EpochInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryEpochInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochInfo(c, req)
	})
}

func Fuzz_Nosy_Keeper_EpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryEpochMsgsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochMsgs(c, req)
	})
}

func Fuzz_Nosy_Keeper_EpochValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryEpochValSetRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochValSet(c, req)
	})
}

func Fuzz_Nosy_Keeper_EpochsInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryEpochsInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochsInfo(c, req)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentEpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetCurrentEpochMsgs(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentQueueLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetCurrentQueueLength(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetCurrentValidatorSet(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentValidatorVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.GetCurrentValidatorVotingPower(ctx, valAddr)
	})
}

func Fuzz_Nosy_Keeper_GetDelegationLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}

		k.GetDelegationLifecycle(ctx, delAddr)
	})
}

func Fuzz_Nosy_Keeper_GetEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetEpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetEpochMsgs(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetEpochNumByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		k.GetEpochNumByHeight(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetHistoricalEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetHistoricalEpoch(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetParams(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetPubKeyByConsAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var consAddr sdk.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}

		k.GetPubKeyByConsAddr(ctx, consAddr)
	})
}

func Fuzz_Nosy_Keeper_GetQueueLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetQueueLength(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetSlashedValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetSlashedValidators(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetSlashedVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetSlashedVotingPower(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetTotalVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetTotalVotingPower(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetValLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.GetValLifecycle(ctx, valAddr)
	})
}

func Fuzz_Nosy_Keeper_GetValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr sdk.ValAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		k.GetValidator(ctx, addr)
	})
}

func Fuzz_Nosy_Keeper_GetValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.GetValidatorSet(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetValidatorVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.GetValidatorVotingPower(ctx, epochNumber, valAddr)
	})
}

func Fuzz_Nosy_Keeper_HandleQueuedMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var qMsg *types.QueuedMessage
		fill_err = tp.Fill(&qMsg)
		if fill_err != nil {
			return
		}
		if qMsg == nil {
			return
		}

		k.HandleQueuedMsg(goCtx, qMsg)
	})
}

func Fuzz_Nosy_Keeper_Hooks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.Hooks()
	})
}

func Fuzz_Nosy_Keeper_IncEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.IncEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_InitEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.InitEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_InitMsgQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.InitMsgQueue(ctx)
	})
}

func Fuzz_Nosy_Keeper_InitSlashedVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.InitSlashedVotingPower(ctx)
	})
}

func Fuzz_Nosy_Keeper_InitValidatorSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.InitValidatorSet(ctx)
	})
}

func Fuzz_Nosy_Keeper_LatestEpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryLatestEpochMsgsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.LatestEpochMsgs(c, req)
	})
}

func Fuzz_Nosy_Keeper_Logger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.Logger(ctx)
	})
}

func Fuzz_Nosy_Keeper_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *epochingtypes.QueryParamsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Params(c, req)
	})
}

func Fuzz_Nosy_Keeper_RecordLastHeaderTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.RecordLastHeaderTime(ctx)
	})
}

func Fuzz_Nosy_Keeper_RecordNewDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		var amount *sdk.Coin
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var state types.BondState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if amount == nil {
			return
		}

		k.RecordNewDelegationState(ctx, delAddr, valAddr, amount, state)
	})
}

func Fuzz_Nosy_Keeper_RecordNewValState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		var state types.BondState
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		k.RecordNewValState(ctx, valAddr, state)
	})
}

func Fuzz_Nosy_Keeper_RecordSealerAppHashForPrevEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.RecordSealerAppHashForPrevEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_RecordSealerBlockHashForEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.RecordSealerBlockHashForEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_SetDelegationLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}
		var lc *types.DelegationLifecycle
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		if lc == nil {
			return
		}

		k.SetDelegationLifecycle(ctx, delAddr, lc)
	})
}

func Fuzz_Nosy_Keeper_SetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p epochingtypes.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_SetValLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		var lc *types.ValidatorLifecycle
		fill_err = tp.Fill(&lc)
		if fill_err != nil {
			return
		}
		if lc == nil {
			return
		}

		k.SetValLifecycle(ctx, valAddr, lc)
	})
}

func Fuzz_Nosy_Keeper_StkMsgCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *stakingtypes.MsgCreateValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		k.StkMsgCreateValidator(ctx, msg)
	})
}

func Fuzz_Nosy_Keeper_ValidatorLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var c context.Context
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var req *types.QueryValidatorLifecycleRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ValidatorLifecycle(c, req)
	})
}

func Fuzz_Nosy_Keeper_delegationLifecycleStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.delegationLifecycleStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_epochInfoStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.epochInfoStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_getAllMatureValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.getAllMatureValidators(ctx)
	})
}

func Fuzz_Nosy_Keeper_getEpochInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.getEpochInfo(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_incCurrentQueueLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.incCurrentQueueLength(ctx)
	})
}

func Fuzz_Nosy_Keeper_msgQueueLengthStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.msgQueueLengthStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_msgQueueStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.msgQueueStore(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_runUnwrappedMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var qMsg *types.QueuedMessage
		fill_err = tp.Fill(&qMsg)
		if fill_err != nil {
			return
		}
		if qMsg == nil {
			return
		}

		k.runUnwrappedMsg(ctx, qMsg)
	})
}

func Fuzz_Nosy_Keeper_setEpochInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epoch *types.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if epoch == nil {
			return
		}

		k.setEpochInfo(ctx, epochNumber, epoch)
	})
}

func Fuzz_Nosy_Keeper_setSlashedVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var power int64
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}

		k.setSlashedVotingPower(ctx, epochNumber, power)
	})
}

func Fuzz_Nosy_Keeper_slashedValSetStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.slashedValSetStore(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_slashedVotingPowerStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.slashedVotingPowerStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_valLifecycleStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.valLifecycleStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_valSetStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}

		k.valSetStore(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_votingPowerStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.votingPowerStore(ctx)
	})
}

func Fuzz_Nosy_msgServer_UpdateParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *types.MsgUpdateParams
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.UpdateParams(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_WrappedBeginRedelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgWrappedBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		ms.WrappedBeginRedelegate(goCtx, msg)
	})
}

func Fuzz_Nosy_msgServer_WrappedCancelUnbondingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		ms.WrappedCancelUnbondingDelegation(goCtx, msg)
	})
}

func Fuzz_Nosy_msgServer_WrappedDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgWrappedDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		ms.WrappedDelegate(goCtx, msg)
	})
}

func Fuzz_Nosy_msgServer_WrappedEditValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msgWrapped *types.MsgWrappedEditValidator
		fill_err = tp.Fill(&msgWrapped)
		if fill_err != nil {
			return
		}
		if msgWrapped == nil {
			return
		}

		ms.WrappedEditValidator(goCtx, msgWrapped)
	})
}

func Fuzz_Nosy_msgServer_WrappedStakingUpdateParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msgWrapped *types.MsgWrappedStakingUpdateParams
		fill_err = tp.Fill(&msgWrapped)
		if fill_err != nil {
			return
		}
		if msgWrapped == nil {
			return
		}

		ms.WrappedStakingUpdateParams(goCtx, msgWrapped)
	})
}

func Fuzz_Nosy_msgServer_WrappedUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgWrappedUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		ms.WrappedUndelegate(goCtx, msg)
	})
}

func Fuzz_Nosy_CalculateEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint64, epochInterval uint64) {
		CalculateEpochNumber(height, epochInterval)
	})
}

func Fuzz_Nosy_cacheTxContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txid []byte
		fill_err = tp.Fill(&txid)
		if fill_err != nil {
			return
		}
		var msgid []byte
		fill_err = tp.Fill(&msgid)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		cacheTxContext(ctx, txid, msgid, height)
	})
}

func Fuzz_Nosy_parseDelValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, delAddr string, valAddr string) {
		parseDelValAddr(delAddr, valAddr)
	})
}
