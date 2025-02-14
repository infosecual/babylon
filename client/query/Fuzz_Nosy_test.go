package query

import (
	"testing"

	"github.com/babylonlabs-io/babylon/client/config"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func Fuzz_Nosy_QueryClient_ActivatedHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ActivatedHeight()
	})
}

func Fuzz_Nosy_QueryClient_ActiveFinalityProvidersAtHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ActiveFinalityProvidersAtHeight(height, pagination)
	})
}

func Fuzz_Nosy_QueryClient_BTCBaseHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCBaseHeader()
	})
}

func Fuzz_Nosy_QueryClient_BTCCheckpointInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCCheckpointInfo(epochNumber)
	})
}

func Fuzz_Nosy_QueryClient_BTCCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCCheckpointParams()
	})
}

func Fuzz_Nosy_QueryClient_BTCCheckpointsInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCCheckpointsInfo(pagination)
	})
}

func Fuzz_Nosy_QueryClient_BTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCDelegation(stakingTxHashHex)
	})
}

func Fuzz_Nosy_QueryClient_BTCDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var status bstypes.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCDelegations(status, pagination)
	})
}

func Fuzz_Nosy_QueryClient_BTCHeaderChainTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCHeaderChainTip()
	})
}

func Fuzz_Nosy_QueryClient_BTCMainChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCMainChain(pagination)
	})
}

func Fuzz_Nosy_QueryClient_BTCStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCStakingParams()
	})
}

func Fuzz_Nosy_QueryClient_BTCStakingParamsByVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BTCStakingParamsByVersion(version)
	})
}

func Fuzz_Nosy_QueryClient_Block__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.Block(height)
	})
}

func Fuzz_Nosy_QueryClient_BlockResults__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BlockResults(height)
	})
}

func Fuzz_Nosy_QueryClient_BlockSearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var events []string
		fill_err = tp.Fill(&events)
		if fill_err != nil {
			return
		}
		var page *int
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		var perPage *int
		fill_err = tp.Fill(&perPage)
		if fill_err != nil {
			return
		}
		var orderBy string
		fill_err = tp.Fill(&orderBy)
		if fill_err != nil {
			return
		}
		if cfg == nil || page == nil || perPage == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BlockSearch(events, page, perPage, orderBy)
	})
}

func Fuzz_Nosy_QueryClient_BlsPublicKeyList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.BlsPublicKeyList(epochNumber, pagination)
	})
}

func Fuzz_Nosy_QueryClient_ContainsBTCBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var blockHash *chainhash.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if cfg == nil || blockHash == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ContainsBTCBlock(blockHash)
	})
}

func Fuzz_Nosy_QueryClient_CurrentEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.CurrentEpoch()
	})
}

func Fuzz_Nosy_QueryClient_DelegationLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var delegator string
		fill_err = tp.Fill(&delegator)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.DelegationLifecycle(delegator)
	})
}

func Fuzz_Nosy_QueryClient_EndedEpochBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.EndedEpochBTCHeight(epochNum)
	})
}

func Fuzz_Nosy_QueryClient_EpochStatusCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var epochCount uint64
		fill_err = tp.Fill(&epochCount)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.EpochStatusCount(epochCount)
	})
}

func Fuzz_Nosy_QueryClient_EpochingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.EpochingParams()
	})
}

func Fuzz_Nosy_QueryClient_EpochsInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.EpochsInfo(pagination)
	})
}

func Fuzz_Nosy_QueryClient_FinalityParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.FinalityParams()
	})
}

func Fuzz_Nosy_QueryClient_FinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.FinalityProvider(fpBtcPkHex)
	})
}

func Fuzz_Nosy_QueryClient_FinalityProviderDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.FinalityProviderDelegations(fpBtcPkHex, pagination)
	})
}

func Fuzz_Nosy_QueryClient_FinalityProviderPowerAtHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.FinalityProviderPowerAtHeight(fpBtcPkHex, height)
	})
}

func Fuzz_Nosy_QueryClient_FinalityProviders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.FinalityProviders(pagination)
	})
}

func Fuzz_Nosy_QueryClient_GetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.GetBlock(height)
	})
}

func Fuzz_Nosy_QueryClient_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.GetStatus()
	})
}

func Fuzz_Nosy_QueryClient_GetTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var hash []byte
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.GetTx(hash)
	})
}

func Fuzz_Nosy_QueryClient_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.IsRunning()
	})
}

func Fuzz_Nosy_QueryClient_LatestEpochFromStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var status checkpointingtypes.CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.LatestEpochFromStatus(status)
	})
}

func Fuzz_Nosy_QueryClient_LatestEpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var endEpoch uint64
		fill_err = tp.Fill(&endEpoch)
		if fill_err != nil {
			return
		}
		var epochCount uint64
		fill_err = tp.Fill(&epochCount)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.LatestEpochMsgs(endEpoch, epochCount, pagination)
	})
}

func Fuzz_Nosy_QueryClient_ListBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var status types.QueriedBlockStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ListBlocks(status, pagination)
	})
}

func Fuzz_Nosy_QueryClient_ListEvidences__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ListEvidences(startHeight, pagination)
	})
}

func Fuzz_Nosy_QueryClient_ListPubRandCommit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ListPubRandCommit(fpBtcPkHex, pagination)
	})
}

// skipping Fuzz_Nosy_QueryClient_QueryBTCCheckpoint__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/btccheckpoint/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryBTCLightclient__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryBTCStaking__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/btcstaking/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryCheckpointing__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryEpoching__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryFinality__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/finality/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryIncentive__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/incentive/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryMonitor__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/babylonlabs-io/babylon/x/monitor/types.QueryClient) error

// skipping Fuzz_Nosy_QueryClient_QueryStaking__ because parameters include func, chan, or unsupported interface: func(ctx context.Context, queryClient github.com/cosmos/cosmos-sdk/x/staking/types.QueryClient) error

func Fuzz_Nosy_QueryClient_RawCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.RawCheckpoint(epochNumber)
	})
}

func Fuzz_Nosy_QueryClient_RawCheckpointList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var status checkpointingtypes.CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.RawCheckpointList(status, pagination)
	})
}

func Fuzz_Nosy_QueryClient_RawCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if cfg == nil || pagination == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.RawCheckpoints(pagination)
	})
}

func Fuzz_Nosy_QueryClient_ReportedCheckpointBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var hashStr string
		fill_err = tp.Fill(&hashStr)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.ReportedCheckpointBTCHeight(hashStr)
	})
}

func Fuzz_Nosy_QueryClient_RewardGauges__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.RewardGauges(address)
	})
}

func Fuzz_Nosy_QueryClient_StakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.StakingParams()
	})
}

func Fuzz_Nosy_QueryClient_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.Start()
	})
}

func Fuzz_Nosy_QueryClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.Stop()
	})
}

func Fuzz_Nosy_QueryClient_Subscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		var q3 string
		fill_err = tp.Fill(&q3)
		if fill_err != nil {
			return
		}
		var outCapacity []int
		fill_err = tp.Fill(&outCapacity)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.Subscribe(subscriber, q3, outCapacity...)
	})
}

func Fuzz_Nosy_QueryClient_TxSearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var events []string
		fill_err = tp.Fill(&events)
		if fill_err != nil {
			return
		}
		var prove bool
		fill_err = tp.Fill(&prove)
		if fill_err != nil {
			return
		}
		var page *int
		fill_err = tp.Fill(&page)
		if fill_err != nil {
			return
		}
		var perPage *int
		fill_err = tp.Fill(&perPage)
		if fill_err != nil {
			return
		}
		var orderBy string
		fill_err = tp.Fill(&orderBy)
		if fill_err != nil {
			return
		}
		if cfg == nil || page == nil || perPage == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.TxSearch(events, prove, page, perPage, orderBy)
	})
}

func Fuzz_Nosy_QueryClient_Unsubscribe__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		var q3 string
		fill_err = tp.Fill(&q3)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.Unsubscribe(subscriber, q3)
	})
}

func Fuzz_Nosy_QueryClient_UnsubscribeAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.UnsubscribeAll(subscriber)
	})
}

func Fuzz_Nosy_QueryClient_VotesAtHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.VotesAtHeight(height)
	})
}

func Fuzz_Nosy_QueryClient_getQueryContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.BabylonQueryConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		c, err := New(cfg)
		if err != nil {
			return
		}
		c.getQueryContext()
	})
}
