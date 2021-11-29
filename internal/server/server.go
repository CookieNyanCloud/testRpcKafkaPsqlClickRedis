package server

import (
	"context"
	"github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
)

type UsersService struct {
	lg *logger.Logger
	api.UnimplementedUsersServer
}

func NewUsersService(lg *logger.Logger) *UsersService {
	return &UsersService{
		lg:                       lg,
		UnimplementedUsersServer: api.UnimplementedUsersServer{},
	}
}

func (u *UsersService) CreateUser(ctx context.Context, r *api.CreateRequest) (*api.CDResponse, error) {
	u.lg.Infof("handle Create User name %s", r.GetName())
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) DeleteUser(ctx context.Context, r *api.DeleteRequest) (*api.CDResponse, error) {
	u.lg.Infof("handle Delete User id %s", r.GetId())
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) FindAll(ctx context.Context, r *api.Empty) (*api.AllResponse, error) {
	u.lg.Info("handle Find All")
	return &api.AllResponse{
		Users: nil,
	}, nil
}
