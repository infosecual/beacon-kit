package types

import (
	"testing"

	"github.com/karalabe/ssz"
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

func Fuzz_Nosy_BlobSidecar_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if b == nil || codec == nil {
			return
		}

		b.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_BlobSidecar_GetBeaconBlockHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBeaconBlockHeader()
	})
}

func Fuzz_Nosy_BlobSidecar_GetBlob__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBlob()
	})
}

func Fuzz_Nosy_BlobSidecar_GetKzgCommitment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetKzgCommitment()
	})
}

func Fuzz_Nosy_BlobSidecar_GetKzgProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetKzgProof()
	})
}

func Fuzz_Nosy_BlobSidecar_HasValidInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var kzgOffset uint64
		fill_err = tp.Fill(&kzgOffset)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.HasValidInclusionProof(kzgOffset)
	})
}

func Fuzz_Nosy_BlobSidecar_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.HashTreeRoot()
	})
}

func Fuzz_Nosy_BlobSidecar_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalSSZ()
	})
}

func Fuzz_Nosy_BlobSidecar_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_BlobSidecar_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SizeSSZ()
	})
}

func Fuzz_Nosy_BlobSidecar_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BlobSidecar
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_BlobSidecars_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if bs == nil || codec == nil {
			return
		}

		bs.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_BlobSidecars_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Empty()
	})
}

func Fuzz_Nosy_BlobSidecars_Get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var index int
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Get(index)
	})
}

func Fuzz_Nosy_BlobSidecars_GetSidecars__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetSidecars()
	})
}

func Fuzz_Nosy_BlobSidecars_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.IsNil()
	})
}

func Fuzz_Nosy_BlobSidecars_Len__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Len()
	})
}

func Fuzz_Nosy_BlobSidecars_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.MarshalSSZ()
	})
}

func Fuzz_Nosy_BlobSidecars_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_BlobSidecars_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var fixed bool
		fill_err = tp.Fill(&fixed)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.SizeSSZ(fixed)
	})
}

func Fuzz_Nosy_BlobSidecars_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_BlobSidecars_ValidateBlockRoots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.ValidateBlockRoots()
	})
}

func Fuzz_Nosy_BlobSidecars_VerifyInclusionProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BlobSidecars
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var kzgOffset uint64
		fill_err = tp.Fill(&kzgOffset)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.VerifyInclusionProofs(kzgOffset)
	})
}
