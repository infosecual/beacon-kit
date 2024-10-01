package bytes

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

func Fuzz_Nosy_B20_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes20(input)
		h.String()
	})
}

func Fuzz_Nosy_B20_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes20(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B20_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes20(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B256_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes256(input)
		h.String()
	})
}

func Fuzz_Nosy_B256_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes256(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B256_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes256(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B32_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes32(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B32_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes32(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B4_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes4(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B4_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes4(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B48_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes48(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B48_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes48(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B8_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes8(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B8_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes8(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_B96_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes96(input)
		h.String()
	})
}

func Fuzz_Nosy_B96_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 []byte, i2 []byte) {
		h := ToBytes96(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_B96_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte, text []byte) {
		h := ToBytes96(input)
		h.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Bytes_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Bytes_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalText(input)
	})
}

func Fuzz_Nosy_B20_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes20(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B20_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes20(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B20_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes20(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B256_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes256(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B256_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes256(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B256_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes256(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B256_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes256(input)
		h.SizeSSZ()
	})
}

func Fuzz_Nosy_B32_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes32(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B32_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes32(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B32_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes32(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B32_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes32(input)
		h.String()
	})
}

func Fuzz_Nosy_B4_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes4(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B4_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes4(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B4_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes4(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B4_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes4(input)
		h.String()
	})
}

func Fuzz_Nosy_B48_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes48(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B48_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes48(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B48_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes48(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B48_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes48(input)
		h.String()
	})
}

func Fuzz_Nosy_B8_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes8(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B8_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes8(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B8_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes8(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_B8_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes8(input)
		h.String()
	})
}

func Fuzz_Nosy_B96_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes96(input)
		h.HashTreeRoot()
	})
}

func Fuzz_Nosy_B96_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes96(input)
		h.MarshalSSZ()
	})
}

func Fuzz_Nosy_B96_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h := ToBytes96(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_Bytes_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_Bytes_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Bytes
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.String()
	})
}

func Fuzz_Nosy_CopyAndReverseEndianess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		CopyAndReverseEndianess(input)
	})
}

func Fuzz_Nosy_ExtendToSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, slice []byte, length int) {
		ExtendToSize(slice, length)
	})
}
