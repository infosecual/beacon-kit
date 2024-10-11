package middleware

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

func Fuzz_Nosy_ABCIMiddlewareMetrics_measurePrepareProposalDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *ABCIMiddlewareMetrics
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

		cm.measurePrepareProposalDuration(start)
	})
}

func Fuzz_Nosy_ABCIMiddlewareMetrics_measureProcessProposalDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cm *ABCIMiddlewareMetrics
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

		cm.measureProcessProposalDuration(start)
	})
}

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_PrepareProposal__ because parameters include func, chan, or unsupported interface: SlotDataT

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForBuiltBeaconBlock__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_FinalizeBlock__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_ProcessProposal__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_handleBuiltBeaconBlockAndSidecars__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_createProcessProposalResponse__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForBeaconBlockVerification__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForBuiltSidecars__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForSidecarVerification__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_InitGenesis__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForFinalValidatorUpdates__ as it appears to be an interface:

// skipping Fuzz_Nosy_ABCIMiddleware[BeaconBlockT BeaconBlock[BeaconBlockT], BlobSidecarsT BlobSidecars[BlobSidecarsT], GenesisT json.Unmarshaler, SlotDataT any]_waitForGenesisProcessed__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconBlock[SelfT any]_NewFromSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/middleware.BeaconBlock[SelfT any]

// skipping Fuzz_Nosy_TelemetrySink_MeasureSince__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/middleware.TelemetrySink
