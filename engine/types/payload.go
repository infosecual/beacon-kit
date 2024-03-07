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

package enginetypes

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/holiman/uint256"
	"github.com/itsdevbear/bolaris/config/version"
	byteslib "github.com/itsdevbear/bolaris/lib/bytes"
	"github.com/itsdevbear/bolaris/math"
)

type ExecutionPayloadEnvelope interface {
	GetExecutionPayload() ExecutionPayload
	GetValue() math.Wei
	GetBlobsBundle() *BlobsBundleV1
	ShouldOverrideBuilder() bool
}

//go:generate go run github.com/fjl/gencodec -type ExecutionPayloadEnvelopeDeneb -field-override executionPayloadEnvelopeMarshaling -out payload_env.json.go
//nolint:lll
type ExecutionPayloadEnvelopeDeneb struct {
	ExecutionPayload *ExecutableDataDeneb `json:"executionPayload"      gencodec:"required"`
	BlockValue       *big.Int             `json:"blockValue"            gencodec:"required"`
	BlobsBundle      *BlobsBundleV1       `json:"blobsBundle"`
	Override         bool                 `json:"shouldOverrideBuilder"`
}

func (e *ExecutionPayloadEnvelopeDeneb) GetExecutionPayload() ExecutionPayload {
	return e.ExecutionPayload
}

func (e *ExecutionPayloadEnvelopeDeneb) GetValue() math.Wei {
	val, ok := uint256.FromBig(e.BlockValue)
	if !ok {
		return math.Wei{}
	}
	return math.Wei{Int: val}
}

func (e *ExecutionPayloadEnvelopeDeneb) GetBlobsBundle() *BlobsBundleV1 {
	return e.BlobsBundle
}

func (e *ExecutionPayloadEnvelopeDeneb) ShouldOverrideBuilder() bool {
	return e.Override
}

func (e *ExecutionPayloadEnvelopeDeneb) String() string {
	return fmt.Sprintf(`
ExecutionPayloadEnvelopeDeneb{
	ExecutionPayload: %s,
	BlockValue: %s,
	BlobsBundle: %s,
	Override: %v,
}`, e.ExecutionPayload.String(),
		e.BlockValue.String(),
		e.GetBlobsBundle().Blobs,
		e.Override,
	)
}

// JSON type overrides for ExecutionPayloadEnvelope.
type executionPayloadEnvelopeMarshaling struct {
	BlockValue *hexutil.Big
}

//go:generate go run github.com/fjl/gencodec -type ExecutableDataDeneb -field-override executableDataDenebMarshaling -out payload.json.go
//nolint:lll
type ExecutableDataDeneb struct {
	ParentHash    common.Hash    `json:"parentHash"    ssz-size:"32"  gencodec:"required"`
	FeeRecipient  common.Address `json:"feeRecipient"  ssz-size:"20"  gencodec:"required"`
	StateRoot     common.Hash    `json:"stateRoot"     ssz-size:"32"  gencodec:"required"`
	ReceiptsRoot  common.Hash    `json:"receiptsRoot"  ssz-size:"32"  gencodec:"required"`
	LogsBloom     []byte         `json:"logsBloom"     ssz-size:"256" gencodec:"required"`
	Random        common.Hash    `json:"prevRandao"    ssz-size:"32"  gencodec:"required"`
	Number        uint64         `json:"blockNumber"                  gencodec:"required"`
	GasLimit      uint64         `json:"gasLimit"                     gencodec:"required"`
	GasUsed       uint64         `json:"gasUsed"                      gencodec:"required"`
	Timestamp     uint64         `json:"timestamp"                    gencodec:"required"`
	ExtraData     []byte         `json:"extraData"                    gencodec:"required" ssz-max:"32"`
	BaseFeePerGas []byte         `json:"baseFeePerGas" ssz-size:"32"  gencodec:"required"`
	BlockHash     common.Hash    `json:"blockHash"     ssz-size:"32"  gencodec:"required"`
	Transactions  [][]byte       `json:"transactions"  ssz-size:"?,?" gencodec:"required" ssz-max:"1048576,1073741824"`
	Withdrawals   []*Withdrawal  `json:"withdrawals"                                      ssz-max:"16"`
	BlobGasUsed   uint64         `json:"blobGasUsed"`
	ExcessBlobGas uint64         `json:"excessBlobGas"`
}

// Version returns the version of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) Version() int {
	return version.Deneb
}

// IsBlinded checks if the ExecutableDataDeneb is blinded.
func (d *ExecutableDataDeneb) IsBlinded() bool {
	return false
}

// GetParentHash returns the parent hash of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetParentHash() common.Hash {
	return d.ParentHash
}

// GetBlockHash returns the block hash of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetBlockHash() common.Hash {
	return d.BlockHash
}

// GetTransactions returns the transactions of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetTransactions() [][]byte {
	return d.Transactions
}

// GetWithdrawals returns the withdrawals of the ExecutableDataDeneb.
func (d *ExecutableDataDeneb) GetWithdrawals() []*Withdrawal {
	return d.Withdrawals
}

func (d *ExecutableDataDeneb) String() string {
	return fmt.Sprintf(
		"ExecutableDataDeneb{\n"+
			"\tParentHash: %s,\n"+
			"\tFeeRecipient: %s,\n"+
			"\tStateRoot: %s,\n"+
			"\tReceiptsRoot: %s,\n"+
			"\tLogsBloom: %x,\n"+
			"\tRandom: %s,\n"+
			"\tNumber: %d,\n"+
			"\tGasLimit: %d,\n"+
			"\tGasUsed: %d,\n"+
			"\tTimestamp: %d,\n"+
			"\tExtraData: %s,\n"+
			"\tBaseFeePerGas: %s,\n"+
			"\tBlockHash: %s,\n"+
			"\tTransactions: %x,\n"+
			"\tWithdrawals: %v,\n"+
			"\tBlobGasUsed: %d,\n"+
			"\tExcessBlobGas: %d,\n"+
			"}",
		d.ParentHash.String(),
		d.FeeRecipient.String(),
		d.StateRoot.String(),
		d.ReceiptsRoot.String(),
		d.LogsBloom,
		d.Random.String(),
		d.Number,
		d.GasLimit,
		d.GasUsed,
		d.Timestamp,
		d.ExtraData,
		big.NewInt(0).
			SetBytes(byteslib.CopyAndReverseEndianess(d.BaseFeePerGas)).
			String(),
		d.BlockHash.String(),
		d.Transactions,
		d.Withdrawals,
		d.BlobGasUsed,
		d.ExcessBlobGas,
	)
}