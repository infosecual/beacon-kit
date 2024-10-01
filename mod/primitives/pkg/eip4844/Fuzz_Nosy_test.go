package eip4844

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

func Fuzz_Nosy_Blob_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b *Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if b == nil {
			return
		}

		b.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_KZGCommitment_UnmarshalJSON__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *KZGCommitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var input []byte
		fill_err = tp.Fill(&input)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.UnmarshalJSON(input)
	})
}

func Fuzz_Nosy_Blob_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var b Blob
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}

		b.MarshalText()
	})
}

func Fuzz_Nosy_KZGCommitment_HashTreeRoot__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c KZGCommitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.HashTreeRoot()
	})
}

func Fuzz_Nosy_KZGCommitment_MarshalText__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c KZGCommitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.MarshalText()
	})
}

func Fuzz_Nosy_KZGCommitment_ToHashChunks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c KZGCommitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ToHashChunks()
	})
}

func Fuzz_Nosy_KZGCommitment_ToVersionedHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c KZGCommitment
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.ToVersionedHash()
	})
}

// skipping Fuzz_Nosy_KZGCommitments[HashT ~[32]byte]_Leafify__ as it appears to be an interface:

// skipping Fuzz_Nosy_KZGCommitments[HashT ~[32]byte]_ToVersionedHashes__ as it appears to be an interface:
