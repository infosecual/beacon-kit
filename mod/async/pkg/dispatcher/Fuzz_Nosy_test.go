package dispatcher

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

func Fuzz_Nosy_Dispatcher_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Dispatcher
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Name()
	})
}

// skipping Fuzz_Nosy_Dispatcher_Publish__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/async.BaseEvent

// skipping Fuzz_Nosy_Dispatcher_RegisterBrokers__ because parameters include func, chan, or unsupported interface: []github.com/berachain/beacon-kit/mod/async/pkg/types.Broker

func Fuzz_Nosy_Dispatcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *Dispatcher
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Start(ctx)
	})
}

// skipping Fuzz_Nosy_Dispatcher_Subscribe__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Dispatcher_Unsubscribe__ because parameters include func, chan, or unsupported interface: any