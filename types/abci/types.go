package types

import (
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/libs/common"
)
const (
	CodeTypeOK uint32 = 0
)
type ResponseBeginBlock struct {
	Tags                 []common.KVPair `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

type ResponseEndBlock struct {
	ValidatorUpdates      []ValidatorUpdate `protobuf:"bytes,1,rep,name=validator_updates,json=validatorUpdates" json:"validator_updates"`
	ConsensusParamUpdates *ConsensusParams  `protobuf:"bytes,2,opt,name=consensus_param_updates,json=consensusParamUpdates" json:"consensus_param_updates,omitempty"`
	Tags                  []common.KVPair   `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}          `json:"-"`
	XXX_unrecognized      []byte            `json:"-"`
	XXX_sizecache         int32             `json:"-"`
}

// ValidatorUpdate
type ValidatorUpdate struct {
	PubKey               PubKey   `protobuf:"bytes,1,opt,name=pub_key,json=pubKey" json:"pub_key"`
	Power                int64    `protobuf:"varint,2,opt,name=power,proto3" json:"power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type PubKey struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// ConsensusParams contains all consensus-relevant parameters
// that can be adjusted by the abci app
type ConsensusParams struct {
	Block                *BlockParams     `protobuf:"bytes,1,opt,name=block" json:"block,omitempty"`
	Evidence             *EvidenceParams  `protobuf:"bytes,2,opt,name=evidence" json:"evidence,omitempty"`
	Validator            *ValidatorParams `protobuf:"bytes,3,opt,name=validator" json:"validator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

// BlockParams contains limits on the block size and timestamp.
type BlockParams struct {
	// Note: must be greater than 0
	MaxBytes int64 `protobuf:"varint,1,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty"`
	// Note: must be greater or equal to -1
	MaxGas               int64    `protobuf:"varint,2,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// EvidenceParams contains limits on the evidence.
type EvidenceParams struct {
	// Note: must be greater than 0
	MaxAge               int64    `protobuf:"varint,1,opt,name=max_age,json=maxAge,proto3" json:"max_age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// ValidatorParams contains limits on validators.
type ValidatorParams struct {
	PubKeyTypes          []string `protobuf:"bytes,1,rep,name=pub_key_types,json=pubKeyTypes" json:"pub_key_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type ResponseCheckTx struct {
	Code                 uint32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log                  string          `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string          `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted            int64           `protobuf:"varint,5,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed              int64           `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Tags                 []common.KVPair `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Codespace            string          `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

// IsOK returns true if Code is OK.
func (r ResponseCheckTx) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseCheckTx) IsErr() bool {
	return r.Code != CodeTypeOK
}

type ResponseDeliverTx struct {
	Code                 uint32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Data                 []byte          `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Log                  string          `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string          `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	GasWanted            int64           `protobuf:"varint,5,opt,name=gas_wanted,json=gasWanted,proto3" json:"gas_wanted,omitempty"`
	GasUsed              int64           `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	Tags                 []common.KVPair `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	Codespace            string          `protobuf:"bytes,8,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}
// IsOK returns true if Code is OK.
func (r ResponseDeliverTx) IsOK() bool {
	return r.Code == CodeTypeOK
}

// IsErr returns true if Code is something other than OK.
func (r ResponseDeliverTx) IsErr() bool {
	return r.Code != CodeTypeOK
}

type ResponseQuery struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// bytes data = 2; // use "value" instead.
	Log                  string        `protobuf:"bytes,3,opt,name=log,proto3" json:"log,omitempty"`
	Info                 string        `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	Index                int64         `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	Key                  []byte        `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte        `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	Proof                *merkle.Proof `protobuf:"bytes,8,opt,name=proof" json:"proof,omitempty"`
	Height               int64         `protobuf:"varint,9,opt,name=height,proto3" json:"height,omitempty"`
	Codespace            string        `protobuf:"bytes,10,opt,name=codespace,proto3" json:"codespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

// IsOK returns true if Code is OK.
func (r ResponseQuery) IsOK() bool {
	return r.Code == CodeTypeOK
}