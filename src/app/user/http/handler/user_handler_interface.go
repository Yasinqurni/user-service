package handler

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	ListUser(c *gin.Context)
	DetailUser(c *gin.Context)
	ListUserByIds(c *gin.Context)
}
