package testutil

import (
	"math/rand"
	"testing"

	btclctypes "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
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

func Fuzz_Nosy_Helper_AddFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var fp *btcstakingtypes.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if h == nil || fp == nil {
			return
		}

		h.AddFinalityProvider(fp)
	})
}

func Fuzz_Nosy_Helper_AddInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var btcHeader *btclctypes.BTCHeaderInfo
		fill_err = tp.Fill(&btcHeader)
		if fill_err != nil {
			return
		}
		var proof *bstypes.InclusionProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var lightClientTipHeight uint32
		fill_err = tp.Fill(&lightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || btcHeader == nil || proof == nil {
			return
		}

		h.AddInclusionProof(stakingTxHash, btcHeader, proof, lightClientTipHeight)
	})
}

func Fuzz_Nosy_Helper_BeginBlocker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.BeginBlocker()
	})
}

func Fuzz_Nosy_Helper_CommitPubRandList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var fpSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpSK)
		if fill_err != nil {
			return
		}
		var fp *btcstakingtypes.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		var timestamped bool
		fill_err = tp.Fill(&timestamped)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || fpSK == nil || fp == nil {
			return
		}

		h.CommitPubRandList(r, fpSK, fp, startHeight, numPubRand, timestamped)
	})
}

func Fuzz_Nosy_Helper_CreateCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var covenantSKs []*btcec.PrivateKey
		fill_err = tp.Fill(&covenantSKs)
		if fill_err != nil {
			return
		}
		var msgCreateBTCDel *types.MsgCreateBTCDelegation
		fill_err = tp.Fill(&msgCreateBTCDel)
		if fill_err != nil {
			return
		}
		var del *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var lightClientTipHeight uint32
		fill_err = tp.Fill(&lightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || msgCreateBTCDel == nil || del == nil {
			return
		}

		h.CreateCovenantSigs(r, covenantSKs, msgCreateBTCDel, del, lightClientTipHeight)
	})
}

func Fuzz_Nosy_Helper_CreateDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var delSK *btcec.PrivateKey
		fill_err = tp.Fill(&delSK)
		if fill_err != nil {
			return
		}
		var fpPK *btcec.PublicKey
		fill_err = tp.Fill(&fpPK)
		if fill_err != nil {
			return
		}
		var stakingValue int64
		fill_err = tp.Fill(&stakingValue)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var unbondingValue int64
		fill_err = tp.Fill(&unbondingValue)
		if fill_err != nil {
			return
		}
		var unbondingTime uint16
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var usePreApproval bool
		fill_err = tp.Fill(&usePreApproval)
		if fill_err != nil {
			return
		}
		var addToAllowList bool
		fill_err = tp.Fill(&addToAllowList)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || delSK == nil || fpPK == nil {
			return
		}

		h.CreateDelegation(r, delSK, fpPK, stakingValue, stakingTime, unbondingValue, unbondingTime, usePreApproval, addToAllowList)
	})
}

func Fuzz_Nosy_Helper_CreateDelegationWithBtcBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var delSK *btcec.PrivateKey
		fill_err = tp.Fill(&delSK)
		if fill_err != nil {
			return
		}
		var fpPK *btcec.PublicKey
		fill_err = tp.Fill(&fpPK)
		if fill_err != nil {
			return
		}
		var stakingValue int64
		fill_err = tp.Fill(&stakingValue)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var unbondingValue int64
		fill_err = tp.Fill(&unbondingValue)
		if fill_err != nil {
			return
		}
		var unbondingTime uint16
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var usePreApproval bool
		fill_err = tp.Fill(&usePreApproval)
		if fill_err != nil {
			return
		}
		var addToAllowList bool
		fill_err = tp.Fill(&addToAllowList)
		if fill_err != nil {
			return
		}
		var stakingTransactionInclusionHeight uint32
		fill_err = tp.Fill(&stakingTransactionInclusionHeight)
		if fill_err != nil {
			return
		}
		var lightClientTipHeight uint32
		fill_err = tp.Fill(&lightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || delSK == nil || fpPK == nil {
			return
		}

		h.CreateDelegationWithBtcBlockHeight(r, delSK, fpPK, stakingValue, stakingTime, unbondingValue, unbondingTime, usePreApproval, addToAllowList, stakingTransactionInclusionHeight, lightClientTipHeight)
	})
}

func Fuzz_Nosy_Helper_CreateFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil {
			return
		}

		h.CreateFinalityProvider(r)
	})
}

// skipping Fuzz_Nosy_Helper_Equal__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Helper_Error__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Helper_GenAndApplyCustomParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var finalizationTimeout uint32
		fill_err = tp.Fill(&finalizationTimeout)
		if fill_err != nil {
			return
		}
		var unbondingTime uint32
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var allowListExpirationHeight uint64
		fill_err = tp.Fill(&allowListExpirationHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil {
			return
		}

		h.GenAndApplyCustomParams(r, finalizationTimeout, unbondingTime, allowListExpirationHeight)
	})
}

func Fuzz_Nosy_Helper_GenAndApplyParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil {
			return
		}

		h.GenAndApplyParams(r)
	})
}

func Fuzz_Nosy_Helper_GenerateCovenantSignaturesMessages__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var covenantSKs []*btcec.PrivateKey
		fill_err = tp.Fill(&covenantSKs)
		if fill_err != nil {
			return
		}
		var msgCreateBTCDel *types.MsgCreateBTCDelegation
		fill_err = tp.Fill(&msgCreateBTCDel)
		if fill_err != nil {
			return
		}
		var del *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || msgCreateBTCDel == nil || del == nil {
			return
		}

		h.GenerateCovenantSignaturesMessages(r, covenantSKs, msgCreateBTCDel, del)
	})
}

// skipping Fuzz_Nosy_Helper_NoError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Helper_SetCtxHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.SetCtxHeight(height)
	})
}

func Fuzz_Nosy_Helper_T__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *Helper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.T()
	})
}
