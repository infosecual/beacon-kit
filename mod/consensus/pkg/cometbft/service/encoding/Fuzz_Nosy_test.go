package encoding

import (
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

// skipping Fuzz_Nosy_ABCIRequest_GetHeight__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest

// skipping Fuzz_Nosy_ABCIRequest_GetTime__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest

// skipping Fuzz_Nosy_ABCIRequest_GetTxs__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest

// skipping Fuzz_Nosy_BeaconBlock[T any]_NewFromSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.BeaconBlock[T any]

// skipping Fuzz_Nosy_ExtractBlobsAndBlockFromRequest__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest

// skipping Fuzz_Nosy_UnmarshalBeaconBlockFromABCIRequest__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest

// skipping Fuzz_Nosy_UnmarshalBlobSidecarsFromABCIRequest__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/consensus/pkg/cometbft/service/encoding.ABCIRequest
