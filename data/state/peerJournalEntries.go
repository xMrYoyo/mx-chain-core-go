package state

import "math/big"

//------- PeerJournalEntryAddress

// PeerJournalEntryAddress is used to revert a round change
type PeerJournalEntryAddress struct {
	account    *PeerAccount
	oldAddress []byte
}

// NewPeerJournalEntryAddress outputs a new PeerJournalEntry implementation used to revert a round change
func NewPeerJournalEntryAddress(account *PeerAccount, oldAddress []byte) (*PeerJournalEntryAddress, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryAddress{
		account:    account,
		oldAddress: oldAddress,
	}, nil
}

// Revert applies undo operation
func (jen *PeerJournalEntryAddress) Revert() (AccountHandler, error) {
	jen.account.Address = jen.oldAddress

	return jen.account, nil
}

//------- PeerJournalEntrySchnorrPublicKey

// PeerJournalEntrySchnorrPublicKey is used to revert a round change
type PeerJournalEntrySchnorrPublicKey struct {
	account          *PeerAccount
	oldSchnorrPubKey []byte
}

// NewPeerJournalEntryAddress outputs a new PeerJournalEntry implementation used to revert a round change
func NewPeerJournalEntrySchnorrPublicKey(account *PeerAccount, oldSchnorrPubKey []byte) (*PeerJournalEntrySchnorrPublicKey, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntrySchnorrPublicKey{
		account:          account,
		oldSchnorrPubKey: oldSchnorrPubKey,
	}, nil
}

// Revert applies undo operation
func (jen *PeerJournalEntrySchnorrPublicKey) Revert() (AccountHandler, error) {
	jen.account.SchnorrPublicKey = jen.oldSchnorrPubKey

	return jen.account, nil
}

//------- PeerJournalEntryBLSPublicKey

// PeerJournalEntryBLSPublicKey is used to revert a round change
type PeerJournalEntryBLSPublicKey struct {
	account      *PeerAccount
	oldBLSPubKey []byte
}

// NewPeerJournalEntryBLSPublicKey outputs a new PeerJournalEntry implementation used to revert a round change
func NewPeerJournalEntryBLSPublicKey(account *PeerAccount, oldPubKey []byte) (*PeerJournalEntryBLSPublicKey, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryBLSPublicKey{
		account:      account,
		oldBLSPubKey: oldPubKey,
	}, nil
}

// Revert applies undo operation
func (jen *PeerJournalEntryBLSPublicKey) Revert() (AccountHandler, error) {
	jen.account.BLSPublicKey = jen.oldBLSPubKey

	return jen.account, nil
}

//------- PeerJournalEntryStake

// PeerJournalEntryStake is used to revert a stake change
type PeerJournalEntryStake struct {
	account  *PeerAccount
	oldStake *big.Int
}

// NewPeerJournalEntryStake outputs a new PeerJournalEntry implementation used to revert a stake change
func NewPeerJournalEntryStake(account *PeerAccount, oldStake *big.Int) (*PeerJournalEntryStake, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryStake{
		account:  account,
		oldStake: oldStake,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryStake) Revert() (AccountHandler, error) {
	jeb.account.Stake = jeb.oldStake

	return jeb.account, nil
}

// PeerJournalEntryJailTime is used to revert a balance change
type PeerJournalEntryJailTime struct {
	account   *PeerAccount
	oldPeriod TimePeriod
}

// NewPeerJournalEntryJailTime outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryJailTime(account *PeerAccount, oldJailTime TimePeriod) (*PeerJournalEntryJailTime, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryJailTime{
		account:   account,
		oldPeriod: oldJailTime,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryJailTime) Revert() (AccountHandler, error) {
	jeb.account.JailTime = jeb.oldPeriod

	return jeb.account, nil
}

// PeerJournalEntryCurrentShardId is used to revert a shardId change
type PeerJournalEntryCurrentShardId struct {
	account    *PeerAccount
	oldShardId uint32
}

// NewPeerJournalEntryCurrentShardId outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryCurrentShardId(account *PeerAccount, oldShId uint32) (*PeerJournalEntryCurrentShardId, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryCurrentShardId{
		account:    account,
		oldShardId: oldShId,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryCurrentShardId) Revert() (AccountHandler, error) {
	jeb.account.CurrentShardId = jeb.oldShardId

	return jeb.account, nil
}

// PeerJournalEntryNextShardId is used to revert a shardId change
type PeerJournalEntryNextShardId struct {
	account    *PeerAccount
	oldShardId uint32
}

// NewPeerJournalEntryNextShardId outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryNextShardId(account *PeerAccount, oldShId uint32) (*PeerJournalEntryNextShardId, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryNextShardId{
		account:    account,
		oldShardId: oldShId,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryNextShardId) Revert() (AccountHandler, error) {
	jeb.account.NextShardId = jeb.oldShardId

	return jeb.account, nil
}

// PeerJournalEntryInWaitingList is used to revert a shardId change
type PeerJournalEntryInWaitingList struct {
	account          *PeerAccount
	oldInWaitingList bool
}

// NewPeerJournalEntryInWaitingList outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryInWaitingList(account *PeerAccount, oldInWaitingList bool) (*PeerJournalEntryInWaitingList, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryInWaitingList{
		account:          account,
		oldInWaitingList: oldInWaitingList,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryInWaitingList) Revert() (AccountHandler, error) {
	jeb.account.NodeInWaitingList = jeb.oldInWaitingList

	return jeb.account, nil
}

// PeerJournalEntryValidatorSuccessRate is used to revert a success rate change
type PeerJournalEntryValidatorSuccessRate struct {
	account        *PeerAccount
	oldSuccessRate SignRate
}

// NewPeerJournalEntryInWaitingList outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryValidatorSuccessRate(account *PeerAccount, oldSuccessRate SignRate) (*PeerJournalEntryValidatorSuccessRate, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryValidatorSuccessRate{
		account:        account,
		oldSuccessRate: oldSuccessRate,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryValidatorSuccessRate) Revert() (AccountHandler, error) {
	jeb.account.ValidatorSuccessRate = jeb.oldSuccessRate

	return jeb.account, nil
}

// PeerJournalEntryLeaderSuccessRate is used to revert a success rate change
type PeerJournalEntryLeaderSuccessRate struct {
	account        *PeerAccount
	oldSuccessRate SignRate
}

// NewPeerJournalEntryInWaitingList outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryLeaderSuccessRate(account *PeerAccount, oldSuccessRate SignRate) (*PeerJournalEntryLeaderSuccessRate, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryLeaderSuccessRate{
		account:        account,
		oldSuccessRate: oldSuccessRate,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryLeaderSuccessRate) Revert() (AccountHandler, error) {
	jeb.account.LeaderSuccessRate = jeb.oldSuccessRate

	return jeb.account, nil
}

// PeerJournalEntryRating is used to revert a rating change
type PeerJournalEntryRating struct {
	account   *PeerAccount
	oldRating uint32
}

// NewPeerJournalEntryRating outputs a new PeerJournalEntry implementation used to revert a state change
func NewPeerJournalEntryRating(account *PeerAccount, oldRating uint32) (*PeerJournalEntryRating, error) {
	if account == nil {
		return nil, ErrNilAccountHandler
	}

	return &PeerJournalEntryRating{
		account:   account,
		oldRating: oldRating,
	}, nil
}

// Revert applies undo operation
func (jeb *PeerJournalEntryRating) Revert() (AccountHandler, error) {
	jeb.account.Rating = jeb.oldRating

	return jeb.account, nil
}
