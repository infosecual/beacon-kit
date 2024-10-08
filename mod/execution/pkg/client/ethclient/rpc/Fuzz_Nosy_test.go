package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/berachain/beacon-kit/mod/primitives/pkg/net/jwt"
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

// skipping Fuzz_Nosy_Client_Call__ because parameters include func, chan, or unsupported interface: any

// skipping Fuzz_Nosy_Client_CallRaw__ because parameters include func, chan, or unsupported interface: []any

func Fuzz_Nosy_Client_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 *Client
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		if r1 == nil {
			return
		}

		r1.Close()
	})
}

func Fuzz_Nosy_Client_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 *Client
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if r1 == nil {
			return
		}

		r1.Start(ctx)
	})
}

func Fuzz_Nosy_Client_updateHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r1 *Client
		fill_err = tp.Fill(&r1)
		if fill_err != nil {
			return
		}
		if r1 == nil {
			return
		}

		r1.updateHeader()
	})
}

func Fuzz_Nosy_Error_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var err Error
		fill_err = tp.Fill(&err)
		if fill_err != nil {
			return
		}

		err.Error()
	})
}

func Fuzz_Nosy_WithJWTRefreshInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var interval time.Duration
		fill_err = tp.Fill(&interval)
		if fill_err != nil {
			return
		}

		WithJWTRefreshInterval(interval)
	})
}

func Fuzz_Nosy_WithJWTSecret__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var secret *jwt.Secret
		fill_err = tp.Fill(&secret)
		if fill_err != nil {
			return
		}
		if secret == nil {
			return
		}

		WithJWTSecret(secret)
	})
}
