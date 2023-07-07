package service

import (
	"github.com/Yasinqurni/be-project/src/app/user/http/request"
	"github.com/Yasinqurni/be-project/src/app/user/model"
)

type UserService interface {
	Create(user *request.UserRequest) error
	Update(id uint, userRequest request.UserRequest) error
	GetUser(id uint) (*model.User, error)
	GetAllUser(ids string) (*[]model.User, error)
	Delete(id uint) error
	ListUser() (*[]model.User, error)
}
