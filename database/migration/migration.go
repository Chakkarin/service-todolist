package main

import (
	"github.com/Chakkarin/service-todolist/config"
	databases "github.com/Chakkarin/service-todolist/database"
	"github.com/Chakkarin/service-todolist/database/entities"
)

func main() {
	cfg := config.LoadConfig()
	db := databases.NewPostgresDatabase(&cfg.Database)

	tx := db.Begin()

	_ = tx.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	if ex := tx.AutoMigrate(
		&entities.Todos{},
	); ex != nil {
		panic(ex)
	}

	if ex := tx.Commit().Error; ex != nil {
		tx.Rollback()
		panic(ex)
	}
}
