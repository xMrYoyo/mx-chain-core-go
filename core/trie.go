package core

// TrieNodeVersion defines the version of the trie node
type TrieNodeVersion uint8

const (
	// NotSpecified means that the value is not populated or is not important
	NotSpecified TrieNodeVersion = iota

	// AutoBalanceEnabled is used for data tries, and only after the activation of AutoBalanceDataTriesEnableEpoch flag
	AutoBalanceEnabled
)

// TrieData holds the data that will be inserted into the trie
type TrieData struct {
	Key     []byte
	Value   []byte
	Version TrieNodeVersion
}
