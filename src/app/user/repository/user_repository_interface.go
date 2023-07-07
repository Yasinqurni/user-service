package repository

import "github.com/Yasinqurni/be-project/src/app/user/model"

type UserRepository interface {
	Create(user *model.User) error
	Update(id uint, user *model.User) error
	GetUser(id uint) (*model.User, error)
	GetAllUser(ids []string) (*[]model.User, error)
	ListUser() (*[]model.User, error)
	Delete(id uint) error
}
