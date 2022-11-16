package busi

import (
	"time"
)

type CirculatingSupply struct {
	StatDate                    time.Time `json:"stat_date"`
	Height                      uint64    `json:"-"`
	ValueHeight                 uint64    `json:"-"`
	CirculatingFil              float64   `json:"circulating_fil"`
	CirculatingFilIncrease      float64   `json:"circulating_fil_increase"`
	MinedFil                    float64   `json:"mined_fil"`
	MinedFilIncrease            float64   `json:"mined_fil_increase"`
	VestedFil                   float64   `json:"vested_fil"`
	VestedFilIncrease           float64   `json:"vested_fil_increase"`
	ReserveDisbursedFil         float64   `json:"reserve_disbursed_fil"`
	ReserveDisbursedFilIncrease float64   `json:"reserve_disbursed_fil_increase"`
	LockedFil                   float64   `json:"locked_fil"`
	LockedFilIncrease           float64   `json:"locked_fil_increase"`
	BurntFil                    float64   `json:"burnt_fil"`
	BurntFilIncrease            float64   `json:"burnt_fil_increase"`
	IsLatest                    bool      `json:"-"`
	CreateAt                    time.Time `json:"-"`
}

type EVMBlockHeader struct {
	Height           int64  `json:"height"`
	Version          int    `json:"version"`
	Number           int64  `json:"number"`
	Hash             string `json:"hash"`
	ParentHash       string `json:"parentHash"`
	Sha3Uncles       string `json:"sha3_uncles"`
	Miner            string `json:"miner"`
	StateRoot        string `json:"state_root"`
	TransactionsRoot string `json:"transactions_root"`
	ReceiptsRoot     string `json:"receipts_root"`
	Difficulty       int64  `json:"difficulty"`
	GasLimit         int64  `json:"gas_limit"`
	GasUsed          int64  `json:"gas_used"`
	Timestamp        int64  `json:"timestamp"`
	ExtraData        string `json:"extra_data"`
	MixHash          string `json:"mix_hash"`
	Nonce            string `json:"nonce"`
	BaseFeePerGas    string `json:"base_fee_per_gas"`
	Size             uint64 `json:"size"`
}

func (b *EVMBlockHeader) TableName() string {
	return "evm_block_header"
}

// Contract evm smart contract
type EVMContract struct {
	Height          int64  `json:"height"`
	Version         int    `json:"version"`
	Address         string `json:"address"`
	FilecoinAddress string `json:"filecoin_address"`
	Balance         string `json:"balance"`
	Nonce           uint64 `json:"nonce"`
	ByteCode        string `json:"byte_code"`
}

func (c *EVMContract) TableName() string {
	return "evm_contract"
}

// InternalTX contract internal transaction
type EVMInternalTX struct {
	Height     int64  `json:"height"`
	Version    int    `json:"version"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parent_hash"`
	Type       uint64 `json:"type"`
	From       string `json:"from"`
	To         string `json:"to"`
	Value      string `json:"value"`
}

func (i *EVMInternalTX) TableName() string {
	return "evm_internal_tx"
}

// Receipt evm transaction receipt
type EVMReceipt struct {
	Height            int64  `json:"height"`
	Version           int    `json:"version"`
	TransactionHash   string `json:"transaction_hash"`
	TransactionIndex  int64  `json:"transaction_index"`
	BlockHash         string `json:"block_hash"`
	BlockNumber       int64  `json:"block_number"`
	From              string `json:"from"`
	To                string `json:"to"`
	StateRoot         string `json:"state_root"`
	Status            int64  `json:"status"`
	ContractAddress   string `json:"contract_address"`
	CumulativeGasUsed int64  `json:"cumulative_gas_used"`
	GasUsed           int64  `json:"gas_used"`
	EffectiveGasPrice int64  `json:"effective_gas_price"`
	LogsBloom         string `json:"logs_bloom"`
	Logs              string `json:"logs"`
}

func (r *EVMReceipt) TableName() string {
	return "evm_receipt"
}

// Transaction evm transaction
type EVMTransaction struct {
	Height               int64  `json:"height"`
	Version              int    `json:"version"`
	Hash                 string `json:"hash"`
	ChainID              uint64 `json:"chain_id"`
	Nonce                uint64 `json:"nonce"`
	BlockHash            string `json:"block_hash"`
	BlockNumber          uint64 `json:"block_number"`
	TransactionIndex     uint64 `json:"transaction_index"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Value                string `json:"value"`
	Type                 uint64 `json:"type"`
	Input                string `json:"input"`
	Gas                  uint64 `json:"gas"`
	GasLimit             uint64 `json:"gas_limit"`
	MaxFeePerGas         string `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
	V                    string `json:"v"`
	R                    string `json:"r"`
	S                    string `json:"s"`
}

func (m *EVMTransaction) TableName() string {
	return "evm_transaction"
}
