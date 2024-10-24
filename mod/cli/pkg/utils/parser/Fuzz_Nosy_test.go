package parser

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

func Fuzz_Nosy_ConvertAmount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, amount string) {
		ConvertAmount(amount)
	})
}

func Fuzz_Nosy_ConvertPubkey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, pubkey string) {
		ConvertPubkey(pubkey)
	})
}

func Fuzz_Nosy_ConvertSignature__(f *testing.F) {
	f.Fuzz(func(t *testing.T, signature string) {
		ConvertSignature(signature)
	})
}

func Fuzz_Nosy_ConvertVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, version string) {
		ConvertVersion(version)
	})
}
