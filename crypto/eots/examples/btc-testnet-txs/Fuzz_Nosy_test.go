package main

import (
	"testing"

	"github.com/babylonlabs-io/babylon/crypto/eots"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
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

func Fuzz_Nosy_generateSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var secretKey *eots.PrivateKey
		fill_err = tp.Fill(&secretKey)
		if fill_err != nil {
			return
		}
		var privateRand *eots.PrivateRand
		fill_err = tp.Fill(&privateRand)
		if fill_err != nil {
			return
		}
		var message1 string
		fill_err = tp.Fill(&message1)
		if fill_err != nil {
			return
		}
		var message2 string
		fill_err = tp.Fill(&message2)
		if fill_err != nil {
			return
		}
		if secretKey == nil || privateRand == nil {
			return
		}

		generateSignatures(secretKey, privateRand, message1, message2)
	})
}

func Fuzz_Nosy_generateTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var senderPk *eots.PrivateKey
		fill_err = tp.Fill(&senderPk)
		if fill_err != nil {
			return
		}
		var amount int64
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var fees int64
		fill_err = tp.Fill(&fees)
		if fill_err != nil {
			return
		}
		var rcvAddress *btcutil.AddressPubKeyHash
		fill_err = tp.Fill(&rcvAddress)
		if fill_err != nil {
			return
		}
		var unspentTxId *chainhash.Hash
		fill_err = tp.Fill(&unspentTxId)
		if fill_err != nil {
			return
		}
		var outIndex uint32
		fill_err = tp.Fill(&outIndex)
		if fill_err != nil {
			return
		}
		if senderPk == nil || rcvAddress == nil || unspentTxId == nil {
			return
		}

		generateTx(senderPk, amount, fees, rcvAddress, unspentTxId, outIndex)
	})
}
