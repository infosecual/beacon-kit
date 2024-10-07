package types

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
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

func Fuzz_Nosy_AttestationData_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if a == nil || codec == nil {
			return
		}

		a.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_AttestationData_GetBeaconBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.GetBeaconBlockRoot()
	})
}

func Fuzz_Nosy_AttestationData_GetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.GetIndex()
	})
}

func Fuzz_Nosy_AttestationData_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.GetSlot()
	})
}

func Fuzz_Nosy_AttestationData_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.GetTree()
	})
}

func Fuzz_Nosy_AttestationData_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_AttestationData_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_AttestationData_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.MarshalSSZ()
	})
}

func Fuzz_Nosy_AttestationData_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_AttestationData_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var slot math.U64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var index math.U64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var beaconBlockRoot common.Root
		fill_err = tp.Fill(&beaconBlockRoot)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.New(slot, index, beaconBlockRoot)
	})
}

func Fuzz_Nosy_AttestationData_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *AttestationData
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_AttestationData_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var a *AttestationData
		fill_err = tp.Fill(&a)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if a == nil {
			return
		}

		a.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_BeaconBlock_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
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

func Fuzz_Nosy_BeaconBlock_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *BeaconBlock
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Empty()
	})
}

func Fuzz_Nosy_BeaconBlock_GetBody__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBody()
	})
}

func Fuzz_Nosy_BeaconBlock_GetExecutionNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetExecutionNumber()
	})
}

func Fuzz_Nosy_BeaconBlock_GetHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetHeader()
	})
}

func Fuzz_Nosy_BeaconBlock_GetParentBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetParentBlockRoot()
	})
}

func Fuzz_Nosy_BeaconBlock_GetProposerIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetProposerIndex()
	})
}

func Fuzz_Nosy_BeaconBlock_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetSlot()
	})
}

func Fuzz_Nosy_BeaconBlock_GetStateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetStateRoot()
	})
}

func Fuzz_Nosy_BeaconBlock_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetTree()
	})
}

func Fuzz_Nosy_BeaconBlock_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
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

// skipping Fuzz_Nosy_BeaconBlock_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_BeaconBlock_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.IsNil()
	})
}

func Fuzz_Nosy_BeaconBlock_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
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

func Fuzz_Nosy_BeaconBlock_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_BeaconBlock_NewFromSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var bz []byte
		fill_err = tp.Fill(&bz)
		if fill_err != nil {
			return
		}
		var forkVersion uint32
		fill_err = tp.Fill(&forkVersion)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.NewFromSSZ(bz, forkVersion)
	})
}

func Fuzz_Nosy_BeaconBlock_NewWithVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var forkVersion uint32
		fill_err = tp.Fill(&forkVersion)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.NewWithVersion(slot, proposerIndex, parentBlockRoot, forkVersion)
	})
}

func Fuzz_Nosy_BeaconBlock_SetStateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var root common.Root
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetStateRoot(root)
	})
}

func Fuzz_Nosy_BeaconBlock_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var fixed bool
		fill_err = tp.Fill(&fixed)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SizeSSZ(fixed)
	})
}

func Fuzz_Nosy_BeaconBlock_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
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

func Fuzz_Nosy_BeaconBlock_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Version()
	})
}

func Fuzz_Nosy_BeaconBlockBody_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
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

func Fuzz_Nosy_BeaconBlockBody_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var forkVersion uint32
		fill_err = tp.Fill(&forkVersion)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Empty(forkVersion)
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetAttestations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetAttestations()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetBlobKzgCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetBlobKzgCommitments()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetDeposits()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetEth1Data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetEth1Data()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetExecutionPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetExecutionPayload()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetGraffiti__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetGraffiti()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetRandaoReveal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetRandaoReveal()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetSlashingInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetSlashingInfo()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetTopLevelRoots__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetTopLevelRoots()
	})
}

func Fuzz_Nosy_BeaconBlockBody_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.GetTree()
	})
}

func Fuzz_Nosy_BeaconBlockBody_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
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

// skipping Fuzz_Nosy_BeaconBlockBody_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_BeaconBlockBody_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.IsNil()
	})
}

func Fuzz_Nosy_BeaconBlockBody_Length__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.Length()
	})
}

func Fuzz_Nosy_BeaconBlockBody_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
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

func Fuzz_Nosy_BeaconBlockBody_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetAttestations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var _x2 []*AttestationData
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetAttestations(_x2)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetBlobKzgCommitments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var commitments eip4844.KZGCommitments[common.ExecutionHash]
		fill_err = tp.Fill(&commitments)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetBlobKzgCommitments(commitments)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetDeposits__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var deposits []*Deposit
		fill_err = tp.Fill(&deposits)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetDeposits(deposits)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetEth1Data__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var eth1Data *Eth1Data
		fill_err = tp.Fill(&eth1Data)
		if fill_err != nil {
			return
		}
		if b == nil || eth1Data == nil {
			return
		}

		b.SetEth1Data(eth1Data)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetExecutionPayload__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var executionData *ExecutionPayload
		fill_err = tp.Fill(&executionData)
		if fill_err != nil {
			return
		}
		if b == nil || executionData == nil {
			return
		}

		b.SetExecutionPayload(executionData)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetGraffiti__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var graffiti common.Bytes32
		fill_err = tp.Fill(&graffiti)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetGraffiti(graffiti)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetRandaoReveal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var reveal crypto.BLSSignature
		fill_err = tp.Fill(&reveal)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetRandaoReveal(reveal)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SetSlashingInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var _x2 []*SlashingInfo
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SetSlashingInfo(_x2)
	})
}

func Fuzz_Nosy_BeaconBlockBody_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var fixed bool
		fill_err = tp.Fill(&fixed)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.SizeSSZ(fixed)
	})
}

func Fuzz_Nosy_BeaconBlockBody_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *BeaconBlockBody
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

func Fuzz_Nosy_BeaconBlockHeader_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		_x1 := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		_x1.Empty()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetBodyRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetBodyRoot()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetParentBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetParentBlockRoot()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetProposerIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetProposerIndex()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetSlot()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetStateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetStateRoot()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.GetTree()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_BeaconBlockHeader_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_BeaconBlockHeader_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.MarshalSSZ()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 math.Slot
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var p2 math.ValidatorIndex
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var p3 common.Root
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		var s4 common.Root
		fill_err = tp.Fill(&s4)
		if fill_err != nil {
			return
		}
		var b5 common.Root
		fill_err = tp.Fill(&b5)
		if fill_err != nil {
			return
		}
		var s6 math.Slot
		fill_err = tp.Fill(&s6)
		if fill_err != nil {
			return
		}
		var p7 math.ValidatorIndex
		fill_err = tp.Fill(&p7)
		if fill_err != nil {
			return
		}
		var p8 common.Root
		fill_err = tp.Fill(&p8)
		if fill_err != nil {
			return
		}
		var s9 common.Root
		fill_err = tp.Fill(&s9)
		if fill_err != nil {
			return
		}
		var b10 common.Root
		fill_err = tp.Fill(&b10)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(s1, p2, p3, s4, b5)
		b.New(s6, p7, p8, s9, b10)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SetBodyRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var b5 common.Root
		fill_err = tp.Fill(&b5)
		if fill_err != nil {
			return
		}
		var b6 common.Root
		fill_err = tp.Fill(&b6)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, b5)
		b.SetBodyRoot(b6)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SetParentBlockRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var p3 common.Root
		fill_err = tp.Fill(&p3)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var p6 common.Root
		fill_err = tp.Fill(&p6)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, p3, stateRoot, bodyRoot)
		b.SetParentBlockRoot(p6)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SetProposerIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var p2 math.ValidatorIndex
		fill_err = tp.Fill(&p2)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var p6 math.ValidatorIndex
		fill_err = tp.Fill(&p6)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, p2, parentBlockRoot, stateRoot, bodyRoot)
		b.SetProposerIndex(p6)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s1 math.Slot
		fill_err = tp.Fill(&s1)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var s6 math.Slot
		fill_err = tp.Fill(&s6)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(s1, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.SetSlot(s6)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SetStateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var s4 common.Root
		fill_err = tp.Fill(&s4)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var s6 common.Root
		fill_err = tp.Fill(&s6)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, s4, bodyRoot)
		b.SetStateRoot(s6)
	})
}

func Fuzz_Nosy_BeaconBlockHeader_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.SizeSSZ()
	})
}

func Fuzz_Nosy_BeaconBlockHeader_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var proposerIndex math.ValidatorIndex
		fill_err = tp.Fill(&proposerIndex)
		if fill_err != nil {
			return
		}
		var parentBlockRoot common.Root
		fill_err = tp.Fill(&parentBlockRoot)
		if fill_err != nil {
			return
		}
		var stateRoot common.Root
		fill_err = tp.Fill(&stateRoot)
		if fill_err != nil {
			return
		}
		var bodyRoot common.Root
		fill_err = tp.Fill(&bodyRoot)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		b := NewBeaconBlockHeader(slot, proposerIndex, parentBlockRoot, stateRoot, bodyRoot)
		b.UnmarshalSSZ(buf)
	})
}

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_New__ because parameters include func, chan, or unsupported interface: ForkT

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_DefineSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_GetTree__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_HashTreeRoot__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_MarshalSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_MarshalSSZTo__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_SizeSSZ__ as it appears to be an interface:

// skipping Fuzz_Nosy_BeaconState[BeaconBlockHeaderT constraints.StaticSSZField[BeaconBlockHeaderT, B], Eth1DataT constraints.StaticSSZField[Eth1DataT, E], ExecutionPayloadHeaderT constraints.DynamicSSZField[ExecutionPayloadHeaderT, P], ForkT constraints.StaticSSZField[ForkT, F], ValidatorT constraints.StaticSSZField[ValidatorT, V], B, E, P, F, V any]_UnmarshalSSZ__ as it appears to be an interface:

func Fuzz_Nosy_Deposit_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var c *ssz.Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.DefineSSZ(c)
	})
}

func Fuzz_Nosy_Deposit_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.Empty()
	})
}

func Fuzz_Nosy_Deposit_GetAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetAmount()
	})
}

func Fuzz_Nosy_Deposit_GetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetIndex()
	})
}

func Fuzz_Nosy_Deposit_GetPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetPubkey()
	})
}

func Fuzz_Nosy_Deposit_GetSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetSignature()
	})
}

func Fuzz_Nosy_Deposit_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetTree()
	})
}

func Fuzz_Nosy_Deposit_GetWithdrawalCredentials__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.GetWithdrawalCredentials()
	})
}

func Fuzz_Nosy_Deposit_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_Deposit_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_Deposit_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.MarshalSSZ()
	})
}

func Fuzz_Nosy_Deposit_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_Deposit_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 crypto.BLSPubkey
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var c2 WithdrawalCredentials
		fill_err = tp.Fill(&c2)
		if fill_err != nil {
			return
		}
		var a3 math.Gwei
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var s4 crypto.BLSSignature
		fill_err = tp.Fill(&s4)
		if fill_err != nil {
			return
		}
		var i5 uint64
		fill_err = tp.Fill(&i5)
		if fill_err != nil {
			return
		}
		var p6 crypto.BLSPubkey
		fill_err = tp.Fill(&p6)
		if fill_err != nil {
			return
		}
		var c7 WithdrawalCredentials
		fill_err = tp.Fill(&c7)
		if fill_err != nil {
			return
		}
		var a8 math.Gwei
		fill_err = tp.Fill(&a8)
		if fill_err != nil {
			return
		}
		var s9 crypto.BLSSignature
		fill_err = tp.Fill(&s9)
		if fill_err != nil {
			return
		}
		var i10 uint64
		fill_err = tp.Fill(&i10)
		if fill_err != nil {
			return
		}

		d := NewDeposit(p1, c2, a3, s4, i5)
		d.New(p6, c7, a8, s9, i10)
	})
}

func Fuzz_Nosy_Deposit_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.SizeSSZ()
	})
}

func Fuzz_Nosy_Deposit_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}
		var index uint64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		d := NewDeposit(pubkey, credentials, amount, signature, index)
		d.UnmarshalSSZ(buf)
	})
}

// skipping Fuzz_Nosy_Deposit_VerifySignature__ because parameters include func, chan, or unsupported interface: func(pubkey github.com/berachain/beacon-kit/mod/primitives/pkg/crypto.BLSPubkey, message []byte, signature github.com/berachain/beacon-kit/mod/primitives/pkg/crypto.BLSSignature) error

func Fuzz_Nosy_DepositMessage_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if dm == nil || codec == nil {
			return
		}

		dm.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_DepositMessage_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		if dm == nil {
			return
		}

		dm.HashTreeRoot()
	})
}

func Fuzz_Nosy_DepositMessage_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		if dm == nil {
			return
		}

		dm.MarshalSSZ()
	})
}

func Fuzz_Nosy_DepositMessage_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if dm == nil {
			return
		}

		dm.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_DepositMessage_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var credentials WithdrawalCredentials
		fill_err = tp.Fill(&credentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		if dm == nil {
			return
		}

		dm.New(pubkey, credentials, amount)
	})
}

func Fuzz_Nosy_DepositMessage_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *DepositMessage
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_DepositMessage_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var dm *DepositMessage
		fill_err = tp.Fill(&dm)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if dm == nil {
			return
		}

		dm.UnmarshalSSZ(buf)
	})
}

// skipping Fuzz_Nosy_DepositMessage_VerifyCreateValidator__ because parameters include func, chan, or unsupported interface: func(pubkey github.com/berachain/beacon-kit/mod/primitives/pkg/crypto.BLSPubkey, message []byte, signature github.com/berachain/beacon-kit/mod/primitives/pkg/crypto.BLSSignature) error

func Fuzz_Nosy_Eth1Data_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if e == nil || codec == nil {
			return
		}

		e.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_Eth1Data_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Eth1Data
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Empty()
	})
}

func Fuzz_Nosy_Eth1Data_GetDepositCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GetDepositCount()
	})
}

func Fuzz_Nosy_Eth1Data_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.GetTree()
	})
}

func Fuzz_Nosy_Eth1Data_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_Eth1Data_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_Eth1Data_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.MarshalSSZ()
	})
}

func Fuzz_Nosy_Eth1Data_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_Eth1Data_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var depositRoot common.Root
		fill_err = tp.Fill(&depositRoot)
		if fill_err != nil {
			return
		}
		var depositCount math.U64
		fill_err = tp.Fill(&depositCount)
		if fill_err != nil {
			return
		}
		var blockHash common.ExecutionHash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.New(depositRoot, depositCount, blockHash)
	})
}

func Fuzz_Nosy_Eth1Data_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *Eth1Data
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_Eth1Data_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *Eth1Data
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_ExecutionPayload_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if p == nil || codec == nil {
			return
		}

		p.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_ExecutionPayload_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 uint32
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Empty(_x2)
	})
}

func Fuzz_Nosy_ExecutionPayload_GetBaseFeePerGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetBaseFeePerGas()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetBlobGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetBlobGasUsed()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetBlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetBlockHash()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetExcessBlobGas__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetExcessBlobGas()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetExtraData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetExtraData()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetFeeRecipient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetFeeRecipient()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetGasLimit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetGasLimit()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetGasUsed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetGasUsed()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetLogsBloom__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetLogsBloom()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetNumber()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetParentHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetParentHash()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetPrevRandao__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetPrevRandao()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetReceiptsRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetReceiptsRoot()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetStateRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetStateRoot()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetTimestamp()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetTransactions__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetTransactions()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetTree()
	})
}

func Fuzz_Nosy_ExecutionPayload_GetWithdrawals__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.GetWithdrawals()
	})
}

func Fuzz_Nosy_ExecutionPayload_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_ExecutionPayload_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_ExecutionPayload_IsBlinded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.IsBlinded()
	})
}

func Fuzz_Nosy_ExecutionPayload_IsNil__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.IsNil()
	})
}

func Fuzz_Nosy_ExecutionPayload_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarshalJSON()
	})
}

func Fuzz_Nosy_ExecutionPayload_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarshalSSZ()
	})
}

func Fuzz_Nosy_ExecutionPayload_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_ExecutionPayload_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var fixed bool
		fill_err = tp.Fill(&fixed)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.SizeSSZ(fixed)
	})
}

func Fuzz_Nosy_ExecutionPayload_ToHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var _x2 uint64
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var eth1ChainID uint64
		fill_err = tp.Fill(&eth1ChainID)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.ToHeader(_x2, eth1ChainID)
	})
}

func Fuzz_Nosy_ExecutionPayload_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_ExecutionPayload_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		var bz []byte
		fill_err = tp.Fill(&bz)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.UnmarshalSSZ(bz)
	})
}

func Fuzz_Nosy_ExecutionPayload_Version__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p *ExecutionPayload
		fill_err = tp.Fill(&p)
		if fill_err != nil {
			return
		}
		if p == nil {
			return
		}

		p.Version()
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.DefineSSZ(codec)
	})
}

// skipping Fuzz_Nosy_ExecutionPayloadHeader_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_ExecutionPayloadHeader_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, dst []byte) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_NewFromJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bz []byte, _x2 uint32) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.NewFromJSON(bz, _x2)
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_NewFromSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bz []byte, _x2 uint32) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.NewFromSSZ(bz, _x2)
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fixed bool) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.SizeSSZ(fixed)
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input []byte) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_ExecutionPayloadHeader_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, bz []byte) {
		h, err := DefaultGenesisExecutionPayloadHeaderDeneb()
		if err != nil {
			return
		}
		h.UnmarshalSSZ(bz)
	})
}

func Fuzz_Nosy_Fork_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if f1 == nil || codec == nil {
			return
		}

		f1.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_Fork_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.Empty()
	})
}

func Fuzz_Nosy_Fork_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.GetTree()
	})
}

func Fuzz_Nosy_Fork_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_Fork_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_Fork_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.MarshalSSZ()
	})
}

func Fuzz_Nosy_Fork_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_Fork_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var previousVersion common.Version
		fill_err = tp.Fill(&previousVersion)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var epoch math.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.New(previousVersion, currentVersion, epoch)
	})
}

func Fuzz_Nosy_Fork_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.SizeSSZ()
	})
}

func Fuzz_Nosy_Fork_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var f1 *Fork
		fill_err = tp.Fill(&f1)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if f1 == nil {
			return
		}

		f1.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_ForkData_ComputeDomain__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		var domainType common.DomainType
		fill_err = tp.Fill(&domainType)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.ComputeDomain(domainType)
	})
}

func Fuzz_Nosy_ForkData_ComputeRandaoSigningRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		var domainType common.DomainType
		fill_err = tp.Fill(&domainType)
		if fill_err != nil {
			return
		}
		var epoch math.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.ComputeRandaoSigningRoot(domainType, epoch)
	})
}

func Fuzz_Nosy_ForkData_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_ForkData_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.HashTreeRoot()
	})
}

func Fuzz_Nosy_ForkData_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.MarshalSSZ()
	})
}

func Fuzz_Nosy_ForkData_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_ForkData_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 common.Version
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var g2 common.Root
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		var c3 common.Version
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var g4 common.Root
		fill_err = tp.Fill(&g4)
		if fill_err != nil {
			return
		}

		fd := NewForkData(c1, g2)
		fd.New(c3, g4)
	})
}

func Fuzz_Nosy_ForkData_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}

		_x1 := NewForkData(currentVersion, genesisValidatorsRoot)
		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_ForkData_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var currentVersion common.Version
		fill_err = tp.Fill(&currentVersion)
		if fill_err != nil {
			return
		}
		var genesisValidatorsRoot common.Root
		fill_err = tp.Fill(&genesisValidatorsRoot)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		fd := NewForkData(currentVersion, genesisValidatorsRoot)
		fd.UnmarshalSSZ(buf)
	})
}

// skipping Fuzz_Nosy_Genesis[DepositT any, ExecutionPayloadHeaderT interface{NewFromJSON([]byte, uint32) (ExecutionPayloadHeaderT, error)}]_GetDeposits__ as it appears to be an interface:

// skipping Fuzz_Nosy_Genesis[DepositT any, ExecutionPayloadHeaderT interface{NewFromJSON([]byte, uint32) (ExecutionPayloadHeaderT, error)}]_GetExecutionPayloadHeader__ as it appears to be an interface:

// skipping Fuzz_Nosy_Genesis[DepositT any, ExecutionPayloadHeaderT interface{NewFromJSON([]byte, uint32) (ExecutionPayloadHeaderT, error)}]_GetForkVersion__ as it appears to be an interface:

// skipping Fuzz_Nosy_Genesis[DepositT any, ExecutionPayloadHeaderT interface{NewFromJSON([]byte, uint32) (ExecutionPayloadHeaderT, error)}]_UnmarshalJSON__ as it appears to be an interface:

func Fuzz_Nosy_SigningData_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningData
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if s == nil || codec == nil {
			return
		}

		s.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_SigningData_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningData
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.HashTreeRoot()
	})
}

func Fuzz_Nosy_SigningData_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningData
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalSSZ()
	})
}

func Fuzz_Nosy_SigningData_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningData
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalSSZTo(buf)
	})
}

func Fuzz_Nosy_SigningData_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *SigningData
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_SigningData_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SigningData
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_SlashingInfo_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if s == nil || codec == nil {
			return
		}

		s.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_SlashingInfo_GetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetIndex()
	})
}

func Fuzz_Nosy_SlashingInfo_GetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetSlot()
	})
}

func Fuzz_Nosy_SlashingInfo_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.GetTree()
	})
}

func Fuzz_Nosy_SlashingInfo_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_SlashingInfo_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_SlashingInfo_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalSSZ()
	})
}

func Fuzz_Nosy_SlashingInfo_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_SlashingInfo_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var slot math.U64
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		var index math.U64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.New(slot, index)
	})
}

func Fuzz_Nosy_SlashingInfo_SetIndex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var index math.U64
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetIndex(index)
	})
}

func Fuzz_Nosy_SlashingInfo_SetSlot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var slot math.Slot
		fill_err = tp.Fill(&slot)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SetSlot(slot)
	})
}

func Fuzz_Nosy_SlashingInfo_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *SlashingInfo
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_SlashingInfo_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SlashingInfo
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_Validator_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var codec *ssz.Codec
		fill_err = tp.Fill(&codec)
		if fill_err != nil {
			return
		}
		if codec == nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.DefineSSZ(codec)
	})
}

func Fuzz_Nosy_Validator_Empty__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		_x1 := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		_x1.Empty()
	})
}

func Fuzz_Nosy_Validator_GetEffectiveBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.GetEffectiveBalance()
	})
}

func Fuzz_Nosy_Validator_GetPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.GetPubkey()
	})
}

func Fuzz_Nosy_Validator_GetTree__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.GetTree()
	})
}

func Fuzz_Nosy_Validator_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.HashTreeRoot()
	})
}

// skipping Fuzz_Nosy_Validator_HashTreeRootWith__ because parameters include func, chan, or unsupported interface: github.com/ferranbt/fastssz.HashWalker

func Fuzz_Nosy_Validator_MarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.MarshalSSZ()
	})
}

func Fuzz_Nosy_Validator_MarshalSSZTo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var dst []byte
		fill_err = tp.Fill(&dst)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.MarshalSSZTo(dst)
	})
}

func Fuzz_Nosy_Validator_New__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var p1 crypto.BLSPubkey
		fill_err = tp.Fill(&p1)
		if fill_err != nil {
			return
		}
		var w2 WithdrawalCredentials
		fill_err = tp.Fill(&w2)
		if fill_err != nil {
			return
		}
		var a3 math.Gwei
		fill_err = tp.Fill(&a3)
		if fill_err != nil {
			return
		}
		var e4 math.Gwei
		fill_err = tp.Fill(&e4)
		if fill_err != nil {
			return
		}
		var m5 math.Gwei
		fill_err = tp.Fill(&m5)
		if fill_err != nil {
			return
		}
		var p6 crypto.BLSPubkey
		fill_err = tp.Fill(&p6)
		if fill_err != nil {
			return
		}
		var w7 WithdrawalCredentials
		fill_err = tp.Fill(&w7)
		if fill_err != nil {
			return
		}
		var a8 math.Gwei
		fill_err = tp.Fill(&a8)
		if fill_err != nil {
			return
		}
		var e9 math.Gwei
		fill_err = tp.Fill(&e9)
		if fill_err != nil {
			return
		}
		var m10 math.Gwei
		fill_err = tp.Fill(&m10)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(p1, w2, a3, e4, m5)
		v.New(p6, w7, a8, e9, m10)
	})
}

func Fuzz_Nosy_Validator_SetEffectiveBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var balance math.Gwei
		fill_err = tp.Fill(&balance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.SetEffectiveBalance(balance)
	})
}

func Fuzz_Nosy_Validator_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		_x1 := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		_x1.SizeSSZ()
	})
}

func Fuzz_Nosy_Validator_UnmarshalSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var buf []byte
		fill_err = tp.Fill(&buf)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.UnmarshalSSZ(buf)
	})
}

func Fuzz_Nosy_WithdrawalCredentials_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}

		wc := NewCredentialsFromExecutionAddress(address)
		wc.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_WithdrawalCredentials_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}
		var text []byte
		fill_err = tp.Fill(&text)
		if fill_err != nil {
			return
		}

		wc := NewCredentialsFromExecutionAddress(address)
		wc.UnmarshalText(text)
	})
}

func Fuzz_Nosy_Deposits_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds Deposits
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var c *ssz.Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		ds.DefineSSZ(c)
	})
}

func Fuzz_Nosy_Deposits_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds Deposits
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}

		ds.HashTreeRoot()
	})
}

func Fuzz_Nosy_Deposits_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ds Deposits
		fill_err = tp.Fill(&ds)
		if fill_err != nil {
			return
		}
		var _x2 bool
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		ds.SizeSSZ(_x2)
	})
}

func Fuzz_Nosy_Validator_GetWithdrawableEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.GetWithdrawableEpoch()
	})
}

func Fuzz_Nosy_Validator_GetWithdrawalCredentials__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.GetWithdrawalCredentials()
	})
}

func Fuzz_Nosy_Validator_HasEth1WithdrawalCredentials__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.HasEth1WithdrawalCredentials()
	})
}

func Fuzz_Nosy_Validator_HasMaxEffectiveBalance__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var m5 math.Gwei
		fill_err = tp.Fill(&m5)
		if fill_err != nil {
			return
		}
		var m6 math.Gwei
		fill_err = tp.Fill(&m6)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, m5)
		v.HasMaxEffectiveBalance(m6)
	})
}

func Fuzz_Nosy_Validator_IsActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var epoch math.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.IsActive(epoch)
	})
}

func Fuzz_Nosy_Validator_IsEligibleForActivation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var finalizedEpoch math.Epoch
		fill_err = tp.Fill(&finalizedEpoch)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.IsEligibleForActivation(finalizedEpoch)
	})
}

func Fuzz_Nosy_Validator_IsEligibleForActivationQueue__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var m5 math.Gwei
		fill_err = tp.Fill(&m5)
		if fill_err != nil {
			return
		}
		var m6 math.Gwei
		fill_err = tp.Fill(&m6)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, m5)
		v.IsEligibleForActivationQueue(m6)
	})
}

func Fuzz_Nosy_Validator_IsFullyWithdrawable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var balance math.Gwei
		fill_err = tp.Fill(&balance)
		if fill_err != nil {
			return
		}
		var epoch math.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.IsFullyWithdrawable(balance, epoch)
	})
}

func Fuzz_Nosy_Validator_IsPartiallyWithdrawable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var m5 math.Gwei
		fill_err = tp.Fill(&m5)
		if fill_err != nil {
			return
		}
		var balance math.Gwei
		fill_err = tp.Fill(&balance)
		if fill_err != nil {
			return
		}
		var m7 math.Gwei
		fill_err = tp.Fill(&m7)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, m5)
		v.IsPartiallyWithdrawable(balance, m7)
	})
}

func Fuzz_Nosy_Validator_IsSlashable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}
		var epoch math.Epoch
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.IsSlashable(epoch)
	})
}

func Fuzz_Nosy_Validator_IsSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pubkey crypto.BLSPubkey
		fill_err = tp.Fill(&pubkey)
		if fill_err != nil {
			return
		}
		var withdrawalCredentials WithdrawalCredentials
		fill_err = tp.Fill(&withdrawalCredentials)
		if fill_err != nil {
			return
		}
		var amount math.Gwei
		fill_err = tp.Fill(&amount)
		if fill_err != nil {
			return
		}
		var effectiveBalanceIncrement math.Gwei
		fill_err = tp.Fill(&effectiveBalanceIncrement)
		if fill_err != nil {
			return
		}
		var maxEffectiveBalance math.Gwei
		fill_err = tp.Fill(&maxEffectiveBalance)
		if fill_err != nil {
			return
		}

		v := NewValidatorFromDeposit(pubkey, withdrawalCredentials, amount, effectiveBalanceIncrement, maxEffectiveBalance)
		v.IsSlashed()
	})
}

func Fuzz_Nosy_Validators_DefineSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs Validators
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}
		var c *ssz.Codec
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		vs.DefineSSZ(c)
	})
}

func Fuzz_Nosy_Validators_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs Validators
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}

		vs.HashTreeRoot()
	})
}

func Fuzz_Nosy_Validators_SizeSSZ__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vs Validators
		fill_err = tp.Fill(&vs)
		if fill_err != nil {
			return
		}
		var _x2 bool
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}

		vs.SizeSSZ(_x2)
	})
}

func Fuzz_Nosy_WithdrawalCredentials_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		wc := NewCredentialsFromExecutionAddress(address)
		wc.MarshalText()
	})
}

func Fuzz_Nosy_WithdrawalCredentials_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		wc := NewCredentialsFromExecutionAddress(address)
		wc.String()
	})
}

func Fuzz_Nosy_WithdrawalCredentials_ToExecutionAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var address common.ExecutionAddress
		fill_err = tp.Fill(&address)
		if fill_err != nil {
			return
		}

		wc := NewCredentialsFromExecutionAddress(address)
		wc.ToExecutionAddress()
	})
}

// skipping Fuzz_Nosy_BlockBodyKZGOffset__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/common.ChainSpec

// skipping Fuzz_Nosy_CreateAndSignDepositMessage__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/primitives/pkg/crypto.BLSSigner
