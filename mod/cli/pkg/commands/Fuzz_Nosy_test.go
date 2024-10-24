package commands

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

func Fuzz_Nosy_Root_Run__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var root *Root
		fill_err = tp.Fill(&root)
		if fill_err != nil {
			return
		}
		var defaultNodeHome string
		fill_err = tp.Fill(&defaultNodeHome)
		if fill_err != nil {
			return
		}
		if root == nil {
			return
		}

		root.Run(defaultNodeHome)
	})
}

// skipping Fuzz_Nosy_DefaultRootCommandSetup__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/cli/pkg/commands/server/types.AppCreator[T, LoggerT]
