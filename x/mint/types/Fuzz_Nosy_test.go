package types

import (
	context "context"
	"testing"
	time "time"

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

func Fuzz_Nosy_GenesisState_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesisState()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesisState()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultGenesisState()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_GenesisState_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := DefaultGenesisState()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_GenesisState_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_GenesisState_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := DefaultGenesisState()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_GenesisTime_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GenesisTime
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

func Fuzz_Nosy_GenesisTime_GetGenesisTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetGenesisTime()
	})
}

func Fuzz_Nosy_GenesisTime_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_GenesisTime_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_GenesisTime_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *GenesisTime
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

func Fuzz_Nosy_GenesisTime_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_GenesisTime_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

// skipping Fuzz_Nosy_GenesisTime_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_GenesisTime_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_GenesisTime_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *GenesisTime
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

func Fuzz_Nosy_Minter_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultMinter()
		m.MarshalTo(dAtA)
	})
}

func Fuzz_Nosy_Minter_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultMinter()
		m.MarshalToSizedBuffer(dAtA)
	})
}

func Fuzz_Nosy_Minter_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		m := DefaultMinter()
		m.Unmarshal(dAtA)
	})
}

func Fuzz_Nosy_Minter_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte, deterministic bool) {
		m := DefaultMinter()
		m.XXX_Marshal(b, deterministic)
	})
}

// skipping Fuzz_Nosy_Minter_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_Minter_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		m := DefaultMinter()
		m.XXX_Unmarshal(b)
	})
}

func Fuzz_Nosy_QueryAnnualProvisionsRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

// skipping Fuzz_Nosy_QueryAnnualProvisionsRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryAnnualProvisionsRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsRequest
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

// skipping Fuzz_Nosy_QueryAnnualProvisionsResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryAnnualProvisionsResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryAnnualProvisionsResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryAnnualProvisionsResponse
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

func Fuzz_Nosy_QueryGenesisTimeRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

// skipping Fuzz_Nosy_QueryGenesisTimeRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryGenesisTimeRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeRequest
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

func Fuzz_Nosy_QueryGenesisTimeResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_GetGenesisTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetGenesisTime()
	})
}

func Fuzz_Nosy_QueryGenesisTimeResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryGenesisTimeResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

// skipping Fuzz_Nosy_QueryGenesisTimeResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryGenesisTimeResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryGenesisTimeResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryGenesisTimeResponse
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

func Fuzz_Nosy_QueryInflationRateRequest_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateRequest_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateRequest_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateRequest_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

// skipping Fuzz_Nosy_QueryInflationRateRequest_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryInflationRateRequest_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateRequest_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateRequest
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

func Fuzz_Nosy_QueryInflationRateResponse_Descriptor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_MarshalTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateResponse_MarshalToSizedBuffer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateResponse_ProtoMessage__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_Reset__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
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

func Fuzz_Nosy_QueryInflationRateResponse_XXX_DiscardUnknown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_XXX_Marshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

// skipping Fuzz_Nosy_QueryInflationRateResponse_XXX_Merge__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/proto.Message

func Fuzz_Nosy_QueryInflationRateResponse_XXX_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_QueryInflationRateResponse_XXX_Unmarshal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *QueryInflationRateResponse
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

func Fuzz_Nosy_UnimplementedQueryServer_AnnualProvisions__(f *testing.F) {
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
		var req *QueryAnnualProvisionsRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.AnnualProvisions(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_GenesisTime__(f *testing.F) {
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
		var req *QueryGenesisTimeRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.GenesisTime(ctx, req)
	})
}

func Fuzz_Nosy_UnimplementedQueryServer_InflationRate__(f *testing.F) {
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
		var req *QueryInflationRateRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if _x1 == nil || req == nil {
			return
		}

		_x1.InflationRate(ctx, req)
	})
}

// skipping Fuzz_Nosy_queryClient_AnnualProvisions__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_GenesisTime__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_queryClient_InflationRate__ because parameters include func, chan, or unsupported interface: []google.golang.org/grpc.CallOption

// skipping Fuzz_Nosy_AccountKeeper_GetModuleAccount__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.AccountKeeper

// skipping Fuzz_Nosy_AccountKeeper_GetModuleAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.AccountKeeper

// skipping Fuzz_Nosy_BankKeeper_MintCoins__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.BankKeeper

// skipping Fuzz_Nosy_BankKeeper_SendCoinsFromModuleToModule__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.BankKeeper

func Fuzz_Nosy_Minter_CalculateBlockProvision__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var current time.Time
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}
		var previous time.Time
		fill_err = tp.Fill(&previous)
		if fill_err != nil {
			return
		}

		m := DefaultMinter()
		m.CalculateBlockProvision(current, previous)
	})
}

func Fuzz_Nosy_Minter_CalculateInflationRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctx sdk.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var genesis time.Time
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}

		m := DefaultMinter()
		m.CalculateInflationRate(ctx, genesis)
	})
}

// skipping Fuzz_Nosy_QueryClient_AnnualProvisions__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_GenesisTime__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryClient

// skipping Fuzz_Nosy_QueryClient_InflationRate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryClient

// skipping Fuzz_Nosy_QueryServer_AnnualProvisions__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_GenesisTime__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryServer

// skipping Fuzz_Nosy_QueryServer_InflationRate__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.QueryServer

// skipping Fuzz_Nosy_StakingKeeper_StakingTokenSupply__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/babylon/x/mint/types.StakingKeeper

// skipping Fuzz_Nosy_RegisterQueryServer__ because parameters include func, chan, or unsupported interface: github.com/cosmos/gogoproto/grpc.Server

// skipping Fuzz_Nosy__Query_AnnualProvisions_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_GenesisTime_Handler__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy__Query_InflationRate_Handler__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_encodeVarintGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintGenesis(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintMint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintMint(dAtA, offset, v)
	})
}

func Fuzz_Nosy_encodeVarintQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte, offset int, v uint64) {
		encodeVarintQuery(dAtA, offset, v)
	})
}

// skipping Fuzz_Nosy_local_request_Query_AnnualProvisions_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_GenesisTime_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_local_request_Query_InflationRate_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_AnnualProvisions_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_GenesisTime_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

// skipping Fuzz_Nosy_request_Query_InflationRate_0__ because parameters include func, chan, or unsupported interface: github.com/grpc-ecosystem/grpc-gateway/runtime.Marshaler

func Fuzz_Nosy_skipGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipGenesis(dAtA)
	})
}

func Fuzz_Nosy_skipMint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipMint(dAtA)
	})
}

func Fuzz_Nosy_skipQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dAtA []byte) {
		skipQuery(dAtA)
	})
}

func Fuzz_Nosy_sovGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovGenesis(x)
	})
}

func Fuzz_Nosy_sovMint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovMint(x)
	})
}

func Fuzz_Nosy_sovQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sovQuery(x)
	})
}

func Fuzz_Nosy_sozGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozGenesis(x)
	})
}

func Fuzz_Nosy_sozMint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozMint(x)
	})
}

func Fuzz_Nosy_sozQuery__(f *testing.F) {
	f.Fuzz(func(t *testing.T, x uint64) {
		sozQuery(x)
	})
}

func Fuzz_Nosy_yearsSinceGenesis__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var genesis time.Time
		fill_err = tp.Fill(&genesis)
		if fill_err != nil {
			return
		}
		var current time.Time
		fill_err = tp.Fill(&current)
		if fill_err != nil {
			return
		}

		yearsSinceGenesis(genesis, current)
	})
}
