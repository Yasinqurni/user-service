package user

import (
	handler "github.com/Yasinqurni/be-project/src/app/user/http/handler"
	"github.com/Yasinqurni/be-project/src/app/user/repository"
	service "github.com/Yasinqurni/be-project/src/app/user/service"
	"gorm.io/gorm"
)

func UserInit(db *gorm.DB) (handler.UserHandler, service.UserService, repository.UserRepository) {

	userRepository := repository.NewUserRepositoryImpl(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandlerImpl(userService)

	return userHandler, userService, userRepository
}
