package client

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/babylon/client/config"
	btcctypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	"github.com/babylonlabs-io/babylon/x/btclightclient/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.uber.org/zap"
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

// skipping Fuzz_Nosy_Client_CalculateGas__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/crypto/types.PubKey

func Fuzz_Nosy_Client_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.GetAddr()
	})
}

func Fuzz_Nosy_Client_GetConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.GetConfig()
	})
}

func Fuzz_Nosy_Client_GetKeyring__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.GetKeyring()
	})
}

func Fuzz_Nosy_Client_InsertBTCSpvProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *btcctypes.MsgInsertBTCSpvProof
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || msg == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.InsertBTCSpvProof(ctx, msg)
	})
}

func Fuzz_Nosy_Client_InsertHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgInsertHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil || msg == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.InsertHeaders(ctx, msg)
	})
}

func Fuzz_Nosy_Client_MustGetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.MustGetAddr()
	})
}

func Fuzz_Nosy_Client_Provider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var logger *zap.Logger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if cfg == nil || logger == nil {
			return
		}

		c, err := New(cfg, logger)
		if err != nil {
			return
		}
		c.Provider()
	})
}

// skipping Fuzz_Nosy_Client_ReliablySendMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Client_ReliablySendMsgs__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Client_ReliablySendMsgsWithSigner__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Client_SendMessageWithSigner__ because parameters include func, chan, or unsupported interface: []github.com/babylonlabs-io/babylon/client/babylonclient.RelayerMessage

// skipping Fuzz_Nosy_Client_SendMsgToMempool__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Client_SendMsgsToMempool__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Client_accessKeyWithLock__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_protoTxProvider_GetProtoTx__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/client/client.protoTxProvider

// skipping Fuzz_Nosy_BuildSimTx__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/crypto/types.PubKey

// skipping Fuzz_Nosy_ToProviderMsgs__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_errorContained__ because parameters include func, chan, or unsupported interface: error
