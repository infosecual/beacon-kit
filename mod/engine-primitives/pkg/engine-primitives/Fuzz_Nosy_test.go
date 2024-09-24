package engineprimitives

import (
	"bytes"
	"io"
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/karalabe/ssz"
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

// skipping Fuzz_Nosy_BlobsBundleV1[C, P ~[48]byte, B ~[131072]byte]_GetBlobs__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlobsBundleV1[C, P ~[48]byte, B ~[131072]byte]_GetCommitments__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlobsBundleV1[C, P ~[48]byte, B ~[131072]byte]_GetProofs__ as it appears to be an interface:

func Fuzz_Nosy_ClientVersionV1_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ClientVersionV1
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.String()
	})
}

// skipping Fuzz_Nosy_ExecutionPayloadEnvelope[ExecutionPayloadT constraints.JSONMarshallable, BlobsBundleT BlobsBundle]_GetBlobsBundle__ as it appears to be an interface:

// skipping Fuzz_Nosy_ExecutionPayloadEnvelope[ExecutionPayloadT constraints.JSONMarshallable, BlobsBundleT BlobsBundle]_GetExecutionPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_ExecutionPayloadEnvelope[ExecutionPayloadT constraints.JSONMarshallable, BlobsBundleT BlobsBundle]_GetValue__ as it appears to be an interface:

// skipping Fuzz_Nosy_ExecutionPayloadEnvelope[ExecutionPayloadT constraints.JSONMarshallable, BlobsBundleT BlobsBundle]_ShouldOverrideBuilder__ as it appears to be an interface:

// skipping Fuzz_Nosy_NewPayloadRequest[ExecutionPayloadT interface{GetBaseFeePerGas() *math.U256; GetBlobGasUsed() math.U64; GetBlockHash() common.ExecutionHash; GetExcessBlobGas() math.U64; GetExtraData() []byte; GetFeeRecipient() common.ExecutionAddress; GetGasLimit() math.U64; GetGasUsed() math.U64; GetLogsBloom() bytes.B256; GetNumber() math.U64; GetParentHash() common.ExecutionHash; GetPrevRandao() common.Bytes32; GetReceiptsRoot() common.Bytes32; GetStateRoot() common.Bytes32; GetTimestamp() math.U64; GetTransactions() Transactions; GetWithdrawals() WithdrawalsT; constraints.ForkTyped[ExecutionPayloadT]}, WithdrawalsT interface{EncodeIndex(int, *bytes.Buffer); Len() int}]_HasValidVersionedAndBlockHashes__ as it appears to be an interface:

// skipping Fuzz_Nosy_PayloadAttributes[WithdrawalT any]_GetSuggestedFeeRecipient__ as it appears to be an interface:

// skipping Fuzz_Nosy_PayloadAttributes[WithdrawalT any]_IsNil__ as it appears to be an interface:

// skipping Fuzz_Nosy_PayloadAttributes[WithdrawalT any]_New__ because parameters include func, chan, or unsupported interface: []WithdrawalT

// skipping Fuzz_Nosy_PayloadAttributes[WithdrawalT any]_Validate__ as it appears to be an interface:

// skipping Fuzz_Nosy_PayloadAttributes[WithdrawalT any]_Version__ as it appears to be an interface:

func Fuzz_Nosy_Withdrawal_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var c *ssz.Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if w == nil || c == nil {
			return
		}

		w.DefineSSZ(c)
	})
}

func Fuzz_Nosy_Withdrawal_Equals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var other *Withdrawal
		fill_err = tp.Fill(&other)
		if fill_err != nil {
			return
		}
		if w == nil || other == nil {
			return
		}

		w.Equals(other)
	})
}

func Fuzz_Nosy_Withdrawal_GetAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.GetAddress()
	})
}

func Fuzz_Nosy_Withdrawal_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.GetAmount()
	})
}

func Fuzz_Nosy_Withdrawal_GetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.GetIndex()
	})
}

func Fuzz_Nosy_Withdrawal_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.GetTree()
	})
}

func Fuzz_Nosy_Withdrawal_GetValidatorIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.GetValidatorIndex()
	})
}

func Fuzz_Nosy_Withdrawal_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_Withdrawal_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_Withdrawal_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.MarshalSSZ()
	})
}

func Fuzz_Nosy_Withdrawal_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_Withdrawal_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var index math.U64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var validator math.ValidatorIndex
		fill_err = tp.Fill(&validator)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.New(index, validator, address, amount)
	})
}

func Fuzz_Nosy_Withdrawal_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Withdrawal
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_Withdrawal_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w *Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if w == nil {
			return
		}

		w.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_BartioTransactions_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txs BartioTransactions
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}

		txs.HashTreeRoot()
	})
}

func Fuzz_Nosy_BartioTx_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx BartioTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		tx.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_BartioTx_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx BartioTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}

		tx.HashTreeRoot()
	})
}

func Fuzz_Nosy_BartioTx_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx BartioTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}

		tx.SizeSSZ()
	})
}

// skipping Fuzz_Nosy_BlobsBundle_GetBlobs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BlobsBundle

// skipping Fuzz_Nosy_BlobsBundle_GetCommitments__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BlobsBundle

// skipping Fuzz_Nosy_BlobsBundle_GetProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BlobsBundle

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT any]_GetBlobsBundle__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BuiltExecutionPayloadEnv[ExecutionPayloadT any]

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT any]_GetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BuiltExecutionPayloadEnv[ExecutionPayloadT any]

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT any]_GetValue__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BuiltExecutionPayloadEnv[ExecutionPayloadT any]

// skipping Fuzz_Nosy_BuiltExecutionPayloadEnv[ExecutionPayloadT any]_ShouldOverrideBuilder__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.BuiltExecutionPayloadEnv[ExecutionPayloadT any]

// skipping Fuzz_Nosy_PayloadAttributer_GetSuggestedFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives.PayloadAttributer

func Fuzz_Nosy_Roots_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var roots Roots
		fill_err = tp.Fill(&roots)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		roots.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_Roots_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var roots Roots
		fill_err = tp.Fill(&roots)
		if fill_err != nil {
			return
		}

		roots.SizeSSZ()
	})
}

func Fuzz_Nosy_Withdrawal_EncodeRLP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawal
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var _w io.Writer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}

		w.EncodeRLP(_w)
	})
}

func Fuzz_Nosy_Withdrawals_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawals
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		w.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_Withdrawals_EncodeIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawals
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}
		var i int
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var _w *bytes.Buffer
		fill_err = tp.Fill(&_w)
		if fill_err != nil {
			return
		}
		if _w == nil {
			return
		}

		w.EncodeIndex(i, _w)
	})
}

func Fuzz_Nosy_Withdrawals_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawals
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		w.HashTreeRoot()
	})
}

func Fuzz_Nosy_Withdrawals_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawals
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		w.Len()
	})
}

func Fuzz_Nosy_Withdrawals_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var w Withdrawals
		fill_err = tp.Fill(&w)
		if fill_err != nil {
			return
		}

		w.SizeSSZ()
	})
}
