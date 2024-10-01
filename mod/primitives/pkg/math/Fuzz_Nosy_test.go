package math

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

func Fuzz_Nosy_U256Hex_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *U256Hex
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.MarshalJSON()
	})
}

func Fuzz_Nosy_U256Hex_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *U256Hex
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.UnmarshalJSON(d2)
	})
}

func Fuzz_Nosy_U64_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_U64_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u *U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if u == nil {
			return
		}

		u.UnmarshalText(input)
	})
}

func Fuzz_Nosy_U64_Base10__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.Base10()
	})
}

func Fuzz_Nosy_U64_ILog2Ceil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.ILog2Ceil()
	})
}

func Fuzz_Nosy_U64_ILog2Floor__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.ILog2Floor()
	})
}

func Fuzz_Nosy_U64_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.MarshalText()
	})
}

func Fuzz_Nosy_U64_NextPowerOfTwo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.NextPowerOfTwo()
	})
}

func Fuzz_Nosy_U64_PrevPowerOfTwo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.PrevPowerOfTwo()
	})
}

func Fuzz_Nosy_U64_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.Unwrap()
	})
}

func Fuzz_Nosy_U64_UnwrapPtr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var u U64
		fill_err = tp.Fill(&u)
		if fill_err != nil {
			return
		}

		u.UnwrapPtr()
	})
}

func Fuzz_Nosy_GweiFromWei__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i *big.Int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		if i == nil {
			return
		}

		GweiFromWei(i)
	})
}

func Fuzz_Nosy_NewU256__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		NewU256(v)
	})
}

func Fuzz_Nosy_NewU256FromBigInt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *big.Int
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		NewU256FromBigInt(b)
	})
}
