package littlejp

import (
	"context"

	pb "github.com/nyogjtrc/littlejp/proto"
)

// Server little jp
type Server struct {
	Pot Pot
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
	reply.Amount = s.Pot.Amount
	return reply, nil
}

func (s *Server) ThrowMoney(c context.Context, req *pb.ThrowRequest) (*pb.ThrowReply, error) {
	s.Pot.Amount += req.Amount
	reply := pb.ThrowReply{}

	if s.Pot.IsWinner() {
		reply.Amount = s.Pot.TakeMoney()
		reply.IsWinner = true
	} else {
		reply.Amount = s.Pot.Amount
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
