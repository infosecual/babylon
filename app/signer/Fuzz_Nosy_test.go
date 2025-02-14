package signer

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

func Fuzz_Nosy_BlsKey_BlsPubKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *BlsKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.BlsPubKey()
	})
}

func Fuzz_Nosy_BlsKey_Save__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *BlsKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var password string
		fill_err = tp.Fill(&password)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.Save(password)
	})
}

func Fuzz_Nosy_BlsKey_SignMsgWithBls__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *BlsKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		k.SignMsgWithBls(msg)
	})
}

func Fuzz_Nosy_DefaultBlsKeyFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, home string) {
		DefaultBlsKeyFile(home)
	})
}

func Fuzz_Nosy_DefaultBlsPasswordFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, home string) {
		DefaultBlsPasswordFile(home)
	})
}

// skipping Fuzz_Nosy_ExportGenBls__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/crypto.PrivKey
