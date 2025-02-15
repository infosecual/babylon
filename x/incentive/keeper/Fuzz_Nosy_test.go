package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/math"
	finality_types "github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/babylonlabs-io/babylon/x/incentive/types"
	incentive_types "github.com/babylonlabs-io/babylon/x/incentive/types"
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

// skipping Fuzz_Nosy_RefundTxDecorator_CheckTxAndClearIndex__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Tx

// skipping Fuzz_Nosy_RefundTxDecorator_PostHandle__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Tx

func Fuzz_Nosy_Keeper_AddFinalityProviderRewardsForBtcDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var rwd sdk.Coins
		fill_err = tp.Fill(&rwd)
		if fill_err != nil {
			return
		}

		k.AddFinalityProviderRewardsForBtcDelegations(ctx, fp, rwd)
	})
}

func Fuzz_Nosy_Keeper_BTCStakingGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *incentive_types.QueryBTCStakingGaugeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.BTCStakingGauge(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_BtcDelegationActivated__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var sat uint64
		fill_err = tp.Fill(&sat)
		if fill_err != nil {
			return
		}

		k.BtcDelegationActivated(ctx, fp, del, sat)
	})
}

func Fuzz_Nosy_Keeper_BtcDelegationUnbonded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var sat uint64
		fill_err = tp.Fill(&sat)
		if fill_err != nil {
			return
		}

		k.BtcDelegationUnbonded(ctx, fp, del, sat)
	})
}

func Fuzz_Nosy_Keeper_CalculateBTCDelegationRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var endPeriod uint64
		fill_err = tp.Fill(&endPeriod)
		if fill_err != nil {
			return
		}

		k.CalculateBTCDelegationRewards(ctx, fp, del, endPeriod)
	})
}

func Fuzz_Nosy_Keeper_CalculateBTCDelegationRewardsAndSendToGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var endPeriod uint64
		fill_err = tp.Fill(&endPeriod)
		if fill_err != nil {
			return
		}

		k.CalculateBTCDelegationRewardsAndSendToGauge(ctx, fp, del, endPeriod)
	})
}

func Fuzz_Nosy_Keeper_DelegatorWithdrawAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *incentive_types.QueryDelegatorWithdrawAddressRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.DelegatorWithdrawAddress(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_FpSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.FpSlashed(ctx, fp)
	})
}

func Fuzz_Nosy_Keeper_GetBTCDelegationRewardsTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		k.GetBTCDelegationRewardsTracker(ctx, fp, del)
	})
}

func Fuzz_Nosy_Keeper_GetBTCStakingGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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

		k.GetBTCStakingGauge(ctx, height)
	})
}

func Fuzz_Nosy_Keeper_GetFinalityProviderCurrentRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.GetFinalityProviderCurrentRewards(ctx, fp)
	})
}

func Fuzz_Nosy_Keeper_GetFinalityProviderHistoricalRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}

		k.GetFinalityProviderHistoricalRewards(ctx, fp, period)
	})
}

func Fuzz_Nosy_Keeper_GetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.GetParams(ctx)
	})
}

func Fuzz_Nosy_Keeper_GetRewardGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		k.GetRewardGauge(ctx, sType, addr)
	})
}

func Fuzz_Nosy_Keeper_GetWithdrawAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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

		k.GetWithdrawAddr(ctx, addr)
	})
}

func Fuzz_Nosy_Keeper_HandleCoinsInFeeCollector__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.HandleCoinsInFeeCollector(ctx)
	})
}

func Fuzz_Nosy_Keeper_HasRefundableMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msgHash []byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}

		k.HasRefundableMsg(ctx, msgHash)
	})
}

func Fuzz_Nosy_Keeper_IncrementFinalityProviderPeriod__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.IncrementFinalityProviderPeriod(ctx, fp)
	})
}

// skipping Fuzz_Nosy_Keeper_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_Keeper_IterateBTCDelegationRewardsTracker__ because parameters include func, chan, or unsupported interface: func(fp github.com/cosmos/cosmos-sdk/sdk.AccAddress, del github.com/cosmos/cosmos-sdk/sdk.AccAddress) error

func Fuzz_Nosy_Keeper_Logger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.Logger(ctx)
	})
}

func Fuzz_Nosy_Keeper_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *incentive_types.QueryParamsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.Params(goCtx, req)
	})
}

// skipping Fuzz_Nosy_Keeper_RefundTx__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.FeeTx

func Fuzz_Nosy_Keeper_RemoveRefundableMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msgHash []byte
		fill_err = tp.Fill(&msgHash)
		if fill_err != nil {
			return
		}

		k.RemoveRefundableMsg(ctx, msgHash)
	})
}

func Fuzz_Nosy_Keeper_RewardBTCStaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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
		var dc *finality_types.VotingPowerDistCache
		fill_err = tp.Fill(&dc)
		if fill_err != nil {
			return
		}
		var voters map[string]struct{}
		fill_err = tp.Fill(&voters)
		if fill_err != nil {
			return
		}
		if dc == nil {
			return
		}

		k.RewardBTCStaking(ctx, height, dc, voters)
	})
}

func Fuzz_Nosy_Keeper_RewardGauges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *types.QueryRewardGaugesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		k.RewardGauges(goCtx, req)
	})
}

func Fuzz_Nosy_Keeper_SetBTCStakingGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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
		var gauge *types.Gauge
		fill_err = tp.Fill(&gauge)
		if fill_err != nil {
			return
		}
		if gauge == nil {
			return
		}

		k.SetBTCStakingGauge(ctx, height, gauge)
	})
}

func Fuzz_Nosy_Keeper_SetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var p incentive_types.Params
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		k.SetParams(ctx, p)
	})
}

func Fuzz_Nosy_Keeper_SetRewardGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var rg *types.RewardGauge
		fill_err = tp.Fill(&rg)
		if fill_err != nil {
			return
		}
		if rg == nil {
			return
		}

		k.SetRewardGauge(ctx, sType, addr, rg)
	})
}

func Fuzz_Nosy_Keeper_SetWithdrawAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
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
		var withdrawAddr sdk.AccAddress
		fill_err = tp.Fill(&withdrawAddr)
		if fill_err != nil {
			return
		}

		k.SetWithdrawAddr(ctx, addr, withdrawAddr)
	})
}

func Fuzz_Nosy_Keeper_accumulateBTCStakingReward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var btcStakingReward sdk.Coins
		fill_err = tp.Fill(&btcStakingReward)
		if fill_err != nil {
			return
		}

		k.accumulateBTCStakingReward(ctx, btcStakingReward)
	})
}

func Fuzz_Nosy_Keeper_accumulateRewardGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var reward sdk.Coins
		fill_err = tp.Fill(&reward)
		if fill_err != nil {
			return
		}

		k.accumulateRewardGauge(ctx, sType, addr, reward)
	})
}

func Fuzz_Nosy_Keeper_addDelegationSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		k.addDelegationSat(ctx, fp, del, amt)
	})
}

func Fuzz_Nosy_Keeper_addFinalityProviderStaked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		k.addFinalityProviderStaked(ctx, fp, amt)
	})
}

func Fuzz_Nosy_Keeper_btcDelegationModified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		k.btcDelegationModified(ctx, fp, del)
	})
}

// skipping Fuzz_Nosy_Keeper_btcDelegationModifiedWithPreInitDel__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, fp github.com/cosmos/cosmos-sdk/sdk.AccAddress, del github.com/cosmos/cosmos-sdk/sdk.AccAddress) error

func Fuzz_Nosy_Keeper_btcStakingGaugeStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.btcStakingGaugeStore(ctx)
	})
}

func Fuzz_Nosy_Keeper_calculateDelegationRewardsBetween__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var btcDelRwdTracker types.BTCDelegationRewardsTracker
		fill_err = tp.Fill(&btcDelRwdTracker)
		if fill_err != nil {
			return
		}
		var endingPeriod uint64
		fill_err = tp.Fill(&endingPeriod)
		if fill_err != nil {
			return
		}

		k.calculateDelegationRewardsBetween(ctx, fp, btcDelRwdTracker, endingPeriod)
	})
}

func Fuzz_Nosy_Keeper_deleteAllFromFinalityProviderRwd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.deleteAllFromFinalityProviderRwd(ctx, fp)
	})
}

func Fuzz_Nosy_Keeper_deleteBTCDelegatorToFP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.deleteBTCDelegatorToFP(ctx, del, fp)
	})
}

func Fuzz_Nosy_Keeper_deleteFinalityProviderCurrentRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.deleteFinalityProviderCurrentRewards(ctx, fp)
	})
}

func Fuzz_Nosy_Keeper_deleteKeysFromBTCDelegationRewardsTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var delKeys [][]byte
		fill_err = tp.Fill(&delKeys)
		if fill_err != nil {
			return
		}

		k.deleteKeysFromBTCDelegationRewardsTracker(ctx, fp, delKeys)
	})
}

func Fuzz_Nosy_Keeper_initializeBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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

		k.initializeBTCDelegation(ctx, fp, del)
	})
}

func Fuzz_Nosy_Keeper_initializeFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.initializeFinalityProvider(ctx, fp)
	})
}

// skipping Fuzz_Nosy_Keeper_iterBtcDelegationsByDelegator__ because parameters include func, chan, or unsupported interface: func(del github.com/cosmos/cosmos-sdk/sdk.AccAddress, fp github.com/cosmos/cosmos-sdk/sdk.AccAddress) error

func Fuzz_Nosy_Keeper_rewardGaugeStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}

		k.rewardGaugeStore(ctx, sType)
	})
}

func Fuzz_Nosy_Keeper_sendAllBtcDelegationTypeToRewardsGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}

		k.sendAllBtcDelegationTypeToRewardsGauge(ctx, sType, del)
	})
}

func Fuzz_Nosy_Keeper_sendAllBtcRewardsToGauge__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}

		k.sendAllBtcRewardsToGauge(ctx, del)
	})
}

func Fuzz_Nosy_Keeper_setBTCDelegationRewardsTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var rwd types.BTCDelegationRewardsTracker
		fill_err = tp.Fill(&rwd)
		if fill_err != nil {
			return
		}

		k.setBTCDelegationRewardsTracker(ctx, fp, del, rwd)
	})
}

func Fuzz_Nosy_Keeper_setBTCDelegatorToFP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}

		k.setBTCDelegatorToFP(ctx, del, fp)
	})
}

func Fuzz_Nosy_Keeper_setFinalityProviderCurrentRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var rwd types.FinalityProviderCurrentRewards
		fill_err = tp.Fill(&rwd)
		if fill_err != nil {
			return
		}

		k.setFinalityProviderCurrentRewards(ctx, fp, rwd)
	})
}

func Fuzz_Nosy_Keeper_setFinalityProviderHistoricalRewards__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var period uint64
		fill_err = tp.Fill(&period)
		if fill_err != nil {
			return
		}
		var rwd types.FinalityProviderHistoricalRewards
		fill_err = tp.Fill(&rwd)
		if fill_err != nil {
			return
		}

		k.setFinalityProviderHistoricalRewards(ctx, fp, period, rwd)
	})
}

func Fuzz_Nosy_Keeper_storeBTCDelegatorToFp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var del sdk.AccAddress
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}

		k.storeBTCDelegatorToFp(ctx, del)
	})
}

func Fuzz_Nosy_Keeper_subDelegationSat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
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
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		k.subDelegationSat(ctx, fp, del, amt)
	})
}

func Fuzz_Nosy_Keeper_subFinalityProviderStaked__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var fp sdk.AccAddress
		fill_err = tp.Fill(&fp)
		if fill_err != nil {
			return
		}
		var amt math.Int
		fill_err = tp.Fill(&amt)
		if fill_err != nil {
			return
		}

		k.subFinalityProviderStaked(ctx, fp, amt)
	})
}

func Fuzz_Nosy_Keeper_withdrawReward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var sType incentive_types.StakeholderType
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}

		k.withdrawReward(ctx, sType, addr)
	})
}

func Fuzz_Nosy_msgServer_SetWithdrawAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var msg *types.MsgSetWithdrawAddress
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		ms.SetWithdrawAddress(ctx, msg)
	})
}

func Fuzz_Nosy_msgServer_UpdateParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *types.MsgUpdateParams
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.UpdateParams(goCtx, req)
	})
}

func Fuzz_Nosy_msgServer_WithdrawReward__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ms msgServer
		fill_err = tp.Fill(&ms)
		if fill_err != nil {
			return
		}
		var goCtx context.Context
		fill_err = tp.Fill(&goCtx)
		if fill_err != nil {
			return
		}
		var req *types.MsgWithdrawReward
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		ms.WithdrawReward(goCtx, req)
	})
}

func Fuzz_Nosy_convertToRewardGaugesResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rgMap map[string]*types.RewardGauge
		fill_err = tp.Fill(&rgMap)
		if fill_err != nil {
			return
		}

		convertToRewardGaugesResponse(rgMap)
	})
}
