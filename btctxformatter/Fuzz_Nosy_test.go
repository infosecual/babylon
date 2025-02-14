package btctxformatter

import (
	"testing"

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

func Fuzz_Nosy_formatHeader_validateHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d1 []byte
		fill_err = tp.Fill(&d1)
		if fill_err != nil {
			return
		}
		var expectedTag BabylonTag
		fill_err = tp.Fill(&expectedTag)
		if fill_err != nil {
			return
		}
		var _x3 FormatVersion
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		var expectedPart uint8
		fill_err = tp.Fill(&expectedPart)
		if fill_err != nil {
			return
		}

		header := parseHeader(d1)
		header.validateHeader(expectedTag, _x3, expectedPart)
	})
}

func Fuzz_Nosy_ConnectParts__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
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

		ConnectParts(version, f2, s)
	})
}

func Fuzz_Nosy_EncodeCheckpointData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var rawBTCCheckpoint *RawBtcCheckpoint
		fill_err = tp.Fill(&rawBTCCheckpoint)
		if fill_err != nil {
			return
		}
		if rawBTCCheckpoint == nil {
			return
		}

		EncodeCheckpointData(tag, version, rawBTCCheckpoint)
	})
}

func Fuzz_Nosy_GetCheckpointData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var partIndex uint8
		fill_err = tp.Fill(&partIndex)
		if fill_err != nil {
			return
		}
		var d4 []byte
		fill_err = tp.Fill(&d4)
		if fill_err != nil {
			return
		}

		GetCheckpointData(tag, version, partIndex, d4)
	})
}

func Fuzz_Nosy_MustEncodeCheckpointData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var rawBTCCheckpoint *RawBtcCheckpoint
		fill_err = tp.Fill(&rawBTCCheckpoint)
		if fill_err != nil {
			return
		}
		if rawBTCCheckpoint == nil {
			return
		}

		MustEncodeCheckpointData(tag, version, rawBTCCheckpoint)
	})
}

func Fuzz_Nosy_U64ToBEBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, u uint64) {
		U64ToBEBytes(u)
	})
}

func Fuzz_Nosy_encodeFirstOpRetrun__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		var appHash []byte
		fill_err = tp.Fill(&appHash)
		if fill_err != nil {
			return
		}
		var bitMap []byte
		fill_err = tp.Fill(&bitMap)
		if fill_err != nil {
			return
		}
		var submitterAddress []byte
		fill_err = tp.Fill(&submitterAddress)
		if fill_err != nil {
			return
		}

		encodeFirstOpRetrun(tag, version, epoch, appHash, bitMap, submitterAddress)
	})
}

func Fuzz_Nosy_encodeHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var halfNumber uint8
		fill_err = tp.Fill(&halfNumber)
		if fill_err != nil {
			return
		}

		encodeHeader(tag, version, halfNumber)
	})
}

func Fuzz_Nosy_encodeSecondOpReturn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var firstOpReturnBytes []byte
		fill_err = tp.Fill(&firstOpReturnBytes)
		if fill_err != nil {
			return
		}
		var blsSig []byte
		fill_err = tp.Fill(&blsSig)
		if fill_err != nil {
			return
		}

		encodeSecondOpReturn(tag, version, firstOpReturnBytes, blsSig)
	})
}

func Fuzz_Nosy_getCheckSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, firstTxBytes []byte) {
		getCheckSum(firstTxBytes)
	})
}

func Fuzz_Nosy_getVerHalf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var version FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var halfNumber uint8
		fill_err = tp.Fill(&halfNumber)
		if fill_err != nil {
			return
		}

		getVerHalf(version, halfNumber)
	})
}
