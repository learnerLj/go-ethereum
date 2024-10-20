// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package t8ntool

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
)

var _ = (*executionResultMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (e ExecutionResult) MarshalJSON() ([]byte, error) {
	type ExecutionResult struct {
		StateRoot            common.Hash           `json:"stateRoot"`
		TxRoot               common.Hash           `json:"txRoot"`
		ReceiptRoot          common.Hash           `json:"receiptsRoot"`
		LogsHash             common.Hash           `json:"logsHash"`
		Bloom                types.Bloom           `json:"logsBloom"        gencodec:"required"`
		Receipts             types.Receipts        `json:"receipts"`
		Rejected             []*rejectedTx         `json:"rejected,omitempty"`
		Difficulty           *math.HexOrDecimal256 `json:"currentDifficulty" gencodec:"required"`
		GasUsed              math.HexOrDecimal64   `json:"gasUsed"`
		BaseFee              *math.HexOrDecimal256 `json:"currentBaseFee,omitempty"`
		WithdrawalsRoot      *common.Hash          `json:"withdrawalsRoot,omitempty"`
		CurrentExcessBlobGas *math.HexOrDecimal64  `json:"currentExcessBlobGas,omitempty"`
		CurrentBlobGasUsed   *math.HexOrDecimal64  `json:"blobGasUsed,omitempty"`
		RequestsHash         *common.Hash          `json:"requestsHash,omitempty"`
		Requests             []hexutil.Bytes       `json:"requests,omitempty"`
	}
	var enc ExecutionResult
	enc.StateRoot = e.StateRoot
	enc.TxRoot = e.TxRoot
	enc.ReceiptRoot = e.ReceiptRoot
	enc.LogsHash = e.LogsHash
	enc.Bloom = e.Bloom
	enc.Receipts = e.Receipts
	enc.Rejected = e.Rejected
	enc.Difficulty = e.Difficulty
	enc.GasUsed = e.GasUsed
	enc.BaseFee = e.BaseFee
	enc.WithdrawalsRoot = e.WithdrawalsRoot
	enc.CurrentExcessBlobGas = e.CurrentExcessBlobGas
	enc.CurrentBlobGasUsed = e.CurrentBlobGasUsed
	enc.RequestsHash = e.RequestsHash
	if e.Requests != nil {
		enc.Requests = make([]hexutil.Bytes, len(e.Requests))
		for k, v := range e.Requests {
			enc.Requests[k] = v
		}
	}
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (e *ExecutionResult) UnmarshalJSON(input []byte) error {
	type ExecutionResult struct {
		StateRoot            *common.Hash          `json:"stateRoot"`
		TxRoot               *common.Hash          `json:"txRoot"`
		ReceiptRoot          *common.Hash          `json:"receiptsRoot"`
		LogsHash             *common.Hash          `json:"logsHash"`
		Bloom                *types.Bloom          `json:"logsBloom"        gencodec:"required"`
		Receipts             *types.Receipts       `json:"receipts"`
		Rejected             []*rejectedTx         `json:"rejected,omitempty"`
		Difficulty           *math.HexOrDecimal256 `json:"currentDifficulty" gencodec:"required"`
		GasUsed              *math.HexOrDecimal64  `json:"gasUsed"`
		BaseFee              *math.HexOrDecimal256 `json:"currentBaseFee,omitempty"`
		WithdrawalsRoot      *common.Hash          `json:"withdrawalsRoot,omitempty"`
		CurrentExcessBlobGas *math.HexOrDecimal64  `json:"currentExcessBlobGas,omitempty"`
		CurrentBlobGasUsed   *math.HexOrDecimal64  `json:"blobGasUsed,omitempty"`
		RequestsHash         *common.Hash          `json:"requestsHash,omitempty"`
		Requests             []hexutil.Bytes       `json:"requests,omitempty"`
	}
	var dec ExecutionResult
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.StateRoot != nil {
		e.StateRoot = *dec.StateRoot
	}
	if dec.TxRoot != nil {
		e.TxRoot = *dec.TxRoot
	}
	if dec.ReceiptRoot != nil {
		e.ReceiptRoot = *dec.ReceiptRoot
	}
	if dec.LogsHash != nil {
		e.LogsHash = *dec.LogsHash
	}
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for ExecutionResult")
	}
	e.Bloom = *dec.Bloom
	if dec.Receipts != nil {
		e.Receipts = *dec.Receipts
	}
	if dec.Rejected != nil {
		e.Rejected = dec.Rejected
	}
	if dec.Difficulty == nil {
		return errors.New("missing required field 'currentDifficulty' for ExecutionResult")
	}
	e.Difficulty = dec.Difficulty
	if dec.GasUsed != nil {
		e.GasUsed = *dec.GasUsed
	}
	if dec.BaseFee != nil {
		e.BaseFee = dec.BaseFee
	}
	if dec.WithdrawalsRoot != nil {
		e.WithdrawalsRoot = dec.WithdrawalsRoot
	}
	if dec.CurrentExcessBlobGas != nil {
		e.CurrentExcessBlobGas = dec.CurrentExcessBlobGas
	}
	if dec.CurrentBlobGasUsed != nil {
		e.CurrentBlobGasUsed = dec.CurrentBlobGasUsed
	}
	if dec.RequestsHash != nil {
		e.RequestsHash = dec.RequestsHash
	}
	if dec.Requests != nil {
		e.Requests = make([][]byte, len(dec.Requests))
		for k, v := range dec.Requests {
			e.Requests[k] = v
		}
	}
	return nil
}