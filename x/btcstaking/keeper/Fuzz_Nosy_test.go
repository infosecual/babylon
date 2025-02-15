package keeper

import (
	"context"
	"testing"

	asig "github.com/babylonlabs-io/babylon/crypto/schnorr-adaptor-signature"
	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_Keeper_AddBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcDel *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		k.AddBTCDelegation(ctx, btcDel)
	})
}

func Fuzz_Nosy_Keeper_AddFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var msg *types.MsgCreateFinalityProvider
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		k.AddFinalityProvider(goCtx, msg)
	})
}

func Fuzz_Nosy_Keeper_BTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryBTCDelegationRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BTCDelegation(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_BTCDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryBTCDelegationsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BTCDelegations(ctx, req)
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

func Fuzz_Nosy_Keeper_ClearPowerDistUpdateEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcHeight uint32
		fill_err = tp.Fill(&btcHeight)
		if fill_err != nil {
			return
		}

		k.ClearPowerDistUpdateEvents(ctx, btcHeight)
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

func Fuzz_Nosy_Keeper_FinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.FinalityProvider(c, req)
	})
}

func Fuzz_Nosy_Keeper_FinalityProviderDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.FinalityProviderDelegations(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_FinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryFinalityProvidersRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.FinalityProviders(c, req)
	})
}

func Fuzz_Nosy_Keeper_GetAllParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetAllParams(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetAllPowerDistUpdateEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var lastBTCTip uint32
		fill_err = tp.Fill(&lastBTCTip)
		if fill_err != nil {
			return
		}
		var curBTCTip uint32
		fill_err = tp.Fill(&curBTCTip)
		if fill_err != nil {
			return
		}

		k.GetAllPowerDistUpdateEvents(ctx, lastBTCTip, curBTCTip)
	})
}

func Fuzz_Nosy_Keeper_GetBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var stakingTxHashStr string
		fill_err = tp.Fill(&stakingTxHashStr)
		if fill_err != nil {
			return
		}

		k.GetBTCDelegation(ctx, stakingTxHashStr)
	})
}

func Fuzz_Nosy_Keeper_GetBTCHeightAtBabylonHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var babylonHeight uint64
		fill_err = tp.Fill(&babylonHeight)
		if fill_err != nil {
			return
		}

		k.GetBTCHeightAtBabylonHeight(ctx, babylonHeight)
	})
}

func Fuzz_Nosy_Keeper_GetCurrentBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetCurrentBTCHeight(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetFinalityProvider(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_GetHeightToVersionMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetHeightToVersionMap(ctx)
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

func Fuzz_Nosy_Keeper_GetParamsByVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var v uint32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		k.GetParamsByVersion(ctx, v)
	})
}

func Fuzz_Nosy_Keeper_GetParamsForBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetParamsForBtcHeight(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetParamsWithVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.GetParamsWithVersion(ctx)
	})
}

func Fuzz_Nosy_Keeper_HasFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.HasFinalityProvider(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_IndexAllowedStakingTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if txHash == nil {
			return
		}

		k.IndexAllowedStakingTransaction(ctx, txHash)
	})
}

func Fuzz_Nosy_Keeper_IndexBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.IndexBTCHeight(ctx)
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
		var gs types.GenesisState
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}

		k.InitGenesis(ctx, gs)
	})
}

func Fuzz_Nosy_Keeper_IsStakingTransactionAllowed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if txHash == nil {
			return
		}

		k.IsStakingTransactionAllowed(ctx, txHash)
	})
}

func Fuzz_Nosy_Keeper_JailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.JailFinalityProvider(ctx, fpBTCPK)
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

func Fuzz_Nosy_Keeper_MinCommissionRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.MinCommissionRate(ctx)
	})
}

func Fuzz_Nosy_Keeper_OverwriteParamsAtVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var v uint32
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var p types.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.OverwriteParamsAtVersion(ctx, v, p)
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
		var req *btcstakingtypes.QueryParamsRequest
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

func Fuzz_Nosy_Keeper_ParamsByBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ParamsByBTCHeight(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_ParamsByVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *types.QueryParamsByVersionRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ParamsByVersion(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_SetHeightToVersionMap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var p *types.HeightToVersionMap
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		k.SetHeightToVersionMap(ctx, p)
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
		var p types.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_SlashFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.SlashFinalityProvider(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_UnjailFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.UnjailFinalityProvider(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_UpdateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var fp *btcstakingtypes.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		k.UpdateFinalityProvider(ctx, fp)
	})
}

func Fuzz_Nosy_Keeper_VerifyInclusionProofAndGetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var stakingTx *btcutil.Tx
		fill_err = tp.Fill(&stakingTx)
		if fill_err != nil {
			return
		}
		var confirmationDepth uint32
		fill_err = tp.Fill(&confirmationDepth)
		if fill_err != nil {
			return
		}
		var stakingTime uint32
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var unbondingTime uint32
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var inclusionProof *types.ParsedProofOfInclusion
		fill_err = tp.Fill(&inclusionProof)
		if fill_err != nil {
			return
		}
		if stakingTx == nil || inclusionProof == nil {
			return
		}

		k.VerifyInclusionProofAndGetHeight(ctx, stakingTx, confirmationDepth, stakingTime, unbondingTime, inclusionProof)
	})
}

func Fuzz_Nosy_Keeper_addCovenantSigsToBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcDel *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var parsedSlashingAdaptorSignatures []asig.AdaptorSignature
		fill_err = tp.Fill(&parsedSlashingAdaptorSignatures)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var parsedUnbondingSlashingAdaptorSignatures []asig.AdaptorSignature
		fill_err = tp.Fill(&parsedUnbondingSlashingAdaptorSignatures)
		if fill_err != nil {
			return
		}
		var params *bstypes.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil || params == nil {
			return
		}

		k.addCovenantSigsToBTCDelegation(ctx, btcDel, covPK, parsedSlashingAdaptorSignatures, unbondingTxSig, parsedUnbondingSlashingAdaptorSignatures, params)
	})
}

func Fuzz_Nosy_Keeper_addPowerDistUpdateEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcHeight uint32
		fill_err = tp.Fill(&btcHeight)
		if fill_err != nil {
			return
		}
		var event *types.EventPowerDistUpdate
		fill_err = tp.Fill(&event)
		if fill_err != nil {
			return
		}
		if event == nil {
			return
		}

		k.addPowerDistUpdateEvent(ctx, btcHeight, event)
	})
}

func Fuzz_Nosy_Keeper_blockHeightChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.blockHeightChains(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcDelegationStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcDelegationStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcDelegations(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcDelegatorFpStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcDelegatorFpStore(ctx, fpBTCPK)
	})
}

func Fuzz_Nosy_Keeper_btcDelegatorStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcDelegatorStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcDelegators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcDelegators(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcHeightStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.btcHeightStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_btcUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcDel *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var u *types.DelegatorUnbondingInfo
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if btcDel == nil || u == nil {
			return
		}

		k.btcUndelegate(ctx, btcDel, u)
	})
}

func Fuzz_Nosy_Keeper_eventIdxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.eventIdxs(ctx)
	})
}

func Fuzz_Nosy_Keeper_finalityProviderStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.finalityProviderStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_finalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.finalityProviders(ctx)
	})
}

func Fuzz_Nosy_Keeper_getBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		k.getBTCDelegation(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_Keeper_getBTCDelegatorDelegationIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var delBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&delBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil || delBTCPK == nil {
			return
		}

		k.getBTCDelegatorDelegationIndex(ctx, fpBTCPK, delBTCPK)
	})
}

func Fuzz_Nosy_Keeper_getBTCDelegatorDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var delBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&delBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil || delBTCPK == nil {
			return
		}

		k.getBTCDelegatorDelegations(ctx, fpBTCPK, delBTCPK)
	})
}

func Fuzz_Nosy_Keeper_getLastParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.getLastParams(ctx)
	})
}

// skipping Fuzz_Nosy_Keeper_iteratePowerDistUpdateEvents__ because parameters include func, chan, or unsupported interface: func(event *github.com/babylonlabs-io/babylon/x/btcstaking/types.EventPowerDistUpdate) bool

func Fuzz_Nosy_Keeper_nextParamsVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.nextParamsVersion(ctx)
	})
}

func Fuzz_Nosy_Keeper_paramsStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.paramsStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_powerDistUpdateEventBtcHeightStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcHeight uint32
		fill_err = tp.Fill(&btcHeight)
		if fill_err != nil {
			return
		}

		k.powerDistUpdateEventBtcHeightStore(ctx, btcHeight)
	})
}

func Fuzz_Nosy_Keeper_powerDistUpdateEventStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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

		k.powerDistUpdateEventStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_setBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var btcDel *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		k.setBTCDelegation(ctx, btcDel)
	})
}

func Fuzz_Nosy_Keeper_setBTCDelegatorDelegationIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var delBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&delBTCPK)
		if fill_err != nil {
			return
		}
		var btcDelIndex *types.BTCDelegatorDelegationIndex
		fill_err = tp.Fill(&btcDelIndex)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil || delBTCPK == nil || btcDelIndex == nil {
			return
		}

		k.setBTCDelegatorDelegationIndex(ctx, fpBTCPK, delBTCPK, btcDelIndex)
	})
}

func Fuzz_Nosy_Keeper_setBlockHeightChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var blocks *types.BlockHeightBbnToBtc
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if blocks == nil {
			return
		}

		k.setBlockHeightChains(ctx, blocks)
	})
}

func Fuzz_Nosy_Keeper_setEventIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var evt *types.EventIndex
		fill_err = tp.Fill(&evt)
		if fill_err != nil {
			return
		}
		if evt == nil {
			return
		}

		k.setEventIdx(ctx, evt)
	})
}

func Fuzz_Nosy_Keeper_setFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var fp *btcstakingtypes.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		k.setFinalityProvider(ctx, fp)
	})
}

func Fuzz_Nosy_msgServer_AddBTCDelegationInclusionProof__(f *testing.F) {
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
		var req *types.MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.AddBTCDelegationInclusionProof(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_AddCovenantSigs__(f *testing.F) {
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
		var req *types.MsgAddCovenantSigs
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.AddCovenantSigs(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_BTCUndelegate__(f *testing.F) {
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
		var req *types.MsgBTCUndelegate
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.BTCUndelegate(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_CreateBTCDelegation__(f *testing.F) {
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
		var req *types.MsgCreateBTCDelegation
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.CreateBTCDelegation(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_CreateFinalityProvider__(f *testing.F) {
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
		var req *types.MsgCreateFinalityProvider
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.CreateFinalityProvider(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_EditFinalityProvider__(f *testing.F) {
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
		var req *types.MsgEditFinalityProvider
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.EditFinalityProvider(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_SelectiveSlashingEvidence__(f *testing.F) {
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
		var req *types.MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.SelectiveSlashingEvidence(goCtx, req)
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

func Fuzz_Nosy_msgServer_getBTCDelWithParams__(f *testing.F) {
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
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		ms.getBTCDelWithParams(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_msgServer_getTimeInfoAndParams__(f *testing.F) {
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
		var parsedMsg *types.ParsedCreateDelegationMessage
		fill_err = tp.Fill(&parsedMsg)
		if fill_err != nil {
			return
		}
		if parsedMsg == nil {
			return
		}

		ms.getTimeInfoAndParams(ctx, parsedMsg)
	})
}

func Fuzz_Nosy_msgServer_isAllowListEnabled__(f *testing.F) {
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
		var p *bstypes.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		ms.isAllowListEnabled(ctx, p)
	})
}

func Fuzz_Nosy_containsInput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var inputHash *chainhash.Hash
		fill_err = tp.Fill(&inputHash)
		if fill_err != nil {
			return
		}
		var inputIdx uint32
		fill_err = tp.Fill(&inputIdx)
		if fill_err != nil {
			return
		}
		if tx == nil || inputHash == nil {
			return
		}

		containsInput(tx, inputHash, inputIdx)
	})
}

func Fuzz_Nosy_mustUint32FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		mustUint32FromBytes(b)
	})
}

func Fuzz_Nosy_parseBIP340PubKeysFromStoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		parseBIP340PubKeysFromStoreKey(key)
	})
}

func Fuzz_Nosy_parseUintsFromStoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		parseUintsFromStoreKey(key)
	})
}

func Fuzz_Nosy_uint32FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		uint32FromBytes(b)
	})
}

func Fuzz_Nosy_uint32ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		uint32ToBytes(v)
	})
}
