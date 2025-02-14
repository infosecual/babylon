package types

import (
	"context"
	"testing"

	"cosmossdk.io/x/staking/types"
	"github.com/babylonlabs-io/babylon/btctxformatter"
	"github.com/babylonlabs-io/babylon/crypto/bls12381"
	"github.com/boljen/go-bitmap"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func Fuzz_Nosy_BlockHash_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		var l BlockHash
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.Equal(l)
	})
}

func Fuzz_Nosy_BlockHash_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.Marshal()
	})
}

func Fuzz_Nosy_BlockHash_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.MustMarshal()
	})
}

func Fuzz_Nosy_BlockHash_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.Size()
	})
}

func Fuzz_Nosy_BlockHash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.String()
	})
}

func Fuzz_Nosy_BlockHash_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		var bz []byte
		fill_err = tp.Fill(&bz)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.Unmarshal(bz)
	})
}

func Fuzz_Nosy_BlockHash_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh *BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		if bh == nil {
			return
		}

		bh.ValidateBasic()
	})
}

func Fuzz_Nosy_BlsKey_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsKey
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

func Fuzz_Nosy_BlsKey_GetPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPop()
	})
}

func Fuzz_Nosy_BlsKey_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsKey
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

func Fuzz_Nosy_BlsKey_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

// skipping Fuzz_Nosy_BlsKey_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BlsKey_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsKey_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsKey
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

func Fuzz_Nosy_BlsPublicKeyListResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_GetBlsPubKeyHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlsPubKeyHex()
	})
}

func Fuzz_Nosy_BlsPublicKeyListResponse_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidatorAddress()
	})
}

func Fuzz_Nosy_BlsPublicKeyListResponse_GetVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVotingPower()
	})
}

func Fuzz_Nosy_BlsPublicKeyListResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

// skipping Fuzz_Nosy_BlsPublicKeyListResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BlsPublicKeyListResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsPublicKeyListResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsPublicKeyListResponse
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

func Fuzz_Nosy_BlsSig_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsSig
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

func Fuzz_Nosy_BlsSig_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNum()
	})
}

func Fuzz_Nosy_BlsSig_GetSignerAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSignerAddress()
	})
}

func Fuzz_Nosy_BlsSig_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidatorAddress()
	})
}

func Fuzz_Nosy_BlsSig_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BlsSig
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

func Fuzz_Nosy_BlsSig_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

// skipping Fuzz_Nosy_BlsSig_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BlsSig_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_BlsSig_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *BlsSig
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

func Fuzz_Nosy_CheckpointStateUpdate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CheckpointStateUpdate
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

// skipping Fuzz_Nosy_CheckpointStateUpdate_Equal__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_CheckpointStateUpdate_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHeight()
	})
}

func Fuzz_Nosy_CheckpointStateUpdate_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockTime()
	})
}

func Fuzz_Nosy_CheckpointStateUpdate_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetState()
	})
}

func Fuzz_Nosy_CheckpointStateUpdate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *CheckpointStateUpdate
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.ToResponse()
	})
}

func Fuzz_Nosy_CheckpointStateUpdate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

// skipping Fuzz_Nosy_CheckpointStateUpdate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_CheckpointStateUpdate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdate
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHeight()
	})
}

func Fuzz_Nosy_CheckpointStateUpdateResponse_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockTime()
	})
}

func Fuzz_Nosy_CheckpointStateUpdateResponse_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetState()
	})
}

func Fuzz_Nosy_CheckpointStateUpdateResponse_GetStatusDesc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatusDesc()
	})
}

func Fuzz_Nosy_CheckpointStateUpdateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

// skipping Fuzz_Nosy_CheckpointStateUpdateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_CheckpointStateUpdateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_CheckpointStateUpdateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *CheckpointStateUpdateResponse
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

func Fuzz_Nosy_EventCheckpointAccumulating_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointAccumulating_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

// skipping Fuzz_Nosy_EventCheckpointAccumulating_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointAccumulating_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointAccumulating_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointAccumulating
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

func Fuzz_Nosy_EventCheckpointConfirmed_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointConfirmed_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

// skipping Fuzz_Nosy_EventCheckpointConfirmed_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointConfirmed_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointConfirmed_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointConfirmed
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

func Fuzz_Nosy_EventCheckpointFinalized_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointFinalized_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

// skipping Fuzz_Nosy_EventCheckpointFinalized_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointFinalized_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointFinalized_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointFinalized
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

func Fuzz_Nosy_EventCheckpointForgotten_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointForgotten_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

// skipping Fuzz_Nosy_EventCheckpointForgotten_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointForgotten_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointForgotten_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointForgotten
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

func Fuzz_Nosy_EventCheckpointSealed_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointSealed_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

// skipping Fuzz_Nosy_EventCheckpointSealed_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointSealed_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSealed_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSealed
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

func Fuzz_Nosy_EventCheckpointSubmitted_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_GetCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCheckpoint()
	})
}

func Fuzz_Nosy_EventCheckpointSubmitted_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

// skipping Fuzz_Nosy_EventCheckpointSubmitted_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventCheckpointSubmitted_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventCheckpointSubmitted_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventCheckpointSubmitted
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

func Fuzz_Nosy_EventConflictingCheckpoint_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_GetConflictingCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetConflictingCheckpoint()
	})
}

func Fuzz_Nosy_EventConflictingCheckpoint_GetLocalCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLocalCheckpoint()
	})
}

func Fuzz_Nosy_EventConflictingCheckpoint_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

// skipping Fuzz_Nosy_EventConflictingCheckpoint_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventConflictingCheckpoint_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_EventConflictingCheckpoint_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventConflictingCheckpoint
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

func Fuzz_Nosy_GenesisKey_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		_x1, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_GenesisKey_GetBlsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.GetBlsKey()
	})
}

func Fuzz_Nosy_GenesisKey_GetValPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.GetValPubkey()
	})
}

func Fuzz_Nosy_GenesisKey_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.GetValidatorAddress()
	})
}

func Fuzz_Nosy_GenesisKey_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_GenesisKey_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, dAtA []byte) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_GenesisKey_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, dAtA []byte) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_GenesisKey_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		_x1, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_GenesisKey_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_GenesisKey_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_GenesisKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_GenesisKey_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, dAtA []byte) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_GenesisKey_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		gk, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		gk.Validate()
	})
}

func Fuzz_Nosy_GenesisKey_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_GenesisKey_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, b []byte, deterministic bool) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_GenesisKey_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_GenesisKey_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_GenesisKey_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, b []byte) {
		m, err := LoadGenesisKeyFromFile(filePath)
		if err != nil {
			return
		}
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

func Fuzz_Nosy_MsgInjectedCheckpoint_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_GetCkpt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCkpt()
	})
}

func Fuzz_Nosy_MsgInjectedCheckpoint_GetExtendedCommitInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetExtendedCommitInfo()
	})
}

func Fuzz_Nosy_MsgInjectedCheckpoint_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

// skipping Fuzz_Nosy_MsgInjectedCheckpoint_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgInjectedCheckpoint_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgInjectedCheckpoint_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInjectedCheckpoint
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

func Fuzz_Nosy_MsgWrappedCreateValidator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		_x1, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		_x1, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.ValidateBasic()
	})
}

// skipping Fuzz_Nosy_MsgWrappedCreateValidator_VerifyPoP__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/crypto/types.PubKey

func Fuzz_Nosy_MsgWrappedCreateValidator_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
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
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedCreateValidator_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedCreateValidator_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidator_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgCreateVal *types.MsgCreateValidator
		fill_err = tp.Fill(&msgCreateVal)
		if fill_err != nil {
			return
		}
		var blsPK *bls12381.PublicKey
		fill_err = tp.Fill(&blsPK)
		if fill_err != nil {
			return
		}
		var pop *ProofOfPossession
		fill_err = tp.Fill(&pop)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msgCreateVal == nil || blsPK == nil || pop == nil {
			return
		}

		m, err := NewMsgWrappedCreateValidator(msgCreateVal, blsPK, pop)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

// skipping Fuzz_Nosy_MsgWrappedCreateValidatorResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_MsgWrappedCreateValidatorResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCreateValidatorResponse
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

func Fuzz_Nosy_ProofOfPossession_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_GetEd25519Sig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEd25519Sig()
	})
}

func Fuzz_Nosy_ProofOfPossession_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

// skipping Fuzz_Nosy_ProofOfPossession_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ProofOfPossession_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_ProofOfPossession_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ProofOfPossession
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNum()
	})
}

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

// skipping Fuzz_Nosy_QueryBlsPublicKeyListRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListRequest
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_GetValidatorWithBlsKeys__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidatorWithBlsKeys()
	})
}

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

// skipping Fuzz_Nosy_QueryBlsPublicKeyListResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryBlsPublicKeyListResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBlsPublicKeyListResponse
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

func Fuzz_Nosy_QueryEpochStatusRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryEpochStatusRequest(epochNum)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.GetEpochNum()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryEpochStatusRequest(epochNum)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.Size()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.String()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte, deterministic bool) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryEpochStatusRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochStatusRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryEpochStatusRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte) {
		m := NewQueryEpochStatusRequest(epochNum)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryEpochStatusResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatus()
	})
}

func Fuzz_Nosy_QueryEpochStatusResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

// skipping Fuzz_Nosy_QueryEpochStatusResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochStatusResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryEpochStatusResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		_x1 := NewQueryLastCheckpointWithStatus(status)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.GetStatus()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		_x1 := NewQueryLastCheckpointWithStatus(status)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.Size()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.String()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
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

		m := NewQueryLastCheckpointWithStatus(status)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewQueryLastCheckpointWithStatus(status)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_GetRawCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRawCheckpoint()
	})
}

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

// skipping Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryLastCheckpointWithStatusResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLastCheckpointWithStatusResponse
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

func Fuzz_Nosy_QueryRawCheckpointListRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		_x1 := NewQueryRawCheckpointListRequest(req, status)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.GetStatus()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		_x1 := NewQueryRawCheckpointListRequest(req, status)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.Size()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.String()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
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
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryRawCheckpointListRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointListRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var req *query.PageRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryRawCheckpointListRequest(req, status)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryRawCheckpointListResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListResponse_GetRawCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRawCheckpoints()
	})
}

func Fuzz_Nosy_QueryRawCheckpointListResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

// skipping Fuzz_Nosy_QueryRawCheckpointListResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointListResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointListResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointListResponse
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

func Fuzz_Nosy_QueryRawCheckpointRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryRawCheckpointRequest(epochNum)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.GetEpochNum()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryRawCheckpointRequest(epochNum)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.Size()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.String()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte, deterministic bool) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryRawCheckpointRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryRawCheckpointRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte) {
		m := NewQueryRawCheckpointRequest(epochNum)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryRawCheckpointResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_GetRawCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRawCheckpoint()
	})
}

func Fuzz_Nosy_QueryRawCheckpointResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

// skipping Fuzz_Nosy_QueryRawCheckpointResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointResponse
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryRawCheckpointsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

// skipping Fuzz_Nosy_QueryRawCheckpointsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsRequest
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryRawCheckpointsResponse_GetRawCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetRawCheckpoints()
	})
}

func Fuzz_Nosy_QueryRawCheckpointsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

// skipping Fuzz_Nosy_QueryRawCheckpointsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRawCheckpointsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRawCheckpointsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRawCheckpointsResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryRecentEpochStatusCountRequest(epochNum)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_GetEpochCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.GetEpochCount()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		_x1 := NewQueryRecentEpochStatusCountRequest(epochNum)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.Size()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.String()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, dAtA []byte) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte, deterministic bool) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryRecentEpochStatusCountRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epochNum uint64, b []byte) {
		m := NewQueryRecentEpochStatusCountRequest(epochNum)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_GetEpochCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochCount()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_GetStatusCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatusCount()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_GetTipEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTipEpoch()
	})
}

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

// skipping Fuzz_Nosy_QueryRecentEpochStatusCountResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_QueryRecentEpochStatusCountResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryRecentEpochStatusCountResponse
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

func Fuzz_Nosy_RawCheckpoint_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		_x1, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

// skipping Fuzz_Nosy_RawCheckpoint_Equal__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_RawCheckpoint_GetBitmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.GetBitmap()
	})
}

func Fuzz_Nosy_RawCheckpoint_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.GetEpochNum()
	})
}

func Fuzz_Nosy_RawCheckpoint_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_RawCheckpoint_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpoint_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpoint_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		_x1, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_RawCheckpoint_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_RawCheckpoint_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_RawCheckpoint_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_RawCheckpoint_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		r, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		r.ToResponse()
	})
}

func Fuzz_Nosy_RawCheckpoint_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpoint_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_RawCheckpoint_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
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
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_RawCheckpoint_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RawCheckpoint_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_RawCheckpoint_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_RawCheckpointResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_GetBitmap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBitmap()
	})
}

func Fuzz_Nosy_RawCheckpointResponse_GetBlockHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlockHashHex()
	})
}

func Fuzz_Nosy_RawCheckpointResponse_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNum()
	})
}

func Fuzz_Nosy_RawCheckpointResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_ToRawCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *RawCheckpointResponse
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ToRawCheckpoint()
	})
}

func Fuzz_Nosy_RawCheckpointResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

// skipping Fuzz_Nosy_RawCheckpointResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RawCheckpointResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointResponse
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

func Fuzz_Nosy_RawCheckpointWithMeta_Accumulate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var vals etypes.ValidatorSet
		fill_err = tp.Fill(&vals)
		if fill_err != nil {
			return
		}
		var signerAddr sdk.ValAddress
		fill_err = tp.Fill(&signerAddr)
		if fill_err != nil {
			return
		}
		var signerBlsKey bls12381.PublicKey
		fill_err = tp.Fill(&signerBlsKey)
		if fill_err != nil {
			return
		}
		var sig bls12381.Signature
		fill_err = tp.Fill(&sig)
		if fill_err != nil {
			return
		}
		var totalPower int64
		fill_err = tp.Fill(&totalPower)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cm := NewCheckpointWithMeta(ckpt, status)
		cm.Accumulate(vals, signerAddr, signerBlsKey, sig, totalPower)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		_x1 := NewCheckpointWithMeta(ckpt, status)
		_x1.Descriptor()
	})
}

// skipping Fuzz_Nosy_RawCheckpointWithMeta_Equal__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_RawCheckpointWithMeta_GetCkpt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.GetCkpt()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_GetLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.GetLifecycle()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_GetPowerSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.GetPowerSum()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.GetStatus()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_IsMoreMatureThanStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var s2 CheckpointStatus
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		var s3 CheckpointStatus
		fill_err = tp.Fill(&s3)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cm := NewCheckpointWithMeta(ckpt, s2)
		cm.IsMoreMatureThanStatus(s3)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.Marshal()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		_x1 := NewCheckpointWithMeta(ckpt, status)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_RecordStateUpdate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var s2 CheckpointStatus
		fill_err = tp.Fill(&s2)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var s4 CheckpointStatus
		fill_err = tp.Fill(&s4)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cm := NewCheckpointWithMeta(ckpt, s2)
		cm.RecordStateUpdate(ctx, s4)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.Reset()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.Size()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.String()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		r := NewCheckpointWithMeta(ckpt, status)
		r.ToResponse()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
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
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_RawCheckpointWithMeta_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RawCheckpointWithMeta_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_RawCheckpointWithMeta_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var status CheckpointStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		m := NewCheckpointWithMeta(ckpt, status)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_GetCkpt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCkpt()
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_GetLifecycle__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLifecycle()
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_GetPowerSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPowerSum()
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_GetStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatus()
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_GetStatusDesc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStatusDesc()
	})
}

func Fuzz_Nosy_RawCheckpointWithMetaResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

// skipping Fuzz_Nosy_RawCheckpointWithMetaResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_RawCheckpointWithMetaResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_RawCheckpointWithMetaResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *RawCheckpointWithMetaResponse
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

func Fuzz_Nosy_UnimplementedMsgServer_WrappedCreateValidator__(f *testing.F) {
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
		var req *MsgWrappedCreateValidator
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedCreateValidator(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_BlsPublicKeyList__(f *testing.F) {
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
		var req *QueryBlsPublicKeyListRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BlsPublicKeyList(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_EpochStatus__(f *testing.F) {
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
		var req *QueryEpochStatusRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EpochStatus(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_LastCheckpointWithStatus__(f *testing.F) {
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
		var req *QueryLastCheckpointWithStatusRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.LastCheckpointWithStatus(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_RawCheckpoint__(f *testing.F) {
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
		var req *QueryRawCheckpointRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.RawCheckpoint(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_RawCheckpointList__(f *testing.F) {
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
		var req *QueryRawCheckpointListRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.RawCheckpointList(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_RawCheckpoints__(f *testing.F) {
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
		var req *QueryRawCheckpointsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.RawCheckpoints(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_RecentEpochStatusCount__(f *testing.F) {
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
		var req *QueryRecentEpochStatusCountRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.RecentEpochStatusCount(ctx, req)
	})
}

func Fuzz_Nosy_ValidatorWithBlsKey_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_GetBlsPubKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBlsPubKey()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKey_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidatorAddress()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKey_GetVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetVotingPower()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKey_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

// skipping Fuzz_Nosy_ValidatorWithBlsKey_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ValidatorWithBlsKey_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKey_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKey
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_FindSubsetWithPowerSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks *ValidatorWithBlsKeySet
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		var bm bitmap.Bitmap
		fill_err = tp.Fill(&bm)
		if fill_err != nil {
			return
		}
		if ks == nil {
			return
		}

		ks.FindSubsetWithPowerSum(bm)
	})
}

func Fuzz_Nosy_ValidatorWithBlsKeySet_GetBLSKeySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks *ValidatorWithBlsKeySet
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		if ks == nil {
			return
		}

		ks.GetBLSKeySet()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKeySet_GetTotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ks *ValidatorWithBlsKeySet
		fill_err = tp.Fill(&ks)
		if fill_err != nil {
			return
		}
		if ks == nil {
			return
		}

		ks.GetTotalPower()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKeySet_GetValSet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValSet()
	})
}

func Fuzz_Nosy_ValidatorWithBlsKeySet_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

// skipping Fuzz_Nosy_ValidatorWithBlsKeySet_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ValidatorWithBlsKeySet_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_ValidatorWithBlsKeySet_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorWithBlsKeySet
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

func Fuzz_Nosy_VoteExtension_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *VoteExtension
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

func Fuzz_Nosy_VoteExtension_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNum()
	})
}

func Fuzz_Nosy_VoteExtension_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSigner()
	})
}

func Fuzz_Nosy_VoteExtension_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidatorAddress()
	})
}

func Fuzz_Nosy_VoteExtension_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *VoteExtension
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

func Fuzz_Nosy_VoteExtension_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_ToBLSSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ve *VoteExtension
		fill_err = tp.Fill(&ve)
		if fill_err != nil {
			return
		}
		if ve == nil {
			return
		}

		ve.ToBLSSig()
	})
}

func Fuzz_Nosy_VoteExtension_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

// skipping Fuzz_Nosy_VoteExtension_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_VoteExtension_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

func Fuzz_Nosy_VoteExtension_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *VoteExtension
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

// skipping Fuzz_Nosy_msgClient_WrappedCreateValidator__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_BlsPublicKeyList__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_EpochStatus__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_LastCheckpointWithStatus__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_RawCheckpoint__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_RawCheckpointList__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_RawCheckpoints__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_RecentEpochStatusCount__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

func Fuzz_Nosy_BlockHash_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bh BlockHash
		fill_err = tp.Fill(&bh)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}

		bh.MarshalTo(d2)
	})
}

func Fuzz_Nosy_BlsSig_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m BlsSig
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Hash()
	})
}

func Fuzz_Nosy_BlsSigHash_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m BlsSigHash
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}

		m.Bytes()
	})
}

// skipping Fuzz_Nosy_BlsSigner_BlsPubKey__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.BlsSigner

// skipping Fuzz_Nosy_BlsSigner_SignMsgWithBls__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.BlsSigner

func Fuzz_Nosy_CheckpointStatus_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 CheckpointStatus
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_CheckpointStatus_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x CheckpointStatus
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

// skipping Fuzz_Nosy_CheckpointingHooks_AfterBlsKeyRegistered__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_CheckpointingHooks_AfterRawCheckpointBlsSigVerified__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_CheckpointingHooks_AfterRawCheckpointConfirmed__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_CheckpointingHooks_AfterRawCheckpointFinalized__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_CheckpointingHooks_AfterRawCheckpointForgotten__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_CheckpointingHooks_AfterRawCheckpointSealed__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.CheckpointingHooks

// skipping Fuzz_Nosy_EpochingKeeper_CheckMsgCreateValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_EnqueueMsg__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetEpochNumByHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetTotalVotingPower__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_GetValidatorSet__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_EpochingKeeper_StkMsgCreateValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.EpochingKeeper

// skipping Fuzz_Nosy_MsgClient_WrappedCreateValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MsgClient

// skipping Fuzz_Nosy_MsgServer_WrappedCreateValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MsgServer

// skipping Fuzz_Nosy_MsgWrappedCreateValidator_UnpackInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.AnyUnpacker

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterBlsKeyRegistered__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterRawCheckpointBlsSigVerified__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterRawCheckpointConfirmed__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterRawCheckpointFinalized__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterRawCheckpointForgotten__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_MultiCheckpointingHooks_AfterRawCheckpointSealed__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.MultiCheckpointingHooks

// skipping Fuzz_Nosy_ProofOfPossession_IsValid__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/crypto/types.PubKey

// skipping Fuzz_Nosy_QueryClient_BlsPublicKeyList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_EpochStatus__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_LastCheckpointWithStatus__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_RawCheckpoint__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_RawCheckpointList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_RawCheckpoints__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_RecentEpochStatusCount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_BlsPublicKeyList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_EpochStatus__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_LastCheckpointWithStatus__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_RawCheckpoint__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_RawCheckpointList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_RawCheckpoints__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_RecentEpochStatusCount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/checkpointing/types.QueryServer

func Fuzz_Nosy_RawCheckpoint_Hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.Hash()
	})
}

func Fuzz_Nosy_RawCheckpoint_HashStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.HashStr()
	})
}

func Fuzz_Nosy_RawCheckpoint_SignedMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		m, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		m.SignedMsg()
	})
}

func Fuzz_Nosy_RawCheckpoint_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var btcCkpt *btctxformatter.RawBtcCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if btcCkpt == nil {
			return
		}

		ckpt, err := FromBTCCkptToRawCkpt(btcCkpt)
		if err != nil {
			return
		}
		ckpt.ValidateBasic()
	})
}

func Fuzz_Nosy_RawCkptHash_Bytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		m, err := FromStringToCkptHash(s)
		if err != nil {
			return
		}
		m.Bytes()
	})
}

func Fuzz_Nosy_RawCkptHash_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s string
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var h RawCkptHash
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}

		m, err := FromStringToCkptHash(s)
		if err != nil {
			return
		}
		m.Equals(h)
	})
}

func Fuzz_Nosy_RawCkptHash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		m, err := FromStringToCkptHash(s)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_AddrToBlsKeyKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		AddrToBlsKeyKey(valAddr)
	})
}

func Fuzz_Nosy_BlsKeyToAddrKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pk bls12381.PublicKey
		fill_err = tp.Fill(&pk)
		if fill_err != nil {
			return
		}

		BlsKeyToAddrKey(pk)
	})
}

// skipping Fuzz_Nosy_CkptWithMetaToBytes__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.BinaryCodec

func Fuzz_Nosy_CkptsObjectKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epoch uint64) {
		CkptsObjectKey(epoch)
	})
}

func Fuzz_Nosy_GetSignBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epoch uint64, hash []byte) {
		GetSignBytes(epoch, hash)
	})
}

func Fuzz_Nosy_KeyPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		KeyPrefix(p)
	})
}

func Fuzz_Nosy_NewCheckpointStateUpdateResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cs []*CheckpointStateUpdate
		fill_err = tp.Fill(&cs)
		if fill_err != nil {
			return
		}

		NewCheckpointStateUpdateResponse(cs)
	})
}

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

func Fuzz_Nosy_ValidatorBlsKeySetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, epoch uint64) {
		ValidatorBlsKeySetKey(epoch)
	})
}

// skipping Fuzz_Nosy_ValidatorBlsKeySetToBytes__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec.BinaryCodec

// skipping Fuzz_Nosy__Msg_WrappedCreateValidator_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_BlsPublicKeyList_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_EpochStatus_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_LastCheckpointWithStatus_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_RawCheckpointList_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_RawCheckpoint_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_RawCheckpoints_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_RecentEpochStatusCount_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintBlsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintBlsKey(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintCheckpoint(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintEvents(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintGenesis(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintQuery(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintTx(dAtA, offset, v)
	})
}

func Fuzz_Nosy_hash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fields [][]byte
		fill_err = tp.Fill(&fields)
		if fill_err != nil {
			return
		}

		hash(fields)
	})
}

// skipping Fuzz_Nosy_local_request_Query_BlsPublicKeyList_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_EpochStatus_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_LastCheckpointWithStatus_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_RawCheckpointList_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_RawCheckpoint_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_RawCheckpoints_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_RecentEpochStatusCount_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_BlsPublicKeyList_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_EpochStatus_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_LastCheckpointWithStatus_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_RawCheckpointList_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_RawCheckpoint_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_RawCheckpoints_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_RecentEpochStatusCount_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipBlsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipBlsKey(dAtA)
	})
}

func Fuzz_Nosy_skipCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipCheckpoint(dAtA)
	})
}

func Fuzz_Nosy_skipEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipEvents(dAtA)
	})
}

func Fuzz_Nosy_skipGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipGenesis(dAtA)
	})
}

func Fuzz_Nosy_skipQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipQuery(dAtA)
	})
}

func Fuzz_Nosy_skipTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipTx(dAtA)
	})
}

func Fuzz_Nosy_sovBlsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovBlsKey(x)
	})
}

func Fuzz_Nosy_sovCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovCheckpoint(x)
	})
}

func Fuzz_Nosy_sovEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovEvents(x)
	})
}

func Fuzz_Nosy_sovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovGenesis(x)
	})
}

func Fuzz_Nosy_sovQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovQuery(x)
	})
}

func Fuzz_Nosy_sovTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovTx(x)
	})
}

func Fuzz_Nosy_sozBlsKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozBlsKey(x)
	})
}

func Fuzz_Nosy_sozCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozCheckpoint(x)
	})
}

func Fuzz_Nosy_sozEvents__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozEvents(x)
	})
}

func Fuzz_Nosy_sozGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozGenesis(x)
	})
}

func Fuzz_Nosy_sozQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozQuery(x)
	})
}

func Fuzz_Nosy_sozTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozTx(x)
	})
}
