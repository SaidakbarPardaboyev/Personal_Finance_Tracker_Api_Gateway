package client

import (
	"api_gateway/configs"
	pb "api_gateway/genproto/learning_service"
	pbp "api_gateway/genproto/progress_service"
	pbu "api_gateway/genproto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IServiceManager interface {
	LessonService() pb.LessonServiceClient
	UserLessonService() pb.UserLessonServiceClient
	UsersService() pbu.UsersServiceClient
	UserProgressService() pbp.UserProgressServiceClient
}

type grpcClients struct {
	lessonService       pb.LessonServiceClient
	userLessonService   pb.UserLessonServiceClient
	usersService        pbu.UsersServiceClient
	userProgressService pbp.UserProgressServiceClient
}

func NewGrpcClients(cfg *configs.Config) (IServiceManager, error) {

	connLessonService, err := grpc.NewClient(cfg.LearingServiceGrpcHost+cfg.LearingServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connUserLessonService, err := grpc.NewClient(cfg.LearingServiceGrpcHost+cfg.LearingServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connUsersService, err := grpc.NewClient(cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connUserProgressService, err := grpc.NewClient(cfg.UserServiceGrpcHost+cfg.UserServiceGrpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		lessonService:       pb.NewLessonServiceClient(connLessonService),
		userLessonService:   pb.NewUserLessonServiceClient(connUserLessonService),
		usersService:        pbu.NewUsersServiceClient(connUsersService),
		userProgressService: pbp.NewUserProgressServiceClient(connUserProgressService),
	}, nil
}

func (g *grpcClients) LessonService() pb.LessonServiceClient {
	return g.lessonService
}

func (g *grpcClients) UserLessonService() pb.UserLessonServiceClient {
	return g.userLessonService
}

func (g *grpcClients) UsersService() pbu.UsersServiceClient {
	return g.usersService
}

func (g *grpcClients) UserProgressService() pbp.UserProgressServiceClient {
	return g.userProgressService
}
