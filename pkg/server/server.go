package server

import (
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/server"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/services"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
	"google.golang.org/grpc"
	"net"
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

func (s *Server) Run(services services.IServices) error {
	gs := grpc.NewServer()
	userServer := server.NewUsersService(services)
	api.RegisterUsersServer(gs, userServer)
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	return gs.Serve(l)
}
