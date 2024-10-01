package mocks

import (
	"testing"

	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
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

func Fuzz_Nosy_BLSSigner_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BLSSigner
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.EXPECT()
	})
}

func Fuzz_Nosy_BLSSigner_PublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BLSSigner
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.PublicKey()
	})
}

func Fuzz_Nosy_BLSSigner_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BLSSigner
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 []byte
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Sign(_a0)
	})
}

func Fuzz_Nosy_BLSSigner_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BLSSigner
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var pubKey bytes.B48
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var signature bytes.B96
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.VerifySignature(pubKey, msg, signature)
	})
}

func Fuzz_Nosy_BLSSigner_Expecter_PublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BLSSigner_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.PublicKey()
	})
}

// skipping Fuzz_Nosy_BLSSigner_Expecter_Sign__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BLSSigner_Expecter_VerifySignature__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_BLSSigner_PublicKey_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BLSSigner_PublicKey_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 bytes.B48
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_BLSSigner_PublicKey_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BLSSigner_PublicKey_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48

// skipping Fuzz_Nosy_BLSSigner_Sign_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BLSSigner_Sign_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 []byte)

// skipping Fuzz_Nosy_BLSSigner_Sign_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) (github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B96, error)

// skipping Fuzz_Nosy_BLSSigner_VerifySignature_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BLSSigner_VerifySignature_Call_Run__ because parameters include func, chan, or unsupported interface: func(pubKey github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48, msg []byte, signature github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B96)

// skipping Fuzz_Nosy_BLSSigner_VerifySignature_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48, []byte, github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B96) error
