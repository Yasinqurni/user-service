package main

import (
	"github.com/Yasinqurni/be-project/database/seeder"
	"github.com/Yasinqurni/be-project/pkg/database"
)

func main() {

	db, _ := database.InitDB()

	seeder.UserSeed(db)
}
