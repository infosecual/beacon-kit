package service

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

// skipping Fuzz_Nosy_Registry_FetchService__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Registry_RegisterService__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Basic

func Fuzz_Nosy_Registry_StartAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Registry
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.StartAll(ctx)
	})
}

// skipping Fuzz_Nosy_Basic_Name__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Basic

// skipping Fuzz_Nosy_Basic_Start__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Basic

// skipping Fuzz_Nosy_Dispatcher_Start__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Dispatcher
