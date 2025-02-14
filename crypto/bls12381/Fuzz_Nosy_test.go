package bls12381

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

func Fuzz_Nosy_PublicKey_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.Unmarshal(d2)
	})
}

func Fuzz_Nosy_Signature_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.Unmarshal(d3)
	})
}

func Fuzz_Nosy_PrivateKey_PubKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, seed []byte) {
		sk := genPrivKeyFromSeed(seed)
		sk.PubKey()
	})
}

func Fuzz_Nosy_PublicKey_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.Bytes()
	})
}

func Fuzz_Nosy_PublicKey_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}
		var k PublicKey
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.Equal(k)
	})
}

func Fuzz_Nosy_PublicKey_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.Marshal()
	})
}

func Fuzz_Nosy_PublicKey_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.MarshalTo(d2)
	})
}

func Fuzz_Nosy_PublicKey_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.MustMarshal()
	})
}

func Fuzz_Nosy_PublicKey_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		pk, err := AggrPKList(pks)
		if err != nil {
			return
		}
		pk.Size()
	})
}

func Fuzz_Nosy_Signature_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.Bytes()
	})
}

func Fuzz_Nosy_Signature_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var s Signature
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.Equal(s)
	})
}

func Fuzz_Nosy_Signature_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.Marshal()
	})
}

func Fuzz_Nosy_Signature_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var d3 []byte
		fill_err = tp.Fill(&d3)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.MarshalTo(d3)
	})
}

func Fuzz_Nosy_Signature_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.MustMarshal()
	})
}

func Fuzz_Nosy_Signature_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.Size()
	})
}

func Fuzz_Nosy_Signature_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.String()
	})
}

func Fuzz_Nosy_Signature_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sk PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		sig := Sign(sk, msg)
		sig.ValidateBasic()
	})
}

func Fuzz_Nosy_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sig Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var pk PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		Verify(sig, pk, msg)
	})
}

func Fuzz_Nosy_VerifyMultiSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sig Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var pks []PublicKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		VerifyMultiSig(sig, pks, msg)
	})
}
