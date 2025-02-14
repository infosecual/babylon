package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/store/prefix"
	bbn "github.com/babylonlabs-io/babylon/types"
	btccheckpointtypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	btcctypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	btclctypes "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	etypes "github.com/babylonlabs-io/babylon/x/epoching/types"
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

func Fuzz_Nosy_Keeper_epochDataStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.epochDataStore(ctx)
	})
}

func Fuzz_Nosy_Hooks_AfterBTCHeaderInserted__(f *testing.F) {
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		h.AfterBTCHeaderInserted(_x2, _x3)
	})
}

func Fuzz_Nosy_Hooks_AfterBTCRollBack__(f *testing.F) {
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
		var _x3 *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		h.AfterBTCRollBack(ctx, _x3)
	})
}

func Fuzz_Nosy_Hooks_AfterBTCRollForward__(f *testing.F) {
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
		var _x3 *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x3 == nil {
			return
		}

		h.AfterBTCRollForward(ctx, _x3)
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}

		h.AfterEpochBegins(_x2, _x3)
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 uint64
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}

		h.AfterEpochEnds(_x2, _x3)
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 etypes.ValidatorSet
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}

		h.BeforeSlashThreshold(_x2, _x3)
	})
}

func Fuzz_Nosy_Keeper_BtcCheckpointInfo__(f *testing.F) {
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
		var req *btccheckpointtypes.QueryBtcCheckpointInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BtcCheckpointInfo(c, req)
	})
}

func Fuzz_Nosy_Keeper_BtcCheckpointsInfo__(f *testing.F) {
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
		var req *btccheckpointtypes.QueryBtcCheckpointsInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BtcCheckpointsInfo(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_BtcLightClientUpdated__(f *testing.F) {
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

		k.BtcLightClientUpdated(ctx)
	})
}

func Fuzz_Nosy_Keeper_EpochSubmissions__(f *testing.F) {
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
		var req *btccheckpointtypes.QueryEpochSubmissionsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.EpochSubmissions(c, req)
	})
}

func Fuzz_Nosy_Keeper_GetBestSubmission__(f *testing.F) {
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

		k.GetBestSubmission(ctx, epochNumber)
	})
}

func Fuzz_Nosy_Keeper_GetBlockHeight__(f *testing.F) {
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
		var b *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		k.GetBlockHeight(ctx, b)
	})
}

func Fuzz_Nosy_Keeper_GetEpochBestSubmissionBtcInfo__(f *testing.F) {
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
		var ed *btccheckpointtypes.EpochData
		fill_err = tp.Fill(&ed)
		if fill_err != nil {
			return
		}
		if ed == nil {
			return
		}

		k.GetEpochBestSubmissionBtcInfo(ctx, ed)
	})
}

func Fuzz_Nosy_Keeper_GetEpochData__(f *testing.F) {
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
		var e uint64
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		k.GetEpochData(ctx, e)
	})
}

func Fuzz_Nosy_Keeper_GetExpectedTag__(f *testing.F) {
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

		k.GetExpectedTag(ctx)
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

func Fuzz_Nosy_Keeper_GetPowLimit__(f *testing.F) {
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

		k.GetPowLimit()
	})
}

func Fuzz_Nosy_Keeper_GetSubmissionBtcInfo__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}

		k.GetSubmissionBtcInfo(ctx, sk)
	})
}

func Fuzz_Nosy_Keeper_GetSubmissionData__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}

		k.GetSubmissionData(ctx, sk)
	})
}

func Fuzz_Nosy_Keeper_HasSubmission__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}

		k.HasSubmission(ctx, sk)
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

func Fuzz_Nosy_Keeper_OnTipChange__(f *testing.F) {
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

		k.OnTipChange(ctx)
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
		var req *btccheckpointtypes.QueryParamsRequest
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
		var p btccheckpointtypes.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_addEpochSubmission__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var sd btccheckpointtypes.SubmissionData
		fill_err = tp.Fill(&sd)
		if fill_err != nil {
			return
		}

		k.addEpochSubmission(ctx, epochNum, sk, sd)
	})
}

func Fuzz_Nosy_Keeper_checkAncestors__(f *testing.F) {
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
		var submisionEpoch uint64
		fill_err = tp.Fill(&submisionEpoch)
		if fill_err != nil {
			return
		}
		var newSubmissionInfo *btccheckpointtypes.SubmissionBtcInfo
		fill_err = tp.Fill(&newSubmissionInfo)
		if fill_err != nil {
			return
		}
		if newSubmissionInfo == nil {
			return
		}

		k.checkAncestors(ctx, submisionEpoch, newSubmissionInfo)
	})
}

func Fuzz_Nosy_Keeper_checkCheckpoints__(f *testing.F) {
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

		k.checkCheckpoints(ctx)
	})
}

func Fuzz_Nosy_Keeper_checkSubmissionStatus__(f *testing.F) {
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
		var info *btccheckpointtypes.SubmissionBtcInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		if info == nil {
			return
		}

		k.checkSubmissionStatus(ctx, info)
	})
}

func Fuzz_Nosy_Keeper_clearEpochData__(f *testing.F) {
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
		var epoch []byte
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		var epochDataStore prefix.Store
		fill_err = tp.Fill(&epochDataStore)
		if fill_err != nil {
			return
		}
		var currentEpoch *btccheckpointtypes.EpochData
		fill_err = tp.Fill(&currentEpoch)
		if fill_err != nil {
			return
		}
		if currentEpoch == nil {
			return
		}

		k.clearEpochData(ctx, epoch, epochDataStore, currentEpoch)
	})
}

func Fuzz_Nosy_Keeper_deleteSubmission__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}

		k.deleteSubmission(ctx, sk)
	})
}

func Fuzz_Nosy_Keeper_getCheckpointInfo__(f *testing.F) {
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
		var epochData *btccheckpointtypes.EpochData
		fill_err = tp.Fill(&epochData)
		if fill_err != nil {
			return
		}
		if epochData == nil {
			return
		}

		k.getCheckpointInfo(ctx, epochNum, epochData)
	})
}

func Fuzz_Nosy_Keeper_getEpochChanges__(f *testing.F) {
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
		var parentEpochBestSubmission *btccheckpointtypes.SubmissionBtcInfo
		fill_err = tp.Fill(&parentEpochBestSubmission)
		if fill_err != nil {
			return
		}
		var ed *btccheckpointtypes.EpochData
		fill_err = tp.Fill(&ed)
		if fill_err != nil {
			return
		}
		if parentEpochBestSubmission == nil || ed == nil {
			return
		}

		k.getEpochChanges(ctx, parentEpochBestSubmission, ed)
	})
}

func Fuzz_Nosy_Keeper_getLastFinalizedEpochNumber__(f *testing.F) {
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

		k.getLastFinalizedEpochNumber(ctx)
	})
}

func Fuzz_Nosy_Keeper_headerDepth__(f *testing.F) {
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
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		if headerHash == nil {
			return
		}

		k.headerDepth(ctx, headerHash)
	})
}

func Fuzz_Nosy_Keeper_saveEpochData__(f *testing.F) {
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
		var e uint64
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var ed *btccheckpointtypes.EpochData
		fill_err = tp.Fill(&ed)
		if fill_err != nil {
			return
		}
		if ed == nil {
			return
		}

		k.saveEpochData(ctx, e, ed)
	})
}

func Fuzz_Nosy_Keeper_saveSubmission__(f *testing.F) {
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
		var sk btccheckpointtypes.SubmissionKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var sd btccheckpointtypes.SubmissionData
		fill_err = tp.Fill(&sd)
		if fill_err != nil {
			return
		}

		k.saveSubmission(ctx, sk, sd)
	})
}

func Fuzz_Nosy_Keeper_setBtcLightClientUpdated__(f *testing.F) {
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

		k.setBtcLightClientUpdated(ctx)
	})
}

func Fuzz_Nosy_Keeper_setLastFinalizedEpochNumber__(f *testing.F) {
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

		k.setLastFinalizedEpochNumber(ctx, epoch)
	})
}

func Fuzz_Nosy_msgServer_InsertBTCSpvProof__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btcctypes.MsgInsertBTCSpvProof
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.InsertBTCSpvProof(ctx, req)
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
		var req *btccheckpointtypes.MsgUpdateParams
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

func Fuzz_Nosy_submissionBtcError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e submissionBtcError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}
