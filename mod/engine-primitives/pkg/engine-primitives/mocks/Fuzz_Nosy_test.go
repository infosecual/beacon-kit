package mocks

import (
	"testing"

	bytes "github.com/berachain/beacon-kit/mod/primitives/pkg/bytes"
	common "github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	eip4844 "github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
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

func Fuzz_Nosy_BlobsBundle_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobsBundle
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

func Fuzz_Nosy_BlobsBundle_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobsBundle
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBlobs()
	})
}

func Fuzz_Nosy_BlobsBundle_GetCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobsBundle
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetCommitments()
	})
}

func Fuzz_Nosy_BlobsBundle_GetProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BlobsBundle
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetProofs()
	})
}

func Fuzz_Nosy_BlobsBundle_Expecter_GetBlobs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlobsBundle_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.GetBlobs()
	})
}

func Fuzz_Nosy_BlobsBundle_Expecter_GetCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlobsBundle_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.GetCommitments()
	})
}

func Fuzz_Nosy_BlobsBundle_Expecter_GetProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BlobsBundle_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.GetProofs()
	})
}

func Fuzz_Nosy_BlobsBundle_GetBlobs_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlobsBundle_GetBlobs_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []*eip4844.Blob
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

// skipping Fuzz_Nosy_BlobsBundle_GetBlobs_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlobsBundle_GetBlobs_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []*github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844.Blob

func Fuzz_Nosy_BlobsBundle_GetCommitments_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlobsBundle_GetCommitments_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []eip4844.KZGCommitment
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

// skipping Fuzz_Nosy_BlobsBundle_GetCommitments_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlobsBundle_GetCommitments_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844.KZGCommitment

func Fuzz_Nosy_BlobsBundle_GetProofs_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BlobsBundle_GetProofs_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 []bytes.B48
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

// skipping Fuzz_Nosy_BlobsBundle_GetProofs_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlobsBundle_GetProofs_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() []github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT interface{}]_GetBlobsBundle__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT interface{}]_GetExecutionPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT interface{}]_GetValue__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT interface{}]_ShouldOverrideBuilder__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_Expecter[ExecutionPayloadT interface{}]_GetBlobsBundle__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_Expecter[ExecutionPayloadT interface{}]_GetExecutionPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_Expecter[ExecutionPayloadT interface{}]_GetValue__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_Expecter[ExecutionPayloadT interface{}]_ShouldOverrideBuilder__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetBlobsBundle_Call[ExecutionPayloadT interface{}]_Return__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BlobsBundle

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetBlobsBundle_Call[ExecutionPayloadT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetBlobsBundle_Call[ExecutionPayloadT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BlobsBundle

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetExecutionPayload_Call[ExecutionPayloadT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ExecutionPayloadT

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetExecutionPayload_Call[ExecutionPayloadT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetExecutionPayload_Call[ExecutionPayloadT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() ExecutionPayloadT

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetValue_Call[ExecutionPayloadT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetValue_Call[ExecutionPayloadT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_GetValue_Call[ExecutionPayloadT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() *github.com/holiman/uint256.Int

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call[ExecutionPayloadT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call[ExecutionPayloadT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv_ShouldOverrideBuilder_Call[ExecutionPayloadT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() bool

func Fuzz_Nosy_PayloadAttributer_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PayloadAttributer
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

func Fuzz_Nosy_PayloadAttributer_GetSuggestedFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PayloadAttributer
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetSuggestedFeeRecipient()
	})
}

func Fuzz_Nosy_PayloadAttributer_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PayloadAttributer
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IsNil()
	})
}

func Fuzz_Nosy_PayloadAttributer_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *PayloadAttributer
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Version()
	})
}

func Fuzz_Nosy_PayloadAttributer_Expecter_GetSuggestedFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *PayloadAttributer_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.GetSuggestedFeeRecipient()
	})
}

func Fuzz_Nosy_PayloadAttributer_Expecter_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *PayloadAttributer_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.IsNil()
	})
}

func Fuzz_Nosy_PayloadAttributer_Expecter_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *PayloadAttributer_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Version()
	})
}

func Fuzz_Nosy_PayloadAttributer_GetSuggestedFeeRecipient_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *PayloadAttributer_GetSuggestedFeeRecipient_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 common.ExecutionAddress
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

// skipping Fuzz_Nosy_PayloadAttributer_GetSuggestedFeeRecipient_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_PayloadAttributer_GetSuggestedFeeRecipient_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/common.ExecutionAddress

func Fuzz_Nosy_PayloadAttributer_IsNil_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *PayloadAttributer_IsNil_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 bool
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

// skipping Fuzz_Nosy_PayloadAttributer_IsNil_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_PayloadAttributer_IsNil_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() bool

func Fuzz_Nosy_PayloadAttributer_Version_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *PayloadAttributer_Version_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 uint32
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

// skipping Fuzz_Nosy_PayloadAttributer_Version_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_PayloadAttributer_Version_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() uint32
