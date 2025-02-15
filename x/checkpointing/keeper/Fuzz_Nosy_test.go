package keeper

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon/btctxformatter"
	"github.com/babylonlabs-io/babylon/crypto/bls12381"
	"github.com/babylonlabs-io/babylon/x/checkpointing/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// skipping Fuzz_Nosy_Keeper_SetEpochingKeeper__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_Keeper_SetHooks__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

func Fuzz_Nosy_CheckpointsState_CreateRawCkptWithMeta__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cs CheckpointsState
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		if ckptWithMeta == nil {
			return
		}

		cs.CreateRawCkptWithMeta(ckptWithMeta)
	})
}

func Fuzz_Nosy_CheckpointsState_GetRawCkptWithMeta__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cs CheckpointsState
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		cs.GetRawCkptWithMeta(epoch)
	})
}

// skipping Fuzz_Nosy_CheckpointsState_GetRawCkptsWithMetaByStatus__ because parameters include func, chan, or unsupported interface: func(*github.com/babylonlabs-io/babylon/x/checkpointing/types.RawCheckpointWithMeta) bool

func Fuzz_Nosy_CheckpointsState_UpdateCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cs CheckpointsState
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cs.UpdateCheckpoint(ckpt)
	})
}

func Fuzz_Nosy_CheckpointsState_UpdateCkptStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cs CheckpointsState
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status checkpointingtypes.CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cs.UpdateCkptStatus(ckpt, status)
	})
}

func Fuzz_Nosy_Keeper_AddRawCheckpoint__(f *testing.F) {
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
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		if ckptWithMeta == nil {
			return
		}

		k.AddRawCheckpoint(ctx, ckptWithMeta)
	})
}

func Fuzz_Nosy_Keeper_AfterBlsKeyRegistered__(f *testing.F) {
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

		k.AfterBlsKeyRegistered(ctx, valAddr)
	})
}

func Fuzz_Nosy_Keeper_AfterRawCheckpointBlsSigVerified__(f *testing.F) {
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
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		k.AfterRawCheckpointBlsSigVerified(ctx, ckpt)
	})
}

func Fuzz_Nosy_Keeper_AfterRawCheckpointConfirmed__(f *testing.F) {
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

		k.AfterRawCheckpointConfirmed(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_AfterRawCheckpointFinalized__(f *testing.F) {
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

		k.AfterRawCheckpointFinalized(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_AfterRawCheckpointForgotten__(f *testing.F) {
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
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		k.AfterRawCheckpointForgotten(ctx, ckpt)
	})
}

func Fuzz_Nosy_Keeper_AfterRawCheckpointSealed__(f *testing.F) {
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

		k.AfterRawCheckpointSealed(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_BlsPublicKeyList__(f *testing.F) {
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
		var req *checkpointingtypes.QueryBlsPublicKeyListRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BlsPublicKeyList(c, req)
	})
}

func Fuzz_Nosy_Keeper_BuildRawCheckpoint__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var blockHash checkpointingtypes.BlockHash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		k.BuildRawCheckpoint(ctx, epochNum, blockHash)
	})
}

func Fuzz_Nosy_Keeper_CheckpointsState__(f *testing.F) {
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

		k.CheckpointsState(ctx)
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

func Fuzz_Nosy_Keeper_CreateRegistration__(f *testing.F) {
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
		var blsPubKey bls12381.PublicKey
		fill_err = tp.Fill(&blsPubKey)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.CreateRegistration(ctx, blsPubKey, valAddr)
	})
}

func Fuzz_Nosy_Keeper_EpochStatus__(f *testing.F) {
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
		var req *checkpointingtypes.QueryEpochStatusRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochStatus(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_GetBLSPubKeySet__(f *testing.F) {
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

		k.GetBLSPubKeySet(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetBlsPubKey__(f *testing.F) {
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
		var address sdk.ValAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		k.GetBlsPubKey(ctx, address)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentValidatorBlsKeySet__(f *testing.F) {
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

		k.GetCurrentValidatorBlsKeySet(ctx)
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

func Fuzz_Nosy_Keeper_GetEpochByHeight__(f *testing.F) {
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

		k.GetEpochByHeight(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetLastCheckpointedEpoch__(f *testing.F) {
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

		k.GetLastCheckpointedEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetLastFinalizedEpoch__(f *testing.F) {
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

		k.GetLastFinalizedEpoch(ctx)
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

func Fuzz_Nosy_Keeper_GetRawCheckpoint__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}

		k.GetRawCheckpoint(ctx, epochNum)
	})
}

func Fuzz_Nosy_Keeper_GetStatus__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}

		k.GetStatus(ctx, epochNum)
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

func Fuzz_Nosy_Keeper_GetValAddr__(f *testing.F) {
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
		var key bls12381.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		k.GetValAddr(ctx, key)
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

func Fuzz_Nosy_Keeper_GetValidatorAddress__(f *testing.F) {
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

		k.GetValidatorAddress(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetValidatorBlsKeySet__(f *testing.F) {
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

		k.GetValidatorBlsKeySet(ctx, epochNumber)
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

func Fuzz_Nosy_Keeper_GetVotingPowerByAddress__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		k.GetVotingPowerByAddress(ctx, epochNum, valAddr)
	})
}

func Fuzz_Nosy_Keeper_InitValidatorBLSSet__(f *testing.F) {
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

		k.InitValidatorBLSSet(ctx)
	})
}

func Fuzz_Nosy_Keeper_LastCheckpointWithStatus__(f *testing.F) {
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
		var req *types.QueryLastCheckpointWithStatusRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.LastCheckpointWithStatus(ctx, req)
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

func Fuzz_Nosy_Keeper_RawCheckpoint__(f *testing.F) {
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
		var req *types.QueryRawCheckpointRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.RawCheckpoint(c, req)
	})
}

func Fuzz_Nosy_Keeper_RawCheckpointList__(f *testing.F) {
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
		var req *types.QueryRawCheckpointListRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.RawCheckpointList(c, req)
	})
}

func Fuzz_Nosy_Keeper_RawCheckpoints__(f *testing.F) {
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
		var req *types.QueryRawCheckpointsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.RawCheckpoints(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_RecentEpochStatusCount__(f *testing.F) {
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
		var req *types.QueryRecentEpochStatusCountRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.RecentEpochStatusCount(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_RegistrationState__(f *testing.F) {
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

		k.RegistrationState(ctx)
	})
}

func Fuzz_Nosy_Keeper_SealCheckpoint__(f *testing.F) {
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
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		if ckptWithMeta == nil {
			return
		}

		k.SealCheckpoint(ctx, ckptWithMeta)
	})
}

func Fuzz_Nosy_Keeper_SetCheckpointConfirmed__(f *testing.F) {
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

		k.SetCheckpointConfirmed(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_SetCheckpointFinalized__(f *testing.F) {
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

		k.SetCheckpointFinalized(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_SetCheckpointForgotten__(f *testing.F) {
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

		k.SetCheckpointForgotten(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_SetCheckpointSubmitted__(f *testing.F) {
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

		k.SetCheckpointSubmitted(ctx, epoch)
	})
}

func Fuzz_Nosy_Keeper_SetGenBlsKeys__(f *testing.F) {
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
		var genKeys []*types.GenesisKey
		fill_err = tp.Fill(&genKeys)
		if fill_err != nil {
			return
		}

		k.SetGenBlsKeys(ctx, genKeys)
	})
}

func Fuzz_Nosy_Keeper_SetLastFinalizedEpoch__(f *testing.F) {
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

		k.SetLastFinalizedEpoch(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_SignBLS__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var blockHash checkpointingtypes.BlockHash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}

		k.SignBLS(epochNum, blockHash)
	})
}

func Fuzz_Nosy_Keeper_UpdateCheckpoint__(f *testing.F) {
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
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		if ckptWithMeta == nil {
			return
		}

		k.UpdateCheckpoint(ctx, ckptWithMeta)
	})
}

func Fuzz_Nosy_Keeper_VerifyBLSSig__(f *testing.F) {
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
		var sig *checkpointingtypes.BlsSig
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if sig == nil {
			return
		}

		k.VerifyBLSSig(ctx, sig)
	})
}

func Fuzz_Nosy_Keeper_VerifyCheckpoint__(f *testing.F) {
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
		var checkpoint btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&checkpoint)
		if fill_err != nil {
			return
		}

		k.VerifyCheckpoint(ctx, checkpoint)
	})
}

func Fuzz_Nosy_Keeper_VerifyRawCheckpoint__(f *testing.F) {
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
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		k.VerifyRawCheckpoint(ctx, ckpt)
	})
}

func Fuzz_Nosy_Keeper_setCheckpointStatus__(f *testing.F) {
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
		var expectedStatus []types.CheckpointStatus
		fill_err = tp.Fill(&expectedStatus)
		if fill_err != nil {
			return
		}
		var to checkpointingtypes.CheckpointStatus
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}

		k.setCheckpointStatus(ctx, epoch, expectedStatus, to)
	})
}

func Fuzz_Nosy_Keeper_valBlsSetStore__(f *testing.F) {
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

		k.valBlsSetStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_validateCheckpointStatus__(f *testing.F) {
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
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		var allowedStatuses []checkpointingtypes.CheckpointStatus
		fill_err = tp.Fill(&allowedStatuses)
		if fill_err != nil {
			return
		}
		if ckptWithMeta == nil {
			return
		}

		k.validateCheckpointStatus(ckptWithMeta, allowedStatuses)
	})
}

func Fuzz_Nosy_Keeper_verifyCkptBytes__(f *testing.F) {
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
		var rawCheckpoint *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&rawCheckpoint)
		if fill_err != nil {
			return
		}
		if rawCheckpoint == nil {
			return
		}

		k.verifyCkptBytes(ctx, rawCheckpoint)
	})
}

func Fuzz_Nosy_RegistrationState_CreateRegistration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs RegistrationState
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var key bls12381.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		rs.CreateRegistration(key, valAddr)
	})
}

func Fuzz_Nosy_RegistrationState_Exists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs RegistrationState
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var addr sdk.ValAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		rs.Exists(addr)
	})
}

func Fuzz_Nosy_RegistrationState_GetBlsPubKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs RegistrationState
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var addr sdk.ValAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		rs.GetBlsPubKey(addr)
	})
}

func Fuzz_Nosy_RegistrationState_GetValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rs RegistrationState
		fill_err = tp.Fill(&rs)
		if fill_err != nil {
			return
		}
		var key bls12381.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		rs.GetValAddr(key)
	})
}

func Fuzz_Nosy_msgServer_WrappedCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m msgServer
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgWrappedCreateValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m.WrappedCreateValidator(goCtx, msg)
	})
}

func Fuzz_Nosy_convertToBlsPublicKeyListResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var valBLSKeys []*types.ValidatorWithBlsKey
		fill_err = tp.Fill(&valBLSKeys)
		if fill_err != nil {
			return
		}

		convertToBlsPublicKeyListResponse(valBLSKeys)
	})
}
