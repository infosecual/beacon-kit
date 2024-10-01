package common

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

func Fuzz_Nosy_ExecutionAddress_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		a := NewExecutionAddressFromHex(i1)
		a.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_ExecutionAddress_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		a := NewExecutionAddressFromHex(i1)
		a.UnmarshalText(i2)
	})
}

func Fuzz_Nosy_ExecutionAddress_checksumHex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		a := NewExecutionAddressFromHex(input)
		a.checksumHex()
	})
}

func Fuzz_Nosy_ExecutionHash_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		h := NewExecutionHashFromHex(i1)
		h.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_ExecutionHash_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		h := NewExecutionHashFromHex(i1)
		h.UnmarshalText(i2)
	})
}

func Fuzz_Nosy_Root_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		r, err := NewRootFromHex(i1)
		if err != nil {
			return
		}
		r.UnmarshalJSON(i2)
	})
}

func Fuzz_Nosy_Root_UnmarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, i1 string, i2 []byte) {
		r, err := NewRootFromHex(i1)
		if err != nil {
			return
		}
		r.UnmarshalText(i2)
	})
}

func Fuzz_Nosy_ExecutionAddress_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		a := NewExecutionAddressFromHex(input)
		a.Hex()
	})
}

func Fuzz_Nosy_ExecutionAddress_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		a := NewExecutionAddressFromHex(input)
		a.MarshalJSON()
	})
}

func Fuzz_Nosy_ExecutionAddress_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		a := NewExecutionAddressFromHex(input)
		a.MarshalText()
	})
}

func Fuzz_Nosy_ExecutionAddress_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		a := NewExecutionAddressFromHex(input)
		a.String()
	})
}

func Fuzz_Nosy_ExecutionHash_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		h := NewExecutionHashFromHex(input)
		h.Hex()
	})
}

func Fuzz_Nosy_ExecutionHash_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		h := NewExecutionHashFromHex(input)
		h.MarshalJSON()
	})
}

func Fuzz_Nosy_ExecutionHash_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		h := NewExecutionHashFromHex(input)
		h.MarshalText()
	})
}

func Fuzz_Nosy_ExecutionHash_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		h := NewExecutionHashFromHex(input)
		h.String()
	})
}

func Fuzz_Nosy_Root_Hex__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		r, err := NewRootFromHex(input)
		if err != nil {
			return
		}
		r.Hex()
	})
}

func Fuzz_Nosy_Root_MarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		r, err := NewRootFromHex(input)
		if err != nil {
			return
		}
		r.MarshalJSON()
	})
}

func Fuzz_Nosy_Root_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		r, err := NewRootFromHex(input)
		if err != nil {
			return
		}
		r.MarshalText()
	})
}

func Fuzz_Nosy_Root_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		r, err := NewRootFromHex(input)
		if err != nil {
			return
		}
		r.String()
	})
}
