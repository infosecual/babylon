package genhelpers

import (
	"testing"

	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
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

// skipping Fuzz_Nosy_CollectTxs__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.JSONCodec

func Fuzz_Nosy_PathsNodeCfg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homeDir string) {
		PathsNodeCfg(homeDir)
	})
}

func Fuzz_Nosy_makeOutputFilepath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rootDir string, nodeID string) {
		makeOutputFilepath(rootDir, nodeID)
	})
}

func Fuzz_Nosy_mapFinalityProvidersByBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fps []*types.FinalityProvider
		fill_err = tp.Fill(&fps)
		if fill_err != nil {
			return
		}

		mapFinalityProvidersByBtcPk(fps)
	})
}
