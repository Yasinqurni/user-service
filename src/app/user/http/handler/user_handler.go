package handler

import (
	"net/http"
	"strconv"

	"github.com/Yasinqurni/be-project/src/app/user/http/request"
	"github.com/Yasinqurni/be-project/src/app/user/http/response"
	"github.com/Yasinqurni/be-project/src/app/user/service"
	"github.com/gin-gonic/gin"
)

type userHandlerImpl struct {
	userService service.UserService
}

func NewUserHandlerImpl(userService service.UserService) UserHandler {
	return &userHandlerImpl{
		userService: userService,
	}
}

func (h *userHandlerImpl) Create(c *gin.Context) {

	var user request.UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {
		data := response.NewErrorResponse("cannot Bind JSON", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	err := h.userService.Create(&user)
	if err != nil {
		data := response.NewErrorResponse("cannot create user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success create data", nil)
	c.JSON(http.StatusCreated, data)
}

func (h *userHandlerImpl) Update(c *gin.Context) {

	var request request.UserRequest

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		data := response.NewErrorResponse("cannot Bind JSON", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot get user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	if user.Name == "" {
		data := response.NewErrorResponse("user not found", nil)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	err = h.userService.Update(uint(id), request)
	if err != nil {
		data := response.NewErrorResponse("cannot update user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success update data", nil)
	c.JSON(http.StatusOK, data)
}

func (h *userHandlerImpl) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot find user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	if user.Name == "" {
		data := response.NewErrorResponse("user not found", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	err = h.userService.Delete(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot delete user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	data := response.NewResponse("success delete data", nil)
	c.JSON(http.StatusOK, data)
}

func (h *userHandlerImpl) ListUser(c *gin.Context) {

	users, err := h.userService.ListUser()
	if err != nil {
		data := response.NewErrorResponse("cannot list user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}
	var userResponse []response.UserResponse
	for _, user := range *users {
		userResponse = append(userResponse, response.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		})
	}

	data := response.NewResponse("success list data", userResponse)
	c.JSON(http.StatusOK, data)
}

func (h *userHandlerImpl) DetailUser(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		data := response.NewErrorResponse("please insert id", err)
		c.JSON(http.StatusBadRequest, data)
		return
	}

	user, err := h.userService.GetUser(uint(id))
	if err != nil {
		data := response.NewErrorResponse("cannot detail user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	userResponse := response.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
	}

	if user.Name == "" {
		data := response.NewErrorResponse("user not found", nil)
		c.JSON(http.StatusBadRequest, data)
		return
	}
	data := response.NewResponse("success detail data", userResponse)
	c.JSON(http.StatusOK, data)
}

func (h *userHandlerImpl) ListUserByIds(c *gin.Context) {

	ids := c.Query("ids")

	users, err := h.userService.GetAllUser(ids)
	if err != nil {
		data := response.NewErrorResponse("cannot list user", err)
		c.JSON(http.StatusInternalServerError, data)
		return
	}

	var result []response.UserResponse
	for _, user := range *users {
		result = append(result, response.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Address: user.Address,
			Email:   user.Email,
		})
	}
	data := response.NewResponse("success detail data", result)
	c.JSON(http.StatusOK, data)
}
