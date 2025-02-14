package types

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/spf13/viper"
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

func Fuzz_Nosy_BIP340PubKey_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 []byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var pk2 *BIP340PubKey
		fill_err = tp.Fill(&pk2)
		if fill_err != nil {
			return
		}
		if pk2 == nil {
			return
		}

		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.Equals(pk2)
	})
}

func Fuzz_Nosy_BIP340PubKey_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.MarshalHex()
	})
}

func Fuzz_Nosy_BIP340PubKey_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, d2 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.Unmarshal(d2)
	})
}

func Fuzz_Nosy_BIP340PubKey_UnmarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, header string) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.UnmarshalHex(header)
	})
}

func Fuzz_Nosy_BIP340PubKey_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, bz []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.UnmarshalJSON(bz)
	})
}

func Fuzz_Nosy_BIP340Signature_ToHexStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.ToHexStr()
	})
}

func Fuzz_Nosy_BIP340Signature_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.Unmarshal(d2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Bits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Bits()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Difficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Difficulty()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Eq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header []byte
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var other *BTCHeaderBytes
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if other == nil {
			return
		}

		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Eq(other)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_FromBlockHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h1 []byte
		fill_err = tp.Fill(&h1)
		if fill_err != nil {
			return
		}
		var h2 *wire.BlockHeader
		fill_err = tp.Fill(&h2)
		if fill_err != nil {
			return
		}
		if h2 == nil {
			return
		}

		m, err := NewBTCHeaderBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.FromBlockHeader(h2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_HasParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h1 []byte
		fill_err = tp.Fill(&h1)
		if fill_err != nil {
			return
		}
		var h2 *BTCHeaderBytes
		fill_err = tp.Fill(&h2)
		if fill_err != nil {
			return
		}
		if h2 == nil {
			return
		}

		m, err := NewBTCHeaderBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.HasParent(h2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Hash()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_ParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.ParentHash()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Time__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Time()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte, d2 []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Unmarshal(d2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_UnmarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, h1 []byte, h2 string) {
		m, err := NewBTCHeaderBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.UnmarshalHex(h2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte, bz []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.UnmarshalJSON(bz)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_Eq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h1 []byte
		fill_err = tp.Fill(&h1)
		if fill_err != nil {
			return
		}
		var h2 *BTCHeaderHashBytes
		fill_err = tp.Fill(&h2)
		if fill_err != nil {
			return
		}
		if h2 == nil {
			return
		}

		m, err := NewBTCHeaderHashBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.Eq(h2)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_FromChainhash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h1 []byte
		fill_err = tp.Fill(&h1)
		if fill_err != nil {
			return
		}
		var h2 *chainhash.Hash
		fill_err = tp.Fill(&h2)
		if fill_err != nil {
			return
		}
		if h2 == nil {
			return
		}

		m, err := NewBTCHeaderHashBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.FromChainhash(h2)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.MarshalHex()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, bz []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.Unmarshal(bz)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_UnmarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, h1 []byte, h2 string) {
		m, err := NewBTCHeaderHashBytesFromBytes(h1)
		if err != nil {
			return
		}
		m.UnmarshalHex(h2)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, bz []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.UnmarshalJSON(bz)
	})
}

func Fuzz_Nosy_BtcConfig_NetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BtcConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.NetParams()
	})
}

func Fuzz_Nosy_BtcConfig_PowLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BtcConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PowLimit()
	})
}

func Fuzz_Nosy_BtcConfig_ReduceMinDifficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BtcConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ReduceMinDifficulty()
	})
}

func Fuzz_Nosy_BtcConfig_RetargetAdjustmentFactor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *BtcConfig
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.RetargetAdjustmentFactor()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sigHex string
		fill_err = tp.Fill(&sigHex)
		if fill_err != nil {
			return
		}
		var sig2 *SchnorrEOTSSig
		fill_err = tp.Fill(&sig2)
		if fill_err != nil {
			return
		}
		if sig2 == nil {
			return
		}

		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.Equals(sig2)
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_ToHexStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.ToHexStr()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string, d2 []byte) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.Unmarshal(d2)
	})
}

func Fuzz_Nosy_SchnorrPubRand_ToHexStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.ToHexStr()
	})
}

func Fuzz_Nosy_SchnorrPubRand_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, d2 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.Unmarshal(d2)
	})
}

func Fuzz_Nosy_BIP340PubKey_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.Marshal()
	})
}

func Fuzz_Nosy_BIP340PubKey_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.MarshalJSON()
	})
}

func Fuzz_Nosy_BIP340PubKey_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, d2 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BIP340PubKey_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.MustMarshal()
	})
}

func Fuzz_Nosy_BIP340PubKey_MustToBTCPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.MustToBTCPK()
	})
}

func Fuzz_Nosy_BIP340PubKey_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.Size()
	})
}

func Fuzz_Nosy_BIP340PubKey_ToBTCPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pk, err := NewBIP340PubKey(d1)
		if err != nil {
			return
		}
		pk.ToBTCPK()
	})
}

func Fuzz_Nosy_BIP340Signature_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.Marshal()
	})
}

func Fuzz_Nosy_BIP340Signature_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BIP340Signature_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.MustMarshal()
	})
}

func Fuzz_Nosy_BIP340Signature_MustToBTCSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.MustToBTCSig()
	})
}

func Fuzz_Nosy_BIP340Signature_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.Size()
	})
}

func Fuzz_Nosy_BIP340Signature_ToBTCSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcSig *schnorr.Signature
		fill_err = tp.Fill(&btcSig)
		if fill_err != nil {
			return
		}
		if btcSig == nil {
			return
		}

		sig := NewBIP340SignatureFromBTCSig(btcSig)
		sig.ToBTCSig()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.MarshalHex()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.MarshalJSON()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte, d2 []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BTCHeaderBytes_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.MustMarshal()
	})
}

func Fuzz_Nosy_BTCHeaderBytes_ToBlockHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, header []byte) {
		m, err := NewBTCHeaderBytesFromBytes(header)
		if err != nil {
			return
		}
		m.ToBlockHeader()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.MarshalJSON()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte, d2 []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.MustMarshal()
	})
}

func Fuzz_Nosy_BTCHeaderHashBytes_ToChainhash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash []byte) {
		m, err := NewBTCHeaderHashBytesFromBytes(hash)
		if err != nil {
			return
		}
		m.ToChainhash()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.Marshal()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string, d2 []byte) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.MarshalTo(d2)
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.MustMarshal()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.Size()
	})
}

func Fuzz_Nosy_SchnorrEOTSSig_ToModNScalar__(f *testing.F) {
	f.Fuzz(func(t *testing.T, sigHex string) {
		sig, err := NewSchnorrEOTSSigFromHex(sigHex)
		if err != nil {
			return
		}
		sig.ToModNScalar()
	})
}

func Fuzz_Nosy_SchnorrPubRand_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.Marshal()
	})
}

func Fuzz_Nosy_SchnorrPubRand_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.MarshalHex()
	})
}

func Fuzz_Nosy_SchnorrPubRand_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte, d2 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.MarshalTo(d2)
	})
}

func Fuzz_Nosy_SchnorrPubRand_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.MustMarshal()
	})
}

func Fuzz_Nosy_SchnorrPubRand_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.Size()
	})
}

func Fuzz_Nosy_SchnorrPubRand_ToFieldVal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		pr, err := NewSchnorrPubRand(d1)
		if err != nil {
			return
		}
		pr.ToFieldVal()
	})
}

func Fuzz_Nosy_GetOutputIdxInBTCTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var output *wire.TxOut
		fill_err = tp.Fill(&output)
		if fill_err != nil {
			return
		}
		if tx == nil || output == nil {
			return
		}

		GetOutputIdxInBTCTx(tx, output)
	})
}

func Fuzz_Nosy_MustGetGasSettings__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var configPath string
		fill_err = tp.Fill(&configPath)
		if fill_err != nil {
			return
		}
		var v *viper.Viper
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		MustGetGasSettings(configPath, v)
	})
}

func Fuzz_Nosy_NewBIP340PKsFromBTCPKs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcPKs []*btcec.PublicKey
		fill_err = tp.Fill(&btcPKs)
		if fill_err != nil {
			return
		}

		NewBIP340PKsFromBTCPKs(btcPKs)
	})
}

func Fuzz_Nosy_NewBTCPKsFromBIP340PKs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pks []BIP340PubKey
		fill_err = tp.Fill(&pks)
		if fill_err != nil {
			return
		}

		NewBTCPKsFromBIP340PKs(pks)
	})
}

func Fuzz_Nosy_NewBTCTxFromHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txHex string) {
		NewBTCTxFromHex(txHex)
	})
}

// skipping Fuzz_Nosy_ParseKeyNameFromConfig__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/server/types.AppOptions

// skipping Fuzz_Nosy_Reverse__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_SerializeBTCTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		SerializeBTCTx(tx)
	})
}

func Fuzz_Nosy_SortBIP340PKs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keys []BIP340PubKey
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}

		SortBIP340PKs(keys)
	})
}

// skipping Fuzz_Nosy_parseGasAdjustmentFromConfig__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/server/types.AppOptions

// skipping Fuzz_Nosy_parseGasPriceFromConfig__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/server/types.AppOptions
