package mocks

import (
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

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_IsDataAvailable__ because parameters include func, chan, or unsupported interface: BeaconBlockBodyT

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Persist__ because parameters include func, chan, or unsupported interface: BlobSidecarsT

// skipping Fuzz_Nosy_AvailabilityStore_Expecter[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_IsDataAvailable__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_AvailabilityStore_Expecter[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Persist__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 context.Context, _a1 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, _a2 BeaconBlockBodyT)

// skipping Fuzz_Nosy_AvailabilityStore_IsDataAvailable_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, BeaconBlockBodyT) bool

// skipping Fuzz_Nosy_AvailabilityStore_Persist_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_AvailabilityStore_Persist_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, _a1 BlobSidecarsT)

// skipping Fuzz_Nosy_AvailabilityStore_Persist_Call[BeaconBlockBodyT interface{}, BlobSidecarsT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, BlobSidecarsT) error

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_GetBodyRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_GetParentBlockRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_GetProposerIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_GetSlot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_GetStateRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_HashTreeRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_MarshalSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_New__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_SetStateRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader[BeaconBlockHeaderT interface{}]_UnmarshalSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_GetBodyRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_GetParentBlockRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_GetProposerIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_GetSlot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_GetStateRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_HashTreeRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_MarshalSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_New__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_SetStateRoot__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconBlockHeader_Expecter[BeaconBlockHeaderT interface{}]_UnmarshalSSZ__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_GetBodyRoot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root

// skipping Fuzz_Nosy_BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_GetParentBlockRoot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root

// skipping Fuzz_Nosy_BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_GetProposerIndex_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64

// skipping Fuzz_Nosy_BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_GetSlot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64

// skipping Fuzz_Nosy_BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_GetStateRoot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root

// skipping Fuzz_Nosy_BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_HashTreeRoot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root

// skipping Fuzz_Nosy_BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlockHeader_MarshalSSZ_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() ([]byte, error)

// skipping Fuzz_Nosy_BeaconBlockHeader_New_Call[BeaconBlockHeaderT interface{}]_Return__ because parameters include func, chan, or unsupported interface: BeaconBlockHeaderT

// skipping Fuzz_Nosy_BeaconBlockHeader_New_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(slot github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, proposerIndex github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, parentBlockRoot github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, stateRoot github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, bodyRoot github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root)

// skipping Fuzz_Nosy_BeaconBlockHeader_New_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root) BeaconBlockHeaderT

// skipping Fuzz_Nosy_BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT interface{}]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root)

// skipping Fuzz_Nosy_BeaconBlockHeader_SetStateRoot_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root)

// skipping Fuzz_Nosy_BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 []byte)

// skipping Fuzz_Nosy_BeaconBlockHeader_UnmarshalSSZ_Call[BeaconBlockHeaderT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) error

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ExpectedWithdrawals__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetBalance__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetBlockRootAtIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetEth1Data__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetEth1DepositIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetFork__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetGenesisValidatorsRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetLatestBlockHeader__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetLatestExecutionPayloadHeader__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetNextWithdrawalIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetNextWithdrawalValidatorIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetRandaoMixAtIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetSlashingAtIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetSlot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalActiveBalances__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalSlashing__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalValidators__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetValidators__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetValidatorsByEffectiveBalance__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_SetSlot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_StateRootAtIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorByIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorIndexByCometBFTAddress__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorIndexByPubkey__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_ExpectedWithdrawals_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: []WithdrawalT

// skipping Fuzz_Nosy_BeaconState_ExpectedWithdrawals_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_ExpectedWithdrawals_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() ([]WithdrawalT, error)

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ExpectedWithdrawals__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetBalance__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetBlockRootAtIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetEth1Data__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetEth1DepositIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetFork__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetGenesisValidatorsRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetLatestBlockHeader__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetLatestExecutionPayloadHeader__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetNextWithdrawalIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetNextWithdrawalValidatorIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetRandaoMixAtIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetSlashingAtIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetSlot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalActiveBalances__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalSlashing__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetTotalValidators__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetValidators__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_GetValidatorsByEffectiveBalance__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_SetSlot__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_StateRootAtIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorByIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorIndexByCometBFTAddress__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_Expecter[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_ValidatorIndexByPubkey__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BeaconState_GetBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_BeaconState_GetBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetBlockRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetBlockRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 uint64)

// skipping Fuzz_Nosy_BeaconState_GetBlockRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64) (github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, error)

// skipping Fuzz_Nosy_BeaconState_GetEth1Data_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: Eth1DataT

// skipping Fuzz_Nosy_BeaconState_GetEth1Data_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetEth1Data_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (Eth1DataT, error)

// skipping Fuzz_Nosy_BeaconState_GetEth1DepositIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetEth1DepositIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetEth1DepositIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (uint64, error)

// skipping Fuzz_Nosy_BeaconState_GetFork_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ForkT

// skipping Fuzz_Nosy_BeaconState_GetFork_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetFork_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (ForkT, error)

// skipping Fuzz_Nosy_BeaconState_GetGenesisValidatorsRoot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetGenesisValidatorsRoot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetGenesisValidatorsRoot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, error)

// skipping Fuzz_Nosy_BeaconState_GetLatestBlockHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: BeaconBlockHeaderT

// skipping Fuzz_Nosy_BeaconState_GetLatestBlockHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetLatestBlockHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (BeaconBlockHeaderT, error)

// skipping Fuzz_Nosy_BeaconState_GetLatestExecutionPayloadHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ExecutionPayloadHeaderT

// skipping Fuzz_Nosy_BeaconState_GetLatestExecutionPayloadHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetLatestExecutionPayloadHeader_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (ExecutionPayloadHeaderT, error)

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (uint64, error)

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalValidatorIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalValidatorIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetNextWithdrawalValidatorIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetRandaoMixAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetRandaoMixAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 uint64)

// skipping Fuzz_Nosy_BeaconState_GetRandaoMixAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64) (github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B32, error)

// skipping Fuzz_Nosy_BeaconState_GetSlashingAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetSlashingAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 uint64)

// skipping Fuzz_Nosy_BeaconState_GetSlashingAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetTotalActiveBalances_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetTotalActiveBalances_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 uint64)

// skipping Fuzz_Nosy_BeaconState_GetTotalActiveBalances_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetTotalSlashing_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetTotalSlashing_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetTotalSlashing_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_GetTotalValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_GetTotalValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetTotalValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (uint64, error)

// skipping Fuzz_Nosy_BeaconState_GetValidatorsByEffectiveBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: []ValidatorT

// skipping Fuzz_Nosy_BeaconState_GetValidatorsByEffectiveBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetValidatorsByEffectiveBalance_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() ([]ValidatorT, error)

// skipping Fuzz_Nosy_BeaconState_GetValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ValidatorsT

// skipping Fuzz_Nosy_BeaconState_GetValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconState_GetValidators_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (ValidatorsT, error)

// skipping Fuzz_Nosy_BeaconState_SetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_SetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_BeaconState_SetSlot_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) error

// skipping Fuzz_Nosy_BeaconState_StateRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_StateRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 uint64)

// skipping Fuzz_Nosy_BeaconState_StateRootAtIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64) (github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root, error)

// skipping Fuzz_Nosy_BeaconState_ValidatorByIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ValidatorT

// skipping Fuzz_Nosy_BeaconState_ValidatorByIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_BeaconState_ValidatorByIndex_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) (ValidatorT, error)

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByCometBFTAddress_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByCometBFTAddress_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(cometBFTAddress []byte)

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByCometBFTAddress_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByPubkey_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByPubkey_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48)

// skipping Fuzz_Nosy_BeaconState_ValidatorIndexByPubkey_Call[BeaconBlockHeaderT interface{}, Eth1DataT interface{}, ExecutionPayloadHeaderT interface{}, ForkT interface{}, ValidatorT interface{}, ValidatorsT interface{}, WithdrawalT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/bytes.B48) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT interface{}]_GetSlotByBlockRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT interface{}]_GetSlotByExecutionNumber__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockStore[BeaconBlockT interface{}]_GetSlotByStateRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockStore_Expecter[BeaconBlockT interface{}]_GetSlotByBlockRoot__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockStore_Expecter[BeaconBlockT interface{}]_GetSlotByExecutionNumber__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockStore_Expecter[BeaconBlockT interface{}]_GetSlotByStateRoot__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(root github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root)

// skipping Fuzz_Nosy_BlockStore_GetSlotByBlockRoot_Call[BeaconBlockT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(executionNumber github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_BlockStore_GetSlotByExecutionNumber_Call[BeaconBlockT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_BlockStore_GetSlotByStateRoot_Call[BeaconBlockT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_BlockStore_GetSlotByStateRoot_Call[BeaconBlockT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(root github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root)

// skipping Fuzz_Nosy_BlockStore_GetSlotByStateRoot_Call[BeaconBlockT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/common.Root) (github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, error)

// skipping Fuzz_Nosy_DepositStore[DepositT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_DepositStore[DepositT interface{}]_EnqueueDeposits__ because parameters include func, chan, or unsupported interface: []DepositT

// skipping Fuzz_Nosy_DepositStore[DepositT interface{}]_GetDepositsByIndex__ as it appears to be an interface:

// skipping Fuzz_Nosy_DepositStore[DepositT interface{}]_Prune__ as it appears to be an interface:

// skipping Fuzz_Nosy_DepositStore_EnqueueDeposits_Call[DepositT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DepositStore_EnqueueDeposits_Call[DepositT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(deposits []DepositT)

// skipping Fuzz_Nosy_DepositStore_EnqueueDeposits_Call[DepositT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]DepositT) error

// skipping Fuzz_Nosy_DepositStore_Expecter[DepositT interface{}]_EnqueueDeposits__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DepositStore_Expecter[DepositT interface{}]_GetDepositsByIndex__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DepositStore_Expecter[DepositT interface{}]_Prune__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DepositStore_GetDepositsByIndex_Call[DepositT interface{}]_Return__ because parameters include func, chan, or unsupported interface: []DepositT

// skipping Fuzz_Nosy_DepositStore_GetDepositsByIndex_Call[DepositT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(startIndex uint64, numView uint64)

// skipping Fuzz_Nosy_DepositStore_GetDepositsByIndex_Call[DepositT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64, uint64) ([]DepositT, error)

// skipping Fuzz_Nosy_DepositStore_Prune_Call[DepositT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DepositStore_Prune_Call[DepositT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(start uint64, end uint64)

// skipping Fuzz_Nosy_DepositStore_Prune_Call[DepositT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64, uint64) error

// skipping Fuzz_Nosy_Node[ContextT interface{}]_CreateQueryContext__ as it appears to be an interface:

// skipping Fuzz_Nosy_Node[ContextT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_Node_CreateQueryContext_Call[ContextT interface{}]_Return__ because parameters include func, chan, or unsupported interface: ContextT

// skipping Fuzz_Nosy_Node_CreateQueryContext_Call[ContextT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(height int64, prove bool)

// skipping Fuzz_Nosy_Node_CreateQueryContext_Call[ContextT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(int64, bool) (ContextT, error)

// skipping Fuzz_Nosy_Node_Expecter[ContextT interface{}]_CreateQueryContext__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_StateProcessor[BeaconStateT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_StateProcessor[BeaconStateT interface{}]_ProcessSlots__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_StateProcessor_Expecter[BeaconStateT interface{}]_ProcessSlots__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_StateProcessor_ProcessSlots_Call[BeaconStateT interface{}]_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_StateProcessor_ProcessSlots_Call[BeaconStateT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 BeaconStateT, _a1 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_StateProcessor_ProcessSlots_Call[BeaconStateT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(BeaconStateT, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) (github.com/berachain/beacon-kit/mod/primitives/pkg/transition.ValidatorUpdates, error)

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_AvailabilityStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_BlockStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_DepositStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_StateFromContext__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend_AvailabilityStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Return__ because parameters include func, chan, or unsupported interface: AvailabilityStoreT

// skipping Fuzz_Nosy_StorageBackend_AvailabilityStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_StorageBackend_AvailabilityStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() AvailabilityStoreT

// skipping Fuzz_Nosy_StorageBackend_BlockStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Return__ because parameters include func, chan, or unsupported interface: BlockStoreT

// skipping Fuzz_Nosy_StorageBackend_BlockStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_StorageBackend_BlockStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() BlockStoreT

// skipping Fuzz_Nosy_StorageBackend_DepositStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Return__ because parameters include func, chan, or unsupported interface: DepositStoreT

// skipping Fuzz_Nosy_StorageBackend_DepositStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_StorageBackend_DepositStore_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() DepositStoreT

// skipping Fuzz_Nosy_StorageBackend_Expecter[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_AvailabilityStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend_Expecter[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_BlockStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend_Expecter[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_DepositStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_StorageBackend_Expecter[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_StateFromContext__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_StorageBackend_StateFromContext_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Return__ because parameters include func, chan, or unsupported interface: BeaconStateT

// skipping Fuzz_Nosy_StorageBackend_StateFromContext_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 context.Context)

// skipping Fuzz_Nosy_StorageBackend_StateFromContext_Call[AvailabilityStoreT interface{}, BeaconStateT interface{}, BlockStoreT interface{}, DepositStoreT interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) BeaconStateT

// skipping Fuzz_Nosy_Validator[WithdrawalCredentialsT backend.WithdrawalCredentials]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator[WithdrawalCredentialsT backend.WithdrawalCredentials]_GetWithdrawalCredentials__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator[WithdrawalCredentialsT backend.WithdrawalCredentials]_IsFullyWithdrawable__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator[WithdrawalCredentialsT backend.WithdrawalCredentials]_IsPartiallyWithdrawable__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator_Expecter[WithdrawalCredentialsT backend.WithdrawalCredentials]_GetWithdrawalCredentials__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator_Expecter[WithdrawalCredentialsT backend.WithdrawalCredentials]_IsFullyWithdrawable__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Validator_Expecter[WithdrawalCredentialsT backend.WithdrawalCredentials]_IsPartiallyWithdrawable__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Return__ because parameters include func, chan, or unsupported interface: WithdrawalCredentialsT

// skipping Fuzz_Nosy_Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Validator_GetWithdrawalCredentials_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() WithdrawalCredentialsT

// skipping Fuzz_Nosy_Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Run__ because parameters include func, chan, or unsupported interface: func(amount github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, epoch github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_Validator_IsFullyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) bool

// skipping Fuzz_Nosy_Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_Run__ because parameters include func, chan, or unsupported interface: func(amount1 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, amount2 github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_Validator_IsPartiallyWithdrawable_Call[WithdrawalCredentialsT backend.WithdrawalCredentials]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) bool

func Fuzz_Nosy_WithdrawalCredentials_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *WithdrawalCredentials
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

func Fuzz_Nosy_WithdrawalCredentials_ToExecutionAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *WithdrawalCredentials
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.ToExecutionAddress()
	})
}

func Fuzz_Nosy_WithdrawalCredentials_Expecter_ToExecutionAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *WithdrawalCredentials_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.ToExecutionAddress()
	})
}

// skipping Fuzz_Nosy_WithdrawalCredentials_ToExecutionAddress_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_WithdrawalCredentials_ToExecutionAddress_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_WithdrawalCredentials_ToExecutionAddress_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() (github.com/berachain/beacon-kit/mod/primitives/pkg/common.ExecutionAddress, error)

// skipping Fuzz_Nosy_Withdrawal[T interface{}]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_Withdrawal[T interface{}]_New__ as it appears to be an interface:

// skipping Fuzz_Nosy_Withdrawal_Expecter[T interface{}]_New__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Withdrawal_New_Call[T interface{}]_Return__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_Withdrawal_New_Call[T interface{}]_Run__ because parameters include func, chan, or unsupported interface: func(index github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, validator github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, address github.com/berachain/beacon-kit/mod/primitives/pkg/common.ExecutionAddress, amount github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64)

// skipping Fuzz_Nosy_Withdrawal_New_Call[T interface{}]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64, github.com/berachain/beacon-kit/mod/primitives/pkg/common.ExecutionAddress, github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64) T
