package builder

import (
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

// skipping Fuzz_Nosy_NodeBuilder[NodeT types.Node, LoggerT interface{log.AdvancedLogger[LoggerT]; log.Configurable[LoggerT, LoggerConfigT]}, LoggerConfigT any]_Build__ because parameters include func, chan, or unsupported interface: LoggerT

// skipping Fuzz_Nosy_DefaultServiceOptions__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/config.AppOptions
