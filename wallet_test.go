package littlejp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction(t *testing.T) {
	wallet := JackpotWallet{}

	r, err := wallet.Transaction(OpCashIn, 10.0)
	if err != nil {
		t.Error(err)
	}

	//assert.True(t, r.ID > 0)

	assert.Equal(t, OpCashIn, r.Opcode)
	assert.Equal(t, 10.0, r.Amount)
	assert.Equal(t, 10.0, r.Balance)

	assert.Equal(t, 10.0, wallet.Balance)

	r, err = wallet.Transaction(OpCashIn, 1.0)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, OpCashIn, r.Opcode)
	assert.Equal(t, 1.0, r.Amount)
	assert.Equal(t, 11.0, r.Balance)

	assert.Equal(t, 11.0, wallet.Balance)
}
