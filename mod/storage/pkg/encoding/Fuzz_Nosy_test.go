package encoding

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

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_Encode__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_SetActiveForkVersion__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_Decode__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_DecodeJSON__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_EncodeJSON__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_Stringify__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZInterfaceCodec[T interface{NewFromSSZ([]byte, uint32) (T, error); constraints.SSZMarshallable; constraints.Versionable}]_ValueType__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_Decode__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_DecodeJSON__ as it appears to be an interface:

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_Encode__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_EncodeJSON__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_Stringify__ because parameters include func, chan, or unsupported interface: T

// skipping Fuzz_Nosy_SSZValueCodec[T interface{constraints.SSZMarshallable; constraints.Empty[T]}]_ValueType__ as it appears to be an interface:
