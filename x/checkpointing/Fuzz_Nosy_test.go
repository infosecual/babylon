package checkpointing

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon/x/checkpointing/keeper"
	"github.com/babylonlabs-io/babylon/x/checkpointing/types"
	ckpttypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
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

func Fuzz_Nosy_ProposalHandler_ExtractInjectedCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.ExtractInjectedCheckpoint(txs)
	})
}

func Fuzz_Nosy_ProposalHandler_PreBlocker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PreBlocker()
	})
}

func Fuzz_Nosy_ProposalHandler_PrepareProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.PrepareProposal()
	})
}

func Fuzz_Nosy_ProposalHandler_ProcessProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.ProcessProposal()
	})
}

func Fuzz_Nosy_ProposalHandler_SetHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var bApp *baseapp.BaseApp
		fill_err = tp.Fill(&bApp)
		if fill_err != nil {
			return
		}
		if h == nil || bApp == nil {
			return
		}

		h.SetHandlers(bApp)
	})
}

func Fuzz_Nosy_ProposalHandler_buildCheckpointFromVoteExtensions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		var extendedVotes []abci.ExtendedVoteInfo
		fill_err = tp.Fill(&extendedVotes)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.buildCheckpointFromVoteExtensions(ctx, epoch, extendedVotes)
	})
}

func Fuzz_Nosy_ProposalHandler_buildInjectedTxBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var injectedCkpt *ckpttypes.MsgInjectedCheckpoint
		fill_err = tp.Fill(&injectedCkpt)
		if fill_err != nil {
			return
		}
		if h == nil || injectedCkpt == nil {
			return
		}

		h.buildInjectedTxBytes(injectedCkpt)
	})
}

func Fuzz_Nosy_ProposalHandler_findLastBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var extendedVotes []abci.ExtendedVoteInfo
		fill_err = tp.Fill(&extendedVotes)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.findLastBlockHash(extendedVotes)
	})
}

func Fuzz_Nosy_ProposalHandler_getValidBlsSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *ProposalHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var extendedVotes []abci.ExtendedVoteInfo
		fill_err = tp.Fill(&extendedVotes)
		if fill_err != nil {
			return
		}
		var blockHash []byte
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.getValidBlsSigs(ctx, extendedVotes, blockHash)
	})
}

func Fuzz_Nosy_VoteExtensionHandler_ExtendVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *VoteExtensionHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.ExtendVote()
	})
}

func Fuzz_Nosy_VoteExtensionHandler_SetHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *VoteExtensionHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var bApp *baseapp.BaseApp
		fill_err = tp.Fill(&bApp)
		if fill_err != nil {
			return
		}
		if h == nil || bApp == nil {
			return
		}

		h.SetHandlers(bApp)
	})
}

func Fuzz_Nosy_VoteExtensionHandler_VerifyVoteExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *VoteExtensionHandler
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.VerifyVoteExtension()
	})
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
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		am.EndBlock(_x2)
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

func Fuzz_Nosy_AppModule_Name__(f *testing.F) {
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

		am.Name()
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

func Fuzz_Nosy_AppModuleBasic_RegisterCodec__(f *testing.F) {
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

		_x1.RegisterCodec(cdc)
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

func Fuzz_Nosy_AppModuleBasic_RegisterRESTRoutes__(f *testing.F) {
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
		var rtr *mux.Router
		fill_err = tp.Fill(&rtr)
		if fill_err != nil {
			return
		}
		if rtr == nil {
			return
		}

		_x1.RegisterRESTRoutes(clientCtx, rtr)
	})
}

// skipping Fuzz_Nosy_AppModuleBasic_ValidateGenesis__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

// skipping Fuzz_Nosy_CheckpointingKeeper_GetBlsPubKey__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_GetEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_GetTotalVotingPower__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_GetValidatorSet__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_SealCheckpoint__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

// skipping Fuzz_Nosy_CheckpointingKeeper_VerifyBLSSig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing.CheckpointingKeeper

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
		var genState types.GenesisState
		fill_err = tp.Fill(&genState)
		if fill_err != nil {
			return
		}

		InitGenesis(ctx, k, genState)
	})
}

func Fuzz_Nosy_removeInjectedTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs [][]byte
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		removeInjectedTx(txs)
	})
}
