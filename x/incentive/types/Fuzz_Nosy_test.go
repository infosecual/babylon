package types

import (
	"context"
	"testing"

	"cosmossdk.io/math"
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

func Fuzz_Nosy_BTCDelegationRewardsTracker_AddTotalActiveSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		f1 := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		f1.AddTotalActiveSat(amt)
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		_x1 := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_GetStartPeriodCumulativeReward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.GetStartPeriodCumulativeReward()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		_x1 := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.Reset()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.Size()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.String()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_SubTotalActiveSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		f1 := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		f1.SubTotalActiveSat(amt)
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
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

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCDelegationRewardsTracker_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCDelegationRewardsTracker_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCDelegationRewardsTracker_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var startPeriod uint64
		fill_err = tp.Fill(&startPeriod)
		if fill_err != nil {
			return
		}
		var totalSat math.Int
		fill_err = tp.Fill(&totalSat)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewBTCDelegationRewardsTracker(startPeriod, totalSat)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCStakingGaugeResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_GetCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCoins()
	})
}

func Fuzz_Nosy_BTCStakingGaugeResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

// skipping Fuzz_Nosy_BTCStakingGaugeResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCStakingGaugeResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_BTCStakingGaugeResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BTCStakingGaugeResponse
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

func Fuzz_Nosy_FinalityProviderCurrentRewards_AddRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var coinsToAdd sdk.Coins
		fill_err = tp.Fill(&coinsToAdd)
		if fill_err != nil {
			return
		}

		f1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		f1.AddRewards(coinsToAdd)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_AddTotalActiveSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		f1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		f1.AddTotalActiveSat(amt)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		_x1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_GetCurrentRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.GetCurrentRewards()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_GetPeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.GetPeriod()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.Marshal()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		_x1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.Reset()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.Size()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.String()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_SubRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var coinsToSubtract sdk.Coins
		fill_err = tp.Fill(&coinsToSubtract)
		if fill_err != nil {
			return
		}

		f1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		f1.SubRewards(coinsToSubtract)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_SubTotalActiveSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		f1 := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		f1.SubTotalActiveSat(amt)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
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

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_FinalityProviderCurrentRewards_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_FinalityProviderCurrentRewards_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_FinalityProviderCurrentRewards_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentRewards sdk.Coins
		fill_err = tp.Fill(&currentRewards)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var totalActiveSatFP math.Int
		fill_err = tp.Fill(&totalActiveSatFP)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderCurrentRewards(currentRewards, period, totalActiveSatFP)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		_x1 := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_GetCumulativeRewardsPerSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.GetCumulativeRewardsPerSat()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.Marshal()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		_x1 := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.Reset()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.Size()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.String()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
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

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_FinalityProviderHistoricalRewards_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_FinalityProviderHistoricalRewards_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_FinalityProviderHistoricalRewards_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cumulativeRewardsPerSat sdk.Coins
		fill_err = tp.Fill(&cumulativeRewardsPerSat)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewFinalityProviderHistoricalRewards(cumulativeRewardsPerSat)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_Gauge_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		_x1 := NewGauge(coins...)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Gauge_GetCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.GetCoins()
	})
}

func Fuzz_Nosy_Gauge_GetCoinsPortion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var portion math.LegacyDec
		fill_err = tp.Fill(&portion)
		if fill_err != nil {
			return
		}

		g := NewGauge(coins...)
		g.GetCoinsPortion(portion)
	})
}

func Fuzz_Nosy_Gauge_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.Marshal()
	})
}

func Fuzz_Nosy_Gauge_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Gauge_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Gauge_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		_x1 := NewGauge(coins...)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Gauge_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.Reset()
	})
}

func Fuzz_Nosy_Gauge_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.Size()
	})
}

func Fuzz_Nosy_Gauge_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.String()
	})
}

func Fuzz_Nosy_Gauge_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Gauge_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_Gauge_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
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

		m := NewGauge(coins...)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Gauge_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Gauge_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_Gauge_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewGauge(coins...)
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

func Fuzz_Nosy_MockAccountKeeper_EXPECT__(f *testing.F) {
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

		m := NewMockAccountKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockAccountKeeper_GetAccount__(f *testing.F) {
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
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockAccountKeeper(ctrl)
		m.GetAccount(ctx, addr)
	})
}

func Fuzz_Nosy_MockAccountKeeper_GetModuleAccount__(f *testing.F) {
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
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockAccountKeeper(ctrl)
		m.GetModuleAccount(ctx, name)
	})
}

// skipping Fuzz_Nosy_MockAccountKeeperMockRecorder_GetAccount__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockAccountKeeperMockRecorder_GetModuleAccount__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBankKeeper_EXPECT__(f *testing.F) {
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

		m := NewMockBankKeeper(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBankKeeper_GetAllBalances__(f *testing.F) {
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
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBankKeeper(ctrl)
		m.GetAllBalances(ctx, addr)
	})
}

func Fuzz_Nosy_MockBankKeeper_SendCoinsFromModuleToAccount__(f *testing.F) {
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
		var senderModule string
		fill_err = tp.Fill(&senderModule)
		if fill_err != nil {
			return
		}
		var recipientAddr sdk.AccAddress
		fill_err = tp.Fill(&recipientAddr)
		if fill_err != nil {
			return
		}
		var amt sdk.Coins
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBankKeeper(ctrl)
		m.SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt)
	})
}

func Fuzz_Nosy_MockBankKeeper_SendCoinsFromModuleToModule__(f *testing.F) {
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
		var senderModule string
		fill_err = tp.Fill(&senderModule)
		if fill_err != nil {
			return
		}
		var recipientModule string
		fill_err = tp.Fill(&recipientModule)
		if fill_err != nil {
			return
		}
		var amt sdk.Coins
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBankKeeper(ctrl)
		m.SendCoinsFromModuleToModule(ctx, senderModule, recipientModule, amt)
	})
}

func Fuzz_Nosy_MockBankKeeper_SpendableCoins__(f *testing.F) {
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
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBankKeeper(ctrl)
		m.SpendableCoins(ctx, addr)
	})
}

// skipping Fuzz_Nosy_MockBankKeeperMockRecorder_GetAllBalances__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBankKeeperMockRecorder_SendCoinsFromModuleToAccount__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBankKeeperMockRecorder_SendCoinsFromModuleToModule__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBankKeeperMockRecorder_SpendableCoins__ because parameters include func, chan, or unsupported interface: interface{}

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

// skipping Fuzz_Nosy_MockEpochingKeeperMockRecorder_GetEpoch__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MsgSetWithdrawAddress_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_GetDelegatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorAddress()
	})
}

func Fuzz_Nosy_MsgSetWithdrawAddress_GetWithdrawAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetWithdrawAddress()
	})
}

func Fuzz_Nosy_MsgSetWithdrawAddress_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

// skipping Fuzz_Nosy_MsgSetWithdrawAddress_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgSetWithdrawAddress_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddress_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddress
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

// skipping Fuzz_Nosy_MsgSetWithdrawAddressResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgSetWithdrawAddressResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgSetWithdrawAddressResponse
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

func Fuzz_Nosy_MsgWithdrawReward_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_GetType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetType()
	})
}

func Fuzz_Nosy_MsgWithdrawReward_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

// skipping Fuzz_Nosy_MsgWithdrawReward_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWithdrawReward_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawReward_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawReward
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_GetCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCoins()
	})
}

func Fuzz_Nosy_MsgWithdrawRewardResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

// skipping Fuzz_Nosy_MsgWithdrawRewardResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWithdrawRewardResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_MsgWithdrawRewardResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWithdrawRewardResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

// skipping Fuzz_Nosy_QueryBTCStakingGaugeRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeRequest
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_GetGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetGauge()
	})
}

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

// skipping Fuzz_Nosy_QueryBTCStakingGaugeResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryBTCStakingGaugeResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBTCStakingGaugeResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

// skipping Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressRequest
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

// skipping Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryDelegatorWithdrawAddressResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegatorWithdrawAddressResponse
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

func Fuzz_Nosy_QueryRewardGaugesRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

// skipping Fuzz_Nosy_QueryRewardGaugesRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRewardGaugesRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesRequest
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

func Fuzz_Nosy_QueryRewardGaugesResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_GetRewardGauges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRewardGauges()
	})
}

func Fuzz_Nosy_QueryRewardGaugesResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

// skipping Fuzz_Nosy_QueryRewardGaugesResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRewardGaugesResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_QueryRewardGaugesResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRewardGaugesResponse
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

func Fuzz_Nosy_RewardGauge_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 []sdk.Coin
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var c2 sdk.Coins
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}

		rg := NewRewardGauge(c1...)
		rg.Add(c2)
	})
}

func Fuzz_Nosy_RewardGauge_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		_x1 := NewRewardGauge(coins...)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_RewardGauge_GetCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.GetCoins()
	})
}

func Fuzz_Nosy_RewardGauge_GetWithdrawableCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		rg := NewRewardGauge(coins...)
		rg.GetWithdrawableCoins()
	})
}

func Fuzz_Nosy_RewardGauge_GetWithdrawnCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.GetWithdrawnCoins()
	})
}

func Fuzz_Nosy_RewardGauge_IsFullyWithdrawn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		rg := NewRewardGauge(coins...)
		rg.IsFullyWithdrawn()
	})
}

func Fuzz_Nosy_RewardGauge_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.Marshal()
	})
}

func Fuzz_Nosy_RewardGauge_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_RewardGauge_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_RewardGauge_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		_x1 := NewRewardGauge(coins...)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_RewardGauge_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.Reset()
	})
}

func Fuzz_Nosy_RewardGauge_SetFullyWithdrawn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		rg := NewRewardGauge(coins...)
		rg.SetFullyWithdrawn()
	})
}

func Fuzz_Nosy_RewardGauge_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.Size()
	})
}

func Fuzz_Nosy_RewardGauge_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.String()
	})
}

func Fuzz_Nosy_RewardGauge_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_RewardGauge_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_RewardGauge_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
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

		m := NewRewardGauge(coins...)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_RewardGauge_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RewardGauge_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_RewardGauge_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var coins []sdk.Coin
		fill_err = tp.Fill(&coins)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewRewardGauge(coins...)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_RewardGaugesResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_GetCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCoins()
	})
}

func Fuzz_Nosy_RewardGaugesResponse_GetWithdrawnCoins__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetWithdrawnCoins()
	})
}

func Fuzz_Nosy_RewardGaugesResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

// skipping Fuzz_Nosy_RewardGaugesResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RewardGaugesResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_RewardGaugesResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RewardGaugesResponse
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

func Fuzz_Nosy_UnimplementedMsgServer_SetWithdrawAddress__(f *testing.F) {
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
		var req *MsgSetWithdrawAddress
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.SetWithdrawAddress(ctx, req)
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

func Fuzz_Nosy_UnimplementedMsgServer_WithdrawReward__(f *testing.F) {
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
		var req *MsgWithdrawReward
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WithdrawReward(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_BTCStakingGauge__(f *testing.F) {
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
		var req *QueryBTCStakingGaugeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BTCStakingGauge(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_DelegatorWithdrawAddress__(f *testing.F) {
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
		var req *QueryDelegatorWithdrawAddressRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.DelegatorWithdrawAddress(ctx, req)
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

func Fuzz_Nosy_UnimplementedQueryServer_RewardGauges__(f *testing.F) {
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
		var req *QueryRewardGaugesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.RewardGauges(ctx, req)
	})
}

// skipping Fuzz_Nosy_msgClient_SetWithdrawAddress__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WithdrawReward__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_BTCStakingGauge__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_DelegatorWithdrawAddress__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Params__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_RewardGauges__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_AccountKeeper_GetAccount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.AccountKeeper

// skipping Fuzz_Nosy_AccountKeeper_GetModuleAccount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.AccountKeeper

// skipping Fuzz_Nosy_BankKeeper_GetAllBalances__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.BankKeeper

// skipping Fuzz_Nosy_BankKeeper_SendCoinsFromModuleToAccount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.BankKeeper

// skipping Fuzz_Nosy_BankKeeper_SendCoinsFromModuleToModule__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.BankKeeper

// skipping Fuzz_Nosy_BankKeeper_SpendableCoins__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.BankKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.EpochingKeeper

// skipping Fuzz_Nosy_MsgClient_SetWithdrawAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WithdrawReward__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgClient

// skipping Fuzz_Nosy_MsgServer_SetWithdrawAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WithdrawReward__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.MsgServer

// skipping Fuzz_Nosy_QueryClient_BTCStakingGauge__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_DelegatorWithdrawAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_RewardGauges__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_BTCStakingGauge__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_DelegatorWithdrawAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_RewardGauges__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/incentive/types.QueryServer

func Fuzz_Nosy_RewardGaugesResponse_ToRewardGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rgr RewardGaugesResponse
		fill_err = tp.Fill(&rgr)
		if fill_err != nil {
			return
		}

		rgr.ToRewardGauge()
	})
}

func Fuzz_Nosy_StakeholderType_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stBytes []byte) {
		st, err := NewStakeHolderType(stBytes)
		if err != nil {
			return
		}
		st.Bytes()
	})
}

func Fuzz_Nosy_StakeholderType_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stBytes []byte) {
		st, err := NewStakeHolderType(stBytes)
		if err != nil {
			return
		}
		st.String()
	})
}

func Fuzz_Nosy_GetWithdrawAddrKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var delAddr sdk.AccAddress
		fill_err = tp.Fill(&delAddr)
		if fill_err != nil {
			return
		}

		GetWithdrawAddrKey(delAddr)
	})
}

// skipping Fuzz_Nosy_HashMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

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

// skipping Fuzz_Nosy__Msg_SetWithdrawAddress_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_UpdateParams_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WithdrawReward_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_BTCStakingGauge_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_DelegatorWithdrawAddress_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Params_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_RewardGauges_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintGenesis(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintIncentive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintIncentive(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintParams(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintQuery(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintRewards(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintTx(dAtA, offset, v)
	})
}

// skipping Fuzz_Nosy_local_request_Query_BTCStakingGauge_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_DelegatorWithdrawAddress_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_RewardGauges_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_BTCStakingGauge_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_DelegatorWithdrawAddress_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_RewardGauges_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipGenesis(dAtA)
	})
}

func Fuzz_Nosy_skipIncentive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipIncentive(dAtA)
	})
}

func Fuzz_Nosy_skipParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipParams(dAtA)
	})
}

func Fuzz_Nosy_skipQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipQuery(dAtA)
	})
}

func Fuzz_Nosy_skipRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipRewards(dAtA)
	})
}

func Fuzz_Nosy_skipTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipTx(dAtA)
	})
}

func Fuzz_Nosy_sovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovGenesis(x)
	})
}

func Fuzz_Nosy_sovIncentive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovIncentive(x)
	})
}

func Fuzz_Nosy_sovParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovParams(x)
	})
}

func Fuzz_Nosy_sovQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovQuery(x)
	})
}

func Fuzz_Nosy_sovRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovRewards(x)
	})
}

func Fuzz_Nosy_sovTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovTx(x)
	})
}

func Fuzz_Nosy_sozGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozGenesis(x)
	})
}

func Fuzz_Nosy_sozIncentive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozIncentive(x)
	})
}

func Fuzz_Nosy_sozParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozParams(x)
	})
}

func Fuzz_Nosy_sozQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozQuery(x)
	})
}

func Fuzz_Nosy_sozRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozRewards(x)
	})
}

func Fuzz_Nosy_sozTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozTx(x)
	})
}
