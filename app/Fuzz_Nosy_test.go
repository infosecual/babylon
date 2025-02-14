package app

import (
	"testing"

	"cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
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

func Fuzz_Nosy_BabylonApp_BeginBlockForks__(f *testing.F) {
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

		a1 := NewTmpBabylonApp()
		a1.BeginBlockForks(ctx)
	})
}

func Fuzz_Nosy_BabylonApp_BeginBlocker__(f *testing.F) {
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

		a1 := NewTmpBabylonApp()
		a1.BeginBlocker(ctx)
	})
}

func Fuzz_Nosy_BabylonApp_EndBlocker__(f *testing.F) {
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

		a1 := NewTmpBabylonApp()
		a1.EndBlocker(ctx)
	})
}

func Fuzz_Nosy_BabylonApp_ExportAppStateAndValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var forZeroHeight bool
		fill_err = tp.Fill(&forZeroHeight)
		if fill_err != nil {
			return
		}
		var jailAllowedAddrs []string
		fill_err = tp.Fill(&jailAllowedAddrs)
		if fill_err != nil {
			return
		}
		var modulesToExport []string
		fill_err = tp.Fill(&modulesToExport)
		if fill_err != nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs, modulesToExport)
	})
}

func Fuzz_Nosy_BabylonApp_InitChainer__(f *testing.F) {
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
		var req *abci.RequestInitChain
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.InitChainer(ctx, req)
	})
}

func Fuzz_Nosy_BabylonApp_LoadHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height int64) {
		a1 := NewTmpBabylonApp()
		a1.LoadHeight(height)
	})
}

func Fuzz_Nosy_BabylonApp_PreBlocker__(f *testing.F) {
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
		var _x2 *abci.RequestFinalizeBlock
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.PreBlocker(ctx, _x2)
	})
}

func Fuzz_Nosy_BabylonApp_RegisterAPIRoutes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var apiSvr *api.Server
		fill_err = tp.Fill(&apiSvr)
		if fill_err != nil {
			return
		}
		var apiConfig config.APIConfig
		fill_err = tp.Fill(&apiConfig)
		if fill_err != nil {
			return
		}
		if apiSvr == nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.RegisterAPIRoutes(apiSvr, apiConfig)
	})
}

func Fuzz_Nosy_BabylonApp_RegisterNodeService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}
		var cfg config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.RegisterNodeService(clientCtx, cfg)
	})
}

func Fuzz_Nosy_BabylonApp_RegisterTendermintService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.RegisterTendermintService(clientCtx)
	})
}

func Fuzz_Nosy_BabylonApp_RegisterTxService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.RegisterTxService(clientCtx)
	})
}

func Fuzz_Nosy_BabylonApp_prepForZeroHeightGenesis__(f *testing.F) {
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
		var jailAllowedAddrs []string
		fill_err = tp.Fill(&jailAllowedAddrs)
		if fill_err != nil {
			return
		}

		a1 := NewTmpBabylonApp()
		a1.prepForZeroHeightGenesis(ctx, jailAllowedAddrs)
	})
}

func Fuzz_Nosy_AddTestAddrs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *BabylonApp
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var accNum int
		fill_err = tp.Fill(&accNum)
		if fill_err != nil {
			return
		}
		var accAmt math.Int
		fill_err = tp.Fill(&accAmt)
		if fill_err != nil {
			return
		}
		if a1 == nil {
			return
		}

		AddTestAddrs(a1, ctx, accNum, accAmt)
	})
}

func Fuzz_Nosy_createRandomAccounts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, accNum int) {
		createRandomAccounts(accNum)
	})
}

func Fuzz_Nosy_initAccountWithCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a1 *BabylonApp
		fill_err = tp.Fill(&a1)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var coins sdk.Coins
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		if a1 == nil {
			return
		}

		initAccountWithCoins(a1, ctx, addr, coins)
	})
}

// skipping Fuzz_Nosy_setup__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.BlsSigner
