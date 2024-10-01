package merkle

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

func Fuzz_Nosy_GeneralizedIndex_GetBranchIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.GetBranchIndices()
	})
}

func Fuzz_Nosy_GeneralizedIndex_GetPathIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.GetPathIndices()
	})
}

func Fuzz_Nosy_GeneralizedIndex_IndexBit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64, position int) {
		g := NewGeneralizedIndex(depth, index)
		g.IndexBit(position)
	})
}

func Fuzz_Nosy_GeneralizedIndex_LeftChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.LeftChild()
	})
}

func Fuzz_Nosy_GeneralizedIndex_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.Length()
	})
}

func Fuzz_Nosy_GeneralizedIndex_Parent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.Parent()
	})
}

func Fuzz_Nosy_GeneralizedIndex_RightChild__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.RightChild()
	})
}

func Fuzz_Nosy_GeneralizedIndex_Sibling__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.Sibling()
	})
}

func Fuzz_Nosy_GeneralizedIndex_Unwrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth uint8, index uint64) {
		g := NewGeneralizedIndex(depth, index)
		g.Unwrap()
	})
}

func Fuzz_Nosy_GeneralizedIndices_Concat__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gs GeneralizedIndices
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}

		gs.Concat()
	})
}

func Fuzz_Nosy_GeneralizedIndices_GetHelperIndices__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gs GeneralizedIndices
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}

		gs.GetHelperIndices()
	})
}

// skipping Fuzz_Nosy_ObjectPath[GeneralizedIndexT ~uint64, RootT ~[32]byte]_GetGeneralizedIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZType

// skipping Fuzz_Nosy_ObjectPath[GeneralizedIndexT ~uint64, RootT ~[32]byte]_Split__ as it appears to be an interface:

// skipping Fuzz_Nosy_BuildProofFromLeaves__ because parameters include func, chan, or unsupported interface: []RootT

// skipping Fuzz_Nosy_CalculateMultiRoot__ because parameters include func, chan, or unsupported interface: []RootT

// skipping Fuzz_Nosy_CalculateRoot__ because parameters include func, chan, or unsupported interface: RootT

func Fuzz_Nosy_GeneralizedIndexReverseComparator__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var i GeneralizedIndex
		fill_err = tp.Fill(&i)
		if fill_err != nil {
			return
		}
		var j GeneralizedIndex
		fill_err = tp.Fill(&j)
		if fill_err != nil {
			return
		}

		GeneralizedIndexReverseComparator(i, j)
	})
}

// skipping Fuzz_Nosy_VerifyMultiproof__ because parameters include func, chan, or unsupported interface: []RootT

// skipping Fuzz_Nosy_VerifyProof__ because parameters include func, chan, or unsupported interface: RootT

// skipping Fuzz_Nosy_buildSingleProofFromTree__ because parameters include func, chan, or unsupported interface: []RootT

// skipping Fuzz_Nosy_getBaseIndex__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/schema.SSZType

// skipping Fuzz_Nosy_hashRoot__ because parameters include func, chan, or unsupported interface: map[github.com/berachain/beacon-kit/mod/primitives/pkg/encoding/ssz/merkle.GeneralizedIndex]RootT

// skipping Fuzz_Nosy_newTree__ because parameters include func, chan, or unsupported interface: []RootT
