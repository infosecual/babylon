package keeper

import (
	"context"
	"testing"

	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/babylonlabs-io/babylon/x/btclightclient/types"
	btclctypes "github.com/babylonlabs-io/babylon/x/btclightclient/types"
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

// skipping Fuzz_Nosy_Keeper_SetHooks__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCLightClientHooks

func Fuzz_Nosy_Keeper_AfterBTCHeaderInserted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.AfterBTCHeaderInserted(ctx, headerInfo)
	})
}

func Fuzz_Nosy_Keeper_AfterBTCRollBack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.AfterBTCRollBack(ctx, headerInfo)
	})
}

func Fuzz_Nosy_Keeper_AfterBTCRollForward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.AfterBTCRollForward(ctx, headerInfo)
	})
}

func Fuzz_Nosy_Keeper_BaseHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.QueryBaseHeaderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BaseHeader(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_BlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
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

		k.BlockHeight(ctx, headerHash)
	})
}

func Fuzz_Nosy_Keeper_Contains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.QueryContainsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Contains(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_ContainsBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.QueryContainsBytesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.ContainsBytes(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_GetBTCNet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.GetBTCNet()
	})
}

func Fuzz_Nosy_Keeper_GetBaseBTCHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetBaseBTCHeader(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hash == nil {
			return
		}

		k.GetHeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_Keeper_GetHeaderByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		k.GetHeaderByHeight(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetMainChainFrom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}

		k.GetMainChainFrom(ctx, startHeight)
	})
}

func Fuzz_Nosy_Keeper_GetMainChainFromWithLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var limit uint32
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}

		k.GetMainChainFromWithLimit(ctx, startHeight, limit)
	})
}

func Fuzz_Nosy_Keeper_GetMainChainReverse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetMainChainReverse(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetMainChainUpTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var depth uint32
		fill_err = tp.Fill(&depth)
		if fill_err != nil {
			return
		}

		k.GetMainChainUpTo(ctx, depth)
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

func Fuzz_Nosy_Keeper_GetTipInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetTipInfo(ctx)
	})
}

func Fuzz_Nosy_Keeper_Hashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.QueryHashesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Hashes(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_HeaderDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.QueryHeaderDepthRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.HeaderDepth(ctx, req)
	})
}

func Fuzz_Nosy_Keeper_InsertHeaderInfos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var infos []*btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&infos)
		if fill_err != nil {
			return
		}

		k.InsertHeaderInfos(ctx, infos)
	})
}

func Fuzz_Nosy_Keeper_InsertHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headers []bbn.BTCHeaderBytes
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		k.InsertHeaders(ctx, headers)
	})
}

func Fuzz_Nosy_Keeper_InsertHeadersWithHookAndEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headers []bbn.BTCHeaderBytes
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		k.InsertHeadersWithHookAndEvents(ctx, headers)
	})
}

func Fuzz_Nosy_Keeper_Logger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 Keeper
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		_x1.Logger(ctx)
	})
}

func Fuzz_Nosy_Keeper_MainChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *btclctypes.QueryMainChainRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.MainChain(c, req)
	})
}

func Fuzz_Nosy_Keeper_MainChainDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerHashBytes *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHashBytes)
		if fill_err != nil {
			return
		}
		if headerHashBytes == nil {
			return
		}

		k.MainChainDepth(ctx, headerHashBytes)
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
		var req *btclctypes.QueryParamsRequest
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

func Fuzz_Nosy_Keeper_SetBaseBTCHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var baseBTCHeader types.BTCHeaderInfo
		fill_err = tp.Fill(&baseBTCHeader)
		if fill_err != nil {
			return
		}

		k.SetBaseBTCHeader(ctx, baseBTCHeader)
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
		var p btclctypes.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_Tip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

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
		var req *btclctypes.QueryTipRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Tip(c, req)
	})
}

// skipping Fuzz_Nosy_Keeper_emitTypedEventWithLog__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Keeper_headersState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.headersState(ctx)
	})
}

func Fuzz_Nosy_Keeper_insertHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.insertHandler()
	})
}

func Fuzz_Nosy_Keeper_insertHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headers []*wire.BlockHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		k.insertHeaders(ctx, headers)
	})
}

// skipping Fuzz_Nosy_Keeper_insertHeadersInternal__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, s github.com/babylonlabs-io/babylon/x/btclightclient/keeper.headersState, result *github.com/babylonlabs-io/babylon/x/btclightclient/types.InsertResult) error

func Fuzz_Nosy_Keeper_insertHeadersWithHookAndEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headers []*wire.BlockHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		k.insertHeadersWithHookAndEvents(ctx, headers)
	})
}

func Fuzz_Nosy_Keeper_triggerEventAndHandleHooksHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		k.triggerEventAndHandleHooksHandler()
	})
}

func Fuzz_Nosy_Keeper_triggerHeaderInserted__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.triggerHeaderInserted(ctx, headerInfo)
	})
}

func Fuzz_Nosy_Keeper_triggerRollBack__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.triggerRollBack(ctx, headerInfo)
	})
}

func Fuzz_Nosy_Keeper_triggerRollForward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var headerInfo *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&headerInfo)
		if fill_err != nil {
			return
		}
		if headerInfo == nil {
			return
		}

		k.triggerRollForward(ctx, headerInfo)
	})
}

func Fuzz_Nosy_headersState_BaseHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.BaseHeader()
	})
}

func Fuzz_Nosy_headersState_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var hash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hash == nil {
			return
		}

		s.GetHeaderByHash(hash)
	})
}

func Fuzz_Nosy_headersState_GetHeaderByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		s.GetHeaderByHeight(height)
	})
}

func Fuzz_Nosy_headersState_GetTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.GetTip()
	})
}

func Fuzz_Nosy_headersState_HeaderExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var hash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hash == nil {
			return
		}

		s.HeaderExists(hash)
	})
}

// skipping Fuzz_Nosy_headersState_IterateForwardHeaders__ because parameters include func, chan, or unsupported interface: func(*github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCHeaderInfo) bool

// skipping Fuzz_Nosy_headersState_IterateReverseHeaders__ because parameters include func, chan, or unsupported interface: func(*github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCHeaderInfo) bool

func Fuzz_Nosy_headersState_TipExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		s.TipExists()
	})
}

func Fuzz_Nosy_headersState_deleteHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var h *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		s.deleteHeader(h)
	})
}

func Fuzz_Nosy_headersState_insertHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var h *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		s.insertHeader(h)
	})
}

func Fuzz_Nosy_headersState_rollBackHeadersUpTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s headersState
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}

		s.rollBackHeadersUpTo(height)
	})
}

func Fuzz_Nosy_msgServer_InsertHeaders__(f *testing.F) {
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *btclctypes.MsgInsertHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m.InsertHeaders(ctx, msg)
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
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *btclctypes.MsgUpdateParams
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.UpdateParams(ctx, req)
	})
}

func Fuzz_Nosy_msgServer_canInsertHeaders__(f *testing.F) {
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
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var reporterAddress sdk.AccAddress
		fill_err = tp.Fill(&reporterAddress)
		if fill_err != nil {
			return
		}

		m.canInsertHeaders(sdkCtx, reporterAddress)
	})
}

func Fuzz_Nosy_BtcHeadersBytesToBlockHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var headers []bbn.BTCHeaderBytes
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		BtcHeadersBytesToBlockHeader(headers)
	})
}
