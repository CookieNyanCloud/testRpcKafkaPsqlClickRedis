package server

import (
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/server"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	lg   *logger.Logger
	port string
}

func NewServer(lg *logger.Logger, port string) *Server {
	return &Server{
		lg:   lg,
		port: port,
	}
}

func (s *Server) Run() error {
	gs := grpc.NewServer()
	userServer := server.NewUsersService(s.lg)
	api.RegisterUsersServer(gs, userServer)
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	return gs.Serve(l)
}
