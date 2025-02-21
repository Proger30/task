package handler

import (
	"Proger30/task/config"
	"Proger30/task/service"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handler struct {
	Config  *config.Configuraiton
	Service *service.Service
	DB      *gorm.DB
	RedisDb *redis.Client
}

func NewHandler(config *config.Configuraiton, service *service.Service, db *gorm.DB, rdb *redis.Client) *Handler {
	return &Handler{
		Config:  config,
		Service: service,
		DB:      db,
		RedisDb: rdb,
	}
}
