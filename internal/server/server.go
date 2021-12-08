package server

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/services"
	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
)

type UsersService struct {
	api.UnimplementedUsersServer
	Srv services.IServices
}

func NewUsersService(services services.IServices) *UsersService {
	return &UsersService{
		UnimplementedUsersServer: api.UnimplementedUsersServer{},
		Srv:                      services,
	}
}

func (u *UsersService) CreateUser(ctx context.Context, r *api.CreateRequest) (*api.CDResponse, error) {
	lg.Infof("handle Create User name %s", r.GetName())
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) DeleteUser(ctx context.Context, r *api.DeleteRequest) (*api.CDResponse, error) {
	lg.Infof("handle Delete User id %s", r.GetId())
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) FindAll(ctx context.Context, r *api.Empty) (*api.AllResponse, error) {
	lg.Info("handle Find All")
	return &api.AllResponse{
		Users: nil,
	}, nil
}
