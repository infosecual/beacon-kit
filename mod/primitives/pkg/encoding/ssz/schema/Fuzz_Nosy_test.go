package schema

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

// skipping Fuzz_Nosy_Field[T any]_GetValue__ as it appears to be an interface:

// skipping Fuzz_Nosy_Field[T any]_GetName__ as it appears to be an interface:

func Fuzz_Nosy_ID_IsBasic__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsBasic()
	})
}

func Fuzz_Nosy_ID_IsComposite__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsComposite()
	})
}

func Fuzz_Nosy_ID_IsContainer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsContainer()
	})
}

func Fuzz_Nosy_ID_IsElements__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsElements()
	})
}

func Fuzz_Nosy_ID_IsEnumerable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsEnumerable()
	})
}

func Fuzz_Nosy_ID_IsList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 ID
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}

		t1.IsList()
	})
}

// skipping Fuzz_Nosy_MerkleizableSSZObject[RootT ~[32]byte]_HashTreeRoot__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.MerkleizableSSZObject[RootT ~[32]byte]

// skipping Fuzz_Nosy_MerkleizableSSZObject[RootT ~[32]byte]_MarshalSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.MerkleizableSSZObject[RootT ~[32]byte]

// skipping Fuzz_Nosy_MerkleizableSSZObject[RootT ~[32]byte]_SizeSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.MerkleizableSSZObject[RootT ~[32]byte]

// skipping Fuzz_Nosy_MinimalSSZObject_IsFixed__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.MinimalSSZObject

// skipping Fuzz_Nosy_MinimalSSZObject_Type__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.MinimalSSZObject

// skipping Fuzz_Nosy_SSZEnumerable[ElementT any]_Elements__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZEnumerable[ElementT any]

// skipping Fuzz_Nosy_SSZEnumerable[ElementT any]_N__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZEnumerable[ElementT any]

// skipping Fuzz_Nosy_SSZObject[T any]_ChunkCount__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZObject[T any]

// skipping Fuzz_Nosy_SSZObject[T any]_NewFromSSZ__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZObject[T any]

func Fuzz_Nosy_SSZType_ElementType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		_x1 := U32()
		_x1.ElementType(p)
	})
}

func Fuzz_Nosy_SSZType_ItemPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, p string) {
		_x1 := U32()
		_x1.ItemPosition(p)
	})
}

func Fuzz_Nosy_basic_ElementType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b basic
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		b.ElementType(_x2)
	})
}

func Fuzz_Nosy_basic_HashChunkCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b basic
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.HashChunkCount()
	})
}

func Fuzz_Nosy_basic_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b basic
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.ID()
	})
}

func Fuzz_Nosy_basic_ItemLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b basic
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.ItemLength()
	})
}

func Fuzz_Nosy_basic_ItemPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b basic
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		b.ItemPosition(_x2)
	})
}

func Fuzz_Nosy_container_ElementType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p string
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		c.ElementType(p)
	})
}

func Fuzz_Nosy_container_HashChunkCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.HashChunkCount()
	})
}

func Fuzz_Nosy_container_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ID()
	})
}

func Fuzz_Nosy_container_ItemLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ItemLength()
	})
}

func Fuzz_Nosy_container_ItemPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var p string
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		c.ItemPosition(p)
	})
}

func Fuzz_Nosy_container_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c container
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Length()
	})
}

func Fuzz_Nosy_list_ElementType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		l.ElementType(_x2)
	})
}

func Fuzz_Nosy_list_HashChunkCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.HashChunkCount()
	})
}

func Fuzz_Nosy_list_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.ID()
	})
}

func Fuzz_Nosy_list_ItemLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.ItemLength()
	})
}

func Fuzz_Nosy_list_ItemPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}
		var p string
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		l.ItemPosition(p)
	})
}

func Fuzz_Nosy_list_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var l list
		fill_err = tp.Fill(&l)
		if fill_err != nil {
			return
		}

		l.Length()
	})
}

func Fuzz_Nosy_vector_ElementType__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var _x2 string
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		v.ElementType(_x2)
	})
}

func Fuzz_Nosy_vector_HashChunkCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.HashChunkCount()
	})
}

func Fuzz_Nosy_vector_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ID()
	})
}

func Fuzz_Nosy_vector_ItemLength__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.ItemLength()
	})
}

func Fuzz_Nosy_vector_ItemPosition__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var p string
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}

		v.ItemPosition(p)
	})
}

func Fuzz_Nosy_vector_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v vector
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}

		v.Length()
	})
}
