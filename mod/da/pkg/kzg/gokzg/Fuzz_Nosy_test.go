package gokzg

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/da/pkg/kzg/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/eip4844"
	gokzg4844 "github.com/crate-crypto/go-kzg-4844"
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

func Fuzz_Nosy_Verifier_GetImplementation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ts *gokzg4844.JSONTrustedSetup
		fill_err = tp.Fill(&ts)
		if fill_err != nil {
			return
		}
		if ts == nil {
			return
		}

		v, err := NewVerifier(ts)
		if err != nil {
			return
		}
		v.GetImplementation()
	})
}

func Fuzz_Nosy_Verifier_VerifyBlobProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ts *gokzg4844.JSONTrustedSetup
		fill_err = tp.Fill(&ts)
		if fill_err != nil {
			return
		}
		var blob *eip4844.Blob
		fill_err = tp.Fill(&blob)
		if fill_err != nil {
			return
		}
		var proof eip4844.KZGProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var commitment eip4844.KZGCommitment
		fill_err = tp.Fill(&commitment)
		if fill_err != nil {
			return
		}
		if ts == nil || blob == nil {
			return
		}

		v, err := NewVerifier(ts)
		if err != nil {
			return
		}
		v.VerifyBlobProof(blob, proof, commitment)
	})
}

func Fuzz_Nosy_Verifier_VerifyBlobProofBatch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ts *gokzg4844.JSONTrustedSetup
		fill_err = tp.Fill(&ts)
		if fill_err != nil {
			return
		}
		var args *types.BlobProofArgs
		fill_err = tp.Fill(&args)
		if fill_err != nil {
			return
		}
		if ts == nil || args == nil {
			return
		}

		v, err := NewVerifier(ts)
		if err != nil {
			return
		}
		v.VerifyBlobProofBatch(args)
	})
}
