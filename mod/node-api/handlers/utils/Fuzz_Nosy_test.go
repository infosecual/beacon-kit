package utils

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

// skipping Fuzz_Nosy_BindAndValidate__ because parameters include func, chan, or unsupported interface: ContextT

func Fuzz_Nosy_IsExecutionNumberPrefix__(f *testing.F) {
	f.Fuzz(func(t *testing.T, executionID string) {
		IsExecutionNumberPrefix(executionID)
	})
}

// skipping Fuzz_Nosy_SlotFromBlockID__ because parameters include func, chan, or unsupported interface: StorageBackendT

// skipping Fuzz_Nosy_SlotFromExecutionID__ because parameters include func, chan, or unsupported interface: StorageBackendT

// skipping Fuzz_Nosy_SlotFromStateID__ because parameters include func, chan, or unsupported interface: StorageBackendT

func Fuzz_Nosy_slotFromStateID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		slotFromStateID(id)
	})
}
