package http

import (
	"github.com/Yasinqurni/be-project/src/app/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(group *gin.RouterGroup, db *gorm.DB) {

	userHandler, _, _ := user.UserInit(db)

	userGroup := group.Group("/user")

	userGroup.POST("/", userHandler.Create)
	userGroup.PATCH("/:id", userHandler.Update)
	userGroup.DELETE("/:id", userHandler.Delete)
	userGroup.GET("/", userHandler.ListUser)
	userGroup.GET("/:id", userHandler.DetailUser)
	userGroup.GET("/ids", userHandler.ListUserByIds)

}
