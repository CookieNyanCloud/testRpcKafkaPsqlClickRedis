package server

import (
	"context"
	"time"

	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"
	"github.com/cookienyancloud/testrpckafkapsqlclick/internal/services"
	lg "github.com/cookienyancloud/testrpckafkapsqlclick/pkg/logger/logger"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
	"github.com/google/uuid"
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
	id := uuid.New().String()
	err := u.Srv.AddUser(&domain.User{
		Id:           id,
		Name:         r.Name,
		PasswordHash: r.Password,
	})
	if err != nil {
		return &api.CDResponse{State: "Error"}, err
	}
	err = u.Srv.SendUserLog(&domain.UserLog{
		Id:        id,
		Name:      r.Name,
		Key:       r.Key,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return &api.CDResponse{State: "Error"}, err
	}
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) DeleteUser(ctx context.Context, r *api.DeleteRequest) (*api.CDResponse, error) {
	lg.Infof("handle Delete User id %s", r.GetId())
	err := u.Srv.DeleteUserById(r.Id)
	if err != nil {
		return &api.CDResponse{State: "Error"}, err
	}
	return &api.CDResponse{
		State: "Success",
	}, nil
}

func (u *UsersService) FindAll(ctx context.Context, r *api.Empty) (*api.AllResponse, error) {
	lg.Info("handle Find All")
	var outUsers []*api.User
	users, err := u.Srv.GetUsersCache()
	if err != nil {
		return &api.AllResponse{}, err
	}
	if len(users) == 0 {
		allUsers, err := u.Srv.GetAllUsers()
		if err != nil {
			return &api.AllResponse{}, err
		}
		err = u.Srv.CacheUsers(allUsers)
		if err != nil {
			return &api.AllResponse{}, err
		}
		for _, user := range allUsers {
			outUsers = append(outUsers, &api.User{
				Id:   user.Id,
				Name: user.Name,
			})
		}
		return &api.AllResponse{
			Users: outUsers,
		}, nil
	}
	for _, user := range users {
		outUsers = append(outUsers, &api.User{
			Id:   user.Id,
			Name: user.Name,
		})
	}
	return &api.AllResponse{
		Users: outUsers,
	}, nil
}
