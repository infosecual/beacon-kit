package filedb

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
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Delete(key)
	})
}

func Fuzz_Nosy_DB_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Get(key)
	})
}

func Fuzz_Nosy_DB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Has(key)
	})
}

func Fuzz_Nosy_DB_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
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
		if db == nil {
			return
		}

		db.Set(key, value)
	})
}

func Fuzz_Nosy_DB_pathForKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *DB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.pathForKey(key)
	})
}

func Fuzz_Nosy_RangeDB_Delete__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Delete(index, key)
	})
}

func Fuzz_Nosy_RangeDB_DeleteRange__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var from uint64
		fill_err = tp.Fill(&from)
		if fill_err != nil {
			return
		}
		var to uint64
		fill_err = tp.Fill(&to)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.DeleteRange(from, to)
	})
}

func Fuzz_Nosy_RangeDB_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Get(index, key)
	})
}

func Fuzz_Nosy_RangeDB_Has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.Has(index, key)
	})
}

func Fuzz_Nosy_RangeDB_Prune__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
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
		if db == nil {
			return
		}

		db.Prune(start, end)
	})
}

func Fuzz_Nosy_RangeDB_Set__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
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
		if db == nil {
			return
		}

		db.Set(index, key, value)
	})
}

func Fuzz_Nosy_RangeDB_prefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var db *RangeDB
		fill_err = tp.Fill(&db)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if db == nil {
			return
		}

		db.prefix(index, key)
	})
}

func Fuzz_Nosy_ExtractIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, prefixedKey []byte) {
		ExtractIndex(prefixedKey)
	})
}
