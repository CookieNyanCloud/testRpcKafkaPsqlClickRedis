package server

import (
	"net"

	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/server"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/services"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
	"google.golang.org/grpc"
)

type Server struct {
	port     string
	services services.IServices
}

func NewServer(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Run(services services.IServices) (*net.Listener, error) {
	gs := grpc.NewServer()
	userServer := server.NewUsersService(services)
	api.RegisterUsersServer(gs, userServer)
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return nil, err
	}
	return &l, gs.Serve(l)
}
