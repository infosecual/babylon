package mocks

import (
	context "context"
	"testing"

	//"cosmossdk.io/x/staking/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	epochtypes "github.com/babylonlabs-io/babylon/x/epoching/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func Fuzz_Nosy_MockCheckpointingHooks_AfterBlsKeyRegistered__(f *testing.F) {
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
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterBlsKeyRegistered(ctx, valAddr)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_AfterRawCheckpointBlsSigVerified__(f *testing.F) {
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
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ctrl == nil || ckpt == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterRawCheckpointBlsSigVerified(ctx, ckpt)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_AfterRawCheckpointConfirmed__(f *testing.F) {
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
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterRawCheckpointConfirmed(ctx, epoch)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_AfterRawCheckpointFinalized__(f *testing.F) {
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
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterRawCheckpointFinalized(ctx, epoch)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_AfterRawCheckpointForgotten__(f *testing.F) {
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
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ctrl == nil || ckpt == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterRawCheckpointForgotten(ctx, ckpt)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_AfterRawCheckpointSealed__(f *testing.F) {
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
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingHooks(ctrl)
		m.AfterRawCheckpointSealed(ctx, epoch)
	})
}

func Fuzz_Nosy_MockCheckpointingHooks_EXPECT__(f *testing.F) {
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

		m := NewMockCheckpointingHooks(ctrl)
		m.EXPECT()
	})
}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterBlsKeyRegistered__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterRawCheckpointBlsSigVerified__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterRawCheckpointConfirmed__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterRawCheckpointFinalized__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterRawCheckpointForgotten__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingHooksMockRecorder_AfterRawCheckpointSealed__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockCheckpointingKeeper_EXPECT__(f *testing.F) {
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

		m := NewMockCheckpointingKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_GetBlsPubKey__(f *testing.F) {
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
		var address sdk.ValAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.GetBlsPubKey(ctx, address)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_GetEpoch__(f *testing.F) {
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

		m := NewMockCheckpointingKeeper(ctrl)
		m.GetEpoch(ctx)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_GetPubKeyByConsAddr__(f *testing.F) {
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
		var arg0 context.Context
		fill_err = tp.Fill(&arg0)
		if fill_err != nil {
			return
		}
		var arg1 sdk.ConsAddress
		fill_err = tp.Fill(&arg1)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.GetPubKeyByConsAddr(arg0, arg1)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_GetTotalVotingPower__(f *testing.F) {
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
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.GetTotalVotingPower(ctx, epochNumber)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_GetValidatorSet__(f *testing.F) {
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
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.GetValidatorSet(ctx, epochNumber)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_SealCheckpoint__(f *testing.F) {
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
		var ckptWithMeta *checkpointingtypes.RawCheckpointWithMeta
		fill_err = tp.Fill(&ckptWithMeta)
		if fill_err != nil {
			return
		}
		if ctrl == nil || ckptWithMeta == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.SealCheckpoint(ctx, ckptWithMeta)
	})
}

func Fuzz_Nosy_MockCheckpointingKeeper_VerifyBLSSig__(f *testing.F) {
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
		var sig *checkpointingtypes.BlsSig
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if ctrl == nil || sig == nil {
			return
		}

		m := NewMockCheckpointingKeeper(ctrl)
		m.VerifyBLSSig(ctx, sig)
	})
}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_GetBlsPubKey__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_GetEpoch__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_GetTotalVotingPower__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_GetValidatorSet__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_SealCheckpoint__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockCheckpointingKeeperMockRecorder_VerifyBLSSig__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockEpochingKeeper_CheckMsgCreateValidator__(f *testing.F) {
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
		var msg *stakingtypes.MsgCreateValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if ctrl == nil || msg == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.CheckMsgCreateValidator(ctx, msg)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_EXPECT__(f *testing.F) {
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

		m := NewMockEpochingKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockEpochingKeeper_EnqueueMsg__(f *testing.F) {
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
		var msg epochtypes.QueuedMessage
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.EnqueueMsg(ctx, msg)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetEpoch__(f *testing.F) {
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

		m := NewMockEpochingKeeper(ctrl)
		m.GetEpoch(ctx)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetEpochNumByHeight__(f *testing.F) {
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
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.GetEpochNumByHeight(ctx, height)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetPubKeyByConsAddr__(f *testing.F) {
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
		var consAddr sdk.ConsAddress
		fill_err = tp.Fill(&consAddr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.GetPubKeyByConsAddr(ctx, consAddr)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetTotalVotingPower__(f *testing.F) {
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
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.GetTotalVotingPower(ctx, epochNumber)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetValidator__(f *testing.F) {
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
		var addr sdk.ValAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.GetValidator(ctx, addr)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_GetValidatorSet__(f *testing.F) {
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
		var epochNumer uint64
		fill_err = tp.Fill(&epochNumer)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.GetValidatorSet(ctx, epochNumer)
	})
}

func Fuzz_Nosy_MockEpochingKeeper_StkMsgCreateValidator__(f *testing.F) {
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
		var msg *stakingtypes.MsgCreateValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if ctrl == nil || msg == nil {
			return
		}

		m := NewMockEpochingKeeper(ctrl)
		m.StkMsgCreateValidator(ctx, msg)
	})
}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_CheckMsgCreateValidator__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_EnqueueMsg__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetEpoch__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetEpochNumByHeight__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetTotalVotingPower__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetValidator__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetValidatorSet__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_StkMsgCreateValidator__ because parameters include func, chan, or unsupported interface: interface{}
