package mocks

import (
	context "context"
	"testing"

	service "github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry"
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

func Fuzz_Nosy_Basic_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Basic
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

func Fuzz_Nosy_Basic_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Basic
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Name()
	})
}

func Fuzz_Nosy_Basic_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Basic
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Start(ctx)
	})
}

func Fuzz_Nosy_Basic_Expecter_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *Basic_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.Name()
	})
}

// skipping Fuzz_Nosy_Basic_Expecter_Start__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Basic_Name_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *Basic_Name_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 string
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _c == nil {
			return
		}

		_c.Return(_a0)
	})
}

// skipping Fuzz_Nosy_Basic_Name_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Basic_Name_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() string

// skipping Fuzz_Nosy_Basic_Start_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Basic_Start_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_Basic_Start_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) error

func Fuzz_Nosy_Dispatcher_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Dispatcher
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

func Fuzz_Nosy_Dispatcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Dispatcher
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Start(ctx)
	})
}

// skipping Fuzz_Nosy_Dispatcher_Expecter_Start__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Dispatcher_Start_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Dispatcher_Start_Call_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_Dispatcher_Start_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context) error

func Fuzz_Nosy_RegistryOption_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *RegistryOption
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

func Fuzz_Nosy_RegistryOption_Execute__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *RegistryOption
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var _a0 *service.Registry
		fill_err = tp.Fill(&_a0)
		if fill_err != nil {
			return
		}
		if _m == nil || _a0 == nil {
			return
		}

		_m.Execute(_a0)
	})
}

// skipping Fuzz_Nosy_RegistryOption_Execute_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_RegistryOption_Execute_Call_Run__ because parameters include func, chan, or unsupported interface: func(_a0 *github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Registry)

// skipping Fuzz_Nosy_RegistryOption_Execute_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(*github.com/berachain/beacon-kit/mod/node-core/pkg/services/registry.Registry) error

// skipping Fuzz_Nosy_RegistryOption_Expecter_Execute__ because parameters include func, chan, or unsupported interface: interface{}
