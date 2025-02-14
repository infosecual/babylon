package cmd

import (
	"testing"

	"cosmossdk.io/client/v2/autocli"
	"github.com/spf13/cobra"
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

func Fuzz_Nosy_EnhanceRootCommandWithoutTxStaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var autoCliOpts autocli.AppOptions
		fill_err = tp.Fill(&autoCliOpts)
		if fill_err != nil {
			return
		}
		var rootCmd *cobra.Command
		fill_err = tp.Fill(&rootCmd)
		if fill_err != nil {
			return
		}
		if rootCmd == nil {
			return
		}

		EnhanceRootCommandWithoutTxStaking(autoCliOpts, rootCmd)
	})
}

// skipping Fuzz_Nosy_PrintDBStats__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-db.DB

func Fuzz_Nosy_addGenesisFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		if c1 == nil {
			return
		}

		addGenesisFlags(c1)
	})
}

func Fuzz_Nosy_addModuleInitFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startCmd *cobra.Command
		fill_err = tp.Fill(&startCmd)
		if fill_err != nil {
			return
		}
		if startCmd == nil {
			return
		}

		addModuleInitFlags(startCmd)
	})
}

// skipping Fuzz_Nosy_automigrate_e2e_upgrade__ because parameters include func, chan, or unsupported interface: cosmossdk.io/log.Logger

func Fuzz_Nosy_blsPassword__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 *cobra.Command
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		if c1 == nil {
			return
		}

		blsPassword(c1)
	})
}

func Fuzz_Nosy_bytesToPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bz []byte, keytype string) {
		bytesToPubkey(bz, keytype)
	})
}

func Fuzz_Nosy_calculateIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, ip string, i int) {
		calculateIP(ip, i)
	})
}

func Fuzz_Nosy_extractStoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fullKey string) {
		extractStoreKey(fullKey)
	})
}

func Fuzz_Nosy_getIP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int, startingIPAddr string) {
		getIP(i, startingIPAddr)
	})
}

// skipping Fuzz_Nosy_initRootCmd__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/client.TxEncodingConfig

func Fuzz_Nosy_printModuleStats__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stats map[string]*ModuleStats
		fill_err = tp.Fill(&stats)
		if fill_err != nil {
			return
		}
		var gs *GlobalStats
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}
		if gs == nil {
			return
		}

		printModuleStats(stats, gs)
	})
}
