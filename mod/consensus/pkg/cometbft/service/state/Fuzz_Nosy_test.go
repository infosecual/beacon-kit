package state

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

func Fuzz_Nosy_Manager_Close__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sm *Manager
		fill_err = tp.Fill(&sm)
		if fill_err != nil {
			return
		}
		if sm == nil {
			return
		}

		sm.Close()
	})
}

func Fuzz_Nosy_Manager_CommitMultiStore__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sm *Manager
		fill_err = tp.Fill(&sm)
		if fill_err != nil {
			return
		}
		if sm == nil {
			return
		}

		sm.CommitMultiStore()
	})
}

func Fuzz_Nosy_Manager_LoadLatestVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sm *Manager
		fill_err = tp.Fill(&sm)
		if fill_err != nil {
			return
		}
		if sm == nil {
			return
		}

		sm.LoadLatestVersion()
	})
}

func Fuzz_Nosy_Manager_LoadVersion__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sm *Manager
		fill_err = tp.Fill(&sm)
		if fill_err != nil {
			return
		}
		var version int64
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if sm == nil {
			return
		}

		sm.LoadVersion(version)
	})
}
