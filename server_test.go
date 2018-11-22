package littlejp

import (
	"context"
	"testing"

	pb "github.com/nyogjtrc/littlejp/proto"
	"github.com/stretchr/testify/assert"
)

func TestThrowMoney(t *testing.T) {
	userId := 1
	amount := 1.0

	s := NewServer()

	throwReply, err := s.ThrowMoney(context.Background(), &pb.ThrowRequest{
		UserId: uint64(userId),
		Amount: amount,
	})
	if err != nil {
		t.Log(err)
	}

	assert.Equal(t, amount, throwReply.Amount)

	statusReply, err := s.GetStatus(context.Background(), &pb.Empty{})
	if err != nil {
		t.Log(err)
	}

	if throwReply.IsWinner {
		assert.Equal(t, 0, statusReply.Amount)
	} else {
		assert.Equal(t, amount, statusReply.Amount)
	}
}
