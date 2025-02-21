package handler

import (
	"Proger30/task/model"
	"Proger30/task/service"
	sr "Proger30/task/utils/service.response"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	redis "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func (h *Handler) PeopleInfoAdd(c *gin.Context) {
	var req *model.PeopleInfoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// log here...
		c.JSON(http.StatusBadRequest, sr.Error(-21, "error of ShouldBindJSON: "+err.Error()))
		return
	}
	isCorrectIin, _, _ := service.IinCheckToCorrect(req.Iin)
	if !isCorrectIin {
		c.JSON(http.StatusUnprocessableEntity, sr.Error(-22, "incorrect IIN"))
		return
	}

	result := h.DB.Create(&model.People{PeopleInfoRequest: *req})
	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, sr.Error(-23, "error adding to db: "+result.Error.Error()))
		return

	}
	c.JSON(http.StatusOK, sr.Ok("success"))
}

func (h *Handler) PeopleInfoGet(c *gin.Context) {
	var peoples []model.People
	var result *gorm.DB
	var sqlScript, arg string
	attribute := c.Param("attribute")
	value := c.Param("value")

	switch attribute {
	case "iin":
		isCorrectIin, _, _ := service.IinCheckToCorrect(value)
		if !isCorrectIin {
			c.JSON(http.StatusUnprocessableEntity, sr.Error(-31, "incorrect IIN"))
			return
		}
		sqlScript = "iin = ?"
		arg = value
	case "phone":
		sqlScript = "LOWER(name) LIKE ?"
		arg = "%" + strings.ToLower(value) + "%"
	default:
		c.JSON(http.StatusUnprocessableEntity, sr.Error(-32, "incorrect attribute"))
		return
	}

	cachedData, err := h.RedisDb.Get(c, value).Bytes()
	if err == nil {
		json.Unmarshal(cachedData, &peoples)
	} else {
		if err != redis.Nil {
			log.Printf("Redis error: %v", err)
		}
		result = h.DB.Where(sqlScript, arg).Find(&peoples)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, sr.Error(-33, "error getting from db: "+result.Error.Error()))
			return
		}
		jsonData, err := json.Marshal(peoples)
		if err != nil {
			c.JSON(500, gin.H{"error": "Data serialization error"})
			return
		}

		go func() {
			if err := h.RedisDb.Set(c, strings.ToLower(value), jsonData, 5*time.Minute).Err(); err != nil {
				log.Printf("Cache update failed: %v", err)
			}
		}()
	}

	if attribute == "iin" {
		if len(peoples) == 0 {
			c.JSON(http.StatusNotFound, sr.Error(-34, "no records found in the database :("))
			return
		}
		c.JSON(http.StatusOK, sr.OkWithData("success", peoples[0]))
		return
	} else if attribute == "phone" {
		c.JSON(http.StatusOK, sr.OkWithData("success", peoples))
		return
	}
}
