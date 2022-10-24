package outport

import (
	"time"

	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/esdt"
)

// AccountTokenData holds the data needed for indexing a token of an altered account
type AccountTokenData struct {
	IsNFTCreate bool           `json:"isNFTCreate"`
	Nonce       uint64         `json:"nonce"`
	Identifier  string         `json:"identifier"`
	Balance     string         `json:"balance"`
	Properties  string         `json:"properties"`
	MetaData    *esdt.MetaData `json:"metadata"`
}

// AlteredAccount holds the data needed of an altered account in a block
type AlteredAccount struct {
	IsSender      bool                `json:"isSender"`
	BalanceChange bool                `json:"balanceChange"`
	Nonce         uint64              `json:"nonce"`
	Address       string              `json:"address"`
	Balance       string              `json:"balance,omitempty"`
	Tokens        []*AccountTokenData `json:"tokens"`
}

// ArgsSaveBlockData will contain all information that are needed to save block data
type ArgsSaveBlockData struct {
	HeaderHash             []byte
	Body                   data.BodyHandler
	Header                 data.HeaderHandler
	SignersIndexes         []uint64
	NotarizedHeadersHashes []string
	HeaderGasConsumption   HeaderGasConsumption
	TransactionsPool       *Pool
	AlteredAccounts        map[string]*AlteredAccount
	NumberOfShards         uint32
	IsImportDB             bool
}

// HeaderGasConsumption holds the data needed to save the gas consumption of a header
type HeaderGasConsumption struct {
	GasProvided    uint64
	GasRefunded    uint64
	GasPenalized   uint64
	MaxGasPerBlock uint64
}

// Pool will hold all types of transaction
type Pool struct {
	Txs      map[string]data.TransactionHandlerWithGasUsedAndFee
	Scrs     map[string]data.TransactionHandlerWithGasUsedAndFee
	Rewards  map[string]data.TransactionHandlerWithGasUsedAndFee
	Invalid  map[string]data.TransactionHandlerWithGasUsedAndFee
	Receipts map[string]data.TransactionHandlerWithGasUsedAndFee
	Logs     []*data.LogData
}

// ValidatorRatingInfo is a structure containing validator rating information
type ValidatorRatingInfo struct {
	PublicKey string
	Rating    float32
}

// RoundInfo is a structure containing block signers and shard id
type RoundInfo struct {
	Index            uint64
	SignersIndexes   []uint64
	BlockWasProposed bool
	ShardId          uint32
	Epoch            uint32
	Timestamp        time.Duration
}
