package types

import (
	"context"
	"math/big"
	"testing"
	"time"

	"cosmossdk.io/math"
	bbn "github.com/babylonlabs-io/babylon/types"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_BTCHeaderInfo_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		_x1 := NewBTCHeaderInfo(header, headerHash, height, work)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Eq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var other *BTCHeaderInfo
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil || other == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Eq(other)
	})
}

func Fuzz_Nosy_BTCHeaderInfo_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.GetHeight()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_HasParent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var parent *BTCHeaderInfo
		fill_err = tp.Fill(&parent)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil || parent == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.HasParent(parent)
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfo_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfo_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		_x1 := NewBTCHeaderInfo(header, headerHash, height, work)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Reset()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Size()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.String()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		b := NewBTCHeaderInfo(header, headerHash, height, work)
		b.ToResponse()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfo_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.Validate()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
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
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCHeaderInfo_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCHeaderInfo_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCHeaderInfo_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfo(header, headerHash, height, work)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		_x1 := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Eq__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var other *BTCHeaderInfo
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil || other == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.Eq(other)
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_GetHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.GetHashHex()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_GetHeaderHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.GetHeaderHex()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.GetHeight()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.Marshal()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		_x1 := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.Reset()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.Size()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.String()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
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
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_BTCHeaderInfoResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_BTCHeaderInfoResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_BTCHeaderInfoResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var header *bbn.BTCHeaderBytes
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var headerHash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var work *math.Uint
		fill_err = tp.Fill(&work)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if header == nil || headerHash == nil || work == nil {
			return
		}

		m := NewBTCHeaderInfoResponse(header, headerHash, height, work)
		m.XXX_Unmarshal(b)
	})
}

// skipping Fuzz_Nosy_BtcLightClient_InsertHeaders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BtcChainReadStore

func Fuzz_Nosy_BtcLightClient_checkHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var ctx *lightChainCtx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var s *storeWithExtensionChain
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var parentHeaderInfo *localHeaderInfo
		fill_err = tp.Fill(&parentHeaderInfo)
		if fill_err != nil {
			return
		}
		var blockHeader *wire.BlockHeader
		fill_err = tp.Fill(&blockHeader)
		if fill_err != nil {
			return
		}
		if params == nil || ctx == nil || s == nil || parentHeaderInfo == nil || blockHeader == nil {
			return
		}

		l := NewBtcLightClient(params, ctx)
		l.checkHeader(s, parentHeaderInfo, blockHeader)
	})
}

func Fuzz_Nosy_BtcLightClient_processNewHeadersChain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var ctx *lightChainCtx
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var chainParent *localHeaderInfo
		fill_err = tp.Fill(&chainParent)
		if fill_err != nil {
			return
		}
		var chain []*wire.BlockHeader
		fill_err = tp.Fill(&chain)
		if fill_err != nil {
			return
		}
		if params == nil || ctx == nil || store == nil || chainParent == nil {
			return
		}

		l := NewBtcLightClient(params, ctx)
		l.processNewHeadersChain(store, chainParent, chain)
	})
}

func Fuzz_Nosy_DisableHeaderInTheFutureValidationTimeSource_AddTimeSample__(f *testing.F) {
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
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 time.Time
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		d := NewDisableHeaderInTheFutureValidationTimeSource(header)
		d.AddTimeSample(_x2, _x3)
	})
}

func Fuzz_Nosy_DisableHeaderInTheFutureValidationTimeSource_AdjustedTime__(f *testing.F) {
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

		d := NewDisableHeaderInTheFutureValidationTimeSource(header)
		d.AdjustedTime()
	})
}

func Fuzz_Nosy_DisableHeaderInTheFutureValidationTimeSource_Offset__(f *testing.F) {
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

		d := NewDisableHeaderInTheFutureValidationTimeSource(header)
		d.Offset()
	})
}

func Fuzz_Nosy_EventBTCHeaderInserted_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeader()
	})
}

func Fuzz_Nosy_EventBTCHeaderInserted_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

// skipping Fuzz_Nosy_EventBTCHeaderInserted_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCHeaderInserted_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCHeaderInserted_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCHeaderInserted
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

func Fuzz_Nosy_EventBTCRollBack_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeader()
	})
}

func Fuzz_Nosy_EventBTCRollBack_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

// skipping Fuzz_Nosy_EventBTCRollBack_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCRollBack_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollBack_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollBack
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

func Fuzz_Nosy_EventBTCRollForward_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeader()
	})
}

func Fuzz_Nosy_EventBTCRollForward_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

// skipping Fuzz_Nosy_EventBTCRollForward_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBTCRollForward_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_EventBTCRollForward_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBTCRollForward
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

func Fuzz_Nosy_MsgInsertHeaders_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		_x1, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_GetSigner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.GetSigner()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgInsertHeaders_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgInsertHeaders_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		_x1, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_ReporterAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		msg, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		msg.ReporterAddress()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgInsertHeaders_ValidateHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}
		var powLimit *big.Int
		fill_err = tp.Fill(&powLimit)
		if fill_err != nil {
			return
		}
		if powLimit == nil {
			return
		}

		msg, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		msg.ValidateHeaders(powLimit)
	})
}

func Fuzz_Nosy_MsgInsertHeaders_ValidateStateless__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		msg, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		msg.ValidateStateless()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
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

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgInsertHeaders_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgInsertHeaders_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgInsertHeaders_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var signer sdk.AccAddress
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var headersHex string
		fill_err = tp.Fill(&headersHex)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m, err := NewMsgInsertHeaders(signer, headersHex)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgInsertHeadersResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

// skipping Fuzz_Nosy_MsgInsertHeadersResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgInsertHeadersResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_MsgInsertHeadersResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgInsertHeadersResponse
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

func Fuzz_Nosy_Params_AllowAllReporters__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		p, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		p.AllowAllReporters()
	})
}

func Fuzz_Nosy_Params_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		_x1, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

// skipping Fuzz_Nosy_Params_Equal__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Params_GetInsertHeadersAllowList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.GetInsertHeadersAllowList()
	})
}

func Fuzz_Nosy_Params_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_Params_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Params_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Params_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		_x1, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Params_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_Params_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_Params_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_Params_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Params_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_Params_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
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

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Params_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Params_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_Params_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBaseHeaderRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryBaseHeaderRequest()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryBaseHeaderRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryBaseHeaderRequest()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryBaseHeaderRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryBaseHeaderRequest()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryBaseHeaderRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := NewQueryBaseHeaderRequest()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryBaseHeaderRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBaseHeaderRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := NewQueryBaseHeaderRequest()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryBaseHeaderResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeader()
	})
}

func Fuzz_Nosy_QueryBaseHeaderResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

// skipping Fuzz_Nosy_QueryBaseHeaderResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryBaseHeaderResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryBaseHeaderResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryBaseHeaderResponse
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

func Fuzz_Nosy_QueryContainsBytesRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_GetHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHash()
	})
}

func Fuzz_Nosy_QueryContainsBytesRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

// skipping Fuzz_Nosy_QueryContainsBytesRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryContainsBytesRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesRequest
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

func Fuzz_Nosy_QueryContainsBytesResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_GetContains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetContains()
	})
}

func Fuzz_Nosy_QueryContainsBytesResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

// skipping Fuzz_Nosy_QueryContainsBytesResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryContainsBytesResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsBytesResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsBytesResponse
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

func Fuzz_Nosy_QueryContainsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		_x1, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryContainsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryContainsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryContainsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryContainsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		_x1, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryContainsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_QueryContainsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_QueryContainsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_QueryContainsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryContainsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryContainsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, b []byte, deterministic bool) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryContainsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryContainsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryContainsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, b []byte) {
		m, err := NewQueryContainsRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryContainsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_GetContains__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetContains()
	})
}

func Fuzz_Nosy_QueryContainsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

// skipping Fuzz_Nosy_QueryContainsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryContainsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryContainsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryContainsResponse
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

func Fuzz_Nosy_QueryHashesRequest_Descriptor__(f *testing.F) {
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
		if req == nil {
			return
		}

		_x1 := NewQueryHashesRequest(req)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryHashesRequest_GetPagination__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryHashesRequest_Marshal__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryHashesRequest_MarshalTo__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryHashesRequest_MarshalToSizedBuffer__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryHashesRequest_ProtoMessage__(f *testing.F) {
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
		if req == nil {
			return
		}

		_x1 := NewQueryHashesRequest(req)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryHashesRequest_Reset__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryHashesRequest_Size__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.Size()
	})
}

func Fuzz_Nosy_QueryHashesRequest_String__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.String()
	})
}

func Fuzz_Nosy_QueryHashesRequest_Unmarshal__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryHashesRequest_XXX_DiscardUnknown__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryHashesRequest_XXX_Marshal__(f *testing.F) {
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

		m := NewQueryHashesRequest(req)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryHashesRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryHashesRequest_XXX_Size__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryHashesRequest_XXX_Unmarshal__(f *testing.F) {
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
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryHashesRequest(req)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryHashesResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

// skipping Fuzz_Nosy_QueryHashesResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryHashesResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHashesResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHashesResponse
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

func Fuzz_Nosy_QueryHeaderDepthRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		_x1, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_GetHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.GetHash()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		_x1, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.Reset()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.Size()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.String()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, dAtA []byte) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, b []byte, deterministic bool) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryHeaderDepthRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryHeaderDepthRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryHeaderDepthRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, hash string, b []byte) {
		m, err := NewQueryHeaderDepthRequest(hash)
		if err != nil {
			return
		}
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryHeaderDepthResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_GetDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDepth()
	})
}

func Fuzz_Nosy_QueryHeaderDepthResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

// skipping Fuzz_Nosy_QueryHeaderDepthResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryHeaderDepthResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryHeaderDepthResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryHeaderDepthResponse
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

func Fuzz_Nosy_QueryMainChainRequest_Descriptor__(f *testing.F) {
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
		if req == nil {
			return
		}

		_x1 := NewQueryMainChainRequest(req)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_GetPagination__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.GetPagination()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_Marshal__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.Marshal()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_MarshalTo__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryMainChainRequest_MarshalToSizedBuffer__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryMainChainRequest_ProtoMessage__(f *testing.F) {
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
		if req == nil {
			return
		}

		_x1 := NewQueryMainChainRequest(req)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_Reset__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.Reset()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_Size__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.Size()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_String__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.String()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_Unmarshal__(f *testing.F) {
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
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryMainChainRequest_XXX_DiscardUnknown__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_XXX_Marshal__(f *testing.F) {
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

		m := NewQueryMainChainRequest(req)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryMainChainRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryMainChainRequest_XXX_Size__(f *testing.F) {
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
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_QueryMainChainRequest_XXX_Unmarshal__(f *testing.F) {
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
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if req == nil {
			return
		}

		m := NewQueryMainChainRequest(req)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryMainChainResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_GetHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeaders()
	})
}

func Fuzz_Nosy_QueryMainChainResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

// skipping Fuzz_Nosy_QueryMainChainResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryMainChainResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryMainChainResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryMainChainResponse
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

func Fuzz_Nosy_QueryTipRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryTipRequest()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_QueryTipRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryTipRequest()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_QueryTipRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := NewQueryTipRequest()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_QueryTipRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := NewQueryTipRequest()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_QueryTipRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryTipRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := NewQueryTipRequest()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryTipResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHeader()
	})
}

func Fuzz_Nosy_QueryTipResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

// skipping Fuzz_Nosy_QueryTipResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryTipResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_QueryTipResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryTipResponse
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

func Fuzz_Nosy_UnimplementedMsgServer_InsertHeaders__(f *testing.F) {
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
		var req *MsgInsertHeaders
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.InsertHeaders(ctx, req)
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

func Fuzz_Nosy_UnimplementedQueryServer_BaseHeader__(f *testing.F) {
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
		var req *QueryBaseHeaderRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.BaseHeader(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_Contains__(f *testing.F) {
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
		var req *QueryContainsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.Contains(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_ContainsBytes__(f *testing.F) {
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
		var req *QueryContainsBytesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.ContainsBytes(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_Hashes__(f *testing.F) {
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
		var req *QueryHashesRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.Hashes(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_HeaderDepth__(f *testing.F) {
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
		var req *QueryHeaderDepthRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.HeaderDepth(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_MainChain__(f *testing.F) {
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
		var req *QueryMainChainRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.MainChain(ctx, req)
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

func Fuzz_Nosy_UnimplementedQueryServer_Tip__(f *testing.F) {
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
		var req *QueryTipRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.Tip(ctx, req)
	})
}

func Fuzz_Nosy_lightChainCtx_BlocksPerRetarget__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.BlocksPerRetarget()
	})
}

func Fuzz_Nosy_lightChainCtx_ChainParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.ChainParams()
	})
}

func Fuzz_Nosy_lightChainCtx_FindPreviousCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.FindPreviousCheckpoint()
	})
}

func Fuzz_Nosy_lightChainCtx_MaxRetargetTimespan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.MaxRetargetTimespan()
	})
}

func Fuzz_Nosy_lightChainCtx_MinRetargetTimespan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.MinRetargetTimespan()
	})
}

func Fuzz_Nosy_lightChainCtx_VerifyCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		var _x2 int32
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *chainhash.Hash
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if params == nil || _x3 == nil {
			return
		}

		l := newLightChainCtxFromParams(params)
		l.VerifyCheckpoint(_x2, _x3)
	})
}

func Fuzz_Nosy_lightHeaderCtx_Bits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		if header == nil || store == nil {
			return
		}

		l := newLightHeaderCtx(height, header, store)
		l.Bits()
	})
}

func Fuzz_Nosy_lightHeaderCtx_Height__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		if header == nil || store == nil {
			return
		}

		l := newLightHeaderCtx(height, header, store)
		l.Height()
	})
}

func Fuzz_Nosy_lightHeaderCtx_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		if header == nil || store == nil {
			return
		}

		l := newLightHeaderCtx(height, header, store)
		l.Parent()
	})
}

func Fuzz_Nosy_lightHeaderCtx_RelativeAncestorCtx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var distance int32
		fill_err = tp.Fill(&distance)
		if fill_err != nil {
			return
		}
		if header == nil || store == nil {
			return
		}

		l := newLightHeaderCtx(height, header, store)
		l.RelativeAncestorCtx(distance)
	})
}

func Fuzz_Nosy_lightHeaderCtx_Timestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		var store *storeWithExtensionChain
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		if header == nil || store == nil {
			return
		}

		l := newLightHeaderCtx(height, header, store)
		l.Timestamp()
	})
}

func Fuzz_Nosy_localHeaderInfo_toBTCHeaderInfo__(f *testing.F) {
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
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var totalWork math.Uint
		fill_err = tp.Fill(&totalWork)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		i := newLocalHeaderInfo(header, height, totalWork)
		i.toBTCHeaderInfo()
	})
}

// skipping Fuzz_Nosy_msgClient_InsertHeaders__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_BaseHeader__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Contains__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_ContainsBytes__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Hashes__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_HeaderDepth__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_MainChain__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Params__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Tip__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

func Fuzz_Nosy_storeWithExtensionChain_addHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *storeWithExtensionChain
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var header *localHeaderInfo
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if s == nil || header == nil {
			return
		}

		s.addHeader(header)
	})
}

func Fuzz_Nosy_storeWithExtensionChain_getHeaderAtHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *storeWithExtensionChain
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getHeaderAtHeight(height)
	})
}

// skipping Fuzz_Nosy_BTCLightClientHooks_AfterBTCHeaderInserted__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCLightClientHooks

// skipping Fuzz_Nosy_BTCLightClientHooks_AfterBTCRollBack__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCLightClientHooks

// skipping Fuzz_Nosy_BTCLightClientHooks_AfterBTCRollForward__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BTCLightClientHooks

// skipping Fuzz_Nosy_BtcChainReadStore_GetHeaderByHash__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BtcChainReadStore

// skipping Fuzz_Nosy_BtcChainReadStore_GetHeaderByHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BtcChainReadStore

// skipping Fuzz_Nosy_BtcChainReadStore_GetTip__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.BtcChainReadStore

// skipping Fuzz_Nosy_IncentiveKeeper_IndexRefundableMsg__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.IncentiveKeeper

// skipping Fuzz_Nosy_MsgClient_InsertHeaders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MsgClient

// skipping Fuzz_Nosy_MsgServer_InsertHeaders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MsgServer

// skipping Fuzz_Nosy_MultiBTCLightClientHooks_AfterBTCHeaderInserted__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MultiBTCLightClientHooks

// skipping Fuzz_Nosy_MultiBTCLightClientHooks_AfterBTCRollBack__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MultiBTCLightClientHooks

// skipping Fuzz_Nosy_MultiBTCLightClientHooks_AfterBTCRollForward__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.MultiBTCLightClientHooks

func Fuzz_Nosy_Params_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var allowedAddresses []string
		fill_err = tp.Fill(&allowedAddresses)
		if fill_err != nil {
			return
		}

		p, err := NewParamsValidate(allowedAddresses)
		if err != nil {
			return
		}
		p.Validate()
	})
}

// skipping Fuzz_Nosy_QueryClient_BaseHeader__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Contains__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_ContainsBytes__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Hashes__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_HeaderDepth__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_MainChain__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Tip__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_BaseHeader__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Contains__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_ContainsBytes__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Hashes__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_HeaderDepth__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_MainChain__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Tip__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/btclightclient/types.QueryServer

func Fuzz_Nosy_BlocksPerRetarget__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if params == nil {
			return
		}

		BlocksPerRetarget(params)
	})
}

func Fuzz_Nosy_HeadersObjectHeightKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var hash *bbn.BTCHeaderHashBytes
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if hash == nil {
			return
		}

		HeadersObjectHeightKey(hash)
	})
}

func Fuzz_Nosy_HeadersObjectKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, height uint32) {
		HeadersObjectKey(height)
	})
}

func Fuzz_Nosy_IsRetargetBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var info *BTCHeaderInfo
		fill_err = tp.Fill(&info)
		if fill_err != nil {
			return
		}
		var params *chaincfg.Params
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if info == nil || params == nil {
			return
		}

		IsRetargetBlock(info, params)
	})
}

func Fuzz_Nosy_ParseBTCHeadersToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var infos []*BTCHeaderInfo
		fill_err = tp.Fill(&infos)
		if fill_err != nil {
			return
		}

		ParseBTCHeadersToResponse(infos)
	})
}

func Fuzz_Nosy_RegisterCodec__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *codec.LegacyAmino
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		RegisterCodec(_x1)
	})
}

// skipping Fuzz_Nosy_RegisterInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.InterfaceRegistry

// skipping Fuzz_Nosy_RegisterMsgServer__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy_RegisterQueryServer__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy__Msg_InsertHeaders_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_UpdateParams_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_BaseHeader_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_ContainsBytes_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Contains_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Hashes_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_HeaderDepth_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_MainChain_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Params_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Tip_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintBtclightclient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintBtclightclient(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintEvent(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintGenesis(dAtA, offset, v)
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

func Fuzz_Nosy_encodeVarintTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintTx(dAtA, offset, v)
	})
}

func Fuzz_Nosy_headersFormChain__(f *testing.F) {
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

		headersFormChain(headers)
	})
}

// skipping Fuzz_Nosy_local_request_Query_BaseHeader_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_ContainsBytes_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Contains_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Hashes_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_HeaderDepth_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_MainChain_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Tip_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_BaseHeader_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_ContainsBytes_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Contains_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Hashes_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_HeaderDepth_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_MainChain_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Tip_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipBtclightclient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipBtclightclient(dAtA)
	})
}

func Fuzz_Nosy_skipEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipEvent(dAtA)
	})
}

func Fuzz_Nosy_skipGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipGenesis(dAtA)
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

func Fuzz_Nosy_skipTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipTx(dAtA)
	})
}

func Fuzz_Nosy_sovBtclightclient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovBtclightclient(x)
	})
}

func Fuzz_Nosy_sovEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovEvent(x)
	})
}

func Fuzz_Nosy_sovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovGenesis(x)
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

func Fuzz_Nosy_sovTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovTx(x)
	})
}

func Fuzz_Nosy_sozBtclightclient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozBtclightclient(x)
	})
}

func Fuzz_Nosy_sozEvent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozEvent(x)
	})
}

func Fuzz_Nosy_sozGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozGenesis(x)
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

func Fuzz_Nosy_sozTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozTx(x)
	})
}

func Fuzz_Nosy_toBTCHeaderInfos__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var infos []*localHeaderInfo
		fill_err = tp.Fill(&infos)
		if fill_err != nil {
			return
		}

		toBTCHeaderInfos(infos)
	})
}
