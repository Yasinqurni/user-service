package routes

import (
	"github.com/Yasinqurni/be-project/src/app/user/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(group *gin.RouterGroup, db *gorm.DB) {

	http.UserRoute(group, db)
}
