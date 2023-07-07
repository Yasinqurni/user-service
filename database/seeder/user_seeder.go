package seeder

import (
	"fmt"
	"log"

	"github.com/Yasinqurni/be-project/src/app/user/model"
	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {

	data := make([]model.User, 3)

	for i := 0; i < 3; i++ {
		user := model.User{
			Name:    fmt.Sprintf("a %d", i+1),
			Address: fmt.Sprintf("jl. prajurit no-%d", i+1),
			Email:   fmt.Sprintf("a%d@example.com", i+1),
		}
		data[i] = user
	}

	err := db.Create(&data).Error
	if err != nil {
		log.Fatal("Error ", err)
	}

}
