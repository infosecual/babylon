package initialization

import (
	"testing"
	"time"

	"github.com/babylonlabs-io/babylon/app/signer"
	btccheckpointtypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	btclighttypes "github.com/babylonlabs-io/babylon/x/btclightclient/types"
	finalitytypes "github.com/babylonlabs-io/babylon/x/finality/types"
	minttypes "github.com/babylonlabs-io/babylon/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	staketypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func Fuzz_Nosy_ChainMeta_configDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *ChainMeta
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.configDir()
	})
}

func Fuzz_Nosy_internalChain_export__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string, dataDir string) {
		c := new(id, dataDir)
		c.export()
	})
}

func Fuzz_Nosy_internalNode_buildCreateValidatorMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		var amount sdk.Coin
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var consensusKey signer.ConsensusKey
		fill_err = tp.Fill(&consensusKey)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.buildCreateValidatorMsg(amount, consensusKey)
	})
}

func Fuzz_Nosy_internalNode_configDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.configDir()
	})
}

func Fuzz_Nosy_internalNode_createAppConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var n2 *NodeConfig
		fill_err = tp.Fill(&n2)
		if fill_err != nil {
			return
		}
		var n3 *NodeConfig
		fill_err = tp.Fill(&n3)
		if fill_err != nil {
			return
		}
		if chain == nil || n2 == nil || n3 == nil {
			return
		}

		n, err := newNode(chain, n2)
		if err != nil {
			return
		}
		n.createAppConfig(n3)
	})
}

func Fuzz_Nosy_internalNode_createConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createConfig()
	})
}

func Fuzz_Nosy_internalNode_createConsensusKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createConsensusKey()
	})
}

func Fuzz_Nosy_internalNode_createKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createKey(name)
	})
}

func Fuzz_Nosy_internalNode_createKeyFromMnemonic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		var mnemonic string
		fill_err = tp.Fill(&mnemonic)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createKeyFromMnemonic(name, mnemonic)
	})
}

func Fuzz_Nosy_internalNode_createMnemonic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createMnemonic()
	})
}

func Fuzz_Nosy_internalNode_createNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.createNodeKey()
	})
}

func Fuzz_Nosy_internalNode_export__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.export()
	})
}

func Fuzz_Nosy_internalNode_getAppGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.getAppGenesis()
	})
}

func Fuzz_Nosy_internalNode_getNodeKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.getNodeKey()
	})
}

func Fuzz_Nosy_internalNode_init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.init()
	})
}

func Fuzz_Nosy_internalNode_initNodeConfigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain *internalChain
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var nodeConfig *NodeConfig
		fill_err = tp.Fill(&nodeConfig)
		if fill_err != nil {
			return
		}
		var persistentPeers []string
		fill_err = tp.Fill(&persistentPeers)
		if fill_err != nil {
			return
		}
		if chain == nil || nodeConfig == nil {
			return
		}

		n, err := newNode(chain, nodeConfig)
		if err != nil {
			return
		}
		n.initNodeConfigs(persistentPeers)
	})
}

// skipping Fuzz_Nosy_internalNode_signMsg__ because parameters include func, chan, or unsupported interface: []github.com/cosmos/cosmos-sdk/types.Msg

func Fuzz_Nosy_updateBankGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bankGenState *banktypes.GenesisState
		fill_err = tp.Fill(&bankGenState)
		if fill_err != nil {
			return
		}
		if bankGenState == nil {
			return
		}

		updateBankGenesis(bankGenState)
	})
}

func Fuzz_Nosy_updateBtcLightClientGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcHeaders []*btclighttypes.BTCHeaderInfo
		fill_err = tp.Fill(&btcHeaders)
		if fill_err != nil {
			return
		}

		updateBtcLightClientGenesis(btcHeaders)
	})
}

func Fuzz_Nosy_updateBtccheckpointGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btccheckpointGenState *btccheckpointtypes.GenesisState
		fill_err = tp.Fill(&btccheckpointGenState)
		if fill_err != nil {
			return
		}
		if btccheckpointGenState == nil {
			return
		}

		updateBtccheckpointGenesis(btccheckpointGenState)
	})
}

func Fuzz_Nosy_updateCrisisGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var crisisGenState *crisistypes.GenesisState
		fill_err = tp.Fill(&crisisGenState)
		if fill_err != nil {
			return
		}
		if crisisGenState == nil {
			return
		}

		updateCrisisGenesis(crisisGenState)
	})
}

func Fuzz_Nosy_updateFinalityGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var finalityGenState *finalitytypes.GenesisState
		fill_err = tp.Fill(&finalityGenState)
		if fill_err != nil {
			return
		}
		if finalityGenState == nil {
			return
		}

		updateFinalityGenesis(finalityGenState)
	})
}

func Fuzz_Nosy_updateGenUtilGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *internalChain
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		updateGenUtilGenesis(c)
	})
}

func Fuzz_Nosy_updateGovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var votingPeriod time.Duration
		fill_err = tp.Fill(&votingPeriod)
		if fill_err != nil {
			return
		}
		var expeditedVotingPeriod time.Duration
		fill_err = tp.Fill(&expeditedVotingPeriod)
		if fill_err != nil {
			return
		}

		updateGovGenesis(votingPeriod, expeditedVotingPeriod)
	})
}

func Fuzz_Nosy_updateMintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mintGenState *minttypes.GenesisState
		fill_err = tp.Fill(&mintGenState)
		if fill_err != nil {
			return
		}
		if mintGenState == nil {
			return
		}

		updateMintGenesis(mintGenState)
	})
}

func Fuzz_Nosy_updateStakeGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakeGenState *staketypes.GenesisState
		fill_err = tp.Fill(&stakeGenState)
		if fill_err != nil {
			return
		}
		if stakeGenState == nil {
			return
		}

		updateStakeGenesis(stakeGenState)
	})
}
