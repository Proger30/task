package main

import (
	"Proger30/task/config"
	"Proger30/task/db"
	"Proger30/task/handler"
	"Proger30/task/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	_cfg := config.NewConfiguration("./config.json")
	_db := db.NewDB(_cfg)
	_rdb := db.NewRedisClient()
	defer _rdb.Close()
	_srv := service.NewService(_cfg)
	_handler := handler.NewHandler(_cfg, _srv, _db, _rdb)

	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/iin_check/:iin", _handler.IinCheck)

	p := r.Group("/people/info")
	p.POST("", _handler.PeopleInfoAdd)
	p.OPTIONS("", func(c *gin.Context) {
		c.Status(204)
	})

	p.GET("/:attribute/:value", _handler.PeopleInfoGet)

	r.Run(_cfg.Port)
}
