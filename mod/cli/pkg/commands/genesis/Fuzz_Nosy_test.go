package genesis

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/genutil/types"
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

func Fuzz_Nosy_CollectValidatorJSONFiles__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var genTxsDir string
		fill_err = tp.Fill(&genTxsDir)
		if fill_err != nil {
			return
		}
		var g2 *types.AppGenesis
		fill_err = tp.Fill(&g2)
		if fill_err != nil {
			return
		}
		if g2 == nil {
			return
		}

		CollectValidatorJSONFiles(genTxsDir, g2)
	})
}

func Fuzz_Nosy_makeOutputFilepath__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rootDir string, pubkey string) {
		makeOutputFilepath(rootDir, pubkey)
	})
}
