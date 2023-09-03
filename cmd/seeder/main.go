package main

import (
	"github.com/Yasinqurni/be-project/database/seeder"
	"github.com/Yasinqurni/be-project/pkg/database"
)

func main() {

	configEnv := database.InitDB()

	seeder.UserSeed(configEnv.DB)
}
