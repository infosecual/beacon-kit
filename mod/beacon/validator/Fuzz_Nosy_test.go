package validator

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

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_buildBlockBody__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_retrieveExecutionPayload__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_computeAndSetStateRoot__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_computeStateRoot__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_getEmptyBeaconBlockForSlot__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_buildBlockAndSidecars__ because parameters include func, chan, or unsupported interface: SlotDataT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_handleNewSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.Event[SlotDataT]

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_buildRandaoReveal__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[AttestationDataT any, BeaconBlockT BeaconBlock[BeaconBlockT, BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT], BeaconStateT BeaconState[ExecutionPayloadHeaderT], BlobSidecarsT, DepositT any, DepositStoreT DepositStore[DepositT], Eth1DataT Eth1Data[Eth1DataT], ExecutionPayloadT any, ExecutionPayloadHeaderT ExecutionPayloadHeader, ForkDataT ForkData[ForkDataT], SlashingInfoT any, SlotDataT SlotData[AttestationDataT, SlashingInfoT]]_eventLoop__ as it appears to be an interface:

// skipping Fuzz_Nosy_validatorMetrics_failedToRetrievePayload__ because parameters include func, chan, or unsupported interface: error

func Fuzz_Nosy_validatorMetrics_measureRequestBlockForProposalTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *validatorMetrics
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

		cm.measureRequestBlockForProposalTime(start)
	})
}

func Fuzz_Nosy_validatorMetrics_measureStateRootComputationTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *validatorMetrics
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

		cm.measureStateRootComputationTime(start)
	})
}

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetAttestations__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetBlobKzgCommitments__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetGraffiti__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetRandaoReveal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetSlashingInfo__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlockBody[AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_GetBody__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_GetParentBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_NewWithVersion__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT any]_SetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconBlock[T, BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_GetBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_GetEth1DepositIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_GetGenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_GetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_HashTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BeaconState[ExecutionPayloadHeaderT any]_ValidatorIndexByPubkey__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BeaconState[ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_BlobFactory[BeaconBlockT, BlobSidecarsT any]_BuildSidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.BlobFactory[BeaconBlockT, BlobSidecarsT any]

// skipping Fuzz_Nosy_DepositStore[DepositT any]_GetDepositsByIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.DepositStore[DepositT any]

// skipping Fuzz_Nosy_Eth1Data[T any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.Eth1Data[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetBlockHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.ExecutionPayloadHeader

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetParentHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.ExecutionPayloadHeader

// skipping Fuzz_Nosy_ExecutionPayloadHeader_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.ExecutionPayloadHeader

// skipping Fuzz_Nosy_ForkData[T any]_ComputeRandaoSigningRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.ForkData[T any]

// skipping Fuzz_Nosy_ForkData[T any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.ForkData[T any]

// skipping Fuzz_Nosy_PayloadBuilder[BeaconStateT, ExecutionPayloadT any]_RequestPayloadSync__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.PayloadBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_PayloadBuilder[BeaconStateT, ExecutionPayloadT any]_RetrievePayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.PayloadBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_SlotData[AttestationDataT, SlashingInfoT any]_GetAttestationData__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.SlotData[AttestationDataT, SlashingInfoT any]

// skipping Fuzz_Nosy_SlotData[AttestationDataT, SlashingInfoT any]_GetSlashingInfo__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.SlotData[AttestationDataT, SlashingInfoT any]

// skipping Fuzz_Nosy_SlotData[AttestationDataT, SlashingInfoT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.SlotData[AttestationDataT, SlashingInfoT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, ExecutionPayloadHeaderT any]_ProcessSlots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, ExecutionPayloadHeaderT any]_Transition__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StorageBackend[BeaconStateT, DepositStoreT any]_DepositStore__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.StorageBackend[BeaconStateT, DepositStoreT any]

// skipping Fuzz_Nosy_StorageBackend[BeaconStateT, DepositStoreT any]_StateFromContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.StorageBackend[BeaconStateT, DepositStoreT any]

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.TelemetrySink

// skipping Fuzz_Nosy_TelemetrySink_MeasureSince__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/beacon/validator.TelemetrySink
