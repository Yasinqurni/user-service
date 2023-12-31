package initproject

import (
	"fmt"

	"github.com/Yasinqurni/be-project/pkg/database"
	"github.com/Yasinqurni/be-project/src/app/user/model"
	"github.com/Yasinqurni/be-project/src/routes"
	"github.com/gin-gonic/gin"
)

func InitProject() {

	db, PORT := database.InitDB()
	db.AutoMigrate(&model.User{})
	r := gin.Default()

	v1 := r.Group("/api/v1")
	routes.Router(v1, db)

	r.Run(fmt.Sprintf(":%s", PORT))
}
