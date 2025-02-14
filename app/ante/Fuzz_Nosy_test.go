package ante

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

// skipping Fuzz_Nosy_BtcValidationDecorator_AnteHandle__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Tx

// skipping Fuzz_Nosy_WrappedAnteHandler_AnteHandle__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Tx

// skipping Fuzz_Nosy_CheckTxFeeWithGlobalMinGasPrices__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Tx

func Fuzz_Nosy_getTxPriority__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fee sdk.Coins
		fill_err = tp.Fill(&fee)
		if fill_err != nil {
			return
		}
		var gas int64
		fill_err = tp.Fill(&gas)
		if fill_err != nil {
			return
		}

		getTxPriority(fee, gas)
	})
}
