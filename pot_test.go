package littlejp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateThrowRecord(t *testing.T) {
	p := Pot{}
	p.CreateThrowRecord(1, 1.0, "")
	assert.Equal(t, 1, len(p.ThrowRecords))
}

func TestCreateWinnerRecord(t *testing.T) {
	p := Pot{}
	p.CreateWinnerRecord(1, 10.0)
	assert.Equal(t, 1, len(p.WinnerRecoards))
}
