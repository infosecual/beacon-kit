package signer

import (
	"testing"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
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

func Fuzz_Nosy_LegacySigner_PublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyBz LegacyKey
		fill_err = tp.Fill(&keyBz)
		if fill_err != nil {
			return
		}

		b, err := NewLegacySigner(keyBz)
		if err != nil {
			return
		}
		b.PublicKey()
	})
}

func Fuzz_Nosy_LegacySigner_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyBz LegacyKey
		fill_err = tp.Fill(&keyBz)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}

		b, err := NewLegacySigner(keyBz)
		if err != nil {
			return
		}
		b.Sign(msg)
	})
}

func Fuzz_Nosy_BLSSigner_PublicKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyFilePath string, stateFilePath string) {
		f1 := NewBLSSigner(keyFilePath, stateFilePath)
		f1.PublicKey()
	})
}

func Fuzz_Nosy_BLSSigner_Sign__(f *testing.F) {
	f.Fuzz(func(t *testing.T, keyFilePath string, stateFilePath string, msg []byte) {
		f1 := NewBLSSigner(keyFilePath, stateFilePath)
		f1.Sign(msg)
	})
}

func Fuzz_Nosy_BLSSigner_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyFilePath string
		fill_err = tp.Fill(&keyFilePath)
		if fill_err != nil {
			return
		}
		var stateFilePath string
		fill_err = tp.Fill(&stateFilePath)
		if fill_err != nil {
			return
		}
		var pubKey crypto.BLSPubkey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}

		f1 := NewBLSSigner(keyFilePath, stateFilePath)
		f1.VerifySignature(pubKey, msg, signature)
	})
}

func Fuzz_Nosy_LegacySigner_VerifySignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var keyBz LegacyKey
		fill_err = tp.Fill(&keyBz)
		if fill_err != nil {
			return
		}
		var pubKey crypto.BLSPubkey
		fill_err = tp.Fill(&pubKey)
		if fill_err != nil {
			return
		}
		var msg []byte
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		var signature crypto.BLSSignature
		fill_err = tp.Fill(&signature)
		if fill_err != nil {
			return
		}

		_x1, err := NewLegacySigner(keyBz)
		if err != nil {
			return
		}
		_x1.VerifySignature(pubKey, msg, signature)
	})
}
