package grpc

import (
	"github.com/Yasinqurni/be-project/src/app/user/grpc/pb"
	"github.com/Yasinqurni/be-project/src/app/user/grpc/server"
	"github.com/Yasinqurni/be-project/src/app/user/repository"
	"github.com/Yasinqurni/be-project/src/app/user/service"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func NewGrpcServer(s *grpc.Server, db *gorm.DB) *grpc.Server {
	{
		repo := repository.NewUserRepositoryImpl(db)
		service := service.NewUserServiceImpl(repo)
		srv := server.NewUserGrpcServer(service)
		pb.RegisterUserServiceServer(s, srv)
	}
	return s
}
