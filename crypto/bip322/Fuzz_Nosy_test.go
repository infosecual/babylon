package bip322

import (
	"io"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
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

func Fuzz_Nosy_GetBIP340TaggedHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, msg []byte) {
		GetBIP340TaggedHash(msg)
	})
}

func Fuzz_Nosy_SerializeWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w wire.TxWitness
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		SerializeWitness(w)
	})
}

func Fuzz_Nosy_SignWithP2TrSpendAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var privKey *btcec.PrivateKey
		fill_err = tp.Fill(&privKey)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if privKey == nil || net == nil {
			return
		}

		SignWithP2TrSpendAddress(msg, privKey, net)
	})
}

func Fuzz_Nosy_SignWithP2WPKHAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var privKey *btcec.PrivateKey
		fill_err = tp.Fill(&privKey)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if privKey == nil || net == nil {
			return
		}

		SignWithP2WPKHAddress(msg, privKey, net)
	})
}

func Fuzz_Nosy_SimpleSigToWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sig []byte) {
		SimpleSigToWitness(sig)
	})
}

func Fuzz_Nosy_readScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r io.Reader
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var pver uint32
		fill_err = tp.Fill(&pver)
		if fill_err != nil {
			return
		}
		var maxAllowed uint32
		fill_err = tp.Fill(&maxAllowed)
		if fill_err != nil {
			return
		}
		var fieldName string
		fill_err = tp.Fill(&fieldName)
		if fill_err != nil {
			return
		}

		readScript(r, pver, maxAllowed, fieldName)
	})
}

func Fuzz_Nosy_toSpendSignatureScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, msg []byte) {
		toSpendSignatureScript(msg)
	})
}
