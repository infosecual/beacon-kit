package mocks

import (
	"testing"

	math "github.com/berachain/beacon-kit/mod/primitives/pkg/math"
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

func Fuzz_Nosy_BeaconBlock_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconBlock
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

func Fuzz_Nosy_BeaconBlock_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BeaconBlock
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetSlot()
	})
}

func Fuzz_Nosy_BeaconBlock_Expecter_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _e *BeaconBlock_Expecter
		fill_err = tp.Fill(&_e)
		if fill_err != nil {
			return
		}
		if _e == nil {
			return
		}

		_e.GetSlot()
	})
}

func Fuzz_Nosy_BeaconBlock_GetSlot_Call_Return__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _c *BeaconBlock_GetSlot_Call
		fill_err = tp.Fill(&_c)
		if fill_err != nil {
			return
		}
		var _a0 math.U64
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

// skipping Fuzz_Nosy_BeaconBlock_GetSlot_Call_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BeaconBlock_GetSlot_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/math.U64

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT pruner.BeaconBlock]_Context__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT pruner.BeaconBlock]_Data__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT pruner.BeaconBlock]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT pruner.BeaconBlock]_ID__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent[BeaconBlockT pruner.BeaconBlock]_Is__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Context_Call[BeaconBlockT pruner.BeaconBlock]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Context_Call[BeaconBlockT pruner.BeaconBlock]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockEvent_Context_Call[BeaconBlockT pruner.BeaconBlock]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() context.Context

// skipping Fuzz_Nosy_BlockEvent_Data_Call[BeaconBlockT pruner.BeaconBlock]_Return__ because parameters include func, chan, or unsupported interface: BeaconBlockT

// skipping Fuzz_Nosy_BlockEvent_Data_Call[BeaconBlockT pruner.BeaconBlock]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockEvent_Data_Call[BeaconBlockT pruner.BeaconBlock]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() BeaconBlockT

// skipping Fuzz_Nosy_BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock]_Context__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock]_Data__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock]_ID__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Expecter[BeaconBlockT pruner.BeaconBlock]_Is__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BlockEvent_ID_Call[BeaconBlockT pruner.BeaconBlock]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_ID_Call[BeaconBlockT pruner.BeaconBlock]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_BlockEvent_ID_Call[BeaconBlockT pruner.BeaconBlock]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() github.com/berachain/beacon-kit/mod/primitives/pkg/async.EventID

// skipping Fuzz_Nosy_BlockEvent_Is_Call[BeaconBlockT pruner.BeaconBlock]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_BlockEvent_Is_Call[BeaconBlockT pruner.BeaconBlock]_Run__ because parameters include func, chan, or unsupported interface: func(_a0 github.com/berachain/beacon-kit/mod/primitives/pkg/async.EventID)

// skipping Fuzz_Nosy_BlockEvent_Is_Call[BeaconBlockT pruner.BeaconBlock]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(github.com/berachain/beacon-kit/mod/primitives/pkg/async.EventID) bool

func Fuzz_Nosy_Prunable_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Prunable
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

func Fuzz_Nosy_Prunable_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *Prunable
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var start uint64
		fill_err = tp.Fill(&start)
		if fill_err != nil {
			return
		}
		var end uint64
		fill_err = tp.Fill(&end)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Prune(start, end)
	})
}

// skipping Fuzz_Nosy_Prunable_Expecter_Prune__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Prunable_Prune_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_Prunable_Prune_Call_Run__ because parameters include func, chan, or unsupported interface: func(start uint64, end uint64)

// skipping Fuzz_Nosy_Prunable_Prune_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(uint64, uint64) error

// skipping Fuzz_Nosy_Pruner[PrunableT pruner.Prunable]_EXPECT__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner[PrunableT pruner.Prunable]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner[PrunableT pruner.Prunable]_Start__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner_Expecter[PrunableT pruner.Prunable]_Name__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner_Expecter[PrunableT pruner.Prunable]_Start__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_Pruner_Name_Call[PrunableT pruner.Prunable]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner_Name_Call[PrunableT pruner.Prunable]_Run__ because parameters include func, chan, or unsupported interface: func()

// skipping Fuzz_Nosy_Pruner_Name_Call[PrunableT pruner.Prunable]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func() string

// skipping Fuzz_Nosy_Pruner_Start_Call[PrunableT pruner.Prunable]_Return__ as it appears to be an interface:

// skipping Fuzz_Nosy_Pruner_Start_Call[PrunableT pruner.Prunable]_Run__ because parameters include func, chan, or unsupported interface: func(ctx context.Context)

// skipping Fuzz_Nosy_Pruner_Start_Call[PrunableT pruner.Prunable]_RunAndReturn__ because parameters include func, chan, or unsupported interface: func(context.Context)
