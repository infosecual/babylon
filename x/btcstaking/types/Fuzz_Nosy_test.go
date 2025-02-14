package types

import (
	"context"
	"testing"

	btcstaking2 "github.com/babylonlabs-io/babylon/btcstaking"
	asig "github.com/babylonlabs-io/babylon/crypto/schnorr-adaptor-signature"
	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
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

func Fuzz_Nosy_BIP322Sig_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BIP322Sig
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BIP322Sig_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAddress()
	})
}

func Fuzz_Nosy_BIP322Sig_GetSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSig()
	})
}

func Fuzz_Nosy_BIP322Sig_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BIP322Sig_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BIP322Sig_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BIP322Sig_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BIP322Sig
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BIP322Sig_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BIP322Sig_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BIP322Sig_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BIP322Sig_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BIP322Sig_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BIP322Sig_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BIP322Sig_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BIP322Sig_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BIP322Sig_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BIP322Sig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegation_AddCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var covPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPk)
		if fill_err != nil {
			return
		}
		var stakingSlashingSigs []asig.AdaptorSignature
		fill_err = tp.Fill(&stakingSlashingSigs)
		if fill_err != nil {
			return
		}
		var unbondingSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingSig)
		if fill_err != nil {
			return
		}
		var unbondingSlashingSigs []asig.AdaptorSignature
		fill_err = tp.Fill(&unbondingSlashingSigs)
		if fill_err != nil {
			return
		}
		if d == nil || covPk == nil || unbondingSig == nil {
			return
		}

		d.AddCovenantSigs(covPk, stakingSlashingSigs, unbondingSig, unbondingSlashingSigs)
	})
}

func Fuzz_Nosy_BTCDelegation_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Address()
	})
}

func Fuzz_Nosy_BTCDelegation_BuildSlashingTxWithWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bsParams *Params
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var btcNet *chaincfg.Params
		fill_err = tp.Fill(&btcNet)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		if d == nil || bsParams == nil || btcNet == nil || fpSK == nil {
			return
		}

		d.BuildSlashingTxWithWitness(bsParams, btcNet, fpSK)
	})
}

func Fuzz_Nosy_BTCDelegation_BuildUnbondingSlashingTxWithWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bsParams *Params
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var btcNet *chaincfg.Params
		fill_err = tp.Fill(&btcNet)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		if d == nil || bsParams == nil || btcNet == nil || fpSK == nil {
			return
		}

		d.BuildUnbondingSlashingTxWithWitness(bsParams, btcNet, fpSK)
	})
}

func Fuzz_Nosy_BTCDelegation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegation_FinalityProviderKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.FinalityProviderKeys()
	})
}

func Fuzz_Nosy_BTCDelegation_GetBtcTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcTipHeight()
	})
}

func Fuzz_Nosy_BTCDelegation_GetBtcUndelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcUndelegation()
	})
}

func Fuzz_Nosy_BTCDelegation_GetCovSlashingAdaptorSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var covBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covBTCPK)
		if fill_err != nil {
			return
		}
		var valIdx int
		fill_err = tp.Fill(&valIdx)
		if fill_err != nil {
			return
		}
		var quorum uint32
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if d == nil || covBTCPK == nil {
			return
		}

		d.GetCovSlashingAdaptorSig(covBTCPK, valIdx, quorum)
	})
}

func Fuzz_Nosy_BTCDelegation_GetCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCovenantSigs()
	})
}

func Fuzz_Nosy_BTCDelegation_GetEndHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEndHeight()
	})
}

func Fuzz_Nosy_BTCDelegation_GetFpIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if d == nil || fpBTCPK == nil {
			return
		}

		d.GetFpIdx(fpBTCPK)
	})
}

func Fuzz_Nosy_BTCDelegation_GetParamsVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParamsVersion()
	})
}

func Fuzz_Nosy_BTCDelegation_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPop()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakerAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakerAddr()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakingInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bsParams *Params
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var btcNet *chaincfg.Params
		fill_err = tp.Fill(&btcNet)
		if fill_err != nil {
			return
		}
		if d == nil || bsParams == nil || btcNet == nil {
			return
		}

		d.GetStakingInfo(bsParams, btcNet)
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakingOutputIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingOutputIdx()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTime()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTx()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GetStakingTxHash()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStartHeight()
	})
}

func Fuzz_Nosy_BTCDelegation_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var btcHeight uint32
		fill_err = tp.Fill(&btcHeight)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.GetStatus(btcHeight, covenantQuorum)
	})
}

func Fuzz_Nosy_BTCDelegation_GetTotalSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTotalSat()
	})
}

func Fuzz_Nosy_BTCDelegation_GetUnbondingInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bsParams *Params
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var btcNet *chaincfg.Params
		fill_err = tp.Fill(&btcNet)
		if fill_err != nil {
			return
		}
		if d == nil || bsParams == nil || btcNet == nil {
			return
		}

		d.GetUnbondingInfo(bsParams, btcNet)
	})
}

func Fuzz_Nosy_BTCDelegation_GetUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTime()
	})
}

func Fuzz_Nosy_BTCDelegation_HasCovenantQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var quorum uint32
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.HasCovenantQuorums(quorum)
	})
}

func Fuzz_Nosy_BTCDelegation_HasInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.HasInclusionProof()
	})
}

func Fuzz_Nosy_BTCDelegation_IsSignedByCovMember__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var covPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPk)
		if fill_err != nil {
			return
		}
		if d == nil || covPk == nil {
			return
		}

		d.IsSignedByCovMember(covPk)
	})
}

func Fuzz_Nosy_BTCDelegation_IsUnbondedEarly__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.IsUnbondedEarly()
	})
}

func Fuzz_Nosy_BTCDelegation_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegation_MustGetStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.MustGetStakingTx()
	})
}

func Fuzz_Nosy_BTCDelegation_MustGetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.MustGetStakingTxHash()
	})
}

func Fuzz_Nosy_BTCDelegation_MustGetValidStakingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.MustGetValidStakingTime()
	})
}

func Fuzz_Nosy_BTCDelegation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegation_SignUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var bsParams *Params
		fill_err = tp.Fill(&bsParams)
		if fill_err != nil {
			return
		}
		var btcNet *chaincfg.Params
		fill_err = tp.Fill(&btcNet)
		if fill_err != nil {
			return
		}
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		if d == nil || bsParams == nil || btcNet == nil || sk == nil {
			return
		}

		d.SignUnbondingTx(bsParams, btcNet, sk)
	})
}

func Fuzz_Nosy_BTCDelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCDelegation_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegation_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.ValidateBasic()
	})
}

func Fuzz_Nosy_BTCDelegation_VotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var btcHeight uint32
		fill_err = tp.Fill(&btcHeight)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.VotingPower(btcHeight, covenantQuorum)
	})
}

func Fuzz_Nosy_BTCDelegation_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegation_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegation_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegation_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegation_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegation_findFPIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *BTCDelegation
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if d == nil || fpBTCPK == nil {
			return
		}

		d.findFPIdx(fpBTCPK)
	})
}

func Fuzz_Nosy_BTCDelegationResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewBTCDelegationResponse(btcDel, status)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetActive()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetCovenantSigs()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetDelegatorSlashSigHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetDelegatorSlashSigHex()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetEndHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetEndHeight()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetParamsVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetParamsVersion()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetSlashingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetSlashingTxHex()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStakerAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStakerAddr()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStakingOutputIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStakingOutputIdx()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStakingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStakingTime()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStakingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStakingTxHex()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStartHeight()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetStatusDesc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetStatusDesc()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetTotalSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetTotalSat()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetUnbondingTime()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_GetUndelegationResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.GetUndelegationResponse()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewBTCDelegationResponse(btcDel, status)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.String()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegationResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegationResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegationResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBTCDelegationResponse(btcDel, status)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegator
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegator_GetIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetIdx()
	})
}

func Fuzz_Nosy_BTCDelegator_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegator
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCDelegator_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegator_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegator_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegator_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegator_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegator_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		i := NewBTCDelegatorDelegationIndex()
		i.Add(stakingTxHash)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		i := NewBTCDelegatorDelegationIndex()
		i.Has(stakingTxHash)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewBTCDelegatorDelegationIndex()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewBTCDelegatorDelegationIndex()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewBTCDelegatorDelegationIndex()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationIndex_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := NewBTCDelegatorDelegationIndex()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegatorDelegationIndex_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegatorDelegationIndex_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := NewBTCDelegatorDelegationIndex()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegatorDelegations
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_GetDels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDels()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegatorDelegations
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegatorDelegations_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegatorDelegations_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegations_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegations
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_GetDels__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDels()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegatorDelegationsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegatorDelegationsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCDelegatorDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCSlashingTx_BuildSlashingTxWithWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		var fpBTCPKs []bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPKs)
		if fill_err != nil {
			return
		}
		var fundingMsgTx *wire.MsgTx
		fill_err = tp.Fill(&fundingMsgTx)
		if fill_err != nil {
			return
		}
		var outputIdx uint32
		fill_err = tp.Fill(&outputIdx)
		if fill_err != nil {
			return
		}
		var delegatorSig *bbn.BIP340Signature
		fill_err = tp.Fill(&delegatorSig)
		if fill_err != nil {
			return
		}
		var covenantSigs []*asig.AdaptorSignature
		fill_err = tp.Fill(&covenantSigs)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var slashingPathSpendInfo *btcstaking2.SpendInfo
		fill_err = tp.Fill(&slashingPathSpendInfo)
		if fill_err != nil {
			return
		}
		if fpSK == nil || fundingMsgTx == nil || delegatorSig == nil || slashingPathSpendInfo == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.BuildSlashingTxWithWitness(fpSK, fpBTCPKs, fundingMsgTx, outputIdx, delegatorSig, covenantSigs, covenantQuorum, slashingPathSpendInfo)
	})
}

func Fuzz_Nosy_BTCSlashingTx_EncSign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingMsgTx *wire.MsgTx
		fill_err = tp.Fill(&fundingMsgTx)
		if fill_err != nil {
			return
		}
		var spendOutputIndex uint32
		fill_err = tp.Fill(&spendOutputIndex)
		if fill_err != nil {
			return
		}
		var slashingPkScriptPath []byte
		fill_err = tp.Fill(&slashingPkScriptPath)
		if fill_err != nil {
			return
		}
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var encKey *schnorr_adaptor_signature.EncryptionKey
		fill_err = tp.Fill(&encKey)
		if fill_err != nil {
			return
		}
		if fundingMsgTx == nil || sk == nil || encKey == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.EncSign(fundingMsgTx, spendOutputIndex, slashingPkScriptPath, sk, encKey)
	})
}

func Fuzz_Nosy_BTCSlashingTx_EncVerifyAdaptorSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingOut *wire.TxOut
		fill_err = tp.Fill(&fundingOut)
		if fill_err != nil {
			return
		}
		var slashingPkScriptPath []byte
		fill_err = tp.Fill(&slashingPkScriptPath)
		if fill_err != nil {
			return
		}
		var pk *btcec.PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var encKey *schnorr_adaptor_signature.EncryptionKey
		fill_err = tp.Fill(&encKey)
		if fill_err != nil {
			return
		}
		var sig *asig.AdaptorSignature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if fundingOut == nil || pk == nil || encKey == nil || sig == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.EncVerifyAdaptorSignature(fundingOut, slashingPkScriptPath, pk, encKey, sig)
	})
}

func Fuzz_Nosy_BTCSlashingTx_EncVerifyAdaptorSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingOut *wire.TxOut
		fill_err = tp.Fill(&fundingOut)
		if fill_err != nil {
			return
		}
		var slashingSpendInfo *btcstaking2.SpendInfo
		fill_err = tp.Fill(&slashingSpendInfo)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var valPKs []types.BIP340PubKey
		fill_err = tp.Fill(&valPKs)
		if fill_err != nil {
			return
		}
		var sigs [][]byte
		fill_err = tp.Fill(&sigs)
		if fill_err != nil {
			return
		}
		if fundingOut == nil || slashingSpendInfo == nil || pk == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.EncVerifyAdaptorSignatures(fundingOut, slashingSpendInfo, pk, valPKs, sigs)
	})
}

func Fuzz_Nosy_BTCSlashingTx_MustGetTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.MustGetTxHash()
	})
}

func Fuzz_Nosy_BTCSlashingTx_ParseEncVerifyAdaptorSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingOut *wire.TxOut
		fill_err = tp.Fill(&fundingOut)
		if fill_err != nil {
			return
		}
		var slashingSpendInfo *btcstaking2.SpendInfo
		fill_err = tp.Fill(&slashingSpendInfo)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var valPKs []types.BIP340PubKey
		fill_err = tp.Fill(&valPKs)
		if fill_err != nil {
			return
		}
		var sigs [][]byte
		fill_err = tp.Fill(&sigs)
		if fill_err != nil {
			return
		}
		if fundingOut == nil || slashingSpendInfo == nil || pk == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.ParseEncVerifyAdaptorSignatures(fundingOut, slashingSpendInfo, pk, valPKs, sigs)
	})
}

func Fuzz_Nosy_BTCSlashingTx_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingTx *wire.MsgTx
		fill_err = tp.Fill(&fundingTx)
		if fill_err != nil {
			return
		}
		var spendOutputIndex uint32
		fill_err = tp.Fill(&spendOutputIndex)
		if fill_err != nil {
			return
		}
		var slashingPkScriptPath []byte
		fill_err = tp.Fill(&slashingPkScriptPath)
		if fill_err != nil {
			return
		}
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		if fundingTx == nil || sk == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.Sign(fundingTx, spendOutputIndex, slashingPkScriptPath, sk)
	})
}

func Fuzz_Nosy_BTCSlashingTx_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.Size()
	})
}

func Fuzz_Nosy_BTCSlashingTx_ToHexStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.ToHexStr()
	})
}

func Fuzz_Nosy_BTCSlashingTx_ToMsgTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.ToMsgTx()
	})
}

func Fuzz_Nosy_BTCSlashingTx_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte, d2 []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.Unmarshal(d2)
	})
}

func Fuzz_Nosy_BTCSlashingTx_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var fundingOut *wire.TxOut
		fill_err = tp.Fill(&fundingOut)
		if fill_err != nil {
			return
		}
		var slashingPkScriptPath []byte
		fill_err = tp.Fill(&slashingPkScriptPath)
		if fill_err != nil {
			return
		}
		var pk *btcec.PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if fundingOut == nil || pk == nil || sig == nil {
			return
		}

		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.VerifySignature(fundingOut, slashingPkScriptPath, pk, sig)
	})
}

func Fuzz_Nosy_BTCUndelegation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCUndelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCUndelegation_GetCovSlashingAdaptorSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covBTCPK)
		if fill_err != nil {
			return
		}
		var valIdx int
		fill_err = tp.Fill(&valIdx)
		if fill_err != nil {
			return
		}
		var quorum uint32
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if ud == nil || covBTCPK == nil {
			return
		}

		ud.GetCovSlashingAdaptorSig(covBTCPK, valIdx, quorum)
	})
}

func Fuzz_Nosy_BTCUndelegation_GetCovenantSlashingSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCovenantSlashingSigs()
	})
}

func Fuzz_Nosy_BTCUndelegation_GetCovenantUnbondingSigList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCovenantUnbondingSigList()
	})
}

func Fuzz_Nosy_BTCUndelegation_GetDelegatorUnbondingInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorUnbondingInfo()
	})
}

func Fuzz_Nosy_BTCUndelegation_GetUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTx()
	})
}

func Fuzz_Nosy_BTCUndelegation_HasCovenantQuorumOnSlashing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var quorum uint32
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if ud == nil {
			return
		}

		ud.HasCovenantQuorumOnSlashing(quorum)
	})
}

func Fuzz_Nosy_BTCUndelegation_HasCovenantQuorumOnUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var quorum uint32
		fill_err = tp.Fill(&quorum)
		if fill_err != nil {
			return
		}
		if ud == nil {
			return
		}

		ud.HasCovenantQuorumOnUnbonding(quorum)
	})
}

func Fuzz_Nosy_BTCUndelegation_HasCovenantQuorums__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		if ud == nil {
			return
		}

		ud.HasCovenantQuorums(covenantQuorum)
	})
}

func Fuzz_Nosy_BTCUndelegation_IsSignedByCovMember__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPk)
		if fill_err != nil {
			return
		}
		if ud == nil || covPk == nil {
			return
		}

		ud.IsSignedByCovMember(covPk)
	})
}

func Fuzz_Nosy_BTCUndelegation_IsSignedByCovMemberOnSlashing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		if ud == nil || covPK == nil {
			return
		}

		ud.IsSignedByCovMemberOnSlashing(covPK)
	})
}

func Fuzz_Nosy_BTCUndelegation_IsSignedByCovMemberOnUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		if ud == nil || covPK == nil {
			return
		}

		ud.IsSignedByCovMemberOnUnbonding(covPK)
	})
}

func Fuzz_Nosy_BTCUndelegation_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCUndelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCUndelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCUndelegation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCUndelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCUndelegation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCUndelegation_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		if ud == nil {
			return
		}

		ud.ToResponse()
	})
}

func Fuzz_Nosy_BTCUndelegation_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegation_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCUndelegation_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCUndelegation_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCUndelegation_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCUndelegation_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCUndelegation_addCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ud *BTCUndelegation
		fill_err = tp.Fill(&ud)
		if fill_err != nil {
			return
		}
		var covPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPk)
		if fill_err != nil {
			return
		}
		var unbondingSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingSig)
		if fill_err != nil {
			return
		}
		var slashingSigs []asig.AdaptorSignature
		fill_err = tp.Fill(&slashingSigs)
		if fill_err != nil {
			return
		}
		if ud == nil || covPk == nil || unbondingSig == nil {
			return
		}

		ud.addCovenantSigs(covPk, unbondingSig, slashingSigs)
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCUndelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetCovenantSlashingSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCovenantSlashingSigs()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetCovenantUnbondingSigList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCovenantUnbondingSigList()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetDelegatorSlashingSigHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorSlashingSigHex()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetDelegatorUnbondingInfoResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorUnbondingInfoResponse()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetSlashingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashingTxHex()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_GetUnbondingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTxHex()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCUndelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCUndelegationResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCUndelegationResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCUndelegationResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCUndelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockHeightBbnToBtc
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_GetBlockHeightBbn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHeightBbn()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_GetBlockHeightBtc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHeightBtc()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlockHeightBbnToBtc
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BlockHeightBbnToBtc_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BlockHeightBbnToBtc_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_BlockHeightBbnToBtc_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlockHeightBbnToBtc
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CovenantAdaptorSignatures
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_GetAdaptorSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAdaptorSigs()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CovenantAdaptorSignatures
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_CovenantAdaptorSignatures_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_CovenantAdaptorSignatures_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_CovenantAdaptorSignatures_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CovenantAdaptorSignatures
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegatorUnbondingInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_GetSpendStakeTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSpendStakeTx()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegatorUnbondingInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_DelegatorUnbondingInfo_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_DelegatorUnbondingInfo_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfo_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegatorUnbondingInfo
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		_x1 := delegatorUnbondingInfoToResponse(ui)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_GetSpendStakeTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.GetSpendStakeTxHex()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.Marshal()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		_x1 := delegatorUnbondingInfoToResponse(ui)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.Reset()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.Size()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.String()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_DelegatorUnbondingInfoResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_DelegatorUnbondingInfoResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ui *DelegatorUnbondingInfo
		fill_err = tp.Fill(&ui)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if ui == nil {
			return
		}

		m := delegatorUnbondingInfoToResponse(ui)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewBtcDelCreationEvent(btcDel)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetFinalityProviderBtcPksHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetFinalityProviderBtcPksHex()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetParamsVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetParamsVersion()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetStakerBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetStakerBtcPkHex()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetStakingOutputIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetStakingOutputIndex()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetStakingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetStakingTime()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetStakingTxHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetStakingTxHex()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetUnbondingTime()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_GetUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.GetUnbondingTx()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewBtcDelCreationEvent(btcDel)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.Reset()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.String()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventBTCDelegationCreated_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCDelegationCreated_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationCreated_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewBtcDelCreationEvent(btcDel)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		_x1 := NewExpiredDelegationEvent(stakingTxHash)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, dAtA []byte) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, dAtA []byte) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		_x1 := NewExpiredDelegationEvent(stakingTxHash)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.Reset()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.String()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, dAtA []byte) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, b []byte, deterministic bool) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventBTCDelegationExpired_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCDelegationExpired_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationExpired_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, b []byte) {
		m := NewExpiredDelegationEvent(stakingTxHash)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		_x1 := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_GetEndHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.GetEndHeight()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_GetStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.GetStartHeight()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		_x1 := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.Reset()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.String()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationInclusionProofReceived_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var endHeight uint32
		fill_err = tp.Fill(&endHeight)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewInclusionProofEvent(stakingTxHash, startHeight, endHeight, state)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetNewState()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventBTCDelegationStateUpdate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCDelegationStateUpdate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventBTCDelegationStateUpdate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCDelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		_x1 := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_GetStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.GetStartHeight()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32, dAtA []byte) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32, dAtA []byte) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		_x1 := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.Reset()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.Size()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.String()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32, dAtA []byte) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32, b []byte, deterministic bool) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventBTCDelgationUnbondedEarly_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventBTCDelgationUnbondedEarly_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, startHeight uint32, b []byte) {
		m := NewDelegationUnbondedEarlyEvent(stakingTxHash, startHeight)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewCovenantQuorumReachedEvent(btcDel, state)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		_x1 := NewCovenantQuorumReachedEvent(btcDel, state)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.Reset()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.Size()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.String()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventCovenantQuorumReached_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCovenantQuorumReached_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventCovenantQuorumReached_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var state BTCDelegationStatus
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcDel == nil {
			return
		}

		m := NewCovenantQuorumReachedEvent(btcDel, state)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		_x1 := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_GetCovenantBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.GetCovenantBtcPkHex()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_GetCovenantUnbondingSignatureHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.GetCovenantUnbondingSignatureHex()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		_x1 := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.Reset()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.Size()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.String()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventCovenantSignatureReceived_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCovenantSignatureReceived_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventCovenantSignatureReceived_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcDel *BTCDelegation
		fill_err = tp.Fill(&btcDel)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcDel == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		m := NewCovenantSignatureReceivedEvent(btcDel, covPK, unbondingTxSig)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		_x1 := NewEventFinalityProviderCreated(fp)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetAddr()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetBtcPkHex()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetCommission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetCommission()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetDetails()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetIdentity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetIdentity()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetMoniker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetMoniker()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetSecurityContact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetSecurityContact()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_GetWebsite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.GetWebsite()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		_x1 := NewEventFinalityProviderCreated(fp)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.Reset()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.String()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventFinalityProviderCreated_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventFinalityProviderCreated_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderCreated_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderCreated(fp)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		_x1 := NewEventFinalityProviderEdited(fp)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetBtcPkHex()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetCommission__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetCommission()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetDetails()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetIdentity__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetIdentity()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetMoniker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetMoniker()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetSecurityContact__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetSecurityContact()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_GetWebsite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.GetWebsite()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		_x1 := NewEventFinalityProviderEdited(fp)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.Reset()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.String()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventFinalityProviderEdited_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventFinalityProviderEdited_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderEdited_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		m := NewEventFinalityProviderEdited(fp)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		_x1 := NewFinalityProviderStatusChangeEvent(fpPk, status)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_GetBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.GetBtcPk()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_GetNewState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.GetNewState()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		_x1 := NewFinalityProviderStatusChangeEvent(fpPk, status)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.Reset()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.String()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventFinalityProviderStatusChange_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventFinalityProviderStatusChange_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventFinalityProviderStatusChange_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPk)
		if fill_err != nil {
			return
		}
		var status FinalityProviderStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fpPk == nil {
			return
		}

		m := NewFinalityProviderStatusChangeEvent(fpPk, status)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventIndex_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventIndex
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventIndex_GetBlockHeightBtc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHeightBtc()
	})
}

func Fuzz_Nosy_EventIndex_GetEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEvent()
	})
}

func Fuzz_Nosy_EventIndex_GetIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetIdx()
	})
}

func Fuzz_Nosy_EventIndex_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventIndex_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventIndex_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventIndex_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventIndex
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventIndex_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventIndex_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventIndex_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventIndex_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventIndex_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventIndex_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventIndex_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventIndex_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventIndex_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventIndex
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		_x1 := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_GetBtcDelStateUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.GetBtcDelStateUpdate()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_GetEv__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.GetEv()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_GetJailedFp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.GetJailedFp()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_GetSlashedFp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.GetSlashedFp()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_GetUnjailedFp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.GetUnjailedFp()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		_x1 := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.Reset()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.String()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventPowerDistUpdate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventPowerDistUpdate_XXX_OneofWrappers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		_x1 := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		_x1.XXX_OneofWrappers()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		m := NewEventPowerDistUpdateWithUnjailedFP(fpBTCPK)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_BtcDelStateUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_BtcDelStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_BtcDelStateUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_BtcDelStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_BtcDelStateUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_BtcDelStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_BtcDelStateUpdate_isEventPowerDistUpdate_Ev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_BtcDelStateUpdate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isEventPowerDistUpdate_Ev()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventJailedFinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventJailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventSlashedFinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventSlashedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_EventUnjailedFinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_EventUnjailedFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_JailedFp_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_JailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_JailedFp_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_JailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_JailedFp_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_JailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_JailedFp_isEventPowerDistUpdate_Ev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_JailedFp
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isEventPowerDistUpdate_Ev()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_SlashedFp_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_SlashedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_SlashedFp_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_SlashedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_SlashedFp_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_SlashedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_SlashedFp_isEventPowerDistUpdate_Ev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_SlashedFp
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isEventPowerDistUpdate_Ev()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_UnjailedFp_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_UnjailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_UnjailedFp_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_UnjailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_UnjailedFp_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventPowerDistUpdate_UnjailedFp
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventPowerDistUpdate_UnjailedFp_isEventPowerDistUpdate_Ev__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventPowerDistUpdate_UnjailedFp
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isEventPowerDistUpdate_Ev()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventSelectiveSlashing
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_GetEvidence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEvidence()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventSelectiveSlashing
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventSelectiveSlashing_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventSelectiveSlashing_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventSelectiveSlashing_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventSelectiveSlashing
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		_x1 := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_GetSpendStakeTxBlockIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.GetSpendStakeTxBlockIndex()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_GetSpendStakeTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.GetSpendStakeTxHash()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_GetSpendStakeTxHeaderHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.GetSpendStakeTxHeaderHash()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32, dAtA []byte) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32, dAtA []byte) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		_x1 := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.Reset()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.Size()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.String()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32, dAtA []byte) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32, b []byte, deterministic bool) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventUnexpectedUnbondingTx_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventUnexpectedUnbondingTx_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventUnexpectedUnbondingTx_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string, spendStakeTxHash string, spendStakeTxHeaderHash string, spendStakeTxBlockIndex uint32, b []byte) {
		m := NewUnexpectedUnbondingTxEvent(stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_FinalityProvider_Address__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.Address()
	})
}

func Fuzz_Nosy_FinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *FinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_FinalityProvider_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAddr()
	})
}

func Fuzz_Nosy_FinalityProvider_GetDescription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDescription()
	})
}

func Fuzz_Nosy_FinalityProvider_GetHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHighestVotedHeight()
	})
}

func Fuzz_Nosy_FinalityProvider_GetJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetJailed()
	})
}

func Fuzz_Nosy_FinalityProvider_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPop()
	})
}

func Fuzz_Nosy_FinalityProvider_GetSlashedBabylonHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashedBabylonHeight()
	})
}

func Fuzz_Nosy_FinalityProvider_GetSlashedBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashedBtcHeight()
	})
}

func Fuzz_Nosy_FinalityProvider_IsJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.IsJailed()
	})
}

func Fuzz_Nosy_FinalityProvider_IsSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.IsSlashed()
	})
}

func Fuzz_Nosy_FinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_FinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_FinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_FinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *FinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_FinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_FinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_FinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_FinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_FinalityProvider_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fp *FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if fp == nil {
			return
		}

		fp.ValidateBasic()
	})
}

func Fuzz_Nosy_FinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_FinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_FinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_FinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_FinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_FinalityProviderResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		_x1 := NewFinalityProviderResponse(f1, bbnBlockHeight)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetAddr()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetDescription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetDescription()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetHeight()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetHighestVotedHeight()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetJailed()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetPop()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetSlashedBabylonHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetSlashedBabylonHeight()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_GetSlashedBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.GetSlashedBtcHeight()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.Marshal()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		_x1 := NewFinalityProviderResponse(f1, bbnBlockHeight)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.Reset()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.Size()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.String()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_FinalityProviderResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_FinalityProviderResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_FinalityProviderResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *FinalityProvider
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var bbnBlockHeight uint64
		fill_err = tp.Fill(&bbnBlockHeight)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		m := NewFinalityProviderResponse(f1, bbnBlockHeight)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *FinalityProviderWithMeta
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeight()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetHighestVotedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHighestVotedHeight()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetJailed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetJailed()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetSlashedBabylonHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashedBabylonHeight()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetSlashedBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashedBtcHeight()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_GetVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVotingPower()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *FinalityProviderWithMeta
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_FinalityProviderWithMeta_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_FinalityProviderWithMeta_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_FinalityProviderWithMeta_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *FinalityProviderWithMeta
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_GenesisState_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesis()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesis()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesis()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := DefaultGenesis()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_GenesisState_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_GenesisState_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := DefaultGenesis()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_HeightToVersionMap_AddNewPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightToVersionMap()
		m.AddNewPair(startHeight, version)
	})
}

func Fuzz_Nosy_HeightToVersionMap_AddPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var newPair *HeightVersionPair
		fill_err = tp.Fill(&newPair)
		if fill_err != nil {
			return
		}
		if newPair == nil {
			return
		}

		m := NewHeightToVersionMap()
		m.AddPair(newPair)
	})
}

func Fuzz_Nosy_HeightToVersionMap_GetVersionForHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint64) {
		m := NewHeightToVersionMap()
		m.GetVersionForHeight(height)
	})
}

func Fuzz_Nosy_HeightToVersionMap_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewHeightToVersionMap()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_HeightToVersionMap_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewHeightToVersionMap()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_HeightToVersionMap_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewHeightToVersionMap()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_HeightToVersionMap_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := NewHeightToVersionMap()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_HeightToVersionMap_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_HeightToVersionMap_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := NewHeightToVersionMap()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_HeightVersionPair_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		_x1 := NewHeightVersionPair(startHeight, version)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_HeightVersionPair_GetStartHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.GetStartHeight()
	})
}

func Fuzz_Nosy_HeightVersionPair_GetVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.GetVersion()
	})
}

func Fuzz_Nosy_HeightVersionPair_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.Marshal()
	})
}

func Fuzz_Nosy_HeightVersionPair_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32, dAtA []byte) {
		m := NewHeightVersionPair(startHeight, version)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_HeightVersionPair_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32, dAtA []byte) {
		m := NewHeightVersionPair(startHeight, version)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_HeightVersionPair_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		_x1 := NewHeightVersionPair(startHeight, version)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_HeightVersionPair_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.Reset()
	})
}

func Fuzz_Nosy_HeightVersionPair_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.Size()
	})
}

func Fuzz_Nosy_HeightVersionPair_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.String()
	})
}

func Fuzz_Nosy_HeightVersionPair_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32, dAtA []byte) {
		m := NewHeightVersionPair(startHeight, version)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_HeightVersionPair_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_HeightVersionPair_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32, b []byte, deterministic bool) {
		m := NewHeightVersionPair(startHeight, version)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_HeightVersionPair_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_HeightVersionPair_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32) {
		m := NewHeightVersionPair(startHeight, version)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_HeightVersionPair_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, startHeight uint64, version uint32, b []byte) {
		m := NewHeightVersionPair(startHeight, version)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_InclusionProof_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		_x1 := NewInclusionProof(txKey, proof)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_InclusionProof_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.GetKey()
	})
}

func Fuzz_Nosy_InclusionProof_GetProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.GetProof()
	})
}

func Fuzz_Nosy_InclusionProof_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.Marshal()
	})
}

func Fuzz_Nosy_InclusionProof_MarshalHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		ip := NewInclusionProof(txKey, proof)
		ip.MarshalHex()
	})
}

func Fuzz_Nosy_InclusionProof_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_InclusionProof_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_InclusionProof_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		_x1 := NewInclusionProof(txKey, proof)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_InclusionProof_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.Reset()
	})
}

func Fuzz_Nosy_InclusionProof_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.Size()
	})
}

func Fuzz_Nosy_InclusionProof_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.String()
	})
}

func Fuzz_Nosy_InclusionProof_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_InclusionProof_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		ip := NewInclusionProof(txKey, proof)
		ip.ValidateBasic()
	})
}

func Fuzz_Nosy_InclusionProof_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_InclusionProof_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_InclusionProof_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_InclusionProof_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_InclusionProof_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txKey *types.TransactionKey
		fill_err = tp.Fill(&txKey)
		if fill_err != nil {
			return
		}
		var proof []byte
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if txKey == nil {
			return
		}

		m := NewInclusionProof(txKey, proof)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MockBTCLightClientKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCLightClientKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBTCLightClientKeeper_GetBaseBTCHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCLightClientKeeper(ctrl)
		m.GetBaseBTCHeader(ctx)
	})
}

func Fuzz_Nosy_MockBTCLightClientKeeper_GetHeaderByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var hash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || hash == nil {
			return
		}

		m := NewMockBTCLightClientKeeper(ctrl)
		m.GetHeaderByHash(ctx, hash)
	})
}

func Fuzz_Nosy_MockBTCLightClientKeeper_GetTipInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCLightClientKeeper(ctrl)
		m.GetTipInfo(ctx)
	})
}

// skipping Fuzz_Nosy_MockBTCLightClientKeeperMockRecorder_GetBaseBTCHeader__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCLightClientKeeperMockRecorder_GetHeaderByHash__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCLightClientKeeperMockRecorder_GetTipInfo__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBtcCheckpointKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBtcCheckpointKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBtcCheckpointKeeper_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBtcCheckpointKeeper(ctrl)
		m.GetParams(ctx)
	})
}

// skipping Fuzz_Nosy_MockBtcCheckpointKeeperMockRecorder_GetParams__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockFinalityKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockFinalityKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockFinalityKeeper_HasTimestampedPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fpBtcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPK)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil || fpBtcPK == nil {
			return
		}

		m := NewMockFinalityKeeper(ctrl)
		m.HasTimestampedPubRand(ctx, fpBtcPK, height)
	})
}

// skipping Fuzz_Nosy_MockFinalityKeeperMockRecorder_HasTimestampedPubRand__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockIncentiveKeeper_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockIncentiveKeeper(ctrl)
		m.EXPECT()
	})
}

// skipping Fuzz_Nosy_MockIncentiveKeeper_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_MockIncentiveKeeperMockRecorder_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSigner()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_GetStakingTxInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxInclusionProof()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProof_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgAddBTCDelegationInclusionProofResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddBTCDelegationInclusionProofResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddCovenantSigs
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSigner()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_GetSlashingTxSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashingTxSigs()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_GetSlashingUnbondingTxSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSlashingUnbondingTxSigs()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddCovenantSigs
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgAddCovenantSigs_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgAddCovenantSigs_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigs_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigs
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgAddCovenantSigsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgAddCovenantSigsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgAddCovenantSigsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgAddCovenantSigsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgBTCUndelegate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSigner()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_GetStakeSpendingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakeSpendingTx()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_GetStakeSpendingTxInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakeSpendingTxInclusionProof()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgBTCUndelegate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgBTCUndelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgBTCUndelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgBTCUndelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgBTCUndelegateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgBTCUndelegateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgBTCUndelegateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgBTCUndelegateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateBTCDelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPop()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetStakerAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakerAddr()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetStakingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTime()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTx()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetStakingTxInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxInclusionProof()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetStakingValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingValue()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTime()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTx()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_GetUnbondingValue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingValue()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateBTCDelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgCreateBTCDelegation_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgCreateBTCDelegation_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegation_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgCreateBTCDelegationResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgCreateBTCDelegationResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAddr()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_GetDescription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDescription()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPop()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgCreateFinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgCreateFinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgCreateFinalityProviderResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgCreateFinalityProviderResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgCreateFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgEditFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAddr()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_GetBtcPk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcPk()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_GetDescription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDescription()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgEditFinalityProvider
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.ValidateBasic()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgEditFinalityProvider_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgEditFinalityProvider_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgEditFinalityProvider_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProvider
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgEditFinalityProviderResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgEditFinalityProviderResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgEditFinalityProviderResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgEditFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_GetRecoveredFpBtcSk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRecoveredFpBtcSk()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSigner()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgSelectiveSlashingEvidence_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidence_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgSelectiveSlashingEvidenceResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSelectiveSlashingEvidenceResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgUpdateParams_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgUpdateParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgUpdateParams_GetAuthority__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAuthority()
	})
}

func Fuzz_Nosy_MsgUpdateParams_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParams()
	})
}

func Fuzz_Nosy_MsgUpdateParams_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgUpdateParams_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParams_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParams_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgUpdateParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgUpdateParams_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgUpdateParams_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgUpdateParams_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgUpdateParams_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParams_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgUpdateParams_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgUpdateParams_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgUpdateParams_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgUpdateParams_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgUpdateParamsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgUpdateParamsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgUpdateParamsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgUpdateParamsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgUpdateParamsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgUpdateParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_Params_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Params_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Params_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Params_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := DefaultParams()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Params_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Params_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := DefaultParams()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_ParsedCreateDelegationMessage_IsIncludedOnBTC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *MsgCreateBTCDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m1, err := ParseCreateDelegationMessage(msg)
		if err != nil {
			return
		}
		m1.IsIncludedOnBTC()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		_x1, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_GetBtcSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.GetBtcSig()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_GetBtcSigType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.GetBtcSigType()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		_x1, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_ToHexStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		pop.ToHexStr()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		pop.ValidateBasic()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_Verify__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var n3 *chaincfg.Params
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		var staker sdk.AccAddress
		fill_err = tp.Fill(&staker)
		if fill_err != nil {
			return
		}
		var bip340PK *bbn.BIP340PubKey
		fill_err = tp.Fill(&bip340PK)
		if fill_err != nil {
			return
		}
		var n6 *chaincfg.Params
		fill_err = tp.Fill(&n6)
		if fill_err != nil {
			return
		}
		if btcSK == nil || n3 == nil || bip340PK == nil || n6 == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, n3)
		if err != nil {
			return
		}
		pop.Verify(staker, bip340PK, n6)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_VerifyBIP322__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var n3 *chaincfg.Params
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var bip340PK *bbn.BIP340PubKey
		fill_err = tp.Fill(&bip340PK)
		if fill_err != nil {
			return
		}
		var n6 *chaincfg.Params
		fill_err = tp.Fill(&n6)
		if fill_err != nil {
			return
		}
		if btcSK == nil || n3 == nil || bip340PK == nil || n6 == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, n3)
		if err != nil {
			return
		}
		pop.VerifyBIP322(addr, bip340PK, n6)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_VerifyBIP340__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var stakerAddr sdk.AccAddress
		fill_err = tp.Fill(&stakerAddr)
		if fill_err != nil {
			return
		}
		var bip340PK *bbn.BIP340PubKey
		fill_err = tp.Fill(&bip340PK)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil || bip340PK == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		pop.VerifyBIP340(stakerAddr, bip340PK)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_VerifyECDSA__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var bip340PK *bbn.BIP340PubKey
		fill_err = tp.Fill(&bip340PK)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil || bip340PK == nil {
			return
		}

		pop, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		pop.VerifyECDSA(addr, bip340PK)
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_ProofOfPossessionBTC_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ProofOfPossessionBTC_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_ProofOfPossessionBTC_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addrToSign sdk.AccAddress
		fill_err = tp.Fill(&addrToSign)
		if fill_err != nil {
			return
		}
		var btcSK *btcec.PrivateKey
		fill_err = tp.Fill(&btcSK)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcSK == nil || net == nil {
			return
		}

		m, err := NewPoPBTCWithBIP322P2TRBIP86Sig(addrToSign, btcSK, net)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_GetStakingTxHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHashHex()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryBTCDelegationRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCDelegationRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_GetBtcDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcDelegation()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryBTCDelegationResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCDelegationResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatus()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryBTCDelegationsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCDelegationsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_GetBtcDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcDelegations()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryBTCDelegationsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCDelegationsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryBTCDelegationsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_GetFpBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetFpBtcPkHex()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_GetBtcDelegatorDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcDelegatorDelegations()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderDelegationsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderDelegationsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_GetFpBtcPkHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetFpBtcPkHex()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProviderRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProviderRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_GetFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetFinalityProvider()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProviderResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProviderResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProviderResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProviderResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProviderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProvidersRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProvidersRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_GetFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetFinalityProviders()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryFinalityProvidersResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryFinalityProvidersResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryFinalityProvidersResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryFinalityProvidersResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_GetBtcHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBtcHeight()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsByBTCHeightRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParams()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_GetVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVersion()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsByBTCHeightResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsByBTCHeightResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByBTCHeightResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByVersionRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_GetVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVersion()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByVersionRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsByVersionRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsByVersionRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsByVersionRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByVersionResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParams()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsByVersionResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsByVersionResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsByVersionResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsByVersionResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsByVersionResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsRequest
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryParamsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryParamsResponse_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParams()
	})
}

func Fuzz_Nosy_QueryParamsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_QueryParamsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryParamsResponse
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryParamsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_QueryParamsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_QueryParamsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_QueryParamsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryParamsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryParamsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryParamsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryParamsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryParamsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryParamsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *SelectiveSlashingEvidence
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_GetRecoveredFpBtcSk__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRecoveredFpBtcSk()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_GetStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStakingTxHash()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *SelectiveSlashingEvidence
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_SelectiveSlashingEvidence_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_SelectiveSlashingEvidence_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_SelectiveSlashingEvidence_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *SelectiveSlashingEvidence
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_SignatureInfo_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		_x1 := NewSignatureInfo(pk, sig)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_SignatureInfo_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.Marshal()
	})
}

func Fuzz_Nosy_SignatureInfo_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_SignatureInfo_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_SignatureInfo_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		_x1 := NewSignatureInfo(pk, sig)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_SignatureInfo_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.Reset()
	})
}

func Fuzz_Nosy_SignatureInfo_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.Size()
	})
}

func Fuzz_Nosy_SignatureInfo_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.String()
	})
}

func Fuzz_Nosy_SignatureInfo_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_SignatureInfo_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_SignatureInfo_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_SignatureInfo_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_SignatureInfo_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_SignatureInfo_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if pk == nil || sig == nil {
			return
		}

		m := NewSignatureInfo(pk, sig)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_StoredParams_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StoredParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Descriptor()
	})
}

func Fuzz_Nosy_StoredParams_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetParams()
	})
}

func Fuzz_Nosy_StoredParams_GetVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVersion()
	})
}

func Fuzz_Nosy_StoredParams_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Marshal()
	})
}

func Fuzz_Nosy_StoredParams_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_StoredParams_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_StoredParams_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *StoredParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_StoredParams_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Reset()
	})
}

func Fuzz_Nosy_StoredParams_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Size()
	})
}

func Fuzz_Nosy_StoredParams_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.String()
	})
}

func Fuzz_Nosy_StoredParams_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_StoredParams_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_StoredParams_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deterministic bool
		fill_err = tp.Fill(&deterministic)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_StoredParams_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_StoredParams_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Size()
	})
}

func Fuzz_Nosy_StoredParams_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *StoredParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_AddBTCDelegationInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgAddBTCDelegationInclusionProof
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.AddBTCDelegationInclusionProof(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_AddCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgAddCovenantSigs
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.AddCovenantSigs(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_BTCUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgBTCUndelegate
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BTCUndelegate(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_CreateBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgCreateBTCDelegation
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.CreateBTCDelegation(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgCreateFinalityProvider
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.CreateFinalityProvider(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_EditFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgEditFinalityProvider
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EditFinalityProvider(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_SelectiveSlashingEvidence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgSelectiveSlashingEvidence
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.SelectiveSlashingEvidence(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_UpdateParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedMsgServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *MsgUpdateParams
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.UpdateParams(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_BTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryBTCDelegationRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BTCDelegation(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_BTCDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryBTCDelegationsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BTCDelegations(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_FinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryFinalityProviderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.FinalityProvider(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_FinalityProviderDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryFinalityProviderDelegationsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.FinalityProviderDelegations(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_FinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryFinalityProvidersRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.FinalityProviders(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryParamsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.Params(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_ParamsByBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryParamsByBTCHeightRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.ParamsByBTCHeight(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_ParamsByVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *UnimplementedQueryServer
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var req *QueryParamsByVersionRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.ParamsByVersion(ctx, req)
	})
}

// skipping Fuzz_Nosy_msgClient_AddBTCDelegationInclusionProof__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_AddCovenantSigs__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_BTCUndelegate__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_CreateBTCDelegation__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_CreateFinalityProvider__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_EditFinalityProvider__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_SelectiveSlashingEvidence__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_BTCDelegation__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_BTCDelegations__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_FinalityProvider__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_FinalityProviderDelegations__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_FinalityProviders__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Params__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_ParamsByBTCHeight__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_ParamsByVersion__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

func Fuzz_Nosy_BTCDelegationStatus_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, statusStr string) {
		_x1, err := NewBTCDelegationStatusFromString(statusStr)
		if err != nil {
			return
		}
		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_BTCDelegationStatus_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, statusStr string) {
		x, err := NewBTCDelegationStatusFromString(statusStr)
		if err != nil {
			return
		}
		x.String()
	})
}

// skipping Fuzz_Nosy_BTCLightClientKeeper_GetBaseBTCHeader__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.BTCLightClientKeeper

// skipping Fuzz_Nosy_BTCLightClientKeeper_GetHeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.BTCLightClientKeeper

// skipping Fuzz_Nosy_BTCLightClientKeeper_GetTipInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.BTCLightClientKeeper

func Fuzz_Nosy_BTCSigType_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 BTCSigType
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_BTCSigType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BTCSigType
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_BTCSlashingTx_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.Marshal()
	})
}

func Fuzz_Nosy_BTCSlashingTx_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte, d2 []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BTCSlashingTx_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txBytes []byte) {
		tx := NewBtcSlashingTxFromBytes(txBytes)
		tx.MustMarshal()
	})
}

// skipping Fuzz_Nosy_BtcCheckpointKeeper_GetParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.BtcCheckpointKeeper

// skipping Fuzz_Nosy_FinalityKeeper_HasTimestampedPubRand__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.FinalityKeeper

func Fuzz_Nosy_FinalityProviderStatus_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 FinalityProviderStatus
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_FinalityProviderStatus_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x FinalityProviderStatus
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

// skipping Fuzz_Nosy_IncentiveKeeper_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.IncentiveKeeper

// skipping Fuzz_Nosy_MsgClient_AddBTCDelegationInclusionProof__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_AddCovenantSigs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_BTCUndelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_CreateBTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_CreateFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_EditFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_SelectiveSlashingEvidence__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgClient

// skipping Fuzz_Nosy_MsgServer_AddBTCDelegationInclusionProof__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_AddCovenantSigs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_BTCUndelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_CreateBTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_CreateFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_EditFinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_SelectiveSlashingEvidence__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.MsgServer

func Fuzz_Nosy_Params_HasCovenantPK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk *bbn.BIP340PubKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}
		if pk == nil {
			return
		}

		p := DefaultParams()
		p.HasCovenantPK(pk)
	})
}

// skipping Fuzz_Nosy_QueryClient_BTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_BTCDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_FinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_FinalityProviderDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_FinalityProviders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_ParamsByBTCHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_ParamsByVersion__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_BTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_BTCDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_FinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_FinalityProviderDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_FinalityProviders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_ParamsByBTCHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_ParamsByVersion__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryServer

// skipping Fuzz_Nosy_isEventPowerDistUpdate_Ev_MarshalTo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.isEventPowerDistUpdate_Ev

// skipping Fuzz_Nosy_isEventPowerDistUpdate_Ev_Size__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.isEventPowerDistUpdate_Ev

// skipping Fuzz_Nosy_isEventPowerDistUpdate_Ev_isEventPowerDistUpdate_Ev__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.isEventPowerDistUpdate_Ev

func Fuzz_Nosy_EmitEarlyUnbondedEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var inclusionHeight uint32
		fill_err = tp.Fill(&inclusionHeight)
		if fill_err != nil {
			return
		}

		EmitEarlyUnbondedEvent(sdkCtx, stakingTxHash, inclusionHeight)
	})
}

func Fuzz_Nosy_EmitExpiredDelegationEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		EmitExpiredDelegationEvent(sdkCtx, stakingTxHash)
	})
}

func Fuzz_Nosy_EmitJailedFPEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		EmitJailedFPEvent(sdkCtx, fpBTCPK)
	})
}

func Fuzz_Nosy_EmitSlashedFPEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		EmitSlashedFPEvent(sdkCtx, fpBTCPK)
	})
}

func Fuzz_Nosy_EmitUnexpectedUnbondingTxEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sdkCtx sdk.Context
		fill_err = tp.Fill(&sdkCtx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var spendStakeTxHash string
		fill_err = tp.Fill(&spendStakeTxHash)
		if fill_err != nil {
			return
		}
		var spendStakeTxHeaderHash string
		fill_err = tp.Fill(&spendStakeTxHeaderHash)
		if fill_err != nil {
			return
		}
		var spendStakeTxBlockIndex uint32
		fill_err = tp.Fill(&spendStakeTxBlockIndex)
		if fill_err != nil {
			return
		}

		EmitUnexpectedUnbondingTxEvent(sdkCtx, stakingTxHash, spendStakeTxHash, spendStakeTxHeaderHash, spendStakeTxBlockIndex)
	})
}

func Fuzz_Nosy_ExistsDup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcPKs []types.BIP340PubKey
		fill_err = tp.Fill(&btcPKs)
		if fill_err != nil {
			return
		}

		ExistsDup(btcPKs)
	})
}

func Fuzz_Nosy_GetOrderedCovenantSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpIdx int
		fill_err = tp.Fill(&fpIdx)
		if fill_err != nil {
			return
		}
		var covSigsList []*CovenantAdaptorSignatures
		fill_err = tp.Fill(&covSigsList)
		if fill_err != nil {
			return
		}
		var params *Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		GetOrderedCovenantSignatures(fpIdx, covSigsList, params)
	})
}

func Fuzz_Nosy_RecordActiveFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, num int) {
		RecordActiveFinalityProviders(num)
	})
}

func Fuzz_Nosy_RecordBTCDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var num int
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		var status BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		RecordBTCDelegations(num, status)
	})
}

func Fuzz_Nosy_RecordInactiveFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, num int) {
		RecordInactiveFinalityProviders(num)
	})
}

func Fuzz_Nosy_RecordMetricsKeyStakedBitcoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, amount float32) {
		RecordMetricsKeyStakedBitcoins(amount)
	})
}

func Fuzz_Nosy_RegisterCodec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cdc *codec.LegacyAmino
		fill_err = tp.Fill(&cdc)
		if fill_err != nil {
			return
		}
		if cdc == nil {
			return
		}

		RegisterCodec(cdc)
	})
}

// skipping Fuzz_Nosy_RegisterInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.InterfaceRegistry

// skipping Fuzz_Nosy_RegisterMsgServer__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy_RegisterQueryServer__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy__Msg_AddBTCDelegationInclusionProof_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_AddCovenantSigs_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_BTCUndelegate_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_CreateBTCDelegation_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_CreateFinalityProvider_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_EditFinalityProvider_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_SelectiveSlashingEvidence_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_UpdateParams_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_BTCDelegation_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_BTCDelegations_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_FinalityProviderDelegations_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_FinalityProvider_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_FinalityProviders_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_ParamsByBTCHeight_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_ParamsByVersion_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Params_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintBtcstaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintBtcstaking(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintEvents(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintGenesis(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintParams(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintPop(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintQuery(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintTx(dAtA, offset, v)
	})
}

func Fuzz_Nosy_findFPIdxInWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		var fpBTCPKs []types.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPKs)
		if fill_err != nil {
			return
		}
		if fpSK == nil {
			return
		}

		findFPIdxInWitness(fpSK, fpBTCPKs)
	})
}

// skipping Fuzz_Nosy_local_request_Query_BTCDelegation_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_BTCDelegations_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_FinalityProviderDelegations_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_FinalityProvider_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_FinalityProviders_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_ParamsByBTCHeight_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_ParamsByVersion_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_newBIP322Sig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btcstaking/types.bip322Sign[A]

// skipping Fuzz_Nosy_request_Query_BTCDelegation_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_BTCDelegations_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_FinalityProviderDelegations_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_FinalityProvider_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_FinalityProviders_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_ParamsByBTCHeight_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_ParamsByVersion_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipBtcstaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipBtcstaking(dAtA)
	})
}

func Fuzz_Nosy_skipEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipEvents(dAtA)
	})
}

func Fuzz_Nosy_skipGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipGenesis(dAtA)
	})
}

func Fuzz_Nosy_skipParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipParams(dAtA)
	})
}

func Fuzz_Nosy_skipPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipPop(dAtA)
	})
}

func Fuzz_Nosy_skipQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipQuery(dAtA)
	})
}

func Fuzz_Nosy_skipTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipTx(dAtA)
	})
}

func Fuzz_Nosy_sovBtcstaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovBtcstaking(x)
	})
}

func Fuzz_Nosy_sovEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovEvents(x)
	})
}

func Fuzz_Nosy_sovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovGenesis(x)
	})
}

func Fuzz_Nosy_sovParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovParams(x)
	})
}

func Fuzz_Nosy_sovPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovPop(x)
	})
}

func Fuzz_Nosy_sovQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovQuery(x)
	})
}

func Fuzz_Nosy_sovTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovTx(x)
	})
}

func Fuzz_Nosy_sozBtcstaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozBtcstaking(x)
	})
}

func Fuzz_Nosy_sozEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozEvents(x)
	})
}

func Fuzz_Nosy_sozGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozGenesis(x)
	})
}

func Fuzz_Nosy_sozParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozParams(x)
	})
}

func Fuzz_Nosy_sozPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozPop(x)
	})
}

func Fuzz_Nosy_sozQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozQuery(x)
	})
}

func Fuzz_Nosy_sozTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozTx(x)
	})
}
