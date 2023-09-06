package account

import "sync"

// Define the Account type here.
type Account struct {
	isOpen bool
	amount int64

	mu sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{isOpen: true, amount: amount}
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen {
		return 0, false
	}

	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen {
		return 0, false
	}

    if a.amount + amount < 0 {
        return 0, false
    }

	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.isOpen {
		return 0, false
	}

	a.isOpen = false
	return a.amount, true
}
