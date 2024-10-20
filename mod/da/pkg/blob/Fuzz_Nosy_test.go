package blob

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

// skipping Fuzz_Nosy_Processor[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT], BeaconBlockBodyT any, BeaconBlockHeaderT BeaconBlockHeader, BlobSidecarT Sidecar[BeaconBlockHeaderT], BlobSidecarsT Sidecars[BlobSidecarT]]_ProcessSidecars__ because parameters include func, chan, or unsupported interface: AvailabilityStoreT

// skipping Fuzz_Nosy_Processor[AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT], BeaconBlockBodyT any, BeaconBlockHeaderT BeaconBlockHeader, BlobSidecarT Sidecar[BeaconBlockHeaderT], BlobSidecarsT Sidecars[BlobSidecarT]]_VerifySidecars__ because parameters include func, chan, or unsupported interface: BlobSidecarsT

// skipping Fuzz_Nosy_SidecarFactory[BeaconBlockT BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT], BeaconBlockBodyT BeaconBlockBody, BeaconBlockHeaderT any]_BuildSidecars__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_SidecarFactory[BeaconBlockT BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT], BeaconBlockBodyT BeaconBlockBody, BeaconBlockHeaderT any]_BuildBlockBodyProof__ because parameters include func, chan, or unsupported interface: BeaconBlockBodyT

// skipping Fuzz_Nosy_SidecarFactory[BeaconBlockT BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT], BeaconBlockBodyT BeaconBlockBody, BeaconBlockHeaderT any]_BuildCommitmentProof__ because parameters include func, chan, or unsupported interface: BeaconBlockBodyT

// skipping Fuzz_Nosy_SidecarFactory[BeaconBlockT BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT], BeaconBlockBodyT BeaconBlockBody, BeaconBlockHeaderT any]_BuildKZGInclusionProof__ because parameters include func, chan, or unsupported interface: BeaconBlockBodyT

// skipping Fuzz_Nosy_Verifier[BeaconBlockHeaderT BeaconBlockHeader, BlobSidecarT Sidecar[BeaconBlockHeaderT], BlobSidecarsT Sidecars[BlobSidecarT]]_VerifyInclusionProofs__ because parameters include func, chan, or unsupported interface: BlobSidecarsT

// skipping Fuzz_Nosy_Verifier[BeaconBlockHeaderT BeaconBlockHeader, BlobSidecarT Sidecar[BeaconBlockHeaderT], BlobSidecarsT Sidecars[BlobSidecarT]]_VerifyKZGProofs__ because parameters include func, chan, or unsupported interface: BlobSidecarsT

// skipping Fuzz_Nosy_Verifier[BeaconBlockHeaderT BeaconBlockHeader, BlobSidecarT Sidecar[BeaconBlockHeaderT], BlobSidecarsT Sidecars[BlobSidecarT]]_VerifySidecars__ because parameters include func, chan, or unsupported interface: BlobSidecarsT

func Fuzz_Nosy_factoryMetrics_measureBuildBlockBodyProofDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fm *factoryMetrics
		fill_err = tp.Fill(&fm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if fm == nil {
			return
		}

		fm.measureBuildBlockBodyProofDuration(startTime)
	})
}

func Fuzz_Nosy_factoryMetrics_measureBuildCommitmentProofDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fm *factoryMetrics
		fill_err = tp.Fill(&fm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if fm == nil {
			return
		}

		fm.measureBuildCommitmentProofDuration(startTime)
	})
}

func Fuzz_Nosy_factoryMetrics_measureBuildKZGInclusionProofDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fm *factoryMetrics
		fill_err = tp.Fill(&fm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		if fm == nil {
			return
		}

		fm.measureBuildKZGInclusionProofDuration(startTime)
	})
}

func Fuzz_Nosy_factoryMetrics_measureBuildSidecarsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fm *factoryMetrics
		fill_err = tp.Fill(&fm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		if fm == nil {
			return
		}

		fm.measureBuildSidecarsDuration(startTime, numSidecars)
	})
}

func Fuzz_Nosy_processorMetrics_measureProcessSidecarsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pm *processorMetrics
		fill_err = tp.Fill(&pm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		if pm == nil {
			return
		}

		pm.measureProcessSidecarsDuration(startTime, numSidecars)
	})
}

func Fuzz_Nosy_processorMetrics_measureVerifySidecarsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pm *processorMetrics
		fill_err = tp.Fill(&pm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		if pm == nil {
			return
		}

		pm.measureVerifySidecarsDuration(startTime, numSidecars)
	})
}

func Fuzz_Nosy_verifierMetrics_measureVerifyInclusionProofsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vm *verifierMetrics
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		if vm == nil {
			return
		}

		vm.measureVerifyInclusionProofsDuration(startTime, numSidecars)
	})
}

func Fuzz_Nosy_verifierMetrics_measureVerifyKZGProofsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vm *verifierMetrics
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		var kzgImplementation string
		fill_err = tp.Fill(&kzgImplementation)
		if fill_err != nil {
			return
		}
		if vm == nil {
			return
		}

		vm.measureVerifyKZGProofsDuration(startTime, numSidecars, kzgImplementation)
	})
}

func Fuzz_Nosy_verifierMetrics_measureVerifySidecarsDuration__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vm *verifierMetrics
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var startTime time.Time
		fill_err = tp.Fill(&startTime)
		if fill_err != nil {
			return
		}
		var numSidecars math.U64
		fill_err = tp.Fill(&numSidecars)
		if fill_err != nil {
			return
		}
		var kzgImplementation string
		fill_err = tp.Fill(&kzgImplementation)
		if fill_err != nil {
			return
		}
		if vm == nil {
			return
		}

		vm.measureVerifySidecarsDuration(startTime, numSidecars, kzgImplementation)
	})
}

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]_IsDataAvailable__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]_Persist__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT any]

// skipping Fuzz_Nosy_BeaconBlockBody_GetBlobKzgCommitments__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlockBody

// skipping Fuzz_Nosy_BeaconBlockBody_GetTopLevelRoots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlockBody

// skipping Fuzz_Nosy_BeaconBlockBody_Length__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlockBody

// skipping Fuzz_Nosy_BeaconBlockHeader_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlockHeader

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT any]_GetBody__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT any]_GetHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BeaconBlock[BeaconBlockBodyT, BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifyInclusionProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifyKZGProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_BlobVerifier[BlobSidecarsT any]_VerifySidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.BlobVerifier[BlobSidecarsT any]

// skipping Fuzz_Nosy_ChainSpec_MaxBlobCommitmentsPerBlock__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.ChainSpec

// skipping Fuzz_Nosy_Sidecar[BeaconBlockHeaderT any]_GetBeaconBlockHeader__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_Sidecar[BeaconBlockHeaderT any]_GetBlob__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_Sidecar[BeaconBlockHeaderT any]_GetKzgCommitment__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_Sidecar[BeaconBlockHeaderT any]_GetKzgProof__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecar[BeaconBlockHeaderT any]

// skipping Fuzz_Nosy_Sidecars[SidecarT any]_Get__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecars[SidecarT any]

// skipping Fuzz_Nosy_Sidecars[SidecarT any]_GetSidecars__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecars[SidecarT any]

// skipping Fuzz_Nosy_Sidecars[SidecarT any]_Len__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecars[SidecarT any]

// skipping Fuzz_Nosy_Sidecars[SidecarT any]_ValidateBlockRoots__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecars[SidecarT any]

// skipping Fuzz_Nosy_Sidecars[SidecarT any]_VerifyInclusionProofs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.Sidecars[SidecarT any]

// skipping Fuzz_Nosy_TelemetrySink_MeasureSince__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/da/pkg/blob.TelemetrySink
