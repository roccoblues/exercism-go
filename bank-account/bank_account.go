// Package account implements a simple bank account.
package account

import "sync"

// Account represents a bank account
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open a new account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit}
}

// Close the account
func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true
	payout := a.balance
	a.balance = 0

	return payout, true
}

// Balance returns the account balance
func (a *Account) Balance() (int64, bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}

	return a.balance, !a.closed
}

// Deposit the amount in the account
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	if a.balance+amount < 0 {
		return a.balance, false
	}
	a.balance += amount

	return a.balance, true
}
