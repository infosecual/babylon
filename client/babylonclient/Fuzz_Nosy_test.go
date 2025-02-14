package babylonclient

import (
	"context"
	"testing"
	"time"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cosmos/cosmos-sdk/client"

	//"github.com/babylonlabs-io/babylon/client"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/bytes"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	tx1 "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//tx3 "github.com/cosmos/cosmos-sdk/types"
	tx2 "github.com/cosmos/cosmos-sdk/types/tx"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"google.golang.org/grpc/metadata"
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

func Fuzz_Nosy_CosmosProvider_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.Address()
	})
}

func Fuzz_Nosy_CosmosProvider_AdjustEstimatedGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var gasUsed uint64
		fill_err = tp.Fill(&gasUsed)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.AdjustEstimatedGas(gasUsed)
	})
}

// skipping Fuzz_Nosy_CosmosProvider_CalculateGas__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

func Fuzz_Nosy_CosmosProvider_ChainId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.ChainId()
	})
}

func Fuzz_Nosy_CosmosProvider_ChainName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.ChainName()
	})
}

func Fuzz_Nosy_CosmosProvider_DecodeBech32AccAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.DecodeBech32AccAddr(addr)
	})
}

func Fuzz_Nosy_CosmosProvider_EncodeBech32AccAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.EncodeBech32AccAddr(addr)
	})
}

func Fuzz_Nosy_CosmosProvider_EnsureExists__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.EnsureExists(clientCtx, addr)
	})
}

func Fuzz_Nosy_CosmosProvider_GetAccount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.GetAccount(clientCtx, addr)
	})
}

func Fuzz_Nosy_CosmosProvider_GetAccountNumberSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var clientCtx client.Context
		fill_err = tp.Fill(&clientCtx)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.GetAccountNumberSequence(clientCtx, addr)
	})
}

func Fuzz_Nosy_CosmosProvider_GetAccountWithHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var _x2 client.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.GetAccountWithHeight(_x2, addr)
	})
}

func Fuzz_Nosy_CosmosProvider_GetKeyAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.GetKeyAddress(key)
	})
}

func Fuzz_Nosy_CosmosProvider_GetKeyAddressForKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var key string
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.GetKeyAddressForKey(key)
	})
}

func Fuzz_Nosy_CosmosProvider_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.Init()
	})
}

// skipping Fuzz_Nosy_CosmosProvider_Invoke__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_CosmosProvider_Key__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.Key()
	})
}

func Fuzz_Nosy_CosmosProvider_MustEncodeAccAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.MustEncodeAccAddr(addr)
	})
}

// skipping Fuzz_Nosy_CosmosProvider_NewStream__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

func Fuzz_Nosy_CosmosProvider_PrepareFactory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var txf tx1.Factory
		fill_err = tp.Fill(&txf)
		if fill_err != nil {
			return
		}
		var signingKey string
		fill_err = tp.Fill(&signingKey)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.PrepareFactory(txf, signingKey)
	})
}

func Fuzz_Nosy_CosmosProvider_ProviderConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.ProviderConfig()
	})
}

func Fuzz_Nosy_CosmosProvider_QueryABCI__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req abci.RequestQuery
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.QueryABCI(ctx, req)
	})
}

// skipping Fuzz_Nosy_CosmosProvider_RunGRPCQuery__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_CosmosProvider_SendMessagesToMempool__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

func Fuzz_Nosy_CosmosProvider_SetRpcAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var rpcAddr string
		fill_err = tp.Fill(&rpcAddr)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.SetRpcAddr(rpcAddr)
	})
}

func Fuzz_Nosy_CosmosProvider_SetSDKContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.SetSDKContext()
	})
}

func Fuzz_Nosy_CosmosProvider_Timeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.Timeout()
	})
}

func Fuzz_Nosy_CosmosProvider_TxFactory__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.TxFactory()
	})
}

func Fuzz_Nosy_CosmosProvider_TxServiceBroadcast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *tx2.BroadcastTxRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if cc == nil || req == nil {
			return
		}

		cc.TxServiceBroadcast(ctx, req)
	})
}

// skipping Fuzz_Nosy_CosmosProvider_broadcastTx__ because parameters include func, chan, or unsupported interface: []func(*github.com/babylonlabs-io/babylon/client/babylonclient.RelayerTxResponse, error)

// skipping Fuzz_Nosy_CosmosProvider_buildMessages__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

// skipping Fuzz_Nosy_CosmosProvider_handleAccountSequenceMismatchError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_CosmosProvider_mkTxResult__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var resTx *coretypes.ResultTx
		fill_err = tp.Fill(&resTx)
		if fill_err != nil {
			return
		}
		if cc == nil || resTx == nil {
			return
		}

		cc.mkTxResult(resTx)
	})
}

func Fuzz_Nosy_CosmosProvider_sdkError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var codeSpace string
		fill_err = tp.Fill(&codeSpace)
		if fill_err != nil {
			return
		}
		var code uint32
		fill_err = tp.Fill(&code)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.sdkError(codeSpace, code)
	})
}

func Fuzz_Nosy_CosmosProvider_updateNextAccountSequence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var sequenceGuard *WalletState
		fill_err = tp.Fill(&sequenceGuard)
		if fill_err != nil {
			return
		}
		var seq uint64
		fill_err = tp.Fill(&seq)
		if fill_err != nil {
			return
		}
		if cc == nil || sequenceGuard == nil {
			return
		}

		cc.updateNextAccountSequence(sequenceGuard, seq)
	})
}

func Fuzz_Nosy_CosmosProvider_waitForBlockInclusion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cc *CosmosProvider
		fill_err = tp.Fill(&cc)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var txHash []byte
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var waitTimeout time.Duration
		fill_err = tp.Fill(&waitTimeout)
		if fill_err != nil {
			return
		}
		if cc == nil {
			return
		}

		cc.waitForBlockInclusion(ctx, txHash, waitTimeout)
	})
}

// skipping Fuzz_Nosy_CosmosProvider_waitForTx__ because parameters include func, chan, or unsupported interface: []func(*github.com/babylonlabs-io/babylon/client/babylonclient.RelayerTxResponse, error)

// skipping Fuzz_Nosy_ChainProvider_Address__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_ChainId__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_ChainName__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_Init__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_Key__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_ProviderConfig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_SendMessagesToMempool__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_SetRpcAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

// skipping Fuzz_Nosy_ChainProvider_Timeout__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ChainProvider

func Fuzz_Nosy_CosmosMessage_MsgBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm CosmosMessage
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}

		cm.MsgBytes()
	})
}

func Fuzz_Nosy_CosmosMessage_Type__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm CosmosMessage
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}

		cm.Type()
	})
}

func Fuzz_Nosy_CosmosProviderConfig_BroadcastMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc CosmosProviderConfig
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}

		pc.BroadcastMode()
	})
}

func Fuzz_Nosy_CosmosProviderConfig_NewProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc CosmosProviderConfig
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}
		var homepath string
		fill_err = tp.Fill(&homepath)
		if fill_err != nil {
			return
		}
		var chainName string
		fill_err = tp.Fill(&chainName)
		if fill_err != nil {
			return
		}

		pc.NewProvider(homepath, chainName)
	})
}

func Fuzz_Nosy_CosmosProviderConfig_SignMode__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc CosmosProviderConfig
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}

		pc.SignMode()
	})
}

func Fuzz_Nosy_CosmosProviderConfig_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pc CosmosProviderConfig
		fill_err = tp.Fill(&pc)
		if fill_err != nil {
			return
		}

		pc.Validate()
	})
}

// skipping Fuzz_Nosy_ProviderConfig_BroadcastMode__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ProviderConfig

// skipping Fuzz_Nosy_ProviderConfig_NewProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ProviderConfig

// skipping Fuzz_Nosy_ProviderConfig_Validate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.ProviderConfig

func Fuzz_Nosy_RPCClient_ABCIInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		r.ABCIInfo(ctx)
	})
}

func Fuzz_Nosy_RPCClient_ABCIQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var d4 bytes.HexBytes
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}

		r.ABCIQuery(ctx, path, d4)
	})
}

func Fuzz_Nosy_RPCClient_ABCIQueryWithOptions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var d4 bytes.HexBytes
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}
		var opts rpcclient.ABCIQueryOptions
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}

		r.ABCIQueryWithOptions(ctx, path, d4, opts)
	})
}

func Fuzz_Nosy_RPCClient_Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height *int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if height == nil {
			return
		}

		r.Block(ctx, height)
	})
}

func Fuzz_Nosy_RPCClient_BlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}

		r.BlockByHash(ctx, hash)
	})
}

func Fuzz_Nosy_RPCClient_BlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height *int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if height == nil {
			return
		}

		r.BlockResults(ctx, height)
	})
}

func Fuzz_Nosy_RPCClient_BlockSearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var page *int
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		var perPage *int
		fill_err = tp.Fill(&perPage)
		if fill_err != nil {
			return
		}
		var orderBy string
		fill_err = tp.Fill(&orderBy)
		if fill_err != nil {
			return
		}
		if page == nil || perPage == nil {
			return
		}

		r.BlockSearch(ctx, query, page, perPage, orderBy)
	})
}

func Fuzz_Nosy_RPCClient_BlockchainInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var minHeight int64
		fill_err = tp.Fill(&minHeight)
		if fill_err != nil {
			return
		}
		var maxHeight int64
		fill_err = tp.Fill(&maxHeight)
		if fill_err != nil {
			return
		}

		r.BlockchainInfo(ctx, minHeight, maxHeight)
	})
}

func Fuzz_Nosy_RPCClient_BroadcastTxAsync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx3 tmtypes.Tx
		fill_err = tp.Fill(&tx3)
		if fill_err != nil {
			return
		}

		r.BroadcastTxAsync(ctx, tx3)
	})
}

func Fuzz_Nosy_RPCClient_BroadcastTxCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx4 tmtypes.Tx
		fill_err = tp.Fill(&tx4)
		if fill_err != nil {
			return
		}

		r.BroadcastTxCommit(ctx, tx4)
	})
}

func Fuzz_Nosy_RPCClient_BroadcastTxSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var tx5 tmtypes.Tx
		fill_err = tp.Fill(&tx5)
		if fill_err != nil {
			return
		}

		r.BroadcastTxSync(ctx, tx5)
	})
}

func Fuzz_Nosy_RPCClient_Commit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height *int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if height == nil {
			return
		}

		r.Commit(ctx, height)
	})
}

func Fuzz_Nosy_RPCClient_Status__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		r.Status(ctx)
	})
}

func Fuzz_Nosy_RPCClient_Tx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		var prove bool
		fill_err = tp.Fill(&prove)
		if fill_err != nil {
			return
		}

		r.Tx(ctx, hash, prove)
	})
}

func Fuzz_Nosy_RPCClient_TxSearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var prove bool
		fill_err = tp.Fill(&prove)
		if fill_err != nil {
			return
		}
		var page *int
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		var perPage *int
		fill_err = tp.Fill(&perPage)
		if fill_err != nil {
			return
		}
		var orderBy string
		fill_err = tp.Fill(&orderBy)
		if fill_err != nil {
			return
		}
		if page == nil || perPage == nil {
			return
		}

		r.TxSearch(ctx, query, prove, page, perPage, orderBy)
	})
}

func Fuzz_Nosy_RPCClient_Validators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r RPCClient
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var height *int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var page *int
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		var perPage *int
		fill_err = tp.Fill(&perPage)
		if fill_err != nil {
			return
		}
		if height == nil || page == nil || perPage == nil {
			return
		}

		r.Validators(ctx, height, page, perPage)
	})
}

// skipping Fuzz_Nosy_RelayerMessage_MsgBytes__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

// skipping Fuzz_Nosy_RelayerMessage_Type__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

func Fuzz_Nosy__err_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e _err
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}

		e.Error()
	})
}

// skipping Fuzz_Nosy_intoAny_AsAny__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.intoAny

// skipping Fuzz_Nosy_protoTxProvider_GetProtoTx__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/babylonclient.protoTxProvider

// skipping Fuzz_Nosy_BuildSimTx__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_CosmosMsgs__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

func Fuzz_Nosy_GetHeightFromMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var md metadata.MD
		fill_err = tp.Fill(&md)
		if fill_err != nil {
			return
		}

		GetHeightFromMetadata(md)
	})
}

func Fuzz_Nosy_GetProveFromMetadata__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var md metadata.MD
		fill_err = tp.Fill(&md)
		if fill_err != nil {
			return
		}

		GetProveFromMetadata(md)
	})
}

func Fuzz_Nosy_SetSDKConfigContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefix string) {
		SetSDKConfigContext(prefix)
	})
}

func Fuzz_Nosy_isQueryStoreWithProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, path string) {
		isQueryStoreWithProof(path)
	})
}

func Fuzz_Nosy_keysDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, home string, chainID string) {
		keysDir(home, chainID)
	})
}

func Fuzz_Nosy_parseEventsFromTxResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var resp *sdk.TxResponse
		fill_err = tp.Fill(&resp)
		if fill_err != nil {
			return
		}
		if resp == nil {
			return
		}

		parseEventsFromTxResponse(resp)
	})
}
