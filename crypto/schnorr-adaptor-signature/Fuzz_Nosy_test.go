package schnorr_adaptor_signature

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
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

func Fuzz_Nosy_AdaptorSignature_Decrypt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var asigHex string
		fill_err = tp.Fill(&asigHex)
		if fill_err != nil {
			return
		}
		var decKey *DecryptionKey
		fill_err = tp.Fill(&decKey)
		if fill_err != nil {
			return
		}
		if decKey == nil {
			return
		}

		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Decrypt(decKey)
	})
}

func Fuzz_Nosy_AdaptorSignature_EncVerify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var asigHex string
		fill_err = tp.Fill(&asigHex)
		if fill_err != nil {
			return
		}
		var pk *btcec.PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var encKey *EncryptionKey
		fill_err = tp.Fill(&encKey)
		if fill_err != nil {
			return
		}
		var msgHash []byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}
		if pk == nil || encKey == nil {
			return
		}

		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.EncVerify(pk, encKey, msgHash)
	})
}

func Fuzz_Nosy_AdaptorSignature_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var asigHex string
		fill_err = tp.Fill(&asigHex)
		if fill_err != nil {
			return
		}
		var sig2 AdaptorSignature
		fill_err = tp.Fill(&sig2)
		if fill_err != nil {
			return
		}

		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Equals(sig2)
	})
}

func Fuzz_Nosy_AdaptorSignature_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Marshal()
	})
}

func Fuzz_Nosy_AdaptorSignature_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.MarshalHex()
	})
}

func Fuzz_Nosy_AdaptorSignature_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string, d2 []byte) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.MarshalTo(d2)
	})
}

func Fuzz_Nosy_AdaptorSignature_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.MustMarshal()
	})
}

func Fuzz_Nosy_AdaptorSignature_Recover__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var asigHex string
		fill_err = tp.Fill(&asigHex)
		if fill_err != nil {
			return
		}
		var decryptedSchnorrSig *schnorr.Signature
		fill_err = tp.Fill(&decryptedSchnorrSig)
		if fill_err != nil {
			return
		}
		if decryptedSchnorrSig == nil {
			return
		}

		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Recover(decryptedSchnorrSig)
	})
}

func Fuzz_Nosy_AdaptorSignature_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Size()
	})
}

func Fuzz_Nosy_AdaptorSignature_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, asigHex string, d2 []byte) {
		sig, err := NewAdaptorSignatureFromHex(asigHex)
		if err != nil {
			return
		}
		sig.Unmarshal(d2)
	})
}

func Fuzz_Nosy_DecryptionKey_GetEncKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		if btcSK == nil {
			return
		}

		dk, err := NewDecyptionKeyFromBTCSK(btcSK)
		if err != nil {
			return
		}
		dk.GetEncKey()
	})
}

func Fuzz_Nosy_DecryptionKey_ToBTCSK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		if btcSK == nil {
			return
		}

		dk, err := NewDecyptionKeyFromBTCSK(btcSK)
		if err != nil {
			return
		}
		dk.ToBTCSK()
	})
}

func Fuzz_Nosy_DecryptionKey_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		if btcSK == nil {
			return
		}

		dk, err := NewDecyptionKeyFromBTCSK(btcSK)
		if err != nil {
			return
		}
		dk.ToBytes()
	})
}

func Fuzz_Nosy_EncryptionKey_ToBTCPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcPK *btcec.PublicKey
		fill_err = tp.Fill(&btcPK)
		if fill_err != nil {
			return
		}
		if btcPK == nil {
			return
		}

		ek, err := NewEncryptionKeyFromBTCPK(btcPK)
		if err != nil {
			return
		}
		ek.ToBTCPK()
	})
}

func Fuzz_Nosy_EncryptionKey_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcPK *btcec.PublicKey
		fill_err = tp.Fill(&btcPK)
		if fill_err != nil {
			return
		}
		if btcPK == nil {
			return
		}

		ek, err := NewEncryptionKeyFromBTCPK(btcPK)
		if err != nil {
			return
		}
		ek.ToBytes()
	})
}

func Fuzz_Nosy_appendAndHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, msgHash []byte, signerPubKeyBytes []byte, encKeyBytes []byte) {
		appendAndHash(msgHash, signerPubKeyBytes, encKeyBytes)
	})
}

func Fuzz_Nosy_intoPointWithEvenY__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var point *btcec.JacobianPoint
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}
		if point == nil {
			return
		}

		intoPointWithEvenY(point)
	})
}

func Fuzz_Nosy_negatePoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var point *btcec.JacobianPoint
		fill_err = tp.Fill(&point)
		if fill_err != nil {
			return
		}
		if point == nil {
			return
		}

		negatePoint(point)
	})
}

func Fuzz_Nosy_unpackSchnorrSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sig *schnorr.Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if sig == nil {
			return
		}

		unpackSchnorrSig(sig)
	})
}
