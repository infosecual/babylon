package keeper

import (
	"testing"

	"github.com/btcsuite/btcd/wire"
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

// skipping Fuzz_Nosy_MockIncentiveKeeper_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_BTCLightClientKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_BTCLightClientKeeperWithCustomParams__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_BTCStakingKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_BTCStakingKeeperWithStore__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_CheckpointingKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_EpochingKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_FinalityKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_FinalityKeeperWithStore__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_IncentiveKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_IncentiveKeeperWithStore__ because parameters include func, chan, or unsupported interface: testing.TB

// skipping Fuzz_Nosy_NewBTCCheckpointKeeper__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_NewBTCHeaderBytesList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain []*wire.BlockHeader
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}

		NewBTCHeaderBytesList(chain)
	})
}
