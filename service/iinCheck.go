package service

import (
	"Proger30/task/model"

	"github.com/gin-gonic/gin"
)

func (s *Service) IinCheck(c *gin.Context, iin string) (*model.IinCheckResponce, error) {
	isCorrect, sex, birthDate := IinCheckToCorrect(iin)
	resp := &model.IinCheckResponce{
		Correct:     isCorrect,
		Sex:         sex,
		DateOfBirth: birthDate,
	}
	return resp, nil
}
