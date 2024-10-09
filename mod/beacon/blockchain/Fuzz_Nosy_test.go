package blockchain

import (
	"testing"
	"time"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
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

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_sendNextFCUWithAttributes__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_executeStateTransition__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_handleOptimisticPayloadBuild__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_optimisticPayloadBuild__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_sendPostBlockFCU__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_verifyStateRoot__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_sendNextFCUWithoutAttributes__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_eventLoop__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_ProcessBeaconBlock__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_VerifyIncomingBlock__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_calculateNextTimestamp__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_handleBeaconBlockFinalization__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.Event[BeaconBlockT]

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_handleBeaconBlockReceived__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.Event[BeaconBlockT]

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_rebuildPayloadForRejectedBlock__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_forceStartupHead__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_handleRebuildPayloadForRejectedBlock__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_ProcessGenesisData__ because parameters include func, chan, or unsupported interface: GenesisT

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_handleGenDataReceived__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.Event[GenesisT]

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[ExecutionPayloadT], BeaconBlockHeaderT BeaconBlockHeader, BeaconStateT ReadOnlyBeaconState[BeaconStateT, BeaconBlockHeaderT, ExecutionPayloadHeaderT], DepositT any, ExecutionPayloadT ExecutionPayload, ExecutionPayloadHeaderT ExecutionPayloadHeader, GenesisT Genesis[DepositT, ExecutionPayloadHeaderT], PayloadAttributesT PayloadAttributes]_shouldBuildOptimisticPayloads__ as it appears to be an interface:

// skipping Fuzz_Nosy_chainMetrics_markOptimisticPayloadBuildFailure__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_chainMetrics_markOptimisticPayloadBuildSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.markOptimisticPayloadBuildSuccess(slot)
	})
}

// skipping Fuzz_Nosy_chainMetrics_markRebuildPayloadForRejectedBlockFailure__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_chainMetrics_markRebuildPayloadForRejectedBlockSuccess__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.markRebuildPayloadForRejectedBlockSuccess(slot)
	})
}

func Fuzz_Nosy_chainMetrics_measureStateRootVerificationTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.measureStateRootVerificationTime(start)
	})
}

func Fuzz_Nosy_chainMetrics_measureStateTransitionDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *chainMetrics
		fill_err = tp.Fill(&cm)
		if fill_err != nil {
			return
		}
		var start time.Time
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		if cm == nil {
			return
		}

		cm.measureStateTransitionDuration(start)
	})
}

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT any]_IsDataAvailable__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.AvailabilityStore[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlockBody[ExecutionPayloadT any]_GetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlockBody[ExecutionPayloadT any]

// skipping Fuzz_Nosy_BeaconBlockHeader_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlockHeader

// skipping Fuzz_Nosy_BeaconBlockHeader_SetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlockHeader

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT any]_GetBody__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlock[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlock[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BeaconBlock[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BlobSidecars_Len__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.BlobSidecars

// skipping Fuzz_Nosy_ExecutionEngine[PayloadAttributesT any]_NotifyForkchoiceUpdate__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ExecutionEngine[PayloadAttributesT any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetBlockHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ExecutionPayloadHeader

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetParentHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ExecutionPayloadHeader

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ExecutionPayloadHeader

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetForkVersion__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT any]_Enabled__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.LocalBuilder[BeaconStateT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT any]_RequestPayloadAsync__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.LocalBuilder[BeaconStateT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT any]_SendForceHeadFCU__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.LocalBuilder[BeaconStateT any]

// skipping Fuzz_Nosy_PayloadAttributes_GetSuggestedFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.PayloadAttributes

// skipping Fuzz_Nosy_PayloadAttributes_IsNil__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.PayloadAttributes

// skipping Fuzz_Nosy_PayloadAttributes_Version__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.PayloadAttributes

// skipping Fuzz_Nosy_ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]_Copy__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]_GetLatestBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]_GetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]_HashTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.ReadOnlyBeaconState[T, BeaconBlockHeaderT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_InitializePreminedBeaconStateFromEth1__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_ProcessSlots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_Transition__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT any]_AvailabilityStore__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.StorageBackend[AvailabilityStoreT, BeaconStateT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT any]_StateFromContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.StorageBackend[AvailabilityStoreT, BeaconStateT any]

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.TelemetrySink

// skipping Fuzz_Nosy_TelemetrySink_MeasureSince__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/blockchain.TelemetrySink
