package containers

import (
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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

func Fuzz_Nosy_Manager_ClearResources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.ClearResources()
	})
}

func Fuzz_Nosy_Manager_CosmosRlyrContainerName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.CosmosRlyrContainerName()
	})
}

func Fuzz_Nosy_Manager_ExecCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var isCosmosRelayer bool
		fill_err = tp.Fill(&isCosmosRelayer)
		if fill_err != nil {
			return
		}
		var isUpgrade bool
		fill_err = tp.Fill(&isUpgrade)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var fullContainerName string
		fill_err = tp.Fill(&fullContainerName)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var success string
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		if t5 == nil {
			return
		}

		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.ExecCmd(t5, fullContainerName, command, success)
	})
}

func Fuzz_Nosy_Manager_ExecHermesCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var isCosmosRelayer bool
		fill_err = tp.Fill(&isCosmosRelayer)
		if fill_err != nil {
			return
		}
		var isUpgrade bool
		fill_err = tp.Fill(&isUpgrade)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var success string
		fill_err = tp.Fill(&success)
		if fill_err != nil {
			return
		}
		if t5 == nil {
			return
		}

		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.ExecHermesCmd(t5, command, success)
	})
}

func Fuzz_Nosy_Manager_ExecTxCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var isCosmosRelayer bool
		fill_err = tp.Fill(&isCosmosRelayer)
		if fill_err != nil {
			return
		}
		var isUpgrade bool
		fill_err = tp.Fill(&isUpgrade)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var nodeName string
		fill_err = tp.Fill(&nodeName)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		if t5 == nil {
			return
		}

		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.ExecTxCmd(t5, chainId, nodeName, command)
	})
}

func Fuzz_Nosy_Manager_ExecTxCmdWithSuccessString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var isCosmosRelayer bool
		fill_err = tp.Fill(&isCosmosRelayer)
		if fill_err != nil {
			return
		}
		var isUpgrade bool
		fill_err = tp.Fill(&isUpgrade)
		if fill_err != nil {
			return
		}
		var t5 *testing.T
		fill_err = tp.Fill(&t5)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerName string
		fill_err = tp.Fill(&containerName)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		var successStr string
		fill_err = tp.Fill(&successStr)
		if fill_err != nil {
			return
		}
		if t5 == nil {
			return
		}

		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.ExecTxCmdWithSuccessString(t5, chainId, containerName, command, successStr)
	})
}

func Fuzz_Nosy_Manager_GetHostPort__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, nodeName string, portId string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.GetHostPort(nodeName, portId)
	})
}

func Fuzz_Nosy_Manager_GetNodeResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, containerName string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.GetNodeResource(containerName)
	})
}

func Fuzz_Nosy_Manager_HermesContainerName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.HermesContainerName()
	})
}

func Fuzz_Nosy_Manager_NetworkName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.NetworkName()
	})
}

func Fuzz_Nosy_Manager_PurgeResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var identifier string
		fill_err = tp.Fill(&identifier)
		if fill_err != nil {
			return
		}
		var isDebugLogEnabled bool
		fill_err = tp.Fill(&isDebugLogEnabled)
		if fill_err != nil {
			return
		}
		var isCosmosRelayer bool
		fill_err = tp.Fill(&isCosmosRelayer)
		if fill_err != nil {
			return
		}
		var isUpgrade bool
		fill_err = tp.Fill(&isUpgrade)
		if fill_err != nil {
			return
		}
		var resource *dockertest.Resource
		fill_err = tp.Fill(&resource)
		if fill_err != nil {
			return
		}
		if resource == nil {
			return
		}

		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.PurgeResource(resource)
	})
}

func Fuzz_Nosy_Manager_RemoveNodeResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, containerName string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.RemoveNodeResource(containerName)
	})
}

func Fuzz_Nosy_Manager_RunChainInitResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, chainId string, chainVotingPeriod int, chainExpeditedVotingPeriod int, validatorInitConfigBytesHexEncoded string, mountDir string, forkHeight int, btcHeaders string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.RunChainInitResource(chainId, chainVotingPeriod, chainExpeditedVotingPeriod, validatorInitConfigBytesHexEncoded, mountDir, forkHeight, btcHeaders)
	})
}

func Fuzz_Nosy_Manager_RunHermesResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, chainAID string, osmoARelayerNodeName string, osmoAValMnemonic string, chainBID string, osmoBRelayerNodeName string, osmoBValMnemonic string, hermesCfgPath string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.RunHermesResource(chainAID, osmoARelayerNodeName, osmoAValMnemonic, chainBID, osmoBRelayerNodeName, osmoBValMnemonic, hermesCfgPath)
	})
}

func Fuzz_Nosy_Manager_RunNodeResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, chainId string, containerName string, valCondifDir string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.RunNodeResource(chainId, containerName, valCondifDir)
	})
}

func Fuzz_Nosy_Manager_RunRlyResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, identifier string, isDebugLogEnabled bool, isCosmosRelayer bool, isUpgrade bool, chainAID string, osmoARelayerNodeName string, osmoAValMnemonic string, chainAIbcPort string, chainBID string, osmoBRelayerNodeName string, osmoBValMnemonic string, chainBIbcPort string, rlyCfgPath string) {
		m, err := NewManager(identifier, isDebugLogEnabled, isCosmosRelayer, isUpgrade)
		if err != nil {
			return
		}
		m.RunRlyResource(chainAID, osmoARelayerNodeName, osmoAValMnemonic, chainAIbcPort, chainBID, osmoBRelayerNodeName, osmoBValMnemonic, chainBIbcPort, rlyCfgPath)
	})
}

func Fuzz_Nosy_noRestart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *docker.HostConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		noRestart(config)
	})
}
