package server

import (
	"context"

	"github.com/Yasinqurni/be-project/src/app/user/grpc/pb"
	"github.com/Yasinqurni/be-project/src/app/user/service"
)

type UserGrpcServer struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserGrpcServer(service service.UserService) *UserGrpcServer {
	return &UserGrpcServer{service: service}
}

func (s *UserGrpcServer) GetAllUser(ctx context.Context, req *pb.UserRequest) (*pb.UserList, error) {
	users, err := s.service.GetAllUser(req.Ids)
	if err != nil {
		return nil, err
	}

	var datas []*pb.UserResponse
	for _, user := range *users {
		data := pb.UserResponse{
			Id:      uint32(user.ID),
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
			Phone:   user.Phone,
		}

		datas = append(datas, &data)
	}

	return &pb.UserList{Users: datas}, nil
}
