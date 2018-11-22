package littlejp

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type ThrowMoneyRecord struct {
	UserID   uint64
	Amount   float64
	FromGame string
	ThrowAt  time.Time
}

type WinnerRecoard struct {
	UserID uint64
	Amount float64
	WinAt  time.Time
}

type PotConfig struct {
	ProbabilityBase int
}

// Pot pool
type Pot struct {
	Config         PotConfig
	ThrowRecords   []ThrowMoneyRecord
	WinnerRecoards []WinnerRecoard
}

func (p *Pot) CreateThrowRecord(userid uint64, amount float64, fromGame string) {
	p.ThrowRecords = append(p.ThrowRecords, ThrowMoneyRecord{
		UserID:   userid,
		Amount:   amount,
		FromGame: fromGame,
		ThrowAt:  time.Now(),
	})
}

func (p *Pot) IsWinner() bool {
	n := rand.Intn(p.Config.ProbabilityBase)
	if n == 1 {
		return true
	}
	return false
}

func (p *Pot) CreateWinnerRecord(userid uint64, amount float64) {
	p.WinnerRecoards = append(p.WinnerRecoards, WinnerRecoard{
		UserID: userid,
		Amount: amount,
		WinAt:  time.Now(),
	})
}
