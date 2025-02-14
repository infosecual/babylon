package configurer

import (
	"testing"

	"github.com/babylonlabs-io/babylon/test/e2e/configurer/chain"
	"github.com/babylonlabs-io/babylon/test/e2e/initialization"
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

func Fuzz_Nosy_CurrentBranchConfigurer_ConfigureChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *CurrentBranchConfigurer
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		var chainConfig *chain.Config
		fill_err = tp.Fill(&chainConfig)
		if fill_err != nil {
			return
		}
		if cb == nil || chainConfig == nil {
			return
		}

		cb.ConfigureChain(chainConfig)
	})
}

func Fuzz_Nosy_CurrentBranchConfigurer_ConfigureChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *CurrentBranchConfigurer
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.ConfigureChains()
	})
}

func Fuzz_Nosy_CurrentBranchConfigurer_RunSetup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cb *CurrentBranchConfigurer
		fill_err = tp.Fill(&cb)
		if fill_err != nil {
			return
		}
		if cb == nil {
			return
		}

		cb.RunSetup()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_ConfigureChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		var chainConfig *chain.Config
		fill_err = tp.Fill(&chainConfig)
		if fill_err != nil {
			return
		}
		if uc == nil || chainConfig == nil {
			return
		}

		uc.ConfigureChain(chainConfig)
	})
}

func Fuzz_Nosy_UpgradeConfigurer_ConfigureChains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.ConfigureChains()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_CreatePreUpgradeState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.CreatePreUpgradeState()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_ParseGovPropFromFile__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.ParseGovPropFromFile()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_RunSetup__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.RunSetup()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_RunUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.RunUpgrade()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_SetGovPropUpgradeHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		var newUpgradeHeight int64
		fill_err = tp.Fill(&newUpgradeHeight)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.SetGovPropUpgradeHeight(newUpgradeHeight)
	})
}

func Fuzz_Nosy_UpgradeConfigurer_UpgradeFilePath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.UpgradeFilePath()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_runForkUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.runForkUpgrade()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_runProposalUpgrade__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		if uc == nil {
			return
		}

		uc.runProposalUpgrade()
	})
}

func Fuzz_Nosy_UpgradeConfigurer_upgradeContainers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var uc *UpgradeConfigurer
		fill_err = tp.Fill(&uc)
		if fill_err != nil {
			return
		}
		var chainConfig *chain.Config
		fill_err = tp.Fill(&chainConfig)
		if fill_err != nil {
			return
		}
		var propHeight int64
		fill_err = tp.Fill(&propHeight)
		if fill_err != nil {
			return
		}
		if uc == nil || chainConfig == nil {
			return
		}

		uc.upgradeContainers(chainConfig, propHeight)
	})
}

func Fuzz_Nosy_baseConfigurer_ClearResources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.ClearResources()
	})
}

func Fuzz_Nosy_baseConfigurer_GetChainConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chainIndex int
		fill_err = tp.Fill(&chainIndex)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.GetChainConfig(chainIndex)
	})
}

func Fuzz_Nosy_baseConfigurer_InstantiateBabylonContract__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.InstantiateBabylonContract()
	})
}

func Fuzz_Nosy_baseConfigurer_RunCosmosRelayerIBC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.RunCosmosRelayerIBC()
	})
}

func Fuzz_Nosy_baseConfigurer_RunHermesRelayerIBC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.RunHermesRelayerIBC()
	})
}

func Fuzz_Nosy_baseConfigurer_RunIBCTransferChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.RunIBCTransferChannel()
	})
}

func Fuzz_Nosy_baseConfigurer_RunValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		if bc == nil {
			return
		}

		bc.RunValidators()
	})
}

func Fuzz_Nosy_baseConfigurer_createIBCTransferChannel__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chainA *chain.Config
		fill_err = tp.Fill(&chainA)
		if fill_err != nil {
			return
		}
		var chainB *chain.Config
		fill_err = tp.Fill(&chainB)
		if fill_err != nil {
			return
		}
		if bc == nil || chainA == nil || chainB == nil {
			return
		}

		bc.createIBCTransferChannel(chainA, chainB)
	})
}

func Fuzz_Nosy_baseConfigurer_initializeChainConfigFromInitChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var initializedChain *initialization.Chain
		fill_err = tp.Fill(&initializedChain)
		if fill_err != nil {
			return
		}
		var chainConfig *chain.Config
		fill_err = tp.Fill(&chainConfig)
		if fill_err != nil {
			return
		}
		if bc == nil || initializedChain == nil || chainConfig == nil {
			return
		}

		bc.initializeChainConfigFromInitChain(initializedChain, chainConfig)
	})
}

func Fuzz_Nosy_baseConfigurer_runCosmosIBCRelayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chainConfigA *chain.Config
		fill_err = tp.Fill(&chainConfigA)
		if fill_err != nil {
			return
		}
		var chainConfigB *chain.Config
		fill_err = tp.Fill(&chainConfigB)
		if fill_err != nil {
			return
		}
		if bc == nil || chainConfigA == nil || chainConfigB == nil {
			return
		}

		bc.runCosmosIBCRelayer(chainConfigA, chainConfigB)
	})
}

func Fuzz_Nosy_baseConfigurer_runHermesIBCRelayer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chainConfigA *chain.Config
		fill_err = tp.Fill(&chainConfigA)
		if fill_err != nil {
			return
		}
		var chainConfigB *chain.Config
		fill_err = tp.Fill(&chainConfigB)
		if fill_err != nil {
			return
		}
		if bc == nil || chainConfigA == nil || chainConfigB == nil {
			return
		}

		bc.runHermesIBCRelayer(chainConfigA, chainConfigB)
	})
}

func Fuzz_Nosy_baseConfigurer_runValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bc *baseConfigurer
		fill_err = tp.Fill(&bc)
		if fill_err != nil {
			return
		}
		var chainConfig *chain.Config
		fill_err = tp.Fill(&chainConfig)
		if fill_err != nil {
			return
		}
		if bc == nil || chainConfig == nil {
			return
		}

		bc.runValidators(chainConfig)
	})
}

func Fuzz_Nosy_Configurer_ClearResources__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.ClearResources()
	})
}

func Fuzz_Nosy_Configurer_ConfigureChains__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.ConfigureChains()
	})
}

func Fuzz_Nosy_Configurer_GetChainConfig__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var chainIndex int
		fill_err = tp.Fill(&chainIndex)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.GetChainConfig(chainIndex)
	})
}

func Fuzz_Nosy_Configurer_InstantiateBabylonContract__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.InstantiateBabylonContract()
	})
}

func Fuzz_Nosy_Configurer_RunCosmosRelayerIBC__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.RunCosmosRelayerIBC()
	})
}

func Fuzz_Nosy_Configurer_RunHermesRelayerIBC__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.RunHermesRelayerIBC()
	})
}

func Fuzz_Nosy_Configurer_RunIBCTransferChannel__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.RunIBCTransferChannel()
	})
}

func Fuzz_Nosy_Configurer_RunSetup__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.RunSetup()
	})
}

func Fuzz_Nosy_Configurer_RunValidators__(f *testing.F) {
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
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		_x1, err := NewBTCTimestampingPhase2Configurer(t1, isDebugLogEnabled)
		if err != nil {
			return
		}
		_x1.RunValidators()
	})
}

func Fuzz_Nosy_identifierName__(f *testing.F) {
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

		identifierName(t1)
	})
}

// skipping Fuzz_Nosy_parseGovPropFromFile__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.Codec

// skipping Fuzz_Nosy_parseSubmitProposal__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.Codec

func Fuzz_Nosy_updateNodeConfigNameWithIdentifier__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfgs []*initialization.NodeConfig
		fill_err = tp.Fill(&cfgs)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}

		updateNodeConfigNameWithIdentifier(cfgs, identifier)
	})
}

// skipping Fuzz_Nosy_updateNodeConfigs__ because parameters include func, chan, or unsupported interface: func(cfg *github.com/babylonlabs-io/babylon/test/e2e/initialization.NodeConfig) *github.com/babylonlabs-io/babylon/test/e2e/initialization.NodeConfig
