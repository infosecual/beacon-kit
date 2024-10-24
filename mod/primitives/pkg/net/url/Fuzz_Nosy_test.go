package url

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

func Fuzz_Nosy_ConnectionURL_IsHTTP__(f *testing.F) {
	f.Fuzz(func(t *testing.T, raw string) {
		d, err := NewFromRaw(raw)
		if err != nil {
			return
		}
		d.IsHTTP()
	})
}

func Fuzz_Nosy_ConnectionURL_IsHTTPS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, raw string) {
		d, err := NewFromRaw(raw)
		if err != nil {
			return
		}
		d.IsHTTPS()
	})
}

func Fuzz_Nosy_ConnectionURL_IsIPC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, raw string) {
		d, err := NewFromRaw(raw)
		if err != nil {
			return
		}
		d.IsIPC()
	})
}