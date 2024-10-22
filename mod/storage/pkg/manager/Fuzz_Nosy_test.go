package manager

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

func Fuzz_Nosy_DBManager_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DBManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Name()
	})
}

func Fuzz_Nosy_DBManager_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *DBManager
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Start(ctx)
	})
}

// skipping Fuzz_Nosy_BeaconBlock_GetSlot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/storage/pkg/manager.BeaconBlock

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT BeaconBlock]_Context__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/storage/pkg/manager.BlockEvent[BeaconBlockT github.com/berachain/beacon-kit/mod/storage/pkg/manager.BeaconBlock]

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT BeaconBlock]_Data__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/storage/pkg/manager.BlockEvent[BeaconBlockT github.com/berachain/beacon-kit/mod/storage/pkg/manager.BeaconBlock]

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT BeaconBlock]_ID__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/storage/pkg/manager.BlockEvent[BeaconBlockT github.com/berachain/beacon-kit/mod/storage/pkg/manager.BeaconBlock]

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT BeaconBlock]_Is__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/storage/pkg/manager.BlockEvent[BeaconBlockT github.com/berachain/beacon-kit/mod/storage/pkg/manager.BeaconBlock]
