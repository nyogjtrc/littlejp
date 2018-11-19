package littlejp

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type WinnerRecoard struct {
	Winner string
	Amount float32
	WinAt  time.Time
}

type PotConfig struct {
	ProbabilityBase int
}

type Pot struct {
	Config  PotConfig
	Amount  float64
	History []WinnerRecoard
}

func (p *Pot) IsWinner() bool {
	n := rand.Intn(p.Config.ProbabilityBase)
	fmt.Println("Amount:", p.Amount, "| is winner?", n)
	if n == 1 {
		return true
	}
	return false
}

func (p *Pot) TakeMoney() float64 {
	amount := p.Amount
	p.Amount -= amount
	return amount
}
