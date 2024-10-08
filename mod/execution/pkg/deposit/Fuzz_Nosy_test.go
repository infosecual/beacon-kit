package deposit

import (
	"testing"

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

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_depositFetcher__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.Event[BeaconBlockT]

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_depositCatchupFetcher__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_eventLoop__ as it appears to be an interface:

// skipping Fuzz_Nosy_Service[BeaconBlockT BeaconBlock[BeaconBlockBodyT], BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], DepositT Deposit[DepositT, WithdrawalCredentialsT], ExecutionPayloadT ExecutionPayload, WithdrawalCredentialsT any]_fetchAndStoreDeposits__ as it appears to be an interface:

// skipping Fuzz_Nosy_WrappedBeaconDepositContract[DepositT Deposit[DepositT, WithdrawalCredentialsT], WithdrawalCredentialsT ~[32]byte]_ReadDeposits__ as it appears to be an interface:

func Fuzz_Nosy_metrics_markFailedToGetBlockLogs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *metrics
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var blockNum math.U64
		fill_err = tp.Fill(&blockNum)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.markFailedToGetBlockLogs(blockNum)
	})
}

// skipping Fuzz_Nosy_BeaconBlockBody[DepositT any, ExecutionPayloadT ExecutionPayload]_GetDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT any, ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_BeaconBlockBody[DepositT any, ExecutionPayloadT ExecutionPayload]_GetExecutionPayload__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT any, ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT any]_GetBody__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT any]_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT any]

// skipping Fuzz_Nosy_BlockEvent[DepositT any, BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT ExecutionPayload]_Context__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BlockEvent[DepositT any, BeaconBlockBodyT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_BlockEvent[DepositT any, BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT ExecutionPayload]_Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BlockEvent[DepositT any, BeaconBlockBodyT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_BlockEvent[DepositT any, BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT ExecutionPayload]_ID__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BlockEvent[DepositT any, BeaconBlockBodyT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_BlockEvent[DepositT any, BeaconBlockBodyT BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT ExecutionPayload]_Is__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BlockEvent[DepositT any, BeaconBlockBodyT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlockBody[DepositT, ExecutionPayloadT], BeaconBlockT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.BeaconBlock[BeaconBlockBodyT], ExecutionPayloadT github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload]

// skipping Fuzz_Nosy_Contract[DepositT any]_ReadDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.Contract[DepositT any]

// skipping Fuzz_Nosy_Deposit[DepositT, WithdrawalCredentialsT any]_GetIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.Deposit[DepositT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_Deposit[DepositT, WithdrawalCredentialsT any]_New__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.Deposit[DepositT, WithdrawalCredentialsT any]

// skipping Fuzz_Nosy_ExecutionPayload_GetNumber__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.ExecutionPayload

// skipping Fuzz_Nosy_Store[DepositT any]_EnqueueDeposits__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.Store[DepositT any]

// skipping Fuzz_Nosy_Store[DepositT any]_Prune__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.Store[DepositT any]

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/execution/pkg/deposit.TelemetrySink

// skipping Fuzz_Nosy_BuildPruneRangeFn__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/common.ChainSpec
