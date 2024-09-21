// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*accountMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (a account) MarshalJSON() ([]byte, error) {
	type account struct {
		Balance *hexutil.Big                `json:"balance,omitempty"`
		Energy  *hexutil.Big                `json:"energy,omitempty"`
		Code    hexutil.Bytes               `json:"code,omitempty"`
		Storage map[common.Hash]common.Hash `json:"storage,omitempty"`
	}
	var enc account
	enc.Balance = (*hexutil.Big)(a.Balance)
	enc.Energy = (*hexutil.Big)(a.Energy)
	enc.Code = a.Code
	enc.Storage = a.Storage
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (a *account) UnmarshalJSON(input []byte) error {
	type account struct {
		Balance *hexutil.Big                `json:"balance,omitempty"`
		Energy  *hexutil.Big                `json:"energy,omitempty"`
		Code    *hexutil.Bytes              `json:"code,omitempty"`
		Storage map[common.Hash]common.Hash `json:"storage,omitempty"`
	}
	var dec account
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Balance != nil {
		a.Balance = (*big.Int)(dec.Balance)
	}
	if dec.Energy != nil {
		a.Energy = (*big.Int)(dec.Energy)
	}
	if dec.Code != nil {
		a.Code = *dec.Code
	}
	if dec.Storage != nil {
		a.Storage = dec.Storage
	}
	return nil
}
