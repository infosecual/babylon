package keeper

import (
	"context"
	"testing"

	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	finality_types "github.com/babylonlabs-io/babylon/x/finality/types"
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

func Fuzz_Nosy_Keeper_ActivatedHeight__(f *testing.F) {
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
		var req *finality_types.QueryActivatedHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ActivatedHeight(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ActiveFinalityProvidersAtHeight__(f *testing.F) {
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
		var req *finality_types.QueryActiveFinalityProvidersAtHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ActiveFinalityProvidersAtHeight(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_BeginBlocker__(f *testing.F) {
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

		k.BeginBlocker(ctx)
	})
}

func Fuzz_Nosy_Keeper_Block__(f *testing.F) {
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
		var req *finality_types.QueryBlockRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Block(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_DeleteMissedBlockBitmap__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.DeleteMissedBlockBitmap(ctx, fpPk)
	})
}

func Fuzz_Nosy_Keeper_Evidence__(f *testing.F) {
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
		var req *finality_types.QueryEvidenceRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Evidence(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ExportGenesis__(f *testing.F) {
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

		k.ExportGenesis(ctx)
	})
}

func Fuzz_Nosy_Keeper_FinalityProviderCurrentPower__(f *testing.F) {
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
		var req *finality_types.QueryFinalityProviderCurrentPowerRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.FinalityProviderCurrentPower(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_FinalityProviderPowerAtHeight__(f *testing.F) {
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
		var req *finality_types.QueryFinalityProviderPowerAtHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.FinalityProviderPowerAtHeight(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_GetBTCStakingActivatedHeight__(f *testing.F) {
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

		k.GetBTCStakingActivatedHeight(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetBlock__(f *testing.F) {
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

		k.GetBlock(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentEpoch__(f *testing.F) {
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

		k.GetCurrentEpoch(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentVotingPower__(f *testing.F) {
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
		var fpBTCPK []byte
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}

		k.GetCurrentVotingPower(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_GetEvidence__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetEvidence(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_GetFinalityProviderMissedBlocks__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.GetFinalityProviderMissedBlocks(ctx, fpPk)
	})
}

func Fuzz_Nosy_Keeper_GetFirstSlashableEvidence__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetFirstSlashableEvidence(ctx, fpBtcPK)
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

func Fuzz_Nosy_Keeper_GetLastPubRand__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetLastPubRand(ctx, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_GetLastPubRandCommit__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetLastPubRandCommit(ctx, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_GetMissedBlockBitmapValue__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var index int64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.GetMissedBlockBitmapValue(ctx, fpPk, index)
	})
}

func Fuzz_Nosy_Keeper_GetNextHeightToReward__(f *testing.F) {
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

		k.GetNextHeightToReward(ctx)
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

func Fuzz_Nosy_Keeper_GetPubRand__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetPubRand(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_GetPubRandCommitForHeight__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetPubRandCommitForHeight(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_GetSig__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetSig(ctx, height, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_GetSigSet__(f *testing.F) {
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

		k.GetSigSet(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetTimestampedPubRandCommitForHeight__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.GetTimestampedPubRandCommitForHeight(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_GetVoters__(f *testing.F) {
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

		k.GetVoters(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetVotingPower__(f *testing.F) {
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
		var fpBTCPK []byte
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		k.GetVotingPower(ctx, fpBTCPK, height)
	})
}

func Fuzz_Nosy_Keeper_GetVotingPowerDistCache__(f *testing.F) {
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

		k.GetVotingPowerDistCache(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetVotingPowerTable__(f *testing.F) {
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

		k.GetVotingPowerTable(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetVotingPowerTableOrdered__(f *testing.F) {
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

		k.GetVotingPowerTableOrdered(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_HandleFinalityProviderLiveness__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var missed bool
		fill_err = tp.Fill(&missed)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.HandleFinalityProviderLiveness(ctx, fpPk, missed, height)
	})
}

func Fuzz_Nosy_Keeper_HandleLiveness__(f *testing.F) {
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
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		k.HandleLiveness(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_HandleResumeFinalityProposal__(f *testing.F) {
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
		var fpPksHex []string
		fill_err = tp.Fill(&fpPksHex)
		if fill_err != nil {
			return
		}
		var haltingHeight uint32
		fill_err = tp.Fill(&haltingHeight)
		if fill_err != nil {
			return
		}

		k.HandleResumeFinalityProposal(ctx, fpPksHex, haltingHeight)
	})
}

func Fuzz_Nosy_Keeper_HandleRewarding__(f *testing.F) {
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
		var targetHeight int64
		fill_err = tp.Fill(&targetHeight)
		if fill_err != nil {
			return
		}

		k.HandleRewarding(ctx, targetHeight)
	})
}

func Fuzz_Nosy_Keeper_HasBlock__(f *testing.F) {
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

		k.HasBlock(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_HasEvidence__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.HasEvidence(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_HasPubRand__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.HasPubRand(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_HasSig__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.HasSig(ctx, height, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_HasTimestampedPubRand__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.HasTimestampedPubRand(ctx, fpBtcPK, height)
	})
}

func Fuzz_Nosy_Keeper_HasVotingPowerTable__(f *testing.F) {
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

		k.HasVotingPowerTable(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_IndexBlock__(f *testing.F) {
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

		k.IndexBlock(ctx)
	})
}

func Fuzz_Nosy_Keeper_InitGenesis__(f *testing.F) {
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
		var gs finality_types.GenesisState
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}

		k.InitGenesis(ctx, gs)
	})
}

func Fuzz_Nosy_Keeper_IsBTCStakingActivated__(f *testing.F) {
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

		k.IsBTCStakingActivated(ctx)
	})
}

func Fuzz_Nosy_Keeper_IsFinalityActive__(f *testing.F) {
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

		k.IsFinalityActive(ctx)
	})
}

// skipping Fuzz_Nosy_Keeper_IterateMissedBlockBitmap__ because parameters include func, chan, or unsupported interface: func(index int64, missed bool) (stop bool)

func Fuzz_Nosy_Keeper_ListBlocks__(f *testing.F) {
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
		var req *finality_types.QueryListBlocksRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ListBlocks(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ListEvidences__(f *testing.F) {
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
		var req *finality_types.QueryListEvidencesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ListEvidences(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ListPubRandCommit__(f *testing.F) {
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
		var req *finality_types.QueryListPubRandCommitRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ListPubRandCommit(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ListPublicRandomness__(f *testing.F) {
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
		var req *finality_types.QueryListPublicRandomnessRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ListPublicRandomness(ctx, req)
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

func Fuzz_Nosy_Keeper_MustProcessBtcDelegationActivated__(f *testing.F) {
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
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var sats uint64
		fill_err = tp.Fill(&sats)
		if fill_err != nil {
			return
		}

		k.MustProcessBtcDelegationActivated(ctx, fp, del, sats)
	})
}

func Fuzz_Nosy_Keeper_MustProcessBtcDelegationUnbonded__(f *testing.F) {
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
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var sats uint64
		fill_err = tp.Fill(&sats)
		if fill_err != nil {
			return
		}

		k.MustProcessBtcDelegationUnbonded(ctx, fp, del, sats)
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
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *finality_types.QueryParamsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Params(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_ProcessAllPowerDistUpdateEvents__(f *testing.F) {
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
		var dc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&dc)
		if fill_err != nil {
			return
		}
		var events []*btcstakingtypes.EventPowerDistUpdate
		fill_err = tp.Fill(&events)
		if fill_err != nil {
			return
		}
		if dc == nil {
			return
		}

		k.ProcessAllPowerDistUpdateEvents(ctx, dc, events)
	})
}

func Fuzz_Nosy_Keeper_RemoveVotingPowerDistCache__(f *testing.F) {
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

		k.RemoveVotingPowerDistCache(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_SetBlock__(f *testing.F) {
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
		var block *finality_types.IndexedBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		k.SetBlock(ctx, block)
	})
}

func Fuzz_Nosy_Keeper_SetEvidence__(f *testing.F) {
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
		var evidence *finality_types.Evidence
		fill_err = tp.Fill(&evidence)
		if fill_err != nil {
			return
		}
		if evidence == nil {
			return
		}

		k.SetEvidence(ctx, evidence)
	})
}

func Fuzz_Nosy_Keeper_SetMissedBlockBitmapChunk__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chunkIndex int64
		fill_err = tp.Fill(&chunkIndex)
		if fill_err != nil {
			return
		}
		var chunk []byte
		fill_err = tp.Fill(&chunk)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.SetMissedBlockBitmapChunk(ctx, fpPk, chunkIndex, chunk)
	})
}

func Fuzz_Nosy_Keeper_SetMissedBlockBitmapValue__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var index int64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var missed bool
		fill_err = tp.Fill(&missed)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.SetMissedBlockBitmapValue(ctx, fpPk, index, missed)
	})
}

func Fuzz_Nosy_Keeper_SetNextHeightToReward__(f *testing.F) {
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

		k.SetNextHeightToReward(ctx, height)
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
		var p finality_types.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_SetPubRand__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var pubRand bbn.SchnorrPubRand
		fill_err = tp.Fill(&pubRand)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.SetPubRand(ctx, fpBtcPK, height, pubRand)
	})
}

func Fuzz_Nosy_Keeper_SetPubRandCommit__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var prCommit *finality_types.PubRandCommit
		fill_err = tp.Fill(&prCommit)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil || prCommit == nil {
			return
		}

		k.SetPubRandCommit(ctx, fpBtcPK, prCommit)
	})
}

func Fuzz_Nosy_Keeper_SetSig__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var sig *bbn.SchnorrEOTSSig
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil || sig == nil {
			return
		}

		k.SetSig(ctx, height, fpBtcPK, sig)
	})
}

func Fuzz_Nosy_Keeper_SetVotingPower__(f *testing.F) {
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
		var fpBTCPK []byte
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var power uint64
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}

		k.SetVotingPower(ctx, fpBTCPK, height, power)
	})
}

func Fuzz_Nosy_Keeper_SetVotingPowerDistCache__(f *testing.F) {
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
		var dc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&dc)
		if fill_err != nil {
			return
		}
		if dc == nil {
			return
		}

		k.SetVotingPowerDistCache(ctx, height, dc)
	})
}

func Fuzz_Nosy_Keeper_SigningInfo__(f *testing.F) {
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
		var req *finality_types.QuerySigningInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.SigningInfo(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_SigningInfos__(f *testing.F) {
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
		var req *finality_types.QuerySigningInfosRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.SigningInfos(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_TallyBlocks__(f *testing.F) {
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

		k.TallyBlocks(ctx)
	})
}

func Fuzz_Nosy_Keeper_UpdatePowerDist__(f *testing.F) {
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

		k.UpdatePowerDist(ctx)
	})
}

func Fuzz_Nosy_Keeper_VotesAtHeight__(f *testing.F) {
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
		var req *finality_types.QueryVotesAtHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.VotesAtHeight(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_blockStore__(f *testing.F) {
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

		k.blockStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_blocks__(f *testing.F) {
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

		k.blocks(ctx)
	})
}

func Fuzz_Nosy_Keeper_evidenceFpStore__(f *testing.F) {
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
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		k.evidenceFpStore(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_evidenceStore__(f *testing.F) {
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

		k.evidenceStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_evidences__(f *testing.F) {
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

		k.evidences(ctx)
	})
}

func Fuzz_Nosy_Keeper_exportPubRandCommit__(f *testing.F) {
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

		k.exportPubRandCommit(ctx)
	})
}

func Fuzz_Nosy_Keeper_finalizeBlock__(f *testing.F) {
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
		var block *finality_types.IndexedBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		k.finalizeBlock(ctx, block)
	})
}

func Fuzz_Nosy_Keeper_fpVotingPowers__(f *testing.F) {
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

		k.fpVotingPowers(ctx)
	})
}

func Fuzz_Nosy_Keeper_getMissedBlockBitmapChunk__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var chunkIndex int64
		fill_err = tp.Fill(&chunkIndex)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.getMissedBlockBitmapChunk(ctx, fpPk, chunkIndex)
	})
}

func Fuzz_Nosy_Keeper_getNextHeightToFinalize__(f *testing.F) {
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

		k.getNextHeightToFinalize(ctx)
	})
}

func Fuzz_Nosy_Keeper_handleActivatedFinalityProvider__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.handleActivatedFinalityProvider(ctx, fpPk)
	})
}

func Fuzz_Nosy_Keeper_handleFPStateUpdates__(f *testing.F) {
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
		var prevDc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&prevDc)
		if fill_err != nil {
			return
		}
		var newDc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&newDc)
		if fill_err != nil {
			return
		}
		if prevDc == nil || newDc == nil {
			return
		}

		k.handleFPStateUpdates(ctx, prevDc, newDc)
	})
}

func Fuzz_Nosy_Keeper_jailSluggishFinalityProvider__(f *testing.F) {
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
		var fpBtcPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPk)
		if fill_err != nil {
			return
		}
		if fpBtcPk == nil {
			return
		}

		k.jailSluggishFinalityProvider(ctx, fpBtcPk)
	})
}

func Fuzz_Nosy_Keeper_loadFP__(f *testing.F) {
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
		var cacheFpByBtcPkHex map[string]*types.FinalityProvider
		fill_err = tp.Fill(&cacheFpByBtcPkHex)
		if fill_err != nil {
			return
		}
		var fpBTCPKHex string
		fill_err = tp.Fill(&fpBTCPKHex)
		if fill_err != nil {
			return
		}

		k.loadFP(ctx, cacheFpByBtcPkHex, fpBTCPKHex)
	})
}

func Fuzz_Nosy_Keeper_processPowerDistUpdateEventUnbond__(f *testing.F) {
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
		var cacheFpByBtcPkHex map[string]*types.FinalityProvider
		fill_err = tp.Fill(&cacheFpByBtcPkHex)
		if fill_err != nil {
			return
		}
		var btcDel *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var unbondedSatsByFpBtcPk map[string][]uint64
		fill_err = tp.Fill(&unbondedSatsByFpBtcPk)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		k.processPowerDistUpdateEventUnbond(ctx, cacheFpByBtcPkHex, btcDel, unbondedSatsByFpBtcPk)
	})
}

// skipping Fuzz_Nosy_Keeper_processRewardTracker__ because parameters include func, chan, or unsupported interface: func(fp github.com/cosmos/cosmos-sdk/sdk.AccAddress, del github.com/cosmos/cosmos-sdk/sdk.AccAddress, sats uint64)

func Fuzz_Nosy_Keeper_pubRandCommitFpStore__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.pubRandCommitFpStore(ctx, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_pubRandCommitStore__(f *testing.F) {
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

		k.pubRandCommitStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_pubRandFpStore__(f *testing.F) {
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
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		if fpBtcPK == nil {
			return
		}

		k.pubRandFpStore(ctx, fpBtcPK)
	})
}

func Fuzz_Nosy_Keeper_pubRandStore__(f *testing.F) {
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

		k.pubRandStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_publicRandomness__(f *testing.F) {
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

		k.publicRandomness(ctx)
	})
}

func Fuzz_Nosy_Keeper_recordMetrics__(f *testing.F) {
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
		var dc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&dc)
		if fill_err != nil {
			return
		}
		if dc == nil {
			return
		}

		k.recordMetrics(dc)
	})
}

func Fuzz_Nosy_Keeper_recordVotingPowerAndCache__(f *testing.F) {
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
		var newDc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&newDc)
		if fill_err != nil {
			return
		}
		if newDc == nil {
			return
		}

		k.recordVotingPowerAndCache(ctx, newDc)
	})
}

func Fuzz_Nosy_Keeper_rewardBTCStaking__(f *testing.F) {
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

		k.rewardBTCStaking(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_setNextHeightToFinalize__(f *testing.F) {
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

		k.setNextHeightToFinalize(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_signingInfosAndMissedBlock__(f *testing.F) {
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

		k.signingInfosAndMissedBlock(ctx)
	})
}

func Fuzz_Nosy_Keeper_slashFinalityProvider__(f *testing.F) {
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
		var fpBtcPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPk)
		if fill_err != nil {
			return
		}
		var evidence *finality_types.Evidence
		fill_err = tp.Fill(&evidence)
		if fill_err != nil {
			return
		}
		if fpBtcPk == nil || evidence == nil {
			return
		}

		k.slashFinalityProvider(ctx, fpBtcPk, evidence)
	})
}

func Fuzz_Nosy_Keeper_updateSigningInfo__(f *testing.F) {
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
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var missed bool
		fill_err = tp.Fill(&missed)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		k.updateSigningInfo(ctx, fpPk, missed, height)
	})
}

func Fuzz_Nosy_Keeper_voteHeightStore__(f *testing.F) {
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

		k.voteHeightStore(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_voteSigs__(f *testing.F) {
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

		k.voteSigs(ctx)
	})
}

func Fuzz_Nosy_Keeper_voteStore__(f *testing.F) {
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

		k.voteStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_votingPowerBbnBlockHeightStore__(f *testing.F) {
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

		k.votingPowerBbnBlockHeightStore(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_votingPowerDistCacheStore__(f *testing.F) {
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

		k.votingPowerDistCacheStore(ctx)
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

func Fuzz_Nosy_Keeper_votingPowersDistCacheBlkHeight__(f *testing.F) {
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

		k.votingPowersDistCacheBlkHeight(ctx)
	})
}

func Fuzz_Nosy_msgServer_AddFinalitySig__(f *testing.F) {
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
		var req *finality_types.MsgAddFinalitySig
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.AddFinalitySig(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_CommitPubRandList__(f *testing.F) {
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
		var req *finality_types.MsgCommitPubRandList
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.CommitPubRandList(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_ResumeFinalityProposal__(f *testing.F) {
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
		var req *finality_types.MsgResumeFinalityProposal
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.ResumeFinalityProposal(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_ShouldAcceptSigForHeight__(f *testing.F) {
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
		var block *finality_types.IndexedBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if block == nil {
			return
		}

		ms.ShouldAcceptSigForHeight(ctx, block)
	})
}

func Fuzz_Nosy_msgServer_UnjailFinalityProvider__(f *testing.F) {
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
		var req *finality_types.MsgUnjailFinalityProvider
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.UnjailFinalityProvider(ctx, req)
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
		var req *finality_types.MsgUpdateParams
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

func Fuzz_Nosy_msgServer_validateActivationHeight__(f *testing.F) {
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
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		ms.validateActivationHeight(ctx, height)
	})
}

func Fuzz_Nosy_convertToActiveFinalityProvidersAtHeightResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var finalityProvidersWithMeta []*bstypes.FinalityProviderWithMeta
		fill_err = tp.Fill(&finalityProvidersWithMeta)
		if fill_err != nil {
			return
		}

		convertToActiveFinalityProvidersAtHeightResponse(finalityProvidersWithMeta)
	})
}

func Fuzz_Nosy_convertToEvidenceListResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var evidences []*finality_types.Evidence
		fill_err = tp.Fill(&evidences)
		if fill_err != nil {
			return
		}

		convertToEvidenceListResponse(evidences)
	})
}

func Fuzz_Nosy_convertToSigningInfosResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signInfos []finality_types.FinalityProviderSigningInfo
		fill_err = tp.Fill(&signInfos)
		if fill_err != nil {
			return
		}

		convertToSigningInfosResponse(signInfos)
	})
}

func Fuzz_Nosy_parsePubKeyAndBlkHeightFromStoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		parsePubKeyAndBlkHeightFromStoreKey(key)
	})
}

func Fuzz_Nosy_tally__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpSet map[string]uint64
		fill_err = tp.Fill(&fpSet)
		if fill_err != nil {
			return
		}
		var voterBTCPKs map[string]struct{}
		fill_err = tp.Fill(&voterBTCPKs)
		if fill_err != nil {
			return
		}

		tally(fpSet, voterBTCPKs)
	})
}
