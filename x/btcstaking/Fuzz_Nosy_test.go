package btcstaking

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon/x/btcstaking/keeper"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

func Fuzz_Nosy_AppModule_BeginBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am AppModule
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		am.BeginBlock(ctx)
	})
}

func Fuzz_Nosy_AppModule_ConsensusVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModule
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.ConsensusVersion()
	})
}

func Fuzz_Nosy_AppModule_EndBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am AppModule
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		am.EndBlock(ctx)
	})
}

// skipping Fuzz_Nosy_AppModule_ExportGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

// skipping Fuzz_Nosy_AppModule_InitGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_AppModule_IsAppModule__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am AppModule
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}

		am.IsAppModule()
	})
}

func Fuzz_Nosy_AppModule_IsOnePerModuleType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var am AppModule
		fill_err = tp.Fill(&am)
		if fill_err != nil {
			return
		}

		am.IsOnePerModuleType()
	})
}

// skipping Fuzz_Nosy_AppModule_RegisterInvariants__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.InvariantRegistry

// skipping Fuzz_Nosy_AppModule_RegisterServices__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types/module.Configurator

// skipping Fuzz_Nosy_AppModuleBasic_DefaultGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_AppModuleBasic_GetQueryCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.GetQueryCmd()
	})
}

func Fuzz_Nosy_AppModuleBasic_GetTxCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a AppModuleBasic
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}

		a.GetTxCmd()
	})
}

func Fuzz_Nosy_AppModuleBasic_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_AppModuleBasic_RegisterGRPCGatewayRoutes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}
		var mux *runtime.ServeMux
		fill_err = tp.Fill(&mux)
		if fill_err != nil {
			return
		}
		if mux == nil {
			return
		}

		_x1.RegisterGRPCGatewayRoutes(clientCtx, mux)
	})
}

// skipping Fuzz_Nosy_AppModuleBasic_RegisterInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.InterfaceRegistry

func Fuzz_Nosy_AppModuleBasic_RegisterLegacyAminoCodec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var cdc *codec.LegacyAmino
		fill_err = tp.Fill(&cdc)
		if fill_err != nil {
			return
		}
		if cdc == nil {
			return
		}

		_x1.RegisterLegacyAminoCodec(cdc)
	})
}

// skipping Fuzz_Nosy_AppModuleBasic_ValidateGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_EndBlocker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var k keeper.Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		EndBlocker(ctx, k)
	})
}

func Fuzz_Nosy_InitGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var k keeper.Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var gs types.GenesisState
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}

		InitGenesis(ctx, k, gs)
	})
}
