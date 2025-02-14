package eots

import (
	"io"
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

func Fuzz_Nosy_Extract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubKey *PublicKey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var r *PublicRand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var message1 []byte
		fill_err = tp.Fill(&message1)
		if fill_err != nil {
			return
		}
		var sig1 *Signature
		fill_err = tp.Fill(&sig1)
		if fill_err != nil {
			return
		}
		var message2 []byte
		fill_err = tp.Fill(&message2)
		if fill_err != nil {
			return
		}
		var sig2 *Signature
		fill_err = tp.Fill(&sig2)
		if fill_err != nil {
			return
		}
		if pubKey == nil || r == nil || sig1 == nil || sig2 == nil {
			return
		}

		Extract(pubKey, r, message1, sig1, message2, sig2)
	})
}

func Fuzz_Nosy_KeyGen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var randSource io.Reader
		fill_err = tp.Fill(&randSource)
		if fill_err != nil {
			return
		}

		KeyGen(randSource)
	})
}

func Fuzz_Nosy_PubGen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k *PrivateKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if k == nil {
			return
		}

		PubGen(k)
	})
}

func Fuzz_Nosy_RandGen__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var randSource io.Reader
		fill_err = tp.Fill(&randSource)
		if fill_err != nil {
			return
		}

		RandGen(randSource)
	})
}

func Fuzz_Nosy_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk *PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var privateRand *PrivateRand
		fill_err = tp.Fill(&privateRand)
		if fill_err != nil {
			return
		}
		var message []byte
		fill_err = tp.Fill(&message)
		if fill_err != nil {
			return
		}
		if sk == nil || privateRand == nil {
			return
		}

		Sign(sk, privateRand, message)
	})
}

func Fuzz_Nosy_extractFromHashes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubKey *PublicKey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var r *PublicRand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var hash1 [32]byte
		fill_err = tp.Fill(&hash1)
		if fill_err != nil {
			return
		}
		var sig1 *Signature
		fill_err = tp.Fill(&sig1)
		if fill_err != nil {
			return
		}
		var hash2 [32]byte
		fill_err = tp.Fill(&hash2)
		if fill_err != nil {
			return
		}
		var sig2 *Signature
		fill_err = tp.Fill(&sig2)
		if fill_err != nil {
			return
		}
		if pubKey == nil || r == nil || sig1 == nil || sig2 == nil {
			return
		}

		extractFromHashes(pubKey, r, hash1, sig1, hash2, sig2)
	})
}

func Fuzz_Nosy_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, message []byte) {
		hash(message)
	})
}

func Fuzz_Nosy_signHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk *PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var privateRand *PrivateRand
		fill_err = tp.Fill(&privateRand)
		if fill_err != nil {
			return
		}
		var hash [32]byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if sk == nil || privateRand == nil {
			return
		}

		signHash(sk, privateRand, hash)
	})
}

func Fuzz_Nosy_signatureError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kind KindOfError
		fill_err = tp.Fill(&kind)
		if fill_err != nil {
			return
		}
		var desc string
		fill_err = tp.Fill(&desc)
		if fill_err != nil {
			return
		}

		signatureError(kind, desc)
	})
}
