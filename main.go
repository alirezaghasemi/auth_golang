package main

import (
	"auth/config"
	"auth/db"
	"auth/routes"
	"auth/utils"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	utils.InitLogger()

	db.Init(cfg)

	r := routes.SetupRouter()
	utils.Logger.Println("Server started on port 6060")
	r.Run(":6060")
}
