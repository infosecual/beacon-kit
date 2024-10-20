package noop

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/da/pkg/kzg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
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

func Fuzz_Nosy_Verifier_VerifyBlobProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *eip4844.Blob
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		var _x2 eip4844.KZGProof
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 eip4844.KZGCommitment
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		v := NewVerifier()
		v.VerifyBlobProof(_x1, _x2, _x3)
	})
}

func Fuzz_Nosy_Verifier_VerifyBlobProofBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *types.BlobProofArgs
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		v := NewVerifier()
		v.VerifyBlobProofBatch(_x1)
	})
}
