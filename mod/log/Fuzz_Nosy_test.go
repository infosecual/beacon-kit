package log

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

// skipping Fuzz_Nosy_AdvancedLogger[LoggerT any]_With__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.AdvancedLogger[LoggerT any]

func Fuzz_Nosy_Color_Raw__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c Color
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.Raw()
	})
}

// skipping Fuzz_Nosy_Colorable_AddKeyColor__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Colorable

// skipping Fuzz_Nosy_Colorable_AddKeyValColor__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Colorable

// skipping Fuzz_Nosy_Configurable[LoggerT, ConfigT any]_WithConfig__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Configurable[LoggerT, ConfigT any]

// skipping Fuzz_Nosy_Logger_Debug__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Logger

// skipping Fuzz_Nosy_Logger_Error__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Logger

// skipping Fuzz_Nosy_Logger_Info__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Logger

// skipping Fuzz_Nosy_Logger_Warn__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/log.Logger
