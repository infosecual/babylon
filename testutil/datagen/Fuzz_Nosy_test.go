package datagen

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/babylonlabs-io/babylon/btctxformatter"
	"github.com/babylonlabs-io/babylon/crypto/bls12381"
	btclightclientk "github.com/babylonlabs-io/babylon/x/btclightclient/keeper"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/cometbft/cometbft/config"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
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

func Fuzz_Nosy_BTCHeaderPartialChain_ChainLength__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.ChainLength()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_ChainToBytes__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.ChainToBytes()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetChainInfo__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetChainInfo()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetChainInfoResponse__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetChainInfoResponse()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetHeadersMap__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetHeadersMap()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetRandomHeaderInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 *rand.Rand
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		var r5 *rand.Rand
		fill_err = tp.Fill(&r5)
		if fill_err != nil {
			return
		}
		if r1 == nil || r5 == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r1, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetRandomHeaderInfo(r5)
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetRandomHeaderInfoNoTip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 *rand.Rand
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		var r5 *rand.Rand
		fill_err = tp.Fill(&r5)
		if fill_err != nil {
			return
		}
		if r1 == nil || r5 == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r1, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetRandomHeaderInfoNoTip(r5)
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_GetTipInfo__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.GetTipInfo()
	})
}

func Fuzz_Nosy_BTCHeaderPartialChain_TipHeader__(f *testing.F) {
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
		var initialHeaderHeight uint32
		fill_err = tp.Fill(&initialHeaderHeight)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork uint32
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}
		var length uint32
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		c := NewBTCHeaderChainWithLength(r, initialHeaderHeight, initialHeaderTotalWork, length)
		c.TipHeader()
	})
}

func Fuzz_Nosy_GenesisValidators_GetBLSPrivKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numVals int) {
		gvs, err := GenesisValidatorSet(numVals)
		if err != nil {
			return
		}
		gvs.GetBLSPrivKeys()
	})
}

func Fuzz_Nosy_GenesisValidators_GetGenesisKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numVals int) {
		gvs, err := GenesisValidatorSet(numVals)
		if err != nil {
			return
		}
		gvs.GetGenesisKeys()
	})
}

func Fuzz_Nosy_GenesisValidators_GetValPrivKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numVals int) {
		gvs, err := GenesisValidatorSet(numVals)
		if err != nil {
			return
		}
		gvs.GetValPrivKeys()
	})
}

func Fuzz_Nosy_TestUnbondingSlashingInfo_GenCovenantSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *TestUnbondingSlashingInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		var covSKs []*btcec.PrivateKey
		fill_err = tp.Fill(&covSKs)
		if fill_err != nil {
			return
		}
		var fpPKs []*btcec.PublicKey
		fill_err = tp.Fill(&fpPKs)
		if fill_err != nil {
			return
		}
		var stakingTx *wire.MsgTx
		fill_err = tp.Fill(&stakingTx)
		if fill_err != nil {
			return
		}
		var unbondingPkScriptPath []byte
		fill_err = tp.Fill(&unbondingPkScriptPath)
		if fill_err != nil {
			return
		}
		if info == nil || stakingTx == nil {
			return
		}

		info.GenCovenantSigs(covSKs, fpPKs, stakingTx, unbondingPkScriptPath)
	})
}

func Fuzz_Nosy_TestUnbondingSlashingInfo_GenDelSlashingTxSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *TestUnbondingSlashingInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		if info == nil || sk == nil {
			return
		}

		info.GenDelSlashingTxSig(sk)
	})
}

func Fuzz_Nosy_PV_SignProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var proposal *cmtproto.Proposal
		fill_err = tp.Fill(&proposal)
		if fill_err != nil {
			return
		}
		if proposal == nil {
			return
		}

		pv := NewPV()
		pv.SignProposal(chainID, proposal)
	})
}

func Fuzz_Nosy_PV_SignVote__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chainID string
		fill_err = tp.Fill(&chainID)
		if fill_err != nil {
			return
		}
		var vote *cmtproto.Vote
		fill_err = tp.Fill(&vote)
		if fill_err != nil {
			return
		}
		if vote == nil {
			return
		}

		pv := NewPV()
		pv.SignVote(chainID, vote)
	})
}

func Fuzz_Nosy_AddRandomSeedsToFuzzer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *testing.F
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var num uint
		fill_err = tp.Fill(&num)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		AddRandomSeedsToFuzzer(f1, num)
	})
}

func Fuzz_Nosy_BlockCreationResultToProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var inputs []*BlockCreationResult
		fill_err = tp.Fill(&inputs)
		if fill_err != nil {
			return
		}

		BlockCreationResultToProofs(inputs)
	})
}

func Fuzz_Nosy_ChainToInfoChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain []*wire.BlockHeader
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var initialHeaderNumber uint32
		fill_err = tp.Fill(&initialHeaderNumber)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork math.Uint
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}

		ChainToInfoChain(chain, initialHeaderNumber, initialHeaderTotalWork)
	})
}

func Fuzz_Nosy_ChainToInfoResponseChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var chain []*wire.BlockHeader
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		var initialHeaderNumber uint32
		fill_err = tp.Fill(&initialHeaderNumber)
		if fill_err != nil {
			return
		}
		var initialHeaderTotalWork math.Uint
		fill_err = tp.Fill(&initialHeaderTotalWork)
		if fill_err != nil {
			return
		}

		ChainToInfoResponseChain(chain, initialHeaderNumber, initialHeaderTotalWork)
	})
}

func Fuzz_Nosy_CreateNFinalityProviders__(f *testing.F) {
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
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if r == nil || t2 == nil {
			return
		}

		CreateNFinalityProviders(r, t2, n)
	})
}

func Fuzz_Nosy_GenChainFromListsOfTransactions__(f *testing.F) {
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
		var txs [][]*wire.MsgTx
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var initHeader *wire.BlockHeader
		fill_err = tp.Fill(&initHeader)
		if fill_err != nil {
			return
		}
		if r == nil || initHeader == nil {
			return
		}

		GenChainFromListsOfTransactions(r, txs, initHeader)
	})
}

func Fuzz_Nosy_GenCovenantAdaptorSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var covenantSKs []*btcec.PrivateKey
		fill_err = tp.Fill(&covenantSKs)
		if fill_err != nil {
			return
		}
		var fpPKs []*btcec.PublicKey
		fill_err = tp.Fill(&fpPKs)
		if fill_err != nil {
			return
		}
		var fundingTx *wire.MsgTx
		fill_err = tp.Fill(&fundingTx)
		if fill_err != nil {
			return
		}
		var pkScriptPath []byte
		fill_err = tp.Fill(&pkScriptPath)
		if fill_err != nil {
			return
		}
		var slashingTx *bstypes.BTCSlashingTx
		fill_err = tp.Fill(&slashingTx)
		if fill_err != nil {
			return
		}
		if fundingTx == nil || slashingTx == nil {
			return
		}

		GenCovenantAdaptorSigs(covenantSKs, fpPKs, fundingTx, pkScriptPath, slashingTx)
	})
}

func Fuzz_Nosy_GenCovenantCommittee__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenCovenantCommittee(r)
	})
}

func Fuzz_Nosy_GenCovenantUnbondingSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var covenantSKs []*btcec.PrivateKey
		fill_err = tp.Fill(&covenantSKs)
		if fill_err != nil {
			return
		}
		var stakingTx *wire.MsgTx
		fill_err = tp.Fill(&stakingTx)
		if fill_err != nil {
			return
		}
		var stakingOutIdx uint32
		fill_err = tp.Fill(&stakingOutIdx)
		if fill_err != nil {
			return
		}
		var unbondingPkScriptPath []byte
		fill_err = tp.Fill(&unbondingPkScriptPath)
		if fill_err != nil {
			return
		}
		var unbondingTx *wire.MsgTx
		fill_err = tp.Fill(&unbondingTx)
		if fill_err != nil {
			return
		}
		if stakingTx == nil || unbondingTx == nil {
			return
		}

		GenCovenantUnbondingSigs(covenantSKs, stakingTx, stakingOutIdx, unbondingPkScriptPath, unbondingTx)
	})
}

func Fuzz_Nosy_GenNEmptyBlocks__(f *testing.F) {
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
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var prevHeader *wire.BlockHeader
		fill_err = tp.Fill(&prevHeader)
		if fill_err != nil {
			return
		}
		if r == nil || prevHeader == nil {
			return
		}

		GenNEmptyBlocks(r, n, prevHeader)
	})
}

func Fuzz_Nosy_GenRandBtcChainInsertingInKeeper__(f *testing.F) {
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
		var k *btclightclientk.Keeper
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var initialHeight uint32
		fill_err = tp.Fill(&initialHeight)
		if fill_err != nil {
			return
		}
		var chainLength uint32
		fill_err = tp.Fill(&chainLength)
		if fill_err != nil {
			return
		}
		if t1 == nil || r == nil || k == nil {
			return
		}

		GenRandBtcChainInsertingInKeeper(t1, r, k, ctx, initialHeight, chainLength)
	})
}

func Fuzz_Nosy_GenRandomAccWithBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		GenRandomAccWithBalance(n)
	})
}

func Fuzz_Nosy_GenRandomAddrAndSat__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomAddrAndSat(r)
	})
}

func Fuzz_Nosy_GenRandomBTCHeaderBits__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBTCHeaderBits(r)
	})
}

func Fuzz_Nosy_GenRandomBTCHeaderNonce__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBTCHeaderNonce(r)
	})
}

func Fuzz_Nosy_GenRandomBTCHeaderVersion__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBTCHeaderVersion(r)
	})
}

func Fuzz_Nosy_GenRandomBTCHeight__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBTCHeight(r)
	})
}

func Fuzz_Nosy_GenRandomBTCKeyPair__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBTCKeyPair(r)
	})
}

func Fuzz_Nosy_GenRandomBTCKeyPairs__(f *testing.F) {
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
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomBTCKeyPairs(r, n)
	})
}

func Fuzz_Nosy_GenRandomBabylonTxPair__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBabylonTxPair(r)
	})
}

func Fuzz_Nosy_GenRandomBitmap__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomBitmap(r)
	})
}

func Fuzz_Nosy_GenRandomBtcdBlock__(f *testing.F) {
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
		var numBabylonTxs int
		fill_err = tp.Fill(&numBabylonTxs)
		if fill_err != nil {
			return
		}
		var prevHash *chainhash.Hash
		fill_err = tp.Fill(&prevHash)
		if fill_err != nil {
			return
		}
		if r == nil || prevHash == nil {
			return
		}

		GenRandomBtcdBlock(r, numBabylonTxs, prevHash)
	})
}

func Fuzz_Nosy_GenRandomBtcdBlockchainWithBabylonTx__(f *testing.F) {
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
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var oneTxThreshold float32
		fill_err = tp.Fill(&oneTxThreshold)
		if fill_err != nil {
			return
		}
		var twoTxThreshold float32
		fill_err = tp.Fill(&twoTxThreshold)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomBtcdBlockchainWithBabylonTx(r, n, oneTxThreshold, twoTxThreshold)
	})
}

func Fuzz_Nosy_GenRandomByteArray__(f *testing.F) {
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
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomByteArray(r, length)
	})
}

func Fuzz_Nosy_GenRandomDenom__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomDenom(r)
	})
}

func Fuzz_Nosy_GenRandomEpochInterval__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomEpochInterval(r)
	})
}

func Fuzz_Nosy_GenRandomEpochNum__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomEpochNum(r)
	})
}

func Fuzz_Nosy_GenRandomFPHistRwdStartAndEnd__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomFPHistRwdStartAndEnd(r)
	})
}

func Fuzz_Nosy_GenRandomFinalityProviderDistInfo__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomFinalityProviderDistInfo(r)
	})
}

func Fuzz_Nosy_GenRandomHexStr__(f *testing.F) {
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
		var length uint64
		fill_err = tp.Fill(&length)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomHexStr(r, length)
	})
}

func Fuzz_Nosy_GenRandomMsgCommitPubRandList__(f *testing.F) {
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
		var sk *btcec.PrivateKey
		fill_err = tp.Fill(&sk)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var numPubRand uint64
		fill_err = tp.Fill(&numPubRand)
		if fill_err != nil {
			return
		}
		if r == nil || sk == nil {
			return
		}

		GenRandomMsgCommitPubRandList(r, sk, startHeight, numPubRand)
	})
}

func Fuzz_Nosy_GenRandomPkHash__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomPkHash(r)
	})
}

func Fuzz_Nosy_GenRandomPubKeyHashScript__(f *testing.F) {
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
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if r == nil || net == nil {
			return
		}

		GenRandomPubKeyHashScript(r, net)
	})
}

func Fuzz_Nosy_GenRandomPubkeysAndSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int, msg []byte) {
		GenRandomPubkeysAndSigs(n, msg)
	})
}

func Fuzz_Nosy_GenRandomSecp256k1KeyPair__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomSecp256k1KeyPair(r)
	})
}

func Fuzz_Nosy_GenRandomSequenceRawCheckpointsWithMeta__(f *testing.F) {
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
		if r == nil {
			return
		}

		GenRandomSequenceRawCheckpointsWithMeta(r)
	})
}

func Fuzz_Nosy_GenRandomValidChainStartingFrom__(f *testing.F) {
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
		var parentHeader *wire.BlockHeader
		fill_err = tp.Fill(&parentHeader)
		if fill_err != nil {
			return
		}
		var timeBetweenBlocks *TimeBetweenBlocksInfo
		fill_err = tp.Fill(&timeBetweenBlocks)
		if fill_err != nil {
			return
		}
		var numHeaders uint32
		fill_err = tp.Fill(&numHeaders)
		if fill_err != nil {
			return
		}
		if r == nil || parentHeader == nil || timeBetweenBlocks == nil {
			return
		}

		GenRandomValidChainStartingFrom(r, parentHeader, timeBetweenBlocks, numHeaders)
	})
}

func Fuzz_Nosy_GenRandomVoteExtension__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var valSet *GenesisValidators
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if valSet == nil || r == nil {
			return
		}

		GenRandomVoteExtension(epochNum, height, blockHash, valSet, r)
	})
}

func Fuzz_Nosy_GenRandomVotingPowerDistCache__(f *testing.F) {
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
		var maxFPs uint32
		fill_err = tp.Fill(&maxFPs)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomVotingPowerDistCache(r, maxFPs)
	})
}

func Fuzz_Nosy_GenSequenceRawCheckpointsWithMeta__(f *testing.F) {
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
		var tipEpoch uint64
		fill_err = tp.Fill(&tipEpoch)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenSequenceRawCheckpointsWithMeta(r, tipEpoch)
	})
}

func Fuzz_Nosy_GenerateBLSSigs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keys []bls12381.PrivateKey
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		GenerateBLSSigs(keys, msg)
	})
}

func Fuzz_Nosy_GenerateValidatorSetWithBLSPrivKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, n int) {
		GenerateValidatorSetWithBLSPrivKeys(n)
	})
}

func Fuzz_Nosy_GenesisValidatorSetWithPrivSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, numVals int) {
		GenesisValidatorSetWithPrivSigner(numVals)
	})
}

func Fuzz_Nosy_HeaderToHeaderBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var headers []*wire.BlockHeader
		fill_err = tp.Fill(&headers)
		if fill_err != nil {
			return
		}

		HeaderToHeaderBytes(headers)
	})
}

func Fuzz_Nosy_InitializeNodeValidatorFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		InitializeNodeValidatorFiles(config, addr)
	})
}

func Fuzz_Nosy_InitializeNodeValidatorFilesFromMnemonic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *config.Config
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		var mnemonic string
		fill_err = tp.Fill(&mnemonic)
		if fill_err != nil {
			return
		}
		var addr sdk.AccAddress
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		InitializeNodeValidatorFilesFromMnemonic(config, mnemonic, addr)
	})
}

func Fuzz_Nosy_OneInN__(f *testing.F) {
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
		var n int
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		OneInN(r, n)
	})
}

func Fuzz_Nosy_RandomInRange__(f *testing.F) {
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
		var min int
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		RandomInRange(r, min, max)
	})
}

func Fuzz_Nosy_RandomInt__(f *testing.F) {
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
		var rng int
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		RandomInt(r, rng)
	})
}

func Fuzz_Nosy_RandomIntOtherThan__(f *testing.F) {
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
		var x int
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}
		var rng int
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		RandomIntOtherThan(r, x, rng)
	})
}

func Fuzz_Nosy_RandomRawCheckpointDataForEpoch__(f *testing.F) {
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
		var e uint64
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		RandomRawCheckpointDataForEpoch(r, e)
	})
}

func Fuzz_Nosy_RandomUInt32__(f *testing.F) {
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
		var rng uint32
		fill_err = tp.Fill(&rng)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		RandomUInt32(r, rng)
	})
}

func Fuzz_Nosy_SolveBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		SolveBlock(header)
	})
}

func Fuzz_Nosy_ValidHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hexStr string, length int) {
		ValidHex(hexStr, length)
	})
}

func Fuzz_Nosy_calculateAdjustedDifficulty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var lastRetargetHeader *wire.BlockHeader
		fill_err = tp.Fill(&lastRetargetHeader)
		if fill_err != nil {
			return
		}
		var currentHeaderTimestamp time.Time
		fill_err = tp.Fill(&currentHeaderTimestamp)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if lastRetargetHeader == nil || params == nil {
			return
		}

		calculateAdjustedDifficulty(lastRetargetHeader, currentHeaderTimestamp, params)
	})
}

func Fuzz_Nosy_getExpectedOpReturn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var f2 []byte
		fill_err = tp.Fill(&f2)
		if fill_err != nil {
			return
		}
		var s []byte
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}

		getExpectedOpReturn(tag, f2, s)
	})
}

func Fuzz_Nosy_getFirstBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNumber uint64, epochInterval uint64) {
		getFirstBlockHeight(epochNumber, epochInterval)
	})
}

func Fuzz_Nosy_numInRange__(f *testing.F) {
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
		var min int
		fill_err = tp.Fill(&min)
		if fill_err != nil {
			return
		}
		var max int
		fill_err = tp.Fill(&max)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		numInRange(r, min, max)
	})
}

func Fuzz_Nosy_opReturnScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, d1 []byte) {
		opReturnScript(d1)
	})
}

func Fuzz_Nosy_standardCoinbaseScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, blockHeight int32, extraNonce uint64) {
		standardCoinbaseScript(blockHeight, extraNonce)
	})
}

func Fuzz_Nosy_uniqueOpReturnScript__(f *testing.F) {
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
		if r == nil {
			return
		}

		uniqueOpReturnScript(r)
	})
}
