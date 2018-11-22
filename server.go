package littlejp

import (
	"context"

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
			ProbabilityBase: 1000,
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
	record, err := s.wallet.Transaction(OpCashIn, req.Amount)
	if err != nil {
		return nil, err
	}
	reply := pb.ThrowReply{}
	reply.Amount = record.Balance

	if s.Pot.IsWinner() {
		reply.IsWinner = true
	} else {
		reply.IsWinner = false
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
