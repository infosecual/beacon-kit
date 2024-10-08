package ethclient

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
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

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_ChainID__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_ExchangeCapabilities__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_FilterLogs__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_ForkchoiceUpdated__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_ForkchoiceUpdatedV3__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_GetClientVersionV1__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_GetPayload__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_GetPayloadV3__ as it appears to be an interface:

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_NewPayload__ because parameters include func, chan, or unsupported interface: ExecutionPayloadT

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_NewPayloadV3__ because parameters include func, chan, or unsupported interface: ExecutionPayloadT

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_SubscribeFilterLogs__ because parameters include func, chan, or unsupported interface: chan<- github.com/ethereum/go-ethereum/core/types.Log

// skipping Fuzz_Nosy_Client[ExecutionPayloadT interface{Empty(uint32) ExecutionPayloadT; Version() uint32; constraints.JSONMarshallable}]_forkchoiceUpdated__ because parameters include func, chan, or unsupported interface: any

func Fuzz_Nosy_toBlockNumArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var number *big.Int
		fill_err = tp.Fill(&number)
		if fill_err != nil {
			return
		}
		if number == nil {
			return
		}

		toBlockNumArg(number)
	})
}

func Fuzz_Nosy_toFilterArg__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var q ethereum.FilterQuery
		fill_err = tp.Fill(&q)
		if fill_err != nil {
			return
		}

		toFilterArg(q)
	})
}
