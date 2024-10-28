package version

import (
	"context"
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

func Fuzz_Nosy_ReportingService_Name__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _x1 *ReportingService
		fill_err = tp.Fill(&_x1)
		if fill_err != nil {
			return
		}
		if _x1 == nil {
			return
		}

		_x1.Name()
	})
}

func Fuzz_Nosy_ReportingService_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v *ReportingService
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if v == nil {
			return
		}

		v.Start(ctx)
	})
}

func Fuzz_Nosy_versionMetrics_reportVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var vm *versionMetrics
		fill_err = tp.Fill(&vm)
		if fill_err != nil {
			return
		}
		var v2 string
		fill_err = tp.Fill(&v2)
		if fill_err != nil {
			return
		}
		if vm == nil {
			return
		}

		vm.reportVersion(v2)
	})
}

// skipping Fuzz_Nosy_TelemetrySink_IncrementCounter__ because parameters include func, chan, or unsupported interface: github.com/berachain/beacon-kit/mod/node-core/pkg/services/version.TelemetrySink
