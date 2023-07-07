package service

import (
	"strings"

	"github.com/Yasinqurni/be-project/src/app/user/http/request"
	"github.com/Yasinqurni/be-project/src/app/user/model"
	"github.com/Yasinqurni/be-project/src/app/user/repository"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (s *userServiceImpl) Create(user *request.UserRequest) error {
	data := model.User{
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}
	return s.userRepository.Create(&data)
}

func (s *userServiceImpl) Update(id uint, userRequest request.UserRequest) error {
	data := model.User{
		Name:    userRequest.Name,
		Address: userRequest.Address,
		Email:   userRequest.Email,
	}
	return s.userRepository.Update(id, &data)

}

func (s *userServiceImpl) GetUser(id uint) (*model.User, error) {

	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userServiceImpl) GetAllUser(ids string) (*[]model.User, error) {

	var idString []string

	if ids != "" {
		idString = strings.Split(ids, ",")
	}

	users, err := s.userRepository.GetAllUser(idString)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userServiceImpl) Delete(id uint) error {
	return s.userRepository.Delete(id)
}

func (s *userServiceImpl) ListUser() (*[]model.User, error) {
	users, err := s.userRepository.ListUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}
