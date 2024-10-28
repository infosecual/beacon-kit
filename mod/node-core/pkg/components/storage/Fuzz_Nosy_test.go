package storage

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

func Fuzz_Nosy_KVStoreProvider_OpenKVStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *KVStoreProvider
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.OpenKVStore(_x2)
	})
}

// skipping Fuzz_Nosy_Backend[AvailabilityStoreT any, BeaconStateT BeaconState[BeaconStateT, KVStoreT], BlockStoreT, DepositStoreT any, KVStoreT KVStore[KVStoreT]]_AvailabilityStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_Backend[AvailabilityStoreT any, BeaconStateT BeaconState[BeaconStateT, KVStoreT], BlockStoreT, DepositStoreT any, KVStoreT KVStore[KVStoreT]]_StateFromContext__ as it appears to be an interface:

// skipping Fuzz_Nosy_Backend[AvailabilityStoreT any, BeaconStateT BeaconState[BeaconStateT, KVStoreT], BlockStoreT, DepositStoreT any, KVStoreT KVStore[KVStoreT]]_BlockStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_Backend[AvailabilityStoreT any, BeaconStateT BeaconState[BeaconStateT, KVStoreT], BlockStoreT, DepositStoreT any, KVStoreT KVStore[KVStoreT]]_DepositStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_Backend[AvailabilityStoreT any, BeaconStateT BeaconState[BeaconStateT, KVStoreT], BlockStoreT, DepositStoreT any, KVStoreT KVStore[KVStoreT]]_BeaconStore__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[T, KVStoreT any]_NewFromDB__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components/storage.BeaconState[T, KVStoreT any]

// skipping Fuzz_Nosy_KVStore[T any]_WithContext__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/components/storage.KVStore[T any]
