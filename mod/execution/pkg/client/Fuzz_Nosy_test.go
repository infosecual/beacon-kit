package client

import (
	"testing"
	"time"

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

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_GetPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_NewPayload__ because parameters include func, chan, or unsupported interface: ExecutionPayloadT

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_ForkchoiceUpdated__ because parameters include func, chan, or unsupported interface: PayloadAttributesT

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_ExchangeCapabilities__ as it appears to be an interface:

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_createContextWithTimeout__ as it appears to be an interface:

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_handleRPCError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_EngineClient[ExecutionPayloadT constraints.EngineType[ExecutionPayloadT], PayloadAttributesT PayloadAttributes]_verifyChainIDAndConnection__ as it appears to be an interface:

func Fuzz_Nosy_clientMetrics_incrementErrorCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var metricName string
		fill_err = tp.Fill(&metricName)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementErrorCounter(metricName)
	})
}

func Fuzz_Nosy_clientMetrics_incrementForkchoiceUpdateTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementForkchoiceUpdateTimeout()
	})
}

func Fuzz_Nosy_clientMetrics_incrementGetPayloadTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementGetPayloadTimeout()
	})
}

func Fuzz_Nosy_clientMetrics_incrementHTTPTimeoutCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementHTTPTimeoutCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInternalErrorCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInternalErrorCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInternalServerErrorCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInternalServerErrorCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInvalidForkchoiceStateCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInvalidForkchoiceStateCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInvalidParamsCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInvalidParamsCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInvalidPayloadAttributesCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInvalidPayloadAttributesCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementInvalidRequestCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementInvalidRequestCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementMethodNotFoundCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementMethodNotFoundCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementNewPayloadTimeout__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementNewPayloadTimeout()
	})
}

func Fuzz_Nosy_clientMetrics_incrementParseErrorCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementParseErrorCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementRequestTooLargeCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementRequestTooLargeCounter()
	})
}

func Fuzz_Nosy_clientMetrics_incrementTimeoutCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var metricName string
		fill_err = tp.Fill(&metricName)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementTimeoutCounter(metricName)
	})
}

func Fuzz_Nosy_clientMetrics_incrementUnknownPayloadErrorCounter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.incrementUnknownPayloadErrorCounter()
	})
}

func Fuzz_Nosy_clientMetrics_measureForkchoiceUpdateDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.measureForkchoiceUpdateDuration(startTime)
	})
}

func Fuzz_Nosy_clientMetrics_measureGetPayloadDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.measureGetPayloadDuration(startTime)
	})
}

func Fuzz_Nosy_clientMetrics_measureNewPayloadDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *clientMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.measureNewPayloadDuration(startTime)
	})
}

// skipping Fuzz_Nosy_PayloadAttributes_GetSuggestedFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/client.PayloadAttributes

// skipping Fuzz_Nosy_PayloadAttributes_IsNil__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/client.PayloadAttributes

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/client.TelemetrySink

// skipping Fuzz_Nosy_TelemetrySink_MeasureSince__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/client.TelemetrySink
