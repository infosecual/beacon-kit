package mocks

import (
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

func Fuzz_Nosy_DB_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DB
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Delete(key)
	})
}

func Fuzz_Nosy_DB_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DB
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

func Fuzz_Nosy_DB_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DB
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Get(key)
	})
}

func Fuzz_Nosy_DB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DB
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Has(key)
	})
}

func Fuzz_Nosy_DB_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DB
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var value []byte
		fill_err = tp.Fill(&value)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Set(key, value)
	})
}

// skipping Fuzz_Nosy_DB_Delete_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DB_Delete_Call_Run__ because parameters include func, chan, or unsupported interface: func(key []byte)

// skipping Fuzz_Nosy_DB_Delete_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) error

// skipping Fuzz_Nosy_DB_Expecter_Delete__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DB_Expecter_Get__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DB_Expecter_Has__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DB_Expecter_Set__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_DB_Get_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DB_Get_Call_Run__ because parameters include func, chan, or unsupported interface: func(key []byte)

// skipping Fuzz_Nosy_DB_Get_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) ([]byte, error)

// skipping Fuzz_Nosy_DB_Has_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DB_Has_Call_Run__ because parameters include func, chan, or unsupported interface: func(key []byte)

// skipping Fuzz_Nosy_DB_Has_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte) (bool, error)

// skipping Fuzz_Nosy_DB_Set_Call_Return__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_DB_Set_Call_Run__ because parameters include func, chan, or unsupported interface: func(key []byte, value []byte)

// skipping Fuzz_Nosy_DB_Set_Call_RunAndReturn__ because parameters include func, chan, or unsupported interface: func([]byte, []byte) error
