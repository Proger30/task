package db

import (
	"Proger30/task/config"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Configuraiton) *gorm.DB {
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.DbCreds.Host, cfg.DbCreds.UserDB, cfg.DbCreds.PasswordDB, cfg.DbCreds.DBName, cfg.DbCreds.Port)
	fmt.Println(connection)
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}
