package keeper

import (
	"context"
	"testing"

	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	epochingtypes "github.com/babylonlabs-io/babylon/x/epoching/types"
	"github.com/babylonlabs-io/babylon/x/monitor/types"
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

func Fuzz_Nosy_Hooks_AfterEpochBegins__(f *testing.F) {
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

		h.AfterEpochBegins(ctx, epoch)
	})
}

func Fuzz_Nosy_Hooks_AfterEpochEnds__(f *testing.F) {
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

		h.AfterEpochEnds(ctx, epoch)
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

func Fuzz_Nosy_Hooks_BeforeSlashThreshold__(f *testing.F) {
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
		var valSet epochingtypes.ValidatorSet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}

		h.BeforeSlashThreshold(ctx, valSet)
	})
}

func Fuzz_Nosy_Keeper_EndedEpochBtcHeight__(f *testing.F) {
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
		var req *types.QueryEndedEpochBtcHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EndedEpochBtcHeight(c, req)
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

func Fuzz_Nosy_Keeper_LightclientHeightAtCheckpointReported__(f *testing.F) {
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
		var hashString string
		fill_err = tp.Fill(&hashString)
		if fill_err != nil {
			return
		}

		k.LightclientHeightAtCheckpointReported(ctx, hashString)
	})
}

func Fuzz_Nosy_Keeper_LightclientHeightAtEpochEnd__(f *testing.F) {
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

		k.LightclientHeightAtEpochEnd(ctx, epoch)
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

func Fuzz_Nosy_Keeper_ReportedCheckpointBtcHeight__(f *testing.F) {
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
		var req *types.QueryReportedCheckpointBtcHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ReportedCheckpointBtcHeight(c, req)
	})
}

func Fuzz_Nosy_Keeper_removeCheckpointRecord__(f *testing.F) {
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

		k.removeCheckpointRecord(ctx, ckpt)
	})
}

func Fuzz_Nosy_Keeper_updateBtcLightClientHeightForCheckpoint__(f *testing.F) {
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

		k.updateBtcLightClientHeightForCheckpoint(ctx, ckpt)
	})
}

func Fuzz_Nosy_Keeper_updateBtcLightClientHeightForEpoch__(f *testing.F) {
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

		k.updateBtcLightClientHeightForEpoch(ctx, epoch)
	})
}

func Fuzz_Nosy_bytesToBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, heightBytes []byte) {
		bytesToBtcHeight(heightBytes)
	})
}
