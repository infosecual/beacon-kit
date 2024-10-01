package hex

import (
	"math/big"
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

func Fuzz_Nosy_String_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.UnmarshalText(text)
	})
}

func Fuzz_Nosy_String_AddQuotes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.AddQuotes()
	})
}

func Fuzz_Nosy_String_Has0xPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.Has0xPrefix()
	})
}

func Fuzz_Nosy_String_IsEmpty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.IsEmpty()
	})
}

func Fuzz_Nosy_String_MustToBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.MustToBigInt()
	})
}

func Fuzz_Nosy_String_MustToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.MustToBytes()
	})
}

func Fuzz_Nosy_String_MustToUInt64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.MustToUInt64()
	})
}

func Fuzz_Nosy_String_ToBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.ToBigInt()
	})
}

func Fuzz_Nosy_String_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.ToBytes()
	})
}

func Fuzz_Nosy_String_ToUint64__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.ToUint64()
	})
}

func Fuzz_Nosy_String_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bigint *big.Int
		fill_err = tp.Fill(&bigint)
		if fill_err != nil {
			return
		}
		if bigint == nil {
			return
		}

		s := FromBigInt(bigint)
		s.Unwrap()
	})
}

// skipping Fuzz_Nosy_EncodeBytes__ because parameters include func, chan, or unsupported interface: B

func Fuzz_Nosy_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b uint64) {
		MarshalText(b)
	})
}

func Fuzz_Nosy_MustToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		MustToBytes(input)
	})
}

func Fuzz_Nosy_ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		ToBytes(input)
	})
}

func Fuzz_Nosy_UnmarshalByteText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		UnmarshalByteText(input)
	})
}

func Fuzz_Nosy_UnmarshalUint64Text__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		UnmarshalUint64Text(input)
	})
}

func Fuzz_Nosy_decodeNibble__(f *testing.F) {
	f.Fuzz(func(t *testing.T, in byte) {
		decodeNibble(in)
	})
}

// skipping Fuzz_Nosy_ensure0xPrefix__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_ensureStringInvariants__(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		ensureStringInvariants(s)
	})
}

// skipping Fuzz_Nosy_formatAndValidateNumber__ because parameters include func, chan, or unsupported interface: T

func Fuzz_Nosy_formatAndValidateText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		formatAndValidateText(input)
	})
}

// skipping Fuzz_Nosy_has0xPrefix__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_isQuotedString__ because parameters include func, chan, or unsupported interface: T
