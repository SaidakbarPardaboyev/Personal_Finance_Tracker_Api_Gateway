package client

import (
	"api_gateway/configs"
	pbu "api_gateway/genproto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IServiceManager interface {
	UsersService() pbu.UsersServiceClient
	AuthService() pbu.AuthServiceClient
}

type grpcClients struct {
	usersService pbu.UsersServiceClient
	authService  pbu.AuthServiceClient
}

func NewGrpcClients(cfg *configs.Config) (IServiceManager, error) {

	connUsersService, err := grpc.NewClient(cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		usersService: pbu.NewUsersServiceClient(connUsersService),
		authService:  pbu.NewAuthServiceClient(connUsersService),
	}, nil
}

func (g *grpcClients) UsersService() pbu.UsersServiceClient {
	return g.usersService
}

func (g *grpcClients) AuthService() pbu.AuthServiceClient {
	return g.authService
}
