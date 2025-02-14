package replay

import (
	"math/rand"
	"testing"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/wire"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// skipping Fuzz_Nosy_BabylonAppDriver_CreateTx__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

func Fuzz_Nosy_BabylonAppDriver_ExtendBTCLcWithNEmptyBlocks__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var n uint32
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || t5 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.ExtendBTCLcWithNEmptyBlocks(r, t5, n)
	})
}

func Fuzz_Nosy_BabylonAppDriver_FinializeCkptForEpoch__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || t5 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.FinializeCkptForEpoch(r, t5, epochNumber)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GenBlockWithTransactions__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var txs []*wire.MsgTx
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || t5 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GenBlockWithTransactions(r, t5, txs)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GenCkptForEpoch__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || t5 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GenCkptForEpoch(r, t5, epochNumber)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GenerateNewBlock__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GenerateNewBlock(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GenerateNewBlockAssertExecutionSuccess__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GenerateNewBlockAssertExecutionSuccess(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetActiveBTCDelegations__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetActiveBTCDelegations(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetActiveFpsAtCurrentHeight__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetActiveFpsAtCurrentHeight(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetActiveFpsAtHeight__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetActiveFpsAtHeight(t4, height)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetAllBTCDelegations__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetAllBTCDelegations(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetBTCCkptParams__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetBTCCkptParams(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetBTCLCTip__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetBTCLCTip()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetBTCStakingParams__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetBTCStakingParams(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetCheckpoint__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetCheckpoint(t4, epochNumber)
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetContextForLastFinalizedBlock__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetContextForLastFinalizedBlock()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetDriverAccountAddress__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetDriverAccountAddress()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetDriverAccountSenderInfo__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetDriverAccountSenderInfo()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetEpoch__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetEpoch()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetEpochingParams__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetEpochingParams()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetLastFinalizedBlock__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetLastFinalizedBlock()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetLastFinalizedEpoch__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetLastFinalizedEpoch()
	})
}

func Fuzz_Nosy_BabylonAppDriver_GetVerifiedBTCDelegations__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.GetVerifiedBTCDelegations(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_ProgressTillFirstBlockTheNextEpoch__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.ProgressTillFirstBlockTheNextEpoch(t4)
	})
}

// skipping Fuzz_Nosy_BabylonAppDriver_SendTxWithMessagesSuccess__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_BabylonAppDriver_SendTxWithMsgsFromDriverAccount__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

func Fuzz_Nosy_BabylonAppDriver_WaitTillAllFpsJailed__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.WaitTillAllFpsJailed(t4)
	})
}

func Fuzz_Nosy_BabylonAppDriver_getDelegationWithStatus__(f *testing.F) {
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
		var dir string
		fill_err = tp.Fill(&dir)
		if fill_err != nil {
			return
		}
		var copyDir string
		fill_err = tp.Fill(&copyDir)
		if fill_err != nil {
			return
		}
		var t4 *testing.T
		fill_err = tp.Fill(&t4)
		if fill_err != nil {
			return
		}
		var status bstypes.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if t1 == nil || t4 == nil {
			return
		}

		d := NewBabylonAppDriver(t1, dir, copyDir)
		d.getDelegationWithStatus(t4, status)
	})
}

func Fuzz_Nosy_BlockReplayer_ReplayBlocks__(f *testing.F) {
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
		var nodeDir string
		fill_err = tp.Fill(&nodeDir)
		if fill_err != nil {
			return
		}
		var t3 *testing.T
		fill_err = tp.Fill(&t3)
		if fill_err != nil {
			return
		}
		var blocks []FinalizedBlock
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if t1 == nil || t3 == nil {
			return
		}

		r := NewBlockReplayer(t1, nodeDir)
		r.ReplayBlocks(t3, blocks)
	})
}

// skipping Fuzz_Nosy_AppOptionsMap_Get__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/test/replay.AppOptionsMap

func Fuzz_Nosy_DelegationInfosToBTCTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delInfos []*datagen.CreateDelegationInfo
		fill_err = tp.Fill(&delInfos)
		if fill_err != nil {
			return
		}

		DelegationInfosToBTCTx(delInfos)
	})
}

func Fuzz_Nosy_FpInfosToMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpInfos []*FinalityProviderInfo
		fill_err = tp.Fill(&fpInfos)
		if fill_err != nil {
			return
		}

		FpInfosToMsgs(fpInfos)
	})
}

func Fuzz_Nosy_GenerateNBTCDelegationsForFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var n uint32
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var senderAddress sdk.AccAddress
		fill_err = tp.Fill(&senderAddress)
		if fill_err != nil {
			return
		}
		var fpInfo *FinalityProviderInfo
		fill_err = tp.Fill(&fpInfo)
		if fill_err != nil {
			return
		}
		var params *bstypes.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if r == nil || t2 == nil || fpInfo == nil || params == nil {
			return
		}

		GenerateNBTCDelegationsForFinalityProvider(r, t2, n, senderAddress, fpInfo, params)
	})
}

func Fuzz_Nosy_GenerateNFinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var n uint32
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var senderAddress sdk.AccAddress
		fill_err = tp.Fill(&senderAddress)
		if fill_err != nil {
			return
		}
		if r == nil || t2 == nil {
			return
		}

		GenerateNFinalityProviders(r, t2, n, senderAddress)
	})
}

// skipping Fuzz_Nosy_MsgsToSdkMsg__ because parameters include func, chan, or unsupported interface: []M

func Fuzz_Nosy_ToCovenantSignaturesMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delInfos []*datagen.CreateDelegationInfo
		fill_err = tp.Fill(&delInfos)
		if fill_err != nil {
			return
		}

		ToCovenantSignaturesMsgs(delInfos)
	})
}

func Fuzz_Nosy_ToCreateBTCDelegationMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delInfos []*datagen.CreateDelegationInfo
		fill_err = tp.Fill(&delInfos)
		if fill_err != nil {
			return
		}

		ToCreateBTCDelegationMsgs(delInfos)
	})
}

func Fuzz_Nosy_blocksWithProofsToHeaderBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var blocks []*datagen.BlockWithProofs
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}

		blocksWithProofsToHeaderBytes(blocks)
	})
}

// skipping Fuzz_Nosy_createTx__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/client.TxConfig

func Fuzz_Nosy_getGenDoc__(f *testing.F) {
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
		var nodeDir string
		fill_err = tp.Fill(&nodeDir)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		getGenDoc(t1, nodeDir)
	})
}

// skipping Fuzz_Nosy_signVoteExtension__ because parameters include func, chan, or unsupported interface: github.com/cometbft/cometbft/crypto.PrivKey
