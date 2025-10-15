package grpc

import (
	"context"

	userpb "github.com/fiqriardiansyah/shopping-proto/gen/go/user"
	"github.com/fiqriardiansyah/user-shopping-api-golang/internal/module/user/usecase"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type UserServer struct {
	userpb.UnimplementedUserServiceServer
	UserUsecase *usecase.UserUseCase
}

func NewUserServer(userUsecase *usecase.UserUseCase) *UserServer {
	return &UserServer{
		UserUsecase: userUsecase,
	}
}

func (s *UserServer) Run(server *grpc.Server) {
	userpb.RegisterUserServiceServer(server, &UserServer{})
}

func (s *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userResponse := &userpb.GetUserResponse{}

	logrus.Info(req.Id)

	userId, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	if user, err := s.UserUsecase.User(userId); err != nil {
		return nil, err
	} else {
		if err := copier.Copy(userResponse, user); err != nil {
			return nil, err
		}
	}

	return userResponse, nil
}
