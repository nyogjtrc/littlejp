package littlejp

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/nyogjtrc/littlejp/proto"
)

// Server little jp
type Server struct {
	Pot    Pot
	wallet JackpotWallet
}

// NewServer creatre Server instance
func NewServer() *Server {
	s := new(Server)
	s.Pot = Pot{
		Config: PotConfig{
			ProbabilityBase: 10,
		},
	}
	return s
}

func (s *Server) GetStatus(c context.Context, req *pb.Empty) (*pb.StatusReply, error) {
	reply := new(pb.StatusReply)
	reply.Amount = s.wallet.GetBalance()
	return reply, nil
}

func (s *Server) ThrowMoney(c context.Context, req *pb.ThrowRequest) (*pb.ThrowReply, error) {
	if req.Amount <= 0 {
		return nil, errors.New("amount should be greater than 0")
	}

	s.Pot.CreateThrowRecord(req.UserId, req.Amount, "")

	record, err := s.wallet.Transaction(OpCashIn, req.Amount)
	if err != nil {
		return nil, err
	}

	reply := pb.ThrowReply{}
	reply.Amount = record.Balance
	reply.IsWinner = s.Pot.IsWinner()

	if reply.IsWinner {
		s.Pot.CreateWinnerRecord(req.UserId, record.Balance)
		winTR, err := s.wallet.Transaction(OpCashOut, -1*record.Balance)
		if err != nil {
			return nil, err
		}
		fmt.Println("winner:", req.UserId, winTR)
	}

	return &reply, nil
}

func (s *Server) Latest10Winner(c context.Context, req *pb.Empty) (*pb.WinnerHistory, error) {
	reply := new(pb.WinnerHistory)
	return reply, nil
}

func (s *Server) Top10Winner(c context.Context, req *pb.Empty) (*pb.WinnerHistory, error) {
	return nil, nil
}
