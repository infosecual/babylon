package chain

import (
	"math/rand"
	"net/url"
	"testing"
	"time"

	govv1 "cosmossdk.io/api/cosmos/gov/v1"
	"cosmossdk.io/math"
	asig "github.com/babylonlabs-io/babylon/crypto/schnorr-adaptor-signature"
	"github.com/babylonlabs-io/babylon/test/e2e/containers"
	"github.com/babylonlabs-io/babylon/test/e2e/initialization"
	bbn "github.com/babylonlabs-io/babylon/types"
	btccheckpointtypes "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/babylon/x/finality/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	cmtcrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
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

func Fuzz_Nosy_Config_BTCHeaderBytesHexJoined__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.BTCHeaderBytesHexJoined()
	})
}

func Fuzz_Nosy_Config_CreateNode__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil || initNode == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.CreateNode(initNode)
	})
}

func Fuzz_Nosy_Config_GetDefaultNode__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.GetDefaultNode()
	})
}

func Fuzz_Nosy_Config_GetNodeAtIndex__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var nodeIndex int
		fill_err = tp.Fill(&nodeIndex)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.GetNodeAtIndex(nodeIndex)
	})
}

func Fuzz_Nosy_Config_GetPersistentPeers__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.GetPersistentPeers()
	})
}

func Fuzz_Nosy_Config_RemoveNode__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var nodeName string
		fill_err = tp.Fill(&nodeName)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.RemoveNode(nodeName)
	})
}

func Fuzz_Nosy_Config_SendIBC__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var dstChain *Config
		fill_err = tp.Fill(&dstChain)
		if fill_err != nil {
			return
		}
		var recipient string
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var token sdk.Coin
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil || dstChain == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.SendIBC(dstChain, recipient, token)
	})
}

func Fuzz_Nosy_Config_TxGovVoteFromAllNodes__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var propID int
		fill_err = tp.Fill(&propID)
		if fill_err != nil {
			return
		}
		var option govv1.VoteOption
		fill_err = tp.Fill(&option)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.TxGovVoteFromAllNodes(propID, option, overallFlags...)
	})
}

func Fuzz_Nosy_Config_WaitForNumHeights__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var heightsToWait int64
		fill_err = tp.Fill(&heightsToWait)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.WaitForNumHeights(heightsToWait)
	})
}

func Fuzz_Nosy_Config_WaitUntilHeight__(f *testing.F) {
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
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var id string
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		var initValidatorConfigs []*initialization.NodeConfig
		fill_err = tp.Fill(&initValidatorConfigs)
		if fill_err != nil {
			return
		}
		var ibcConfig *ibctesting.ChannelConfig
		fill_err = tp.Fill(&ibcConfig)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || containerManager == nil || ibcConfig == nil {
			return
		}

		c := New(t1, containerManager, id, initValidatorConfigs, ibcConfig)
		c.WaitUntilHeight(height)
	})
}

func Fuzz_Nosy_NodeConfig_AddBTCDelegationInclusionProof__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var stakingTxHash *chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var inclusionProof *bstypes.InclusionProof
		fill_err = tp.Fill(&inclusionProof)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || stakingTxHash == nil || inclusionProof == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddBTCDelegationInclusionProof(stakingTxHash, inclusionProof)
	})
}

func Fuzz_Nosy_NodeConfig_AddCovenantSigs__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fromWalletName string
		fill_err = tp.Fill(&fromWalletName)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var slashingSigs [][]byte
		fill_err = tp.Fill(&slashingSigs)
		if fill_err != nil {
			return
		}
		var unbondingSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingSig)
		if fill_err != nil {
			return
		}
		var unbondingSlashingSigs [][]byte
		fill_err = tp.Fill(&unbondingSlashingSigs)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || covPK == nil || unbondingSig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddCovenantSigs(fromWalletName, covPK, stakingTxHash, slashingSigs, unbondingSig, unbondingSlashingSigs)
	})
}

func Fuzz_Nosy_NodeConfig_AddCovenantSigsFromVal__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var slashingSigs [][]byte
		fill_err = tp.Fill(&slashingSigs)
		if fill_err != nil {
			return
		}
		var unbondingSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingSig)
		if fill_err != nil {
			return
		}
		var unbondingSlashingSigs [][]byte
		fill_err = tp.Fill(&unbondingSlashingSigs)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || covPK == nil || unbondingSig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddCovenantSigsFromVal(covPK, stakingTxHash, slashingSigs, unbondingSig, unbondingSlashingSigs)
	})
}

func Fuzz_Nosy_NodeConfig_AddCovenantUnbondingSigs__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var covPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&covPK)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var unbondingTxSig *bbn.BIP340Signature
		fill_err = tp.Fill(&unbondingTxSig)
		if fill_err != nil {
			return
		}
		var slashUnbondingTxSigs []*asig.AdaptorSignature
		fill_err = tp.Fill(&slashUnbondingTxSigs)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || covPK == nil || unbondingTxSig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddCovenantUnbondingSigs(covPK, stakingTxHash, unbondingTxSig, slashUnbondingTxSigs)
	})
}

func Fuzz_Nosy_NodeConfig_AddFinalitySig__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var blockHeight uint64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		var pubRand *bbn.SchnorrPubRand
		fill_err = tp.Fill(&pubRand)
		if fill_err != nil {
			return
		}
		var proof cmtcrypto.Proof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var appHash []byte
		fill_err = tp.Fill(&appHash)
		if fill_err != nil {
			return
		}
		var finalitySig *bbn.SchnorrEOTSSig
		fill_err = tp.Fill(&finalitySig)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCPK == nil || pubRand == nil || finalitySig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddFinalitySig(fpBTCPK, blockHeight, pubRand, proof, appHash, finalitySig, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_AddFinalitySigFromVal__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var blockHeight uint64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		var pubRand *bbn.SchnorrPubRand
		fill_err = tp.Fill(&pubRand)
		if fill_err != nil {
			return
		}
		var proof cmtcrypto.Proof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var appHash []byte
		fill_err = tp.Fill(&appHash)
		if fill_err != nil {
			return
		}
		var finalitySig *bbn.SchnorrEOTSSig
		fill_err = tp.Fill(&finalitySig)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCPK == nil || pubRand == nil || finalitySig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddFinalitySigFromVal(fpBTCPK, blockHeight, pubRand, proof, appHash, finalitySig, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_AddFinalitySignatureToBlock__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCSK *secp256k1.PrivateKey
		fill_err = tp.Fill(&fpBTCSK)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var blockHeight uint64
		fill_err = tp.Fill(&blockHeight)
		if fill_err != nil {
			return
		}
		var privateRand *secp256k1.ModNScalar
		fill_err = tp.Fill(&privateRand)
		if fill_err != nil {
			return
		}
		var pubRand *bbn.SchnorrPubRand
		fill_err = tp.Fill(&pubRand)
		if fill_err != nil {
			return
		}
		var proof cmtcrypto.Proof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCSK == nil || fpBTCPK == nil || privateRand == nil || pubRand == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.AddFinalitySignatureToBlock(fpBTCSK, fpBTCPK, blockHeight, privateRand, pubRand, proof, overallFlags...)
	})
}

// skipping Fuzz_Nosy_NodeConfig_BTCStakingUnbondSlashInfo__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_NodeConfig_BTCUndelegate__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var stakingTxHash *chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var spendStakeTx *wire.MsgTx
		fill_err = tp.Fill(&spendStakeTx)
		if fill_err != nil {
			return
		}
		var spendStakeTxInclusionProof *bstypes.InclusionProof
		fill_err = tp.Fill(&spendStakeTxInclusionProof)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || stakingTxHash == nil || spendStakeTx == nil || spendStakeTxInclusionProof == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BTCUndelegate(stakingTxHash, spendStakeTx, spendStakeTxInclusionProof)
	})
}

func Fuzz_Nosy_NodeConfig_BankMultiSend__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fromWallet string
		fill_err = tp.Fill(&fromWallet)
		if fill_err != nil {
			return
		}
		var receivers []string
		fill_err = tp.Fill(&receivers)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BankMultiSend(fromWallet, receivers, amount, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_BankMultiSendFromNode__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var addresses []string
		fill_err = tp.Fill(&addresses)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BankMultiSendFromNode(addresses, amount)
	})
}

func Fuzz_Nosy_NodeConfig_BankSend__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fromWallet string
		fill_err = tp.Fill(&fromWallet)
		if fill_err != nil {
			return
		}
		var to string
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BankSend(fromWallet, to, amount, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_BankSendFromNode__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var receiveAddress string
		fill_err = tp.Fill(&receiveAddress)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BankSendFromNode(receiveAddress, amount)
	})
}

func Fuzz_Nosy_NodeConfig_BankSendOutput__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fromWallet string
		fill_err = tp.Fill(&fromWallet)
		if fill_err != nil {
			return
		}
		var to string
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.BankSendOutput(fromWallet, to, amount, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_CommitPubRandList__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var numPubrand uint64
		fill_err = tp.Fill(&numPubrand)
		if fill_err != nil {
			return
		}
		var commitment []byte
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		var sig *bbn.BIP340Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCPK == nil || sig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.CommitPubRandList(fpBTCPK, startHeight, numPubrand, commitment, sig)
	})
}

func Fuzz_Nosy_NodeConfig_CreateBTCDelegation__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var btcPk *bbn.BIP340PubKey
		fill_err = tp.Fill(&btcPk)
		if fill_err != nil {
			return
		}
		var pop *bstypes.ProofOfPossessionBTC
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var stakingTx []byte
		fill_err = tp.Fill(&stakingTx)
		if fill_err != nil {
			return
		}
		var inclusionProof *bstypes.InclusionProof
		fill_err = tp.Fill(&inclusionProof)
		if fill_err != nil {
			return
		}
		var fpPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpPK)
		if fill_err != nil {
			return
		}
		var stakingTimeBlocks uint16
		fill_err = tp.Fill(&stakingTimeBlocks)
		if fill_err != nil {
			return
		}
		var stakingValue btcutil.Amount
		fill_err = tp.Fill(&stakingValue)
		if fill_err != nil {
			return
		}
		var slashingTx *bstypes.BTCSlashingTx
		fill_err = tp.Fill(&slashingTx)
		if fill_err != nil {
			return
		}
		var delegatorSig *bbn.BIP340Signature
		fill_err = tp.Fill(&delegatorSig)
		if fill_err != nil {
			return
		}
		var unbondingTx *wire.MsgTx
		fill_err = tp.Fill(&unbondingTx)
		if fill_err != nil {
			return
		}
		var unbondingSlashingTx *bstypes.BTCSlashingTx
		fill_err = tp.Fill(&unbondingSlashingTx)
		if fill_err != nil {
			return
		}
		var unbondingTime uint16
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var unbondingValue btcutil.Amount
		fill_err = tp.Fill(&unbondingValue)
		if fill_err != nil {
			return
		}
		var delUnbondingSlashingSig *bbn.BIP340Signature
		fill_err = tp.Fill(&delUnbondingSlashingSig)
		if fill_err != nil {
			return
		}
		var fromWalletName string
		fill_err = tp.Fill(&fromWalletName)
		if fill_err != nil {
			return
		}
		var generateOnly bool
		fill_err = tp.Fill(&generateOnly)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || btcPk == nil || pop == nil || inclusionProof == nil || fpPK == nil || slashingTx == nil || delegatorSig == nil || unbondingTx == nil || unbondingSlashingTx == nil || delUnbondingSlashingSig == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.CreateBTCDelegation(btcPk, pop, stakingTx, inclusionProof, fpPK, stakingTimeBlocks, stakingValue, slashingTx, delegatorSig, unbondingTx, unbondingSlashingTx, unbondingTime, unbondingValue, delUnbondingSlashingSig, fromWalletName, generateOnly, overallFlags...)
	})
}

// skipping Fuzz_Nosy_NodeConfig_CreateBTCDelegationAndCheck__ because parameters include func, chan, or unsupported interface: testing.TB

func Fuzz_Nosy_NodeConfig_CreateFinalityProvider__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletAddrOrName string
		fill_err = tp.Fill(&walletAddrOrName)
		if fill_err != nil {
			return
		}
		var btcPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&btcPK)
		if fill_err != nil {
			return
		}
		var pop *bstypes.ProofOfPossessionBTC
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var moniker string
		fill_err = tp.Fill(&moniker)
		if fill_err != nil {
			return
		}
		var identity string
		fill_err = tp.Fill(&identity)
		if fill_err != nil {
			return
		}
		var website string
		fill_err = tp.Fill(&website)
		if fill_err != nil {
			return
		}
		var securityContract string
		fill_err = tp.Fill(&securityContract)
		if fill_err != nil {
			return
		}
		var details string
		fill_err = tp.Fill(&details)
		if fill_err != nil {
			return
		}
		var commission *math.LegacyDec
		fill_err = tp.Fill(&commission)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || btcPK == nil || pop == nil || commission == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.CreateFinalityProvider(walletAddrOrName, btcPK, pop, moniker, identity, website, securityContract, details, commission)
	})
}

func Fuzz_Nosy_NodeConfig_FailIBCTransfer__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var recipient string
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var amount string
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.FailIBCTransfer(from, recipient, amount)
	})
}

func Fuzz_Nosy_NodeConfig_FinalizeSealedEpochs__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var startEpoch uint64
		fill_err = tp.Fill(&startEpoch)
		if fill_err != nil {
			return
		}
		var lastEpoch uint64
		fill_err = tp.Fill(&lastEpoch)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.FinalizeSealedEpochs(startEpoch, lastEpoch)
	})
}

func Fuzz_Nosy_NodeConfig_FlagChainID__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.FlagChainID()
	})
}

func Fuzz_Nosy_NodeConfig_GetHostPort__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var portId string
		fill_err = tp.Fill(&portId)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.GetHostPort(portId)
	})
}

func Fuzz_Nosy_NodeConfig_GetWallet__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.GetWallet(walletName)
	})
}

func Fuzz_Nosy_NodeConfig_InsertHeader__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var h *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || h == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.InsertHeader(h)
	})
}

func Fuzz_Nosy_NodeConfig_InsertNewEmptyBtcHeader__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || r == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.InsertNewEmptyBtcHeader(r)
	})
}

func Fuzz_Nosy_NodeConfig_InsertProofs__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var p1 *btccheckpointtypes.BTCSpvProof
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var p2 *btccheckpointtypes.BTCSpvProof
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || p1 == nil || p2 == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.InsertProofs(p1, p2)
	})
}

func Fuzz_Nosy_NodeConfig_InstantiateWasmContract__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var codeId string
		fill_err = tp.Fill(&codeId)
		if fill_err != nil {
			return
		}
		var initMsg string
		fill_err = tp.Fill(&initMsg)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.InstantiateWasmContract(codeId, initMsg, from)
	})
}

func Fuzz_Nosy_NodeConfig_KeysAdd__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.KeysAdd(walletName, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_LatestBlockNumber__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.LatestBlockNumber()
	})
}

// skipping Fuzz_Nosy_NodeConfig_LogActionF__ because parameters include func, chan, or unsupported interface: []interface{}

func Fuzz_Nosy_NodeConfig_QueryAccount__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryAccount(address)
	})
}

func Fuzz_Nosy_NodeConfig_QueryActivatedHeight__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryActivatedHeight()
	})
}

func Fuzz_Nosy_NodeConfig_QueryActiveDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryActiveDelegations()
	})
}

func Fuzz_Nosy_NodeConfig_QueryActiveFinalityProvidersAtHeight__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryActiveFinalityProvidersAtHeight(height)
	})
}

func Fuzz_Nosy_NodeConfig_QueryAppliedPlan__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var planName string
		fill_err = tp.Fill(&planName)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryAppliedPlan(planName)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBTCStakingGauge__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBTCStakingGauge(height)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBTCStakingParams__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBTCStakingParams()
	})
}

func Fuzz_Nosy_NodeConfig_QueryBTCStakingParamsByVersion__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBTCStakingParamsByVersion(version)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBalance__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var denom string
		fill_err = tp.Fill(&denom)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBalance(address, denom)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBalances__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var address string
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBalances(address)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBlock__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBlock(height)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBtcBaseHeader__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBtcBaseHeader()
	})
}

func Fuzz_Nosy_NodeConfig_QueryBtcDelegation__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBtcDelegation(stakingTxHash)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBtcDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var status bstypes.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBtcDelegations(status)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBtcLightClientMainchain__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || pagination == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBtcLightClientMainchain(pagination)
	})
}

func Fuzz_Nosy_NodeConfig_QueryBtcLightClientMainchainAll__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryBtcLightClientMainchainAll()
	})
}

func Fuzz_Nosy_NodeConfig_QueryChannelClientState__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var channelID string
		fill_err = tp.Fill(&channelID)
		if fill_err != nil {
			return
		}
		var portID string
		fill_err = tp.Fill(&portID)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryChannelClientState(channelID, portID)
	})
}

func Fuzz_Nosy_NodeConfig_QueryContractsFromId__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var codeId int
		fill_err = tp.Fill(&codeId)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryContractsFromId(codeId)
	})
}

func Fuzz_Nosy_NodeConfig_QueryCurrentEpoch__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryCurrentEpoch()
	})
}

func Fuzz_Nosy_NodeConfig_QueryCurrentHeight__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryCurrentHeight()
	})
}

func Fuzz_Nosy_NodeConfig_QueryFinalityParams__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryFinalityParams()
	})
}

func Fuzz_Nosy_NodeConfig_QueryFinalityProviderDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK string
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryFinalityProviderDelegations(fpBTCPK)
	})
}

func Fuzz_Nosy_NodeConfig_QueryFinalityProviders__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryFinalityProviders()
	})
}

func Fuzz_Nosy_NodeConfig_QueryFinalityProvidersDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpsBTCPK []string
		fill_err = tp.Fill(&fpsBTCPK)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryFinalityProvidersDelegations(fpsBTCPK...)
	})
}

func Fuzz_Nosy_NodeConfig_QueryGRPCGateway__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var path string
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var queryParams url.Values
		fill_err = tp.Fill(&queryParams)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryGRPCGateway(path, queryParams)
	})
}

func Fuzz_Nosy_NodeConfig_QueryHashFromBlock__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryHashFromBlock(height)
	})
}

func Fuzz_Nosy_NodeConfig_QueryHeaderDepth__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var hash string
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryHeaderDepth(hash)
	})
}

func Fuzz_Nosy_NodeConfig_QueryIBCChannels__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryIBCChannels()
	})
}

func Fuzz_Nosy_NodeConfig_QueryIncentiveParams__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryIncentiveParams()
	})
}

func Fuzz_Nosy_NodeConfig_QueryIndexedBlock__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryIndexedBlock(height)
	})
}

func Fuzz_Nosy_NodeConfig_QueryLastFinalizedEpoch__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryLastFinalizedEpoch()
	})
}

func Fuzz_Nosy_NodeConfig_QueryLatestBlockTime__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryLatestBlockTime()
	})
}

func Fuzz_Nosy_NodeConfig_QueryLatestWasmCodeID__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryLatestWasmCodeID()
	})
}

func Fuzz_Nosy_NodeConfig_QueryLightClientHeightCheckpointReported__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var ckptHash []byte
		fill_err = tp.Fill(&ckptHash)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryLightClientHeightCheckpointReported(ckptHash)
	})
}

func Fuzz_Nosy_NodeConfig_QueryLightClientHeightEpochEnd__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryLightClientHeightEpochEnd(epoch)
	})
}

func Fuzz_Nosy_NodeConfig_QueryListBlocks__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var status types.QueriedBlockStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryListBlocks(status)
	})
}

func Fuzz_Nosy_NodeConfig_QueryListPubRandCommit__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCPK == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryListPubRandCommit(fpBTCPK)
	})
}

func Fuzz_Nosy_NodeConfig_QueryListPublicRandomness__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || fpBTCPK == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryListPublicRandomness(fpBTCPK)
	})
}

func Fuzz_Nosy_NodeConfig_QueryListSnapshots__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryListSnapshots()
	})
}

func Fuzz_Nosy_NodeConfig_QueryModuleAddress__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var name string
		fill_err = tp.Fill(&name)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryModuleAddress(name)
	})
}

func Fuzz_Nosy_NodeConfig_QueryNextSequenceReceive__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var channelID string
		fill_err = tp.Fill(&channelID)
		if fill_err != nil {
			return
		}
		var portID string
		fill_err = tp.Fill(&portID)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryNextSequenceReceive(channelID, portID)
	})
}

func Fuzz_Nosy_NodeConfig_QueryNextSequenceSend__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var channelID string
		fill_err = tp.Fill(&channelID)
		if fill_err != nil {
			return
		}
		var portID string
		fill_err = tp.Fill(&portID)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryNextSequenceSend(channelID, portID)
	})
}

func Fuzz_Nosy_NodeConfig_QueryPacketAcknowledgement__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var channelID string
		fill_err = tp.Fill(&channelID)
		if fill_err != nil {
			return
		}
		var portID string
		fill_err = tp.Fill(&portID)
		if fill_err != nil {
			return
		}
		var sequence uint64
		fill_err = tp.Fill(&sequence)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryPacketAcknowledgement(channelID, portID, sequence)
	})
}

// skipping Fuzz_Nosy_NodeConfig_QueryParams__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_NodeConfig_QueryProposal__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var proposalNumber int
		fill_err = tp.Fill(&proposalNumber)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryProposal(proposalNumber)
	})
}

func Fuzz_Nosy_NodeConfig_QueryProposals__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryProposals()
	})
}

func Fuzz_Nosy_NodeConfig_QueryRawCheckpoint__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryRawCheckpoint(epoch)
	})
}

func Fuzz_Nosy_NodeConfig_QueryRawCheckpoints__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil || pagination == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryRawCheckpoints(pagination)
	})
}

func Fuzz_Nosy_NodeConfig_QueryRewardGauge__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var sAddr sdk.AccAddress
		fill_err = tp.Fill(&sAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryRewardGauge(sAddr)
	})
}

func Fuzz_Nosy_NodeConfig_QuerySupplyOf__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var denom string
		fill_err = tp.Fill(&denom)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QuerySupplyOf(denom)
	})
}

func Fuzz_Nosy_NodeConfig_QueryTip__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryTip()
	})
}

func Fuzz_Nosy_NodeConfig_QueryTx__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var txHash string
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryTx(txHash, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_QueryUnbondedDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryUnbondedDelegations()
	})
}

func Fuzz_Nosy_NodeConfig_QueryVerifiedDelegations__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryVerifiedDelegations()
	})
}

func Fuzz_Nosy_NodeConfig_QueryVotesAtHeight__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryVotesAtHeight(height)
	})
}

// skipping Fuzz_Nosy_NodeConfig_QueryWasmSmart__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_NodeConfig_QueryWasmSmartObject__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.QueryWasmSmartObject(contract, msg)
	})
}

func Fuzz_Nosy_NodeConfig_Run__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.Run()
	})
}

func Fuzz_Nosy_NodeConfig_SendHeaderHex__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var headerHex string
		fill_err = tp.Fill(&headerHex)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.SendHeaderHex(headerHex)
	})
}

func Fuzz_Nosy_NodeConfig_SendIBCTransfer__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var recipient string
		fill_err = tp.Fill(&recipient)
		if fill_err != nil {
			return
		}
		var memo string
		fill_err = tp.Fill(&memo)
		if fill_err != nil {
			return
		}
		var token sdk.Coin
		fill_err = tp.Fill(&token)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.SendIBCTransfer(from, recipient, memo, token)
	})
}

func Fuzz_Nosy_NodeConfig_Status__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.Status()
	})
}

func Fuzz_Nosy_NodeConfig_Stop__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.Stop()
	})
}

func Fuzz_Nosy_NodeConfig_StoreWasmCode__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var wasmFile string
		fill_err = tp.Fill(&wasmFile)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.StoreWasmCode(wasmFile, from)
	})
}

// skipping Fuzz_Nosy_NodeConfig_SubmitRefundableTxWithAssertion__ because parameters include func, chan, or unsupported interface: func()

func Fuzz_Nosy_NodeConfig_TxBroadcast__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var txSignedFileFullPath string
		fill_err = tp.Fill(&txSignedFileFullPath)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxBroadcast(txSignedFileFullPath, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxFeeGrant__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var granter string
		fill_err = tp.Fill(&granter)
		if fill_err != nil {
			return
		}
		var grantee string
		fill_err = tp.Fill(&grantee)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxFeeGrant(granter, grantee, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxGovPropSubmitProposal__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var proposalJsonFilePath string
		fill_err = tp.Fill(&proposalJsonFilePath)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxGovPropSubmitProposal(proposalJsonFilePath, from, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxGovVote__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var propID int
		fill_err = tp.Fill(&propID)
		if fill_err != nil {
			return
		}
		var option govv1.VoteOption
		fill_err = tp.Fill(&option)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxGovVote(from, propID, option, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxMultisigSign__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		var multisigAddr string
		fill_err = tp.Fill(&multisigAddr)
		if fill_err != nil {
			return
		}
		var txFileFullPath string
		fill_err = tp.Fill(&txFileFullPath)
		if fill_err != nil {
			return
		}
		var fileName string
		fill_err = tp.Fill(&fileName)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxMultisigSign(walletName, multisigAddr, txFileFullPath, fileName, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxMultisign__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletNameMultisig string
		fill_err = tp.Fill(&walletNameMultisig)
		if fill_err != nil {
			return
		}
		var txFileFullPath string
		fill_err = tp.Fill(&txFileFullPath)
		if fill_err != nil {
			return
		}
		var outputFileName string
		fill_err = tp.Fill(&outputFileName)
		if fill_err != nil {
			return
		}
		var signedFiles []string
		fill_err = tp.Fill(&signedFiles)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxMultisign(walletNameMultisig, txFileFullPath, outputFileName, signedFiles, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxMultisignBroadcast__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletNameMultisig string
		fill_err = tp.Fill(&walletNameMultisig)
		if fill_err != nil {
			return
		}
		var txFileFullPath string
		fill_err = tp.Fill(&txFileFullPath)
		if fill_err != nil {
			return
		}
		var walleNameSigners []string
		fill_err = tp.Fill(&walleNameSigners)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxMultisignBroadcast(walletNameMultisig, txFileFullPath, walleNameSigners)
	})
}

func Fuzz_Nosy_NodeConfig_TxSign__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		var txFileFullPath string
		fill_err = tp.Fill(&txFileFullPath)
		if fill_err != nil {
			return
		}
		var fileName string
		fill_err = tp.Fill(&fileName)
		if fill_err != nil {
			return
		}
		var overallFlags []string
		fill_err = tp.Fill(&overallFlags)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxSign(walletName, txFileFullPath, fileName, overallFlags...)
	})
}

func Fuzz_Nosy_NodeConfig_TxSignBroadcast__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var walletName string
		fill_err = tp.Fill(&walletName)
		if fill_err != nil {
			return
		}
		var txFileFullPath string
		fill_err = tp.Fill(&txFileFullPath)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.TxSignBroadcast(walletName, txFileFullPath)
	})
}

func Fuzz_Nosy_NodeConfig_WaitFinalityIsActivated__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitFinalityIsActivated()
	})
}

// skipping Fuzz_Nosy_NodeConfig_WaitForCondition__ because parameters include func, chan, or unsupported interface: func() bool

// skipping Fuzz_Nosy_NodeConfig_WaitForConditionWithPause__ because parameters include func, chan, or unsupported interface: func() bool

func Fuzz_Nosy_NodeConfig_WaitForNextBlock__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitForNextBlock()
	})
}

func Fuzz_Nosy_NodeConfig_WaitForNextBlockWithSleep50ms__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitForNextBlockWithSleep50ms()
	})
}

func Fuzz_Nosy_NodeConfig_WaitForNextBlocks__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var numberOfBlocks uint64
		fill_err = tp.Fill(&numberOfBlocks)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitForNextBlocks(numberOfBlocks)
	})
}

// skipping Fuzz_Nosy_NodeConfig_WaitUntil__ because parameters include func, chan, or unsupported interface: func(syncInfo github.com/cometbft/cometbft/rpc/core/types.SyncInfo) bool

func Fuzz_Nosy_NodeConfig_WaitUntilBtcHeight__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitUntilBtcHeight(height)
	})
}

func Fuzz_Nosy_NodeConfig_WaitUntilCurrentEpochIsSealedAndFinalized__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var startEpoch uint64
		fill_err = tp.Fill(&startEpoch)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WaitUntilCurrentEpochIsSealedAndFinalized(startEpoch)
	})
}

func Fuzz_Nosy_NodeConfig_WasmExecute__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var contract string
		fill_err = tp.Fill(&contract)
		if fill_err != nil {
			return
		}
		var execMsg string
		fill_err = tp.Fill(&execMsg)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WasmExecute(contract, execMsg, from)
	})
}

func Fuzz_Nosy_NodeConfig_WithSetupTime__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var t6 time.Time
		fill_err = tp.Fill(&t6)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WithSetupTime(t6)
	})
}

func Fuzz_Nosy_NodeConfig_WithdrawReward__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var sType string
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var from string
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WithdrawReward(sType, from)
	})
}

func Fuzz_Nosy_NodeConfig_WithdrawRewardCheckingBalances__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var sType string
		fill_err = tp.Fill(&sType)
		if fill_err != nil {
			return
		}
		var fromAddr string
		fill_err = tp.Fill(&fromAddr)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WithdrawRewardCheckingBalances(sType, fromAddr)
	})
}

func Fuzz_Nosy_NodeConfig_WriteFile__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		var fileName string
		fill_err = tp.Fill(&fileName)
		if fill_err != nil {
			return
		}
		var content string
		fill_err = tp.Fill(&content)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.WriteFile(fileName, content)
	})
}

func Fuzz_Nosy_NodeConfig_extractOperatorAddressIfValidator__(f *testing.F) {
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
		var initNode *initialization.Node
		fill_err = tp.Fill(&initNode)
		if fill_err != nil {
			return
		}
		var initConfig *initialization.NodeConfig
		fill_err = tp.Fill(&initConfig)
		if fill_err != nil {
			return
		}
		var chainId string
		fill_err = tp.Fill(&chainId)
		if fill_err != nil {
			return
		}
		var containerManager *containers.Manager
		fill_err = tp.Fill(&containerManager)
		if fill_err != nil {
			return
		}
		if t1 == nil || initNode == nil || initConfig == nil || containerManager == nil {
			return
		}

		n := NewNodeConfig(t1, initNode, initConfig, chainId, containerManager)
		n.extractOperatorAddressIfValidator()
	})
}

func Fuzz_Nosy_CovenantBTCPKs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *bstypes.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		CovenantBTCPKs(params)
	})
}

func Fuzz_Nosy_GetTxHashFromOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, txOutput string) {
		GetTxHashFromOutput(txOutput)
	})
}
