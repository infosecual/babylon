package types

import (
	"context"
	"testing"
	"time"

	"github.com/boljen/go-bitmap"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
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

func Fuzz_Nosy_DelegationLifecycle_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_GetDelAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelAddr()
	})
}

func Fuzz_Nosy_DelegationLifecycle_GetDelLife__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelLife()
	})
}

func Fuzz_Nosy_DelegationLifecycle_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

// skipping Fuzz_Nosy_DelegationLifecycle_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_DelegationLifecycle_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationLifecycle_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationLifecycle
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

func Fuzz_Nosy_DelegationStateUpdate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAmount()
	})
}

func Fuzz_Nosy_DelegationStateUpdate_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_GetValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValAddr()
	})
}

func Fuzz_Nosy_DelegationStateUpdate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

// skipping Fuzz_Nosy_DelegationStateUpdate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_DelegationStateUpdate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_DelegationStateUpdate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DelegationStateUpdate
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

func Fuzz_Nosy_Epoch_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		_x1 := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_Epoch_GetCurrentEpochInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetCurrentEpochInterval()
	})
}

func Fuzz_Nosy_Epoch_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_Epoch_GetFirstBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetFirstBlockHeight()
	})
}

func Fuzz_Nosy_Epoch_GetLastBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetLastBlockTime()
	})
}

func Fuzz_Nosy_Epoch_GetSealerAppHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetSealerAppHash()
	})
}

func Fuzz_Nosy_Epoch_GetSealerBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.GetSealerBlockHash()
	})
}

func Fuzz_Nosy_Epoch_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.Marshal()
	})
}

func Fuzz_Nosy_Epoch_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Epoch_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Epoch_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		_x1 := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_Epoch_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.Reset()
	})
}

func Fuzz_Nosy_Epoch_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.Size()
	})
}

func Fuzz_Nosy_Epoch_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.String()
	})
}

func Fuzz_Nosy_Epoch_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.ToResponse()
	})
}

func Fuzz_Nosy_Epoch_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Epoch_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_Epoch_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
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
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Epoch_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Epoch_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_Epoch_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		m := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EpochResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EpochResponse
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

func Fuzz_Nosy_EpochResponse_GetCurrentEpochInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCurrentEpochInterval()
	})
}

func Fuzz_Nosy_EpochResponse_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_EpochResponse_GetFirstBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetFirstBlockHeight()
	})
}

func Fuzz_Nosy_EpochResponse_GetLastBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLastBlockTime()
	})
}

func Fuzz_Nosy_EpochResponse_GetSealerAppHashHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSealerAppHashHex()
	})
}

func Fuzz_Nosy_EpochResponse_GetSealerBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSealerBlockHash()
	})
}

func Fuzz_Nosy_EpochResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EpochResponse
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

func Fuzz_Nosy_EpochResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

// skipping Fuzz_Nosy_EpochResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EpochResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EpochResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EpochResponse
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

func Fuzz_Nosy_EventBeginEpoch_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_EventBeginEpoch_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

// skipping Fuzz_Nosy_EventBeginEpoch_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventBeginEpoch_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventBeginEpoch_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventBeginEpoch
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

func Fuzz_Nosy_EventEndEpoch_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_EventEndEpoch_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

// skipping Fuzz_Nosy_EventEndEpoch_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventEndEpoch_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventEndEpoch_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventEndEpoch
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

func Fuzz_Nosy_EventHandleQueuedMsg_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_EventHandleQueuedMsg_GetError__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetError()
	})
}

func Fuzz_Nosy_EventHandleQueuedMsg_GetHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_GetMsgId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgId()
	})
}

func Fuzz_Nosy_EventHandleQueuedMsg_GetOriginalEventType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetOriginalEventType()
	})
}

func Fuzz_Nosy_EventHandleQueuedMsg_GetTxId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTxId()
	})
}

func Fuzz_Nosy_EventHandleQueuedMsg_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

// skipping Fuzz_Nosy_EventHandleQueuedMsg_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventHandleQueuedMsg_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventHandleQueuedMsg_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventHandleQueuedMsg
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

func Fuzz_Nosy_EventSlashThreshold_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		_x1 := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_EventSlashThreshold_GetSlashedValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.GetSlashedValidators()
	})
}

func Fuzz_Nosy_EventSlashThreshold_GetSlashedVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.GetSlashedVotingPower()
	})
}

func Fuzz_Nosy_EventSlashThreshold_GetTotalVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.GetTotalVotingPower()
	})
}

func Fuzz_Nosy_EventSlashThreshold_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.Marshal()
	})
}

func Fuzz_Nosy_EventSlashThreshold_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_EventSlashThreshold_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_EventSlashThreshold_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		_x1 := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_EventSlashThreshold_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.Reset()
	})
}

func Fuzz_Nosy_EventSlashThreshold_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.Size()
	})
}

func Fuzz_Nosy_EventSlashThreshold_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.String()
	})
}

func Fuzz_Nosy_EventSlashThreshold_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_EventSlashThreshold_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_EventSlashThreshold_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
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

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_EventSlashThreshold_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventSlashThreshold_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_EventSlashThreshold_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slashedVotingPower int64
		fill_err = tp.Fill(&slashedVotingPower)
		if fill_err != nil {
			return
		}
		var totalVotingPower int64
		fill_err = tp.Fill(&totalVotingPower)
		if fill_err != nil {
			return
		}
		var slashedValSet ValidatorSet
		fill_err = tp.Fill(&slashedValSet)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		m := NewEventSlashThreshold(slashedVotingPower, totalVotingPower, slashedValSet)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAmount()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetDelegatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetDenom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDenom()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetDestinationValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDestinationValidatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_GetSourceValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetSourceValidatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedBeginRedelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

// skipping Fuzz_Nosy_EventWrappedBeginRedelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedBeginRedelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedBeginRedelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedBeginRedelegate
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAmount()
	})
}

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_GetCreationHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCreationHeight()
	})
}

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_GetDelegatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

// skipping Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedCancelUnbondingDelegation_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedCancelUnbondingDelegation
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

func Fuzz_Nosy_EventWrappedDelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAmount()
	})
}

func Fuzz_Nosy_EventWrappedDelegate_GetDelegatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedDelegate_GetDenom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDenom()
	})
}

func Fuzz_Nosy_EventWrappedDelegate_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedDelegate_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

// skipping Fuzz_Nosy_EventWrappedDelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedDelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedDelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedDelegate
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

func Fuzz_Nosy_EventWrappedEditValidator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedEditValidator_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

// skipping Fuzz_Nosy_EventWrappedEditValidator_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedEditValidator_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedEditValidator_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedEditValidator
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetBondDenom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetBondDenom()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetHistoricalEntries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetHistoricalEntries()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetMaxEntries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMaxEntries()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetMaxValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMaxValidators()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetMinCommissionRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMinCommissionRate()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_GetUnbondingTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetUnbondingTime()
	})
}

func Fuzz_Nosy_EventWrappedStakingUpdateParams_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

// skipping Fuzz_Nosy_EventWrappedStakingUpdateParams_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedStakingUpdateParams_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedStakingUpdateParams_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedStakingUpdateParams
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

func Fuzz_Nosy_EventWrappedUndelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAmount()
	})
}

func Fuzz_Nosy_EventWrappedUndelegate_GetDelegatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelegatorAddress()
	})
}

func Fuzz_Nosy_EventWrappedUndelegate_GetDenom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDenom()
	})
}

func Fuzz_Nosy_EventWrappedUndelegate_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_EventWrappedUndelegate_GetValidatorAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

// skipping Fuzz_Nosy_EventWrappedUndelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_EventWrappedUndelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_EventWrappedUndelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *EventWrappedUndelegate
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

func Fuzz_Nosy_MsgWrappedBeginRedelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedBeginRedelegate(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedBeginRedelegate(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedBeginRedelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedBeginRedelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgBeginRedelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedBeginRedelegate(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

// skipping Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedBeginRedelegateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedBeginRedelegateResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedCancelUnbondingDelegation(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedCancelUnbondingDelegation(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegation_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedCancelUnbondingDelegation(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

// skipping Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedCancelUnbondingDelegationResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedCancelUnbondingDelegationResponse
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

func Fuzz_Nosy_MsgWrappedDelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedDelegate(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedDelegate(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedDelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedDelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedDelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgDelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedDelegate(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedDelegateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

// skipping Fuzz_Nosy_MsgWrappedDelegateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedDelegateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedDelegateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedDelegateResponse
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

func Fuzz_Nosy_MsgWrappedEditValidator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedEditValidator(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedEditValidator(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedEditValidator_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedEditValidator_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedEditValidator_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgEditValidator
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedEditValidator(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

// skipping Fuzz_Nosy_MsgWrappedEditValidatorResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedEditValidatorResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedEditValidatorResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedStakingUpdateParams(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedStakingUpdateParams(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedStakingUpdateParams_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParams_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUpdateParams
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedStakingUpdateParams(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

// skipping Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedStakingUpdateParamsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedStakingUpdateParamsResponse
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

func Fuzz_Nosy_MsgWrappedUndelegate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedUndelegate(msg)
		_x1.Descriptor()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.Marshal()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		_x1 := NewMsgWrappedUndelegate(msg)
		_x1.ProtoMessage()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.Reset()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.Size()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.String()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var dAtA []byte
		fill_err = tp.Fill(&dAtA)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.XXX_DiscardUnknown()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
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
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_MsgWrappedUndelegate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedUndelegate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.XXX_Size()
	})
}

func Fuzz_Nosy_MsgWrappedUndelegate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *types.MsgUndelegate
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var b []byte
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		m := NewMsgWrappedUndelegate(msg)
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_MsgWrappedUndelegateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

// skipping Fuzz_Nosy_MsgWrappedUndelegateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_MsgWrappedUndelegateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

func Fuzz_Nosy_MsgWrappedUndelegateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *MsgWrappedUndelegateResponse
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

// skipping Fuzz_Nosy_Params_Equal__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Params_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Params_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Params_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultParams()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Params_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := DefaultParams()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Params_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Params_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := DefaultParams()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryCurrentEpochRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

// skipping Fuzz_Nosy_QueryCurrentEpochRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryCurrentEpochRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochRequest
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

func Fuzz_Nosy_QueryCurrentEpochResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_GetCurrentEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCurrentEpoch()
	})
}

func Fuzz_Nosy_QueryCurrentEpochResponse_GetEpochBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochBoundary()
	})
}

func Fuzz_Nosy_QueryCurrentEpochResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

// skipping Fuzz_Nosy_QueryCurrentEpochResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryCurrentEpochResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryCurrentEpochResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryCurrentEpochResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_GetDelAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelAddr()
	})
}

func Fuzz_Nosy_QueryDelegationLifecycleRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

// skipping Fuzz_Nosy_QueryDelegationLifecycleRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryDelegationLifecycleRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleRequest
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_GetDelLife__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetDelLife()
	})
}

func Fuzz_Nosy_QueryDelegationLifecycleResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

// skipping Fuzz_Nosy_QueryDelegationLifecycleResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryDelegationLifecycleResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryDelegationLifecycleResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryDelegationLifecycleResponse
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

func Fuzz_Nosy_QueryEpochInfoRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

// skipping Fuzz_Nosy_QueryEpochInfoRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochInfoRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoRequest
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

func Fuzz_Nosy_QueryEpochInfoResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_GetEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpoch()
	})
}

func Fuzz_Nosy_QueryEpochInfoResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

// skipping Fuzz_Nosy_QueryEpochInfoResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochInfoResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochInfoResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochInfoResponse
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

func Fuzz_Nosy_QueryEpochMsgsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

// skipping Fuzz_Nosy_QueryEpochMsgsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochMsgsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsRequest
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

func Fuzz_Nosy_QueryEpochMsgsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_GetMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgs()
	})
}

func Fuzz_Nosy_QueryEpochMsgsResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

// skipping Fuzz_Nosy_QueryEpochMsgsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochMsgsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochMsgsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochMsgsResponse
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

func Fuzz_Nosy_QueryEpochValSetRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_GetEpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

// skipping Fuzz_Nosy_QueryEpochValSetRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochValSetRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetRequest
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

func Fuzz_Nosy_QueryEpochValSetResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_GetTotalVotingPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTotalVotingPower()
	})
}

func Fuzz_Nosy_QueryEpochValSetResponse_GetValidators__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValidators()
	})
}

func Fuzz_Nosy_QueryEpochValSetResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

// skipping Fuzz_Nosy_QueryEpochValSetResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochValSetResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochValSetResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochValSetResponse
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

func Fuzz_Nosy_QueryEpochsInfoRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

// skipping Fuzz_Nosy_QueryEpochsInfoRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochsInfoRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoRequest
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

func Fuzz_Nosy_QueryEpochsInfoResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_GetEpochs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochs()
	})
}

func Fuzz_Nosy_QueryEpochsInfoResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

// skipping Fuzz_Nosy_QueryEpochsInfoResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryEpochsInfoResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryEpochsInfoResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryEpochsInfoResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_GetEndEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEndEpoch()
	})
}

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_GetEpochCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

// skipping Fuzz_Nosy_QueryLatestEpochMsgsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsRequest
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_GetLatestEpochMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetLatestEpochMsgs()
	})
}

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_GetPagination__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

// skipping Fuzz_Nosy_QueryLatestEpochMsgsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryLatestEpochMsgsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryLatestEpochMsgsResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_GetValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValAddr()
	})
}

func Fuzz_Nosy_QueryValidatorLifecycleRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

// skipping Fuzz_Nosy_QueryValidatorLifecycleRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryValidatorLifecycleRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleRequest
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_GetValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValAddr()
	})
}

func Fuzz_Nosy_QueryValidatorLifecycleResponse_GetValLife__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValLife()
	})
}

func Fuzz_Nosy_QueryValidatorLifecycleResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

// skipping Fuzz_Nosy_QueryValidatorLifecycleResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryValidatorLifecycleResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueryValidatorLifecycleResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryValidatorLifecycleResponse
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

func Fuzz_Nosy_QueuedMessage_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_GetMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsg()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgBeginRedelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgBeginRedelegate()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgCancelUnbondingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgCancelUnbondingDelegation()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgCreateValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgCreateValidator()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgDelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgDelegate()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgEditValidator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgEditValidator()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgId()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgUndelegate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgUndelegate()
	})
}

func Fuzz_Nosy_QueuedMessage_GetMsgUpdateParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgUpdateParams()
	})
}

func Fuzz_Nosy_QueuedMessage_GetTxId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTxId()
	})
}

func Fuzz_Nosy_QueuedMessage_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q *QueuedMessage
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}
		if q == nil {
			return
		}

		q.ToResponse()
	})
}

func Fuzz_Nosy_QueuedMessage_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_UnwrapToSdkMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var qm *QueuedMessage
		fill_err = tp.Fill(&qm)
		if fill_err != nil {
			return
		}
		if qm == nil {
			return
		}

		qm.UnwrapToSdkMsg()
	})
}

func Fuzz_Nosy_QueuedMessage_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

// skipping Fuzz_Nosy_QueuedMessage_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueuedMessage_XXX_OneofWrappers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.XXX_OneofWrappers()
	})
}

func Fuzz_Nosy_QueuedMessage_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessage_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage
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

func Fuzz_Nosy_QueuedMessageList_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetEpochNumber()
	})
}

func Fuzz_Nosy_QueuedMessageList_GetMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgs()
	})
}

func Fuzz_Nosy_QueuedMessageList_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

// skipping Fuzz_Nosy_QueuedMessageList_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueuedMessageList_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageList_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageList
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

func Fuzz_Nosy_QueuedMessageResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_GetMsg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsg()
	})
}

func Fuzz_Nosy_QueuedMessageResponse_GetMsgId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetMsgId()
	})
}

func Fuzz_Nosy_QueuedMessageResponse_GetTxId__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetTxId()
	})
}

func Fuzz_Nosy_QueuedMessageResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

// skipping Fuzz_Nosy_QueuedMessageResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueuedMessageResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessageResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessageResponse
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

func Fuzz_Nosy_QueuedMessage_MsgBeginRedelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgBeginRedelegate
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

func Fuzz_Nosy_QueuedMessage_MsgBeginRedelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgBeginRedelegate
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

func Fuzz_Nosy_QueuedMessage_MsgBeginRedelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgBeginRedelegate
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

func Fuzz_Nosy_QueuedMessage_MsgBeginRedelegate_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgBeginRedelegate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgCancelUnbondingDelegation_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCancelUnbondingDelegation
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

func Fuzz_Nosy_QueuedMessage_MsgCancelUnbondingDelegation_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCancelUnbondingDelegation
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

func Fuzz_Nosy_QueuedMessage_MsgCancelUnbondingDelegation_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCancelUnbondingDelegation
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

func Fuzz_Nosy_QueuedMessage_MsgCancelUnbondingDelegation_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgCancelUnbondingDelegation
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgCreateValidator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCreateValidator
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

func Fuzz_Nosy_QueuedMessage_MsgCreateValidator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCreateValidator
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

func Fuzz_Nosy_QueuedMessage_MsgCreateValidator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgCreateValidator
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

func Fuzz_Nosy_QueuedMessage_MsgCreateValidator_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgCreateValidator
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgDelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgDelegate
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

func Fuzz_Nosy_QueuedMessage_MsgDelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgDelegate
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

func Fuzz_Nosy_QueuedMessage_MsgDelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgDelegate
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

func Fuzz_Nosy_QueuedMessage_MsgDelegate_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgDelegate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgEditValidator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgEditValidator
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

func Fuzz_Nosy_QueuedMessage_MsgEditValidator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgEditValidator
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

func Fuzz_Nosy_QueuedMessage_MsgEditValidator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgEditValidator
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

func Fuzz_Nosy_QueuedMessage_MsgEditValidator_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgEditValidator
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgUndelegate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUndelegate
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

func Fuzz_Nosy_QueuedMessage_MsgUndelegate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUndelegate
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

func Fuzz_Nosy_QueuedMessage_MsgUndelegate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUndelegate
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

func Fuzz_Nosy_QueuedMessage_MsgUndelegate_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgUndelegate
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
	})
}

func Fuzz_Nosy_QueuedMessage_MsgUpdateParams_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUpdateParams
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

func Fuzz_Nosy_QueuedMessage_MsgUpdateParams_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUpdateParams
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

func Fuzz_Nosy_QueuedMessage_MsgUpdateParams_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueuedMessage_MsgUpdateParams
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

func Fuzz_Nosy_QueuedMessage_MsgUpdateParams_isQueuedMessage_Msg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueuedMessage_MsgUpdateParams
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.isQueuedMessage_Msg()
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

func Fuzz_Nosy_UnimplementedMsgServer_WrappedBeginRedelegate__(f *testing.F) {
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
		var req *MsgWrappedBeginRedelegate
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedBeginRedelegate(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_WrappedCancelUnbondingDelegation__(f *testing.F) {
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
		var req *MsgWrappedCancelUnbondingDelegation
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedCancelUnbondingDelegation(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_WrappedDelegate__(f *testing.F) {
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
		var req *MsgWrappedDelegate
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedDelegate(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_WrappedEditValidator__(f *testing.F) {
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
		var req *MsgWrappedEditValidator
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedEditValidator(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_WrappedStakingUpdateParams__(f *testing.F) {
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
		var req *MsgWrappedStakingUpdateParams
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedStakingUpdateParams(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedMsgServer_WrappedUndelegate__(f *testing.F) {
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
		var req *MsgWrappedUndelegate
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.WrappedUndelegate(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_CurrentEpoch__(f *testing.F) {
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
		var req *QueryCurrentEpochRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.CurrentEpoch(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_DelegationLifecycle__(f *testing.F) {
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
		var req *QueryDelegationLifecycleRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.DelegationLifecycle(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_EpochInfo__(f *testing.F) {
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
		var req *QueryEpochInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EpochInfo(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_EpochMsgs__(f *testing.F) {
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
		var req *QueryEpochMsgsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EpochMsgs(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_EpochValSet__(f *testing.F) {
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
		var req *QueryEpochValSetRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EpochValSet(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_EpochsInfo__(f *testing.F) {
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
		var req *QueryEpochsInfoRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.EpochsInfo(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_LatestEpochMsgs__(f *testing.F) {
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
		var req *QueryLatestEpochMsgsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.LatestEpochMsgs(ctx, req)
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

func Fuzz_Nosy_UnimplementedQueryServer_ValidatorLifecycle__(f *testing.F) {
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
		var req *QueryValidatorLifecycleRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.ValidatorLifecycle(ctx, req)
	})
}

func Fuzz_Nosy_ValStateUpdate_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_GetState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_ToResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValStateUpdate
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.ToResponse()
	})
}

func Fuzz_Nosy_ValStateUpdate_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

// skipping Fuzz_Nosy_ValStateUpdate_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ValStateUpdate_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdate_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdate
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

func Fuzz_Nosy_ValStateUpdateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_GetBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_GetBlockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_GetStateDesc__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetStateDesc()
	})
}

func Fuzz_Nosy_ValStateUpdateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

// skipping Fuzz_Nosy_ValStateUpdateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ValStateUpdateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_ValStateUpdateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValStateUpdateResponse
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

func Fuzz_Nosy_Validator_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Validator
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

func Fuzz_Nosy_Validator_GetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetAddr()
	})
}

func Fuzz_Nosy_Validator_GetPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetPower()
	})
}

func Fuzz_Nosy_Validator_GetValAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.GetValAddress()
	})
}

func Fuzz_Nosy_Validator_GetValAddressStr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *Validator
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.GetValAddressStr()
	})
}

func Fuzz_Nosy_Validator_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Validator
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

func Fuzz_Nosy_Validator_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

// skipping Fuzz_Nosy_Validator_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Validator_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_Validator_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Validator
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

func Fuzz_Nosy_ValidatorLifecycle_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_GetValAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValAddr()
	})
}

func Fuzz_Nosy_ValidatorLifecycle_GetValLife__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetValLife()
	})
}

func Fuzz_Nosy_ValidatorLifecycle_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

// skipping Fuzz_Nosy_ValidatorLifecycle_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_ValidatorLifecycle_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

func Fuzz_Nosy_ValidatorLifecycle_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *ValidatorLifecycle
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

// skipping Fuzz_Nosy_msgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedBeginRedelegate__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedCancelUnbondingDelegation__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedDelegate__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedEditValidator__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedStakingUpdateParams__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_msgClient_WrappedUndelegate__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_CurrentEpoch__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_DelegationLifecycle__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_EpochInfo__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_EpochMsgs__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_EpochValSet__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_EpochsInfo__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_LatestEpochMsgs__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_Params__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_ValidatorLifecycle__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_AccountKeeper_GetAccount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.AccountKeeper

// skipping Fuzz_Nosy_BankKeeper_GetBalance__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.BankKeeper

// skipping Fuzz_Nosy_BankKeeper_SpendableCoins__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.BankKeeper

func Fuzz_Nosy_BondState_EnumDescriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 BondState
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}

		_x1.EnumDescriptor()
	})
}

func Fuzz_Nosy_BondState_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var x BondState
		fill_err = tp.Fill(&x)
		if fill_err != nil {
			return
		}

		x.String()
	})
}

func Fuzz_Nosy_Epoch_GetLastBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.GetLastBlockHeight()
	})
}

func Fuzz_Nosy_Epoch_GetSealerBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.GetSealerBlockHeight()
	})
}

func Fuzz_Nosy_Epoch_GetSecondBlockHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.GetSecondBlockHeight()
	})
}

func Fuzz_Nosy_Epoch_IsFirstBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsFirstBlock(ctx)
	})
}

func Fuzz_Nosy_Epoch_IsFirstBlockOfNextEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsFirstBlockOfNextEpoch(ctx)
	})
}

func Fuzz_Nosy_Epoch_IsLastBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsLastBlock(ctx)
	})
}

func Fuzz_Nosy_Epoch_IsLastBlockByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var height int64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsLastBlockByHeight(height)
	})
}

func Fuzz_Nosy_Epoch_IsSecondBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsSecondBlock(ctx)
	})
}

func Fuzz_Nosy_Epoch_IsVoteExtensionProposal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.IsVoteExtensionProposal(ctx)
	})
}

func Fuzz_Nosy_Epoch_ValidateBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.ValidateBasic()
	})
}

func Fuzz_Nosy_Epoch_WithinBoundary__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		var epochInterval uint64
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		var firstBlockHeight uint64
		fill_err = tp.Fill(&firstBlockHeight)
		if fill_err != nil {
			return
		}
		var lastBlockTime *time.Time
		fill_err = tp.Fill(&lastBlockTime)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if lastBlockTime == nil {
			return
		}

		e := NewEpoch(epochNumber, epochInterval, firstBlockHeight, lastBlockTime)
		e.WithinBoundary(height)
	})
}

// skipping Fuzz_Nosy_EpochingHooks_AfterEpochBegins__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.EpochingHooks

// skipping Fuzz_Nosy_EpochingHooks_AfterEpochEnds__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.EpochingHooks

// skipping Fuzz_Nosy_EpochingHooks_BeforeSlashThreshold__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.EpochingHooks

// skipping Fuzz_Nosy_MsgClient_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedBeginRedelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedCancelUnbondingDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedDelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedEditValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedStakingUpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgClient_WrappedUndelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgClient

// skipping Fuzz_Nosy_MsgServer_UpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedBeginRedelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedCancelUnbondingDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedDelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedEditValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedStakingUpdateParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MsgServer_WrappedUndelegate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MsgServer

// skipping Fuzz_Nosy_MultiEpochingHooks_AfterEpochBegins__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MultiEpochingHooks

// skipping Fuzz_Nosy_MultiEpochingHooks_AfterEpochEnds__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MultiEpochingHooks

// skipping Fuzz_Nosy_MultiEpochingHooks_BeforeSlashThreshold__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.MultiEpochingHooks

// skipping Fuzz_Nosy_QueryClient_CurrentEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_DelegationLifecycle__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_EpochInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_EpochMsgs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_EpochValSet__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_EpochsInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_LatestEpochMsgs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_ValidatorLifecycle__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_CurrentEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_DelegationLifecycle__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_EpochInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_EpochMsgs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_EpochValSet__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_EpochsInfo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_LatestEpochMsgs__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_ValidatorLifecycle__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.QueryServer

// skipping Fuzz_Nosy_QueuedMessage_UnpackInterfaces__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/codec/types.AnyUnpacker

// skipping Fuzz_Nosy_StakingHooks_AfterValidatorBeginUnbonding__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingHooks

// skipping Fuzz_Nosy_StakingHooks_AfterValidatorBonded__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingHooks

// skipping Fuzz_Nosy_StakingHooks_AfterValidatorCreated__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingHooks

// skipping Fuzz_Nosy_StakingHooks_AfterValidatorRemoved__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingHooks

// skipping Fuzz_Nosy_StakingHooks_BeforeValidatorSlashed__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingHooks

// skipping Fuzz_Nosy_StakingKeeper_ApplyAndReturnValidatorSetUpdates__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_BondDenom__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_CompleteRedelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_CompleteUnbonding__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_DequeueAllMatureRedelegationQueue__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_DequeueAllMatureUBDQueue__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_GetParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_GetPubKeyByConsAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_GetValidator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_GetValidatorByConsAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_GetValidatorDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_HasMaxRedelegationEntries__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_HasMaxUnbondingDelegationEntries__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_HasReceivingRedelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_IterateLastValidatorPowers__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_UnbondAllMatureValidators__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_ValidateUnbondAmount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

// skipping Fuzz_Nosy_StakingKeeper_ValidatorQueueIterator__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.StakingKeeper

func Fuzz_Nosy_ValidatorSet_FindSubset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vsBytes []byte
		fill_err = tp.Fill(&vsBytes)
		if fill_err != nil {
			return
		}
		var bm bitmap.Bitmap
		fill_err = tp.Fill(&bm)
		if fill_err != nil {
			return
		}

		vs, err := NewValidatorSetFromBytes(vsBytes)
		if err != nil {
			return
		}
		vs.FindSubset(bm)
	})
}

func Fuzz_Nosy_ValidatorSet_FindValidatorWithIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vsBytes []byte
		fill_err = tp.Fill(&vsBytes)
		if fill_err != nil {
			return
		}
		var valAddr sdk.ValAddress
		fill_err = tp.Fill(&valAddr)
		if fill_err != nil {
			return
		}

		vs, err := NewValidatorSetFromBytes(vsBytes)
		if err != nil {
			return
		}
		vs.FindValidatorWithIndex(valAddr)
	})
}

func Fuzz_Nosy_ValidatorSet_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vsBytes []byte) {
		vs, err := NewValidatorSetFromBytes(vsBytes)
		if err != nil {
			return
		}
		vs.Marshal()
	})
}

func Fuzz_Nosy_ValidatorSet_MustMarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, vsBytes []byte) {
		vs, err := NewValidatorSetFromBytes(vsBytes)
		if err != nil {
			return
		}
		vs.MustMarshal()
	})
}

func Fuzz_Nosy_ValidatorSet_binarySearch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vsBytes []byte
		fill_err = tp.Fill(&vsBytes)
		if fill_err != nil {
			return
		}
		var targetAddr sdk.ValAddress
		fill_err = tp.Fill(&targetAddr)
		if fill_err != nil {
			return
		}

		vs, err := NewValidatorSetFromBytes(vsBytes)
		if err != nil {
			return
		}
		vs.binarySearch(targetAddr)
	})
}

// skipping Fuzz_Nosy_isQueuedMessage_Msg_MarshalTo__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.isQueuedMessage_Msg

// skipping Fuzz_Nosy_isQueuedMessage_Msg_Size__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.isQueuedMessage_Msg

// skipping Fuzz_Nosy_isQueuedMessage_Msg_isQueuedMessage_Msg__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/epoching/types.isQueuedMessage_Msg

func Fuzz_Nosy_KeyPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		KeyPrefix(p)
	})
}

func Fuzz_Nosy_NewQueuedMessagesResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msgs []*QueuedMessage
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}

		NewQueuedMessagesResponse(msgs)
	})
}

func Fuzz_Nosy_NewValsetUpdateResponses__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs []*ValStateUpdate
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}

		NewValsetUpdateResponses(vs)
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

// skipping Fuzz_Nosy__Msg_UpdateParams_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedBeginRedelegate_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedCancelUnbondingDelegation_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedDelegate_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedEditValidator_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedStakingUpdateParams_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Msg_WrappedUndelegate_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_CurrentEpoch_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_DelegationLifecycle_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_EpochInfo_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_EpochMsgs_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_EpochValSet_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_EpochsInfo_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_LatestEpochMsgs_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_Params_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_ValidatorLifecycle_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintEpoching__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintEpoching(dAtA, offset, v)
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

// skipping Fuzz_Nosy_local_request_Query_CurrentEpoch_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_DelegationLifecycle_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_EpochInfo_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_EpochMsgs_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_EpochValSet_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_EpochsInfo_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_LatestEpochMsgs_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_ValidatorLifecycle_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_CurrentEpoch_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_DelegationLifecycle_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_EpochInfo_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_EpochMsgs_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_EpochValSet_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_EpochsInfo_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_LatestEpochMsgs_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_Params_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_ValidatorLifecycle_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipEpoching__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipEpoching(dAtA)
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

func Fuzz_Nosy_sovEpoching__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovEpoching(x)
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

func Fuzz_Nosy_sozEpoching__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozEpoching(x)
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
