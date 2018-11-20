package littlejp

import "time"

// transaction opcode
const (
	OpCashIn uint = iota + 1
	OpCashOut
)

// TransactionID unique id
type TransactionID uint64

// TransactionRecord logging each transaction
type TransactionRecord struct {
	ID        TransactionID
	Opcode    uint
	Amount    float64
	Balance   float64
	CreatedAt time.Time
}

// JackpotWallet store jackpot money
type JackpotWallet struct {
	Balance float64
}

// Transaction with jackpot wallet
func (w *JackpotWallet) Transaction(opcode uint, amount float64) (TransactionRecord, error) {
	r := TransactionRecord{}
	r.Opcode = opcode
	r.Amount = amount
	r.CreatedAt = time.Now()

	w.Balance += amount

	r.Balance = w.Balance

	return r, nil
}
