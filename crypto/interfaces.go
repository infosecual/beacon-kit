// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package crypto

import (
	bls12381 "github.com/berachain/beacon-kit/crypto/bls12-381"
)

// Signer defines an interface for cryptographic signing operations.
// It uses generic type parameters Signature and Pubkey, both of which are
// slices of bytes.
type Signer[Signature, Pubkey any] interface {
	// Sign takes a message as a slice of bytes and returns a signature as a
	// slice of bytes and an error.
	Sign(msg []byte) Signature
}

// NewBLS12381Signer creates a new BLS12-381 signer instance given a secret key.
func NewBLS12381Signer(
	secretKey [bls12381.SecretKeyLength]byte,
) (Signer[[bls12381.SignatureLength]byte, [bls12381.PubKeyLength]byte], error) {
	return bls12381.NewSigner(secretKey)
}