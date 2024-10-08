package engine

import (
	"testing"

	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
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

// skipping Fuzz_Nosy_Engine[ExecutionPayloadT ExecutionPayload[ExecutionPayloadT, WithdrawalsT], PayloadAttributesT engineprimitives.PayloadAttributer, PayloadIDT ~[8]byte, WithdrawalsT interface{EncodeIndex(int, *bytes.Buffer); Len() int}]_VerifyAndNotifyNewPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_Engine[ExecutionPayloadT ExecutionPayload[ExecutionPayloadT, WithdrawalsT], PayloadAttributesT engineprimitives.PayloadAttributer, PayloadIDT ~[8]byte, WithdrawalsT interface{EncodeIndex(int, *bytes.Buffer); Len() int}]_GetPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_Engine[ExecutionPayloadT ExecutionPayload[ExecutionPayloadT, WithdrawalsT], PayloadAttributesT engineprimitives.PayloadAttributer, PayloadIDT ~[8]byte, WithdrawalsT interface{EncodeIndex(int, *bytes.Buffer); Len() int}]_NotifyForkchoiceUpdate__ as it appears to be an interface:

// skipping Fuzz_Nosy_Engine[ExecutionPayloadT ExecutionPayload[ExecutionPayloadT, WithdrawalsT], PayloadAttributesT engineprimitives.PayloadAttributer, PayloadIDT ~[8]byte, WithdrawalsT interface{EncodeIndex(int, *bytes.Buffer); Len() int}]_Start__ as it appears to be an interface:

func Fuzz_Nosy_engineMetrics_errorLoggerFn__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var isOptimistic bool
		fill_err = tp.Fill(&isOptimistic)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.errorLoggerFn(isOptimistic)
	})
}

// skipping Fuzz_Nosy_engineMetrics_markForkchoiceUpdateAcceptedSyncing__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_engineMetrics_markForkchoiceUpdateInvalid__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_engineMetrics_markForkchoiceUpdateJSONRPCError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_engineMetrics_markForkchoiceUpdateUndefinedError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_engineMetrics_markForkchoiceUpdateValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var state *engineprimitives.ForkchoiceStateV1
		fill_err = tp.Fill(&state)
		if fill_err != nil {
			return
		}
		var hasPayloadAttributes bool
		fill_err = tp.Fill(&hasPayloadAttributes)
		if fill_err != nil {
			return
		}
		var payloadID *engineprimitives.PayloadID
		fill_err = tp.Fill(&payloadID)
		if fill_err != nil {
			return
		}
		if em == nil || state == nil || payloadID == nil {
			return
		}

		em.markForkchoiceUpdateValid(state, hasPayloadAttributes, payloadID)
	})
}

func Fuzz_Nosy_engineMetrics_markNewPayloadAcceptedSyncingPayloadStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var payloadHash common.ExecutionHash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}
		var parentHash common.ExecutionHash
		fill_err = tp.Fill(&parentHash)
		if fill_err != nil {
			return
		}
		var isOptimistic bool
		fill_err = tp.Fill(&isOptimistic)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.markNewPayloadAcceptedSyncingPayloadStatus(payloadHash, parentHash, isOptimistic)
	})
}

func Fuzz_Nosy_engineMetrics_markNewPayloadCalled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var payloadHash common.ExecutionHash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}
		var parentHash common.ExecutionHash
		fill_err = tp.Fill(&parentHash)
		if fill_err != nil {
			return
		}
		var isOptimistic bool
		fill_err = tp.Fill(&isOptimistic)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.markNewPayloadCalled(payloadHash, parentHash, isOptimistic)
	})
}

func Fuzz_Nosy_engineMetrics_markNewPayloadInvalidPayloadStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var payloadHash common.ExecutionHash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}
		var isOptimistic bool
		fill_err = tp.Fill(&isOptimistic)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.markNewPayloadInvalidPayloadStatus(payloadHash, isOptimistic)
	})
}

// skipping Fuzz_Nosy_engineMetrics_markNewPayloadJSONRPCError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_engineMetrics_markNewPayloadUndefinedError__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_engineMetrics_markNewPayloadValid__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var payloadHash common.ExecutionHash
		fill_err = tp.Fill(&payloadHash)
		if fill_err != nil {
			return
		}
		var parentHash common.ExecutionHash
		fill_err = tp.Fill(&parentHash)
		if fill_err != nil {
			return
		}
		var isOptimistic bool
		fill_err = tp.Fill(&isOptimistic)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.markNewPayloadValid(payloadHash, parentHash, isOptimistic)
	})
}

func Fuzz_Nosy_engineMetrics_markNotifyForkchoiceUpdateCalled__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var em *engineMetrics
		fill_err = tp.Fill(&em)
		if fill_err != nil {
			return
		}
		var hasPayloadAttributes bool
		fill_err = tp.Fill(&hasPayloadAttributes)
		if fill_err != nil {
			return
		}
		if em == nil {
			return
		}

		em.markNotifyForkchoiceUpdateCalled(hasPayloadAttributes)
	})
}

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetBaseFeePerGas__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetBlobGasUsed__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetBlockHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetExcessBlobGas__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetExtraData__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetGasLimit__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetGasUsed__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetLogsBloom__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetParentHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetPrevRandao__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetReceiptsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetTransactions__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]_GetWithdrawals__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.ExecutionPayload[ExecutionPayloadT, WithdrawalsT any]

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.TelemetrySink

// skipping Fuzz_Nosy_Withdrawal[WithdrawalT any]_GetAddress__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.Withdrawal[WithdrawalT any]

// skipping Fuzz_Nosy_Withdrawal[WithdrawalT any]_GetAmount__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.Withdrawal[WithdrawalT any]

// skipping Fuzz_Nosy_Withdrawal[WithdrawalT any]_GetIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.Withdrawal[WithdrawalT any]

// skipping Fuzz_Nosy_Withdrawal[WithdrawalT any]_GetValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/engine.Withdrawal[WithdrawalT any]
