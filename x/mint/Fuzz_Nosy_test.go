package mint

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
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

// skipping Fuzz_Nosy_AppModule_ExportGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_AppModule_GenerateGenesisState__(f *testing.F) {
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
		var _x2 *module.SimulationState
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		_x1.GenerateGenesisState(_x2)
	})
}

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

func Fuzz_Nosy_AppModule_Name__(f *testing.F) {
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

		_x1.Name()
	})
}

func Fuzz_Nosy_AppModule_ProposalContents__(f *testing.F) {
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
		var _x2 module.SimulationState
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.ProposalContents(_x2)
	})
}

func Fuzz_Nosy_AppModule_QuerierRoute__(f *testing.F) {
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

		_x1.QuerierRoute()
	})
}

// skipping Fuzz_Nosy_AppModule_RegisterInvariants__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.InvariantRegistry

// skipping Fuzz_Nosy_AppModule_RegisterServices__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types/module.Configurator

func Fuzz_Nosy_AppModule_WeightedOperations__(f *testing.F) {
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
		var _x2 module.SimulationState
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		_x1.WeightedOperations(_x2)
	})
}

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
		var _x1 AppModuleBasic
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.GetTxCmd()
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
		var _x2 *codec.LegacyAmino
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if _x2 == nil {
			return
		}

		_x1.RegisterLegacyAminoCodec(_x2)
	})
}

// skipping Fuzz_Nosy_AppModuleBasic_ValidateGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_BeginBlocker__(f *testing.F) {
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

		BeginBlocker(ctx, k)
	})
}

func Fuzz_Nosy_maybeUpdateMinter__(f *testing.F) {
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

		maybeUpdateMinter(ctx, k)
	})
}

func Fuzz_Nosy_mintBlockProvision__(f *testing.F) {
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

		mintBlockProvision(ctx, k)
	})
}

func Fuzz_Nosy_setPreviousBlockTime__(f *testing.F) {
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

		setPreviousBlockTime(ctx, k)
	})
}
