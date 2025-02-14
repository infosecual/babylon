package helper

import (
	"math/rand"
	"testing"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/testutil/datagen"
	btcstakingtypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func Fuzz_Nosy_Helper_AddDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var del *btcstakingtypes.BTCDelegation
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		if t1 == nil || del == nil {
			return
		}

		h := NewHelper(t1)
		h.AddDelegation(del)
	})
}

func Fuzz_Nosy_Helper_AddFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var fp *btcstakingtypes.FinalityProvider
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		if t1 == nil || fp == nil {
			return
		}

		h := NewHelper(t1)
		h.AddFinalityProvider(fp)
	})
}

func Fuzz_Nosy_Helper_ApplyEmptyBlockWithInvalidVoteExtensions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil {
			return
		}

		h := NewHelper(t1)
		h.ApplyEmptyBlockWithInvalidVoteExtensions(r)
	})
}

func Fuzz_Nosy_Helper_ApplyEmptyBlockWithSomeInvalidVoteExtensions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil {
			return
		}

		h := NewHelper(t1)
		h.ApplyEmptyBlockWithSomeInvalidVoteExtensions(r)
	})
}

func Fuzz_Nosy_Helper_ApplyEmptyBlockWithValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var valSetWithKeys *datagen.GenesisValidators
		fill_err = tp.Fill(&valSetWithKeys)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || valSetWithKeys == nil {
			return
		}

		h := NewHelper(t1)
		h.ApplyEmptyBlockWithValSet(r, valSetWithKeys)
	})
}

func Fuzz_Nosy_Helper_ApplyEmptyBlockWithVoteExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil {
			return
		}

		h := NewHelper(t1)
		h.ApplyEmptyBlockWithVoteExtension(r)
	})
}

func Fuzz_Nosy_Helper_CheckDelegator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var val sdk.ValAddress
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var found bool
		fill_err = tp.Fill(&found)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.CheckDelegator(delegator, val, found)
	})
}

func Fuzz_Nosy_Helper_CheckValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var addr sdk.ValAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var status stakingtypes.BondStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var jailed bool
		fill_err = tp.Fill(&jailed)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.CheckValidator(addr, status, jailed)
	})
}

// skipping Fuzz_Nosy_Helper_EqualError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Helper_Error__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Helper_Handle__ because parameters include func, chan, or unsupported interface: func(github.com/cosmos/cosmos-sdk/types.Context) (github.com/cosmos/gogoproto/proto.Message, error)

// skipping Fuzz_Nosy_Helper_NoError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_Helper_WrappedBeginRedelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var srcVal sdk.ValAddress
		fill_err = tp.Fill(&srcVal)
		if fill_err != nil {
			return
		}
		var dstVal sdk.ValAddress
		fill_err = tp.Fill(&dstVal)
		if fill_err != nil {
			return
		}
		var amount math.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.WrappedBeginRedelegate(delegator, srcVal, dstVal, amount)
	})
}

func Fuzz_Nosy_Helper_WrappedCancelUnbondingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var val sdk.ValAddress
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var amount math.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var creationHeight int64
		fill_err = tp.Fill(&creationHeight)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.WrappedCancelUnbondingDelegation(delegator, val, amount, creationHeight)
	})
}

func Fuzz_Nosy_Helper_WrappedDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var val sdk.ValAddress
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var amount math.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.WrappedDelegate(delegator, val, amount)
	})
}

func Fuzz_Nosy_Helper_WrappedDelegateWithPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var val sdk.ValAddress
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var power int64
		fill_err = tp.Fill(&power)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.WrappedDelegateWithPower(delegator, val, power)
	})
}

func Fuzz_Nosy_Helper_WrappedUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var delegator sdk.AccAddress
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		var val sdk.ValAddress
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var amount math.Int
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.WrappedUndelegate(delegator, val, amount)
	})
}

func Fuzz_Nosy_Helper_genAndApplyEmptyBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		h := NewHelper(t1)
		h.genAndApplyEmptyBlock()
	})
}

func Fuzz_Nosy_Helper_getExtendedVotesFromValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var blockHash checkpointingtypes.BlockHash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		var valSet *datagen.GenesisValidators
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}
		var numInvalidVotes int
		fill_err = tp.Fill(&numInvalidVotes)
		if fill_err != nil {
			return
		}
		if t1 == nil || valSet == nil {
			return
		}

		h := NewHelper(t1)
		h.getExtendedVotesFromValSet(epochNum, height, blockHash, valSet, numInvalidVotes)
	})
}

func Fuzz_Nosy_CalculateValHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var valSet []stakingtypes.Validator
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}

		CalculateValHash(valSet)
	})
}

// skipping Fuzz_Nosy_DiffKVStores__ because parameters include func, chan, or unsupported interface: cosmossdk.io/core/store.KVStore

func Fuzz_Nosy_ExtendedCommitToLastCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ec types.ExtendedCommitInfo
		fill_err = tp.Fill(&ec)
		if fill_err != nil {
			return
		}

		ExtendedCommitToLastCommit(ec)
	})
}

func Fuzz_Nosy_getDiffFromKVPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var kvAs []kv.Pair
		fill_err = tp.Fill(&kvAs)
		if fill_err != nil {
			return
		}
		var kvBs []kv.Pair
		fill_err = tp.Fill(&kvBs)
		if fill_err != nil {
			return
		}

		getDiffFromKVPair(kvAs, kvBs)
	})
}

// skipping Fuzz_Nosy_getKVPairs__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-db.Iterator
