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

func TestLatest10Winner(t *testing.T) {
	p := Pot{}
	assert.Equal(t, 0, len(p.Latest10Winner()))

	p.CreateWinnerRecord(1, 10.0)
	p.CreateWinnerRecord(2, 11.0)
	assert.Equal(t, 2, len(p.Latest10Winner()))

	p.CreateWinnerRecord(3, 12.0)
	p.CreateWinnerRecord(4, 13.0)
	p.CreateWinnerRecord(5, 14.0)
	p.CreateWinnerRecord(6, 15.0)
	p.CreateWinnerRecord(7, 16.0)
	p.CreateWinnerRecord(8, 17.0)
	p.CreateWinnerRecord(9, 18.0)
	p.CreateWinnerRecord(10, 19.0)
	p.CreateWinnerRecord(11, 20.0)
	assert.Equal(t, 10, len(p.Latest10Winner()))
	assert.Equal(t, uint64(2), p.Latest10Winner()[0].UserID)
	assert.Equal(t, uint64(11), p.Latest10Winner()[9].UserID)

}
func TestTop10Winner(t *testing.T) {
	p := Pot{}
	assert.True(t, len(p.Top10Winner()) == 0)
	assert.Equal(t, 0, len(p.Top10Winner()))

	p.CreateWinnerRecord(1, 10.0)
	p.CreateWinnerRecord(2, 11.0)
	assert.Equal(t, 2, len(p.Top10Winner()))

	p.CreateWinnerRecord(3, 12.0)
	p.CreateWinnerRecord(4, 13.0)
	p.CreateWinnerRecord(5, 14.0)
	p.CreateWinnerRecord(6, 15.0)
	p.CreateWinnerRecord(7, 16.0)
	p.CreateWinnerRecord(8, 17.0)
	p.CreateWinnerRecord(9, 18.0)
	p.CreateWinnerRecord(10, 19.0)
	p.CreateWinnerRecord(11, 20.0)
	assert.Equal(t, 10, len(p.Top10Winner()))
	assert.Equal(t, 20.0, p.Top10Winner()[0].Amount)
	assert.Equal(t, 11.0, p.Top10Winner()[9].Amount)
}
