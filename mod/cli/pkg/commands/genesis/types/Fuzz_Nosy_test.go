package types

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

func Fuzz_Nosy_ValidatorsMarshaling_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.GetTree()
	})
}

func Fuzz_Nosy_ValidatorsMarshaling_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_ValidatorsMarshaling_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_ValidatorsMarshaling_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.MarshalSSZ()
	})
}

func Fuzz_Nosy_ValidatorsMarshaling_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_ValidatorsMarshaling_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.SizeSSZ()
	})
}

func Fuzz_Nosy_ValidatorsMarshaling_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ValidatorsMarshaling
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.UnmarshalSSZ(buf)
	})
}
