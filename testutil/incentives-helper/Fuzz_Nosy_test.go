package testutil

import (
	"math/rand"
	"testing"

	//"github.com/babylonlabs-io/babylon/testutil"
	testutil "github.com/babylonlabs-io/babylon/testutil/btcstaking-helper"
	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
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

func Fuzz_Nosy_IncentiveHelper_BtcUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var del *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var unbondingInfo *testutil.UnbondingTxInfo
		fill_err = tp.Fill(&unbondingInfo)
		if fill_err != nil {
			return
		}
		var btcLightClientTipHeight uint32
		fill_err = tp.Fill(&btcLightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || del == nil || unbondingInfo == nil {
			return
		}

		h.BtcUndelegate(stakingTxHash, del, unbondingInfo, btcLightClientTipHeight)
	})
}

func Fuzz_Nosy_IncentiveHelper_CreateActiveBtcDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var covenantSKs []*secp256k1.PrivateKey
		fill_err = tp.Fill(&covenantSKs)
		if fill_err != nil {
			return
		}
		var fpPK *secp256k1.PublicKey
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
		var btcLightClientTipHeight uint32
		fill_err = tp.Fill(&btcLightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || fpPK == nil {
			return
		}

		h.CreateActiveBtcDelegation(r, covenantSKs, fpPK, stakingValue, stakingTime, btcLightClientTipHeight)
	})
}

func Fuzz_Nosy_IncentiveHelper_CreateBtcDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var fpPK *secp256k1.PublicKey
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
		var btcLightClientTipHeight uint32
		fill_err = tp.Fill(&btcLightClientTipHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || fpPK == nil {
			return
		}

		h.CreateBtcDelegation(r, fpPK, stakingValue, stakingTime, btcLightClientTipHeight)
	})
}

func Fuzz_Nosy_IncentiveHelper_CtxAddBlkHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var blocksToAdd int64
		fill_err = tp.Fill(&blocksToAdd)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.CtxAddBlkHeight(blocksToAdd)
	})
}

func Fuzz_Nosy_IncentiveHelper_EqualBtcDelRwdTrackerActiveSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var expectedSatAmount uint64
		fill_err = tp.Fill(&expectedSatAmount)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.EqualBtcDelRwdTrackerActiveSat(fp, del, expectedSatAmount)
	})
}

func Fuzz_Nosy_IncentiveHelper_EqualBtcDelegationStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var stakingTxHashStr string
		fill_err = tp.Fill(&stakingTxHashStr)
		if fill_err != nil {
			return
		}
		var tipHeight uint32
		fill_err = tp.Fill(&tipHeight)
		if fill_err != nil {
			return
		}
		var expectedStatus bstypes.BTCDelegationStatus
		fill_err = tp.Fill(&expectedStatus)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.EqualBtcDelegationStatus(stakingTxHashStr, tipHeight, expectedStatus)
	})
}

func Fuzz_Nosy_IncentiveHelper_FpAddPubRand__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if h == nil || r == nil || sk == nil {
			return
		}

		h.FpAddPubRand(r, sk, startHeight)
	})
}

func Fuzz_Nosy_IncentiveHelper_GenerateAndSendCovenantSignatures__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
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

		h.GenerateAndSendCovenantSignatures(r, covenantSKs, msgCreateBTCDel, del)
	})
}

func Fuzz_Nosy_IncentiveHelper_SetFinalityActivationHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var h *IncentiveHelper
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		var newActivationHeight uint64
		fill_err = tp.Fill(&newActivationHeight)
		if fill_err != nil {
			return
		}
		if h == nil {
			return
		}

		h.SetFinalityActivationHeight(newActivationHeight)
	})
}
