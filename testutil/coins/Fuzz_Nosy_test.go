package coins

import (
	"testing"

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

func Fuzz_Nosy_CoinsDiffInPointOnePercentMargin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 sdk.Coins
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var c2 sdk.Coins
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}

		CoinsDiffInPointOnePercentMargin(c1, c2)
	})
}

func Fuzz_Nosy_RequireCoinsDiffInPointOnePercentMargin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var c1 sdk.Coins
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var c2 sdk.Coins
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		RequireCoinsDiffInPointOnePercentMargin(t1, c1, c2)
	})
}
