package main

import (
	"log"

	"github.com/Chakkarin/service-todolist/config"
	databases "github.com/Chakkarin/service-todolist/database"

	_ "github.com/Chakkarin/service-todolist/docs"
	"github.com/Chakkarin/service-todolist/server"
	"github.com/joho/godotenv"
)

func init() {
	ex := godotenv.Load()
	if ex != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title Todo API
func main() {
	cfg := config.LoadConfig()
	db := databases.NewPostgresDatabase(&cfg.Database)

	server := server.NewServer(&cfg, db)

	server.Start()
}
