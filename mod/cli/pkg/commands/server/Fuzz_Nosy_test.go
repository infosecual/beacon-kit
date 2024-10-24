package server

import (
	"testing"

	"github.com/spf13/cobra"
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

func Fuzz_Nosy_addStartNodeFlags__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cmd *cobra.Command
		fill_err = tp.Fill(&cmd)
		if fill_err != nil {
			return
		}
		var opts StartCmdOptions[T]
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		if cmd == nil {
			return
		}

		addStartNodeFlags(cmd, opts)
	})
}