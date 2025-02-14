package keepers

import (
	"testing"

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

func Fuzz_Nosy_AppKeepers_GetKVStoreKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetKVStoreKeys()
	})
}

func Fuzz_Nosy_AppKeepers_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		var storeKey string
		fill_err = tp.Fill(&storeKey)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetKey(storeKey)
	})
}

func Fuzz_Nosy_AppKeepers_GetMemKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		var storeKey string
		fill_err = tp.Fill(&storeKey)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetMemKey(storeKey)
	})
}

func Fuzz_Nosy_AppKeepers_GetMemoryStoreKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetMemoryStoreKeys()
	})
}

func Fuzz_Nosy_AppKeepers_GetSubspace__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		var moduleName string
		fill_err = tp.Fill(&moduleName)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetSubspace(moduleName)
	})
}

func Fuzz_Nosy_AppKeepers_GetTKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		var storeKey string
		fill_err = tp.Fill(&storeKey)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetTKey(storeKey)
	})
}

func Fuzz_Nosy_AppKeepers_GetTransientStoreKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var appKeepers *AppKeepers
		fill_err = tp.Fill(&appKeepers)
		if fill_err != nil {
			return
		}
		if appKeepers == nil {
			return
		}

		appKeepers.GetTransientStoreKeys()
	})
}

// skipping Fuzz_Nosy_AppKeepers_InitKeepers__ because parameters include func, chan, or unsupported interface: cosmossdk.io/log.Logger
