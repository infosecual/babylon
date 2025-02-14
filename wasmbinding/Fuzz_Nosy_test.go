package wasmbinding

import (
	"testing"

	lcKeeper "github.com/babylonlabs-io/babylon/x/btclightclient/keeper"
	checkpointingkeeper "github.com/babylonlabs-io/babylon/x/checkpointing/keeper"
	epochingkeeper "github.com/babylonlabs-io/babylon/x/epoching/keeper"
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

func Fuzz_Nosy_CustomQuerier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var qp *QueryPlugin
		fill_err = tp.Fill(&qp)
		if fill_err != nil {
			return
		}
		if qp == nil {
			return
		}

		CustomQuerier(qp)
	})
}

func Fuzz_Nosy_RegisterCustomPlugins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ek *epochingkeeper.Keeper
		fill_err = tp.Fill(&ek)
		if fill_err != nil {
			return
		}
		var ck *checkpointingkeeper.Keeper
		fill_err = tp.Fill(&ck)
		if fill_err != nil {
			return
		}
		var lcKeeper *lcKeeper.Keeper
		fill_err = tp.Fill(&lcKeeper)
		if fill_err != nil {
			return
		}
		if ek == nil || ck == nil || lcKeeper == nil {
			return
		}

		RegisterCustomPlugins(ek, ck, lcKeeper)
	})
}
