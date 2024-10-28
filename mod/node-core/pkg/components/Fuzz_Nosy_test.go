package components

import (
	"context"
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

// skipping Fuzz_Nosy_AttributesFactory[BeaconStateT, PayloadAttributesT any]_BuildPayloadAttributes__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.AttributesFactory[BeaconStateT, PayloadAttributesT any]

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]_IsDataAvailable__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]_Persist__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_GetBlobKzgCommitments__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_GetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_GetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_GetRandaoReveal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_GetTopLevelRoots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_Length__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetAttestations__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetBlobKzgCommitments__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetGraffiti__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetRandaoReveal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]_SetSlashingInfo__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockBody[T, AttestationDataT, DepositT, Eth1DataT, ExecutionPayloadT, SlashingInfoT any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetBodyRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetParentBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetProposerIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_GetTree__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlockHeader[T any]_SetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlockHeader[T any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetBody__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetExecutionNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetParentBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetProposerIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_NewFromSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_NewWithVersion__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]_SetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconBlock[T, BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconStateMarshallable[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_GetTree__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStateMarshallable[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_BeaconStateMarshallable[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStateMarshallable[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]_Context__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]_Copy__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]_GetMarshallable__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]_HashTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]_NewFromDB__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconState[T, BeaconBlockHeaderT, BeaconStateMarshallableT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, KVStoreT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_AddValidator__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_AddValidatorBartio__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_Context__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_Copy__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBalances__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetEth1DepositIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetFork__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetGenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetLatestBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetNextWithdrawalIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetNextWithdrawalValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetRandaoMixAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetSlashingAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetSlashings__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalActiveBalances__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalSlashing__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalValidators__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetValidators__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetValidatorsByEffectiveBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetEth1DepositIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetFork__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetGenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetLatestBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetNextWithdrawalIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetNextWithdrawalValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetSlashingAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_SetTotalSlashing__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_StateRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_UpdateBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_UpdateRandaoMixAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_UpdateStateRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_UpdateValidatorAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_ValidatorByIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_ValidatorIndexByCometBFTAddress__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_ValidatorIndexByPubkey__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_WithContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BeaconStore[T, BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_BlobProcessor[AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT any]_ProcessSidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobProcessor[AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobProcessor[AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT any]_VerifySidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobProcessor[AvailabilityStoreT, BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobSidecar[BeaconBlockHeaderT any]_GetBeaconBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlobSidecar[BeaconBlockHeaderT any]_GetBlob__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlobSidecar[BeaconBlockHeaderT any]_GetKzgCommitment__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlobSidecar[BeaconBlockHeaderT any]_GetKzgProof__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlobSidecars[T, BlobSidecarT any]_Get__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecars[T, BlobSidecarT any]

// skipping Fuzz_Nosy_BlobSidecars[T, BlobSidecarT any]_GetSidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecars[T, BlobSidecarT any]

// skipping Fuzz_Nosy_BlobSidecars[T, BlobSidecarT any]_Len__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecars[T, BlobSidecarT any]

// skipping Fuzz_Nosy_BlobSidecars[T, BlobSidecarT any]_ValidateBlockRoots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecars[T, BlobSidecarT any]

// skipping Fuzz_Nosy_BlobSidecars[T, BlobSidecarT any]_VerifyInclusionProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobSidecars[T, BlobSidecarT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifyInclusionProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifyKZGProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifySidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_BlockBackend[BeaconBlockHeaderT any]_BlockHeaderAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockBackend[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlockBackend[BeaconBlockHeaderT any]_BlockRewardsAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockBackend[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlockBackend[BeaconBlockHeaderT any]_BlockRootAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockBackend[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT any]_GetSlotByBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockStore[BeaconBlockT any]

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT any]_GetSlotByExecutionNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockStore[BeaconBlockT any]

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT any]_GetSlotByStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockStore[BeaconBlockT any]

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT any]_Set__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.BlockStore[BeaconBlockT any]

// skipping Fuzz_Nosy_ConsensusEngine_PrepareProposal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ConsensusEngine

// skipping Fuzz_Nosy_ConsensusEngine_ProcessProposal__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ConsensusEngine

// skipping Fuzz_Nosy_DepositStore[DepositT any]_EnqueueDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.DepositStore[DepositT any]

// skipping Fuzz_Nosy_DepositStore[DepositT any]_GetDepositsByIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.DepositStore[DepositT any]

// skipping Fuzz_Nosy_DepositStore[DepositT any]_Prune__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.DepositStore[DepositT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_GetAmount__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_GetIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_GetPubkey__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_GetWithdrawalCredentials__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[T, ForkDataT, WithdrawalCredentialsT any]_VerifySignature__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Deposit[T, ForkDataT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_GetBlockHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_GetFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_GetNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_GetParentHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayloadHeader[T any]_NewFromSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayloadHeader[T any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetBaseFeePerGas__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetBlobGasUsed__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetBlockHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetExcessBlobGas__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetExtraData__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetFeeRecipient__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetGasLimit__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetGasUsed__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetLogsBloom__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetParentHash__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetPrevRandao__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetReceiptsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetTimestamp__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetTransactions__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_GetWithdrawals__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]_ToHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ExecutionPayload[ExecutionPayloadT, ExecutionPayloadHeaderT, WithdrawalsT any]

// skipping Fuzz_Nosy_GenesisBackend_GenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.GenesisBackend

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_Genesis[DepositT, ExecutionPayloadHeaderT any]_GetForkVersion__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Genesis[DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_HistoricalBackend[ForkT any]_StateForkAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.HistoricalBackend[ForkT any]

// skipping Fuzz_Nosy_HistoricalBackend[ForkT any]_StateRootAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.HistoricalBackend[ForkT any]

// skipping Fuzz_Nosy_IndexDB_Has__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.IndexDB

// skipping Fuzz_Nosy_IndexDB_Prune__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.IndexDB

// skipping Fuzz_Nosy_IndexDB_Set__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.IndexDB

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT, ExecutionPayloadT any]_Enabled__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.LocalBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT, ExecutionPayloadT any]_RequestPayloadAsync__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.LocalBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT, ExecutionPayloadT any]_RequestPayloadSync__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.LocalBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT, ExecutionPayloadT any]_RetrievePayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.LocalBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_LocalBuilder[BeaconStateT, ExecutionPayloadT any]_SendForceHeadFCU__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.LocalBuilder[BeaconStateT, ExecutionPayloadT any]

// skipping Fuzz_Nosy_NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]_AttachQueryBackend__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]_ChainSpec__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]_GetSlotByBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]_GetSlotByExecutionNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]_GetSlotByStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, NodeT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBeaconBackend[BeaconStateT, BeaconBlockHeaderT, ForkT, ValidatorT any]_GetSlotByBlockRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBeaconBackend[BeaconStateT, BeaconBlockHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIBeaconBackend[BeaconStateT, BeaconBlockHeaderT, ForkT, ValidatorT any]_GetSlotByStateRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIBeaconBackend[BeaconStateT, BeaconBlockHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_NodeAPIContext_Bind__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIContext

// skipping Fuzz_Nosy_NodeAPIContext_Validate__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIContext

// skipping Fuzz_Nosy_NodeAPIEngine[ContextT NodeAPIContext]_RegisterRoutes__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIEngine[ContextT github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIContext]

// skipping Fuzz_Nosy_NodeAPIEngine[ContextT NodeAPIContext]_Run__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIEngine[ContextT github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIContext]

// skipping Fuzz_Nosy_NodeAPIProofBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, ValidatorT any]_GetSlotByExecutionNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.NodeAPIProofBackend[BeaconBlockHeaderT, BeaconStateT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_RandaoBackend_RandaoAtEpoch__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.RandaoBackend

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBalances__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetFork__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetGenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetLatestBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetNextWithdrawalIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetNextWithdrawalValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetSlashingAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalActiveBalances__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalSlashing__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetTotalValidators__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetValidators__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_GetValidatorsByEffectiveBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]_ValidatorIndexByCometBFTAddress__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT, ValidatorsT, WithdrawalT any]

// skipping Fuzz_Nosy_ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_GetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_GetEth1DepositIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_GetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_ReadOnlyRandaoMixes_GetRandaoMixAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyRandaoMixes

// skipping Fuzz_Nosy_ReadOnlyStateRoots_StateRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyStateRoots

// skipping Fuzz_Nosy_ReadOnlyValidators[ValidatorT any]_ValidatorByIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyValidators[ValidatorT any]

// skipping Fuzz_Nosy_ReadOnlyValidators[ValidatorT any]_ValidatorIndexByPubkey__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyValidators[ValidatorT any]

// skipping Fuzz_Nosy_ReadOnlyWithdrawals[WithdrawalT any]_ExpectedWithdrawals__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ReadOnlyWithdrawals[WithdrawalT any]

// skipping Fuzz_Nosy_SidecarFactory[BeaconBlockT, BlobSidecarsT any]_BuildSidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.SidecarFactory[BeaconBlockT, BlobSidecarsT any]

// skipping Fuzz_Nosy_StateBackend[BeaconStateT, ForkT any]_StateForkAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateBackend[BeaconStateT, ForkT any]

// skipping Fuzz_Nosy_StateBackend[BeaconStateT, ForkT any]_StateFromSlotForProof__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateBackend[BeaconStateT, ForkT any]

// skipping Fuzz_Nosy_StateBackend[BeaconStateT, ForkT any]_StateRootAtSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateBackend[BeaconStateT, ForkT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_InitializePreminedBeaconStateFromEth1__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_ProcessSlots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]_Transition__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StateProcessor[BeaconBlockT, BeaconStateT, ContextT, DepositT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]_AvailabilityStore__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]_BlockStore__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]_DepositStore__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]_StateFromContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.StorageBackend[AvailabilityStoreT, BeaconStateT, BlockStoreT, DepositStoreT any]

// skipping Fuzz_Nosy_ValidatorBackend[ValidatorT any]_ValidatorBalancesByIDs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ValidatorBackend[ValidatorT any]

// skipping Fuzz_Nosy_ValidatorBackend[ValidatorT any]_ValidatorByID__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ValidatorBackend[ValidatorT any]

// skipping Fuzz_Nosy_ValidatorBackend[ValidatorT any]_ValidatorsByIDs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.ValidatorBackend[ValidatorT any]

// skipping Fuzz_Nosy_Withdrawal[T any]_Equals__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawal[T any]_GetAddress__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawal[T any]_GetAmount__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawal[T any]_GetIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawal[T any]_GetValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawal[T any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawal[T any]

// skipping Fuzz_Nosy_Withdrawals[WithdrawalT any]_EncodeIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawals[WithdrawalT any]

// skipping Fuzz_Nosy_Withdrawals[WithdrawalT any]_Len__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.Withdrawals[WithdrawalT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_DecreaseBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_IncreaseBalance__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetFork__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetGenesisValidatorsRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetLatestBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetNextWithdrawalIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetNextWithdrawalValidatorIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_SetTotalSlashing__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_UpdateBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]_UpdateSlashingAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyBeaconState[BeaconBlockHeaderT, Eth1DataT, ExecutionPayloadHeaderT, ForkT, ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_SetEth1Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_SetEth1DepositIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]_SetLatestExecutionPayloadHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyEth1Data[Eth1DataT, ExecutionPayloadHeaderT any]

// skipping Fuzz_Nosy_WriteOnlyRandaoMixes_UpdateRandaoMixAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyRandaoMixes

// skipping Fuzz_Nosy_WriteOnlyStateRoots_UpdateStateRootAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyStateRoots

// skipping Fuzz_Nosy_WriteOnlyValidators[ValidatorT any]_AddValidator__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyValidators[ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyValidators[ValidatorT any]_AddValidatorBartio__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyValidators[ValidatorT any]

// skipping Fuzz_Nosy_WriteOnlyValidators[ValidatorT any]_UpdateValidatorAtIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components.WriteOnlyValidators[ValidatorT any]

func Fuzz_Nosy_coreKVStore_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		store.Delete(key)
	})
}

func Fuzz_Nosy_coreKVStore_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		store.Get(key)
	})
}

func Fuzz_Nosy_coreKVStore_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}

		store.Has(key)
	})
}

func Fuzz_Nosy_coreKVStore_Iterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end []byte
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}

		store.Iterator(start, end)
	})
}

func Fuzz_Nosy_coreKVStore_ReverseIterator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var start []byte
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end []byte
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}

		store.ReverseIterator(start, end)
	})
}

func Fuzz_Nosy_coreKVStore_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var store coreKVStore
		fill_err = tp.Fill(&store)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}

		store.Set(key, value)
	})
}

func Fuzz_Nosy_kvStoreService_OpenKVStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var k kvStoreService
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}

		k.OpenKVStore(ctx)
	})
}

func Fuzz_Nosy_ProvideDispatcher__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in DispatcherInput[LoggerT]
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		ProvideDispatcher(in)
	})
}

func Fuzz_Nosy_ProvideNodeAPIHandlers__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in NodeAPIHandlersInput[BeaconBlockHeaderT, BeaconStateT, BeaconStateMarshallableT, ExecutionPayloadHeaderT, KVStoreT, NodeAPIContextT, WithdrawalT]
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		ProvideNodeAPIHandlers(in)
	})
}

func Fuzz_Nosy_ProvideReportingService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var in ReportingServiceInput[LoggerT]
		fill_err = tp.Fill(&in)
		if fill_err != nil {
			return
		}

		ProvideReportingService(in)
	})
}
