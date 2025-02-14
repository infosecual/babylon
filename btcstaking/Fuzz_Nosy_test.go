package btcstaking

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_IdentifiableStakingInfo_SlashingPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag []byte
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKey *btcec.PublicKey
		fill_err = tp.Fill(&fpKey)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || fpKey == nil || net == nil {
			return
		}

		i, err := BuildV0IdentifiableStakingOutputs(tag, stakerKey, fpKey, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.SlashingPathSpendInfo()
	})
}

func Fuzz_Nosy_IdentifiableStakingInfo_TimeLockPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag []byte
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKey *btcec.PublicKey
		fill_err = tp.Fill(&fpKey)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || fpKey == nil || net == nil {
			return
		}

		i, err := BuildV0IdentifiableStakingOutputs(tag, stakerKey, fpKey, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.TimeLockPathSpendInfo()
	})
}

func Fuzz_Nosy_IdentifiableStakingInfo_UnbondingPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag []byte
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKey *btcec.PublicKey
		fill_err = tp.Fill(&fpKey)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || fpKey == nil || net == nil {
			return
		}

		i, err := BuildV0IdentifiableStakingOutputs(tag, stakerKey, fpKey, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.UnbondingPathSpendInfo()
	})
}

func Fuzz_Nosy_SpendInfo_CreateSlashingPathWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var revealedScript []byte
		fill_err = tp.Fill(&revealedScript)
		if fill_err != nil {
			return
		}
		var internalKey *btcec.PublicKey
		fill_err = tp.Fill(&internalKey)
		if fill_err != nil {
			return
		}
		var tree *txscript.IndexedTapScriptTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var covenantSigs []*schnorr.Signature
		fill_err = tp.Fill(&covenantSigs)
		if fill_err != nil {
			return
		}
		var fpSigs []*schnorr.Signature
		fill_err = tp.Fill(&fpSigs)
		if fill_err != nil {
			return
		}
		var delegatorSig *schnorr.Signature
		fill_err = tp.Fill(&delegatorSig)
		if fill_err != nil {
			return
		}
		if internalKey == nil || tree == nil || delegatorSig == nil {
			return
		}

		si, err := SpendInfoFromRevealedScript(revealedScript, internalKey, tree)
		if err != nil {
			return
		}
		si.CreateSlashingPathWitness(covenantSigs, fpSigs, delegatorSig)
	})
}

func Fuzz_Nosy_SpendInfo_CreateTimeLockPathWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var revealedScript []byte
		fill_err = tp.Fill(&revealedScript)
		if fill_err != nil {
			return
		}
		var internalKey *btcec.PublicKey
		fill_err = tp.Fill(&internalKey)
		if fill_err != nil {
			return
		}
		var tree *txscript.IndexedTapScriptTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var delegatorSig *schnorr.Signature
		fill_err = tp.Fill(&delegatorSig)
		if fill_err != nil {
			return
		}
		if internalKey == nil || tree == nil || delegatorSig == nil {
			return
		}

		si, err := SpendInfoFromRevealedScript(revealedScript, internalKey, tree)
		if err != nil {
			return
		}
		si.CreateTimeLockPathWitness(delegatorSig)
	})
}

func Fuzz_Nosy_SpendInfo_CreateUnbondingPathWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var revealedScript []byte
		fill_err = tp.Fill(&revealedScript)
		if fill_err != nil {
			return
		}
		var internalKey *btcec.PublicKey
		fill_err = tp.Fill(&internalKey)
		if fill_err != nil {
			return
		}
		var tree *txscript.IndexedTapScriptTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		var covenantSigs []*schnorr.Signature
		fill_err = tp.Fill(&covenantSigs)
		if fill_err != nil {
			return
		}
		var delegatorSig *schnorr.Signature
		fill_err = tp.Fill(&delegatorSig)
		if fill_err != nil {
			return
		}
		if internalKey == nil || tree == nil || delegatorSig == nil {
			return
		}

		si, err := SpendInfoFromRevealedScript(revealedScript, internalKey, tree)
		if err != nil {
			return
		}
		si.CreateUnbondingPathWitness(covenantSigs, delegatorSig)
	})
}

func Fuzz_Nosy_SpendInfo_GetPkScriptPath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var revealedScript []byte
		fill_err = tp.Fill(&revealedScript)
		if fill_err != nil {
			return
		}
		var internalKey *btcec.PublicKey
		fill_err = tp.Fill(&internalKey)
		if fill_err != nil {
			return
		}
		var tree *txscript.IndexedTapScriptTree
		fill_err = tp.Fill(&tree)
		if fill_err != nil {
			return
		}
		if internalKey == nil || tree == nil {
			return
		}

		si, err := SpendInfoFromRevealedScript(revealedScript, internalKey, tree)
		if err != nil {
			return
		}
		si.GetPkScriptPath()
	})
}

func Fuzz_Nosy_StakingInfo_GetOutputFetcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		sti, err := BuildStakingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		sti.GetOutputFetcher()
	})
}

func Fuzz_Nosy_StakingInfo_GetPkScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		sti, err := BuildStakingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		sti.GetPkScript()
	})
}

func Fuzz_Nosy_StakingInfo_SlashingPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		i, err := BuildStakingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.SlashingPathSpendInfo()
	})
}

func Fuzz_Nosy_StakingInfo_TimeLockPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		i, err := BuildStakingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.TimeLockPathSpendInfo()
	})
}

func Fuzz_Nosy_StakingInfo_UnbondingPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		i, err := BuildStakingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
		if err != nil {
			return
		}
		i.UnbondingPathSpendInfo()
	})
}

func Fuzz_Nosy_UnbondingInfo_SlashingPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var unbondingTime uint16
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var unbondingAmount btcutil.Amount
		fill_err = tp.Fill(&unbondingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		i, err := BuildUnbondingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, unbondingTime, unbondingAmount, net)
		if err != nil {
			return
		}
		i.SlashingPathSpendInfo()
	})
}

func Fuzz_Nosy_UnbondingInfo_TimeLockPathSpendInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKeys []*btcec.PublicKey
		fill_err = tp.Fill(&fpKeys)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var unbondingTime uint16
		fill_err = tp.Fill(&unbondingTime)
		if fill_err != nil {
			return
		}
		var unbondingAmount btcutil.Amount
		fill_err = tp.Fill(&unbondingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || net == nil {
			return
		}

		i, err := BuildUnbondingInfo(stakerKey, fpKeys, covenantKeys, covenantQuorum, unbondingTime, unbondingAmount, net)
		if err != nil {
			return
		}
		i.TimeLockPathSpendInfo()
	})
}

func Fuzz_Nosy_V0OpReturnData_Marshall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, tag []byte, stakerPublicKey []byte, finalityProviderPublicKey []byte, stakingTime []byte) {
		d, err := NewV0OpReturnData(tag, stakerPublicKey, finalityProviderPublicKey, stakingTime)
		if err != nil {
			return
		}
		d.Marshall()
	})
}

func Fuzz_Nosy_V0OpReturnData_ToTxOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, tag []byte, stakerPublicKey []byte, finalityProviderPublicKey []byte, stakingTime []byte) {
		d, err := NewV0OpReturnData(tag, stakerPublicKey, finalityProviderPublicKey, stakingTime)
		if err != nil {
			return
		}
		d.ToTxOutput()
	})
}

func Fuzz_Nosy_XonlyPubKey_Marshall__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pkBytes []byte) {
		p, err := XOnlyPublicKeyFromBytes(pkBytes)
		if err != nil {
			return
		}
		p.Marshall()
	})
}

func Fuzz_Nosy_taprootScriptHolder_scriptSpendInfoByName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var internalPubKey *btcec.PublicKey
		fill_err = tp.Fill(&internalPubKey)
		if fill_err != nil {
			return
		}
		var scripts [][]byte
		fill_err = tp.Fill(&scripts)
		if fill_err != nil {
			return
		}
		var leafHash chainhash.Hash
		fill_err = tp.Fill(&leafHash)
		if fill_err != nil {
			return
		}
		if internalPubKey == nil {
			return
		}

		t1, err := newTaprootScriptHolder(internalPubKey, scripts)
		if err != nil {
			return
		}
		t1.scriptSpendInfoByName(leafHash)
	})
}

func Fuzz_Nosy_taprootScriptHolder_taprootPkScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var internalPubKey *btcec.PublicKey
		fill_err = tp.Fill(&internalPubKey)
		if fill_err != nil {
			return
		}
		var scripts [][]byte
		fill_err = tp.Fill(&scripts)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if internalPubKey == nil || net == nil {
			return
		}

		t1, err := newTaprootScriptHolder(internalPubKey, scripts)
		if err != nil {
			return
		}
		t1.taprootPkScript(net)
	})
}

func Fuzz_Nosy_BuildV0IdentifiableStakingOutputsAndTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag []byte
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var stakerKey *btcec.PublicKey
		fill_err = tp.Fill(&stakerKey)
		if fill_err != nil {
			return
		}
		var fpKey *btcec.PublicKey
		fill_err = tp.Fill(&fpKey)
		if fill_err != nil {
			return
		}
		var covenantKeys []*btcec.PublicKey
		fill_err = tp.Fill(&covenantKeys)
		if fill_err != nil {
			return
		}
		var covenantQuorum uint32
		fill_err = tp.Fill(&covenantQuorum)
		if fill_err != nil {
			return
		}
		var stakingTime uint16
		fill_err = tp.Fill(&stakingTime)
		if fill_err != nil {
			return
		}
		var stakingAmount btcutil.Amount
		fill_err = tp.Fill(&stakingAmount)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if stakerKey == nil || fpKey == nil || net == nil {
			return
		}

		BuildV0IdentifiableStakingOutputsAndTx(tag, stakerKey, fpKey, covenantKeys, covenantQuorum, stakingTime, stakingAmount, net)
	})
}

func Fuzz_Nosy_DeriveTaprootPkScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tapScriptTree *txscript.IndexedTapScriptTree
		fill_err = tp.Fill(&tapScriptTree)
		if fill_err != nil {
			return
		}
		var internalPubKey *btcec.PublicKey
		fill_err = tp.Fill(&internalPubKey)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if tapScriptTree == nil || internalPubKey == nil || net == nil {
			return
		}

		DeriveTaprootPkScript(tapScriptTree, internalPubKey, net)
	})
}

func Fuzz_Nosy_IsPossibleV0StakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var expectedTag []byte
		fill_err = tp.Fill(&expectedTag)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		IsPossibleV0StakingTx(tx, expectedTag)
	})
}

func Fuzz_Nosy_IsSlashingRateValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rate math.LegacyDec
		fill_err = tp.Fill(&rate)
		if fill_err != nil {
			return
		}

		IsSlashingRateValid(rate)
	})
}

func Fuzz_Nosy_ParseBlkHeightAndPubKeyFromStoreKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, key []byte) {
		ParseBlkHeightAndPubKeyFromStoreKey(key)
	})
}

func Fuzz_Nosy_SortKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keys []*btcec.PublicKey
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}

		SortKeys(keys)
	})
}

func Fuzz_Nosy_aggregateScripts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var scripts [][]byte
		fill_err = tp.Fill(&scripts)
		if fill_err != nil {
			return
		}

		aggregateScripts(scripts...)
	})
}

func Fuzz_Nosy_assembleMultiSigScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkeys []*btcec.PublicKey
		fill_err = tp.Fill(&pubkeys)
		if fill_err != nil {
			return
		}
		var threshold uint32
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		var withVerify bool
		fill_err = tp.Fill(&withVerify)
		if fill_err != nil {
			return
		}

		assembleMultiSigScript(pubkeys, threshold, withVerify)
	})
}

func Fuzz_Nosy_buildMultiSigScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keys []*btcec.PublicKey
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}
		var threshold uint32
		fill_err = tp.Fill(&threshold)
		if fill_err != nil {
			return
		}
		var withVerify bool
		fill_err = tp.Fill(&withVerify)
		if fill_err != nil {
			return
		}

		buildMultiSigScript(keys, threshold, withVerify)
	})
}

func Fuzz_Nosy_buildSingleKeySigScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubKey *btcec.PublicKey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var withVerify bool
		fill_err = tp.Fill(&withVerify)
		if fill_err != nil {
			return
		}
		if pubKey == nil {
			return
		}

		buildSingleKeySigScript(pubKey, withVerify)
	})
}

func Fuzz_Nosy_buildTimeLockScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubKey *btcec.PublicKey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var lockTime uint16
		fill_err = tp.Fill(&lockTime)
		if fill_err != nil {
			return
		}
		if pubKey == nil {
			return
		}

		buildTimeLockScript(pubKey, lockTime)
	})
}

func Fuzz_Nosy_getV0OpReturnBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var out *wire.TxOut
		fill_err = tp.Fill(&out)
		if fill_err != nil {
			return
		}
		if out == nil {
			return
		}

		getV0OpReturnBytes(out)
	})
}

func Fuzz_Nosy_keyToString__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key *btcec.PublicKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		keyToString(key)
	})
}

func Fuzz_Nosy_prepareKeysForMultisigScript__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keys []*btcec.PublicKey
		fill_err = tp.Fill(&keys)
		if fill_err != nil {
			return
		}

		prepareKeysForMultisigScript(keys)
	})
}

func Fuzz_Nosy_tryToGetOpReturnDataFromOutputs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var outputs []*wire.TxOut
		fill_err = tp.Fill(&outputs)
		if fill_err != nil {
			return
		}

		tryToGetOpReturnDataFromOutputs(outputs)
	})
}

func Fuzz_Nosy_tryToGetStakingOutput__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var outputs []*wire.TxOut
		fill_err = tp.Fill(&outputs)
		if fill_err != nil {
			return
		}
		var stakingOutputPkScript []byte
		fill_err = tp.Fill(&stakingOutputPkScript)
		if fill_err != nil {
			return
		}

		tryToGetStakingOutput(outputs, stakingOutputPkScript)
	})
}

func Fuzz_Nosy_uint16FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		uint16FromBytes(b)
	})
}

func Fuzz_Nosy_uint16ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint16) {
		uint16ToBytes(v)
	})
}

func Fuzz_Nosy_unspendableKeyPathInternalPubKeyInternal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyHex string) {
		unspendableKeyPathInternalPubKeyInternal(keyHex)
	})
}
