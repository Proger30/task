package service

import (
	"Proger30/task/config"
	"regexp"
	"strconv"
	"time"
)

const (
	Male   = "male"
	Female = "female"
)

var W2 = [11]int{3, 4, 5, 6, 7, 8, 9, 10, 11, 1, 2}

type Service struct {
	Config *config.Configuraiton
}

func NewService(config *config.Configuraiton) *Service {
	return &Service{
		Config: config,
	}
}

func IinCheckToCorrect(iin string) (isCorrect bool, sex string, birthDate string) {
	if m, _ := regexp.MatchString(`^[0-9]{12}$`, iin); len(iin) != 12 || !m {
		return
	}
	controll := 0
	for i := 0; i < 11; i++ {
		rank, _ := strconv.Atoi(iin[i : i+1])
		controll += (i + 1) * rank
	}
	controll %= 11

	if controll%11 == 10 {
		controll = 0
		for i := 0; i < 11; i++ {
			controll += controll * W2[i]
		}
		if controll%11 == 10 {
			return
		}
	}

	if rank12, _ := strconv.Atoi(iin[11:]); controll != rank12 {
		return
	}

	if t, err := time.Parse("060102", iin[:6]); err != nil {
		return
	} else {
		birthDate = t.Format("02.01.2006")
	}

	if rank7, _ := strconv.Atoi(iin[6:7]); rank7%2 == 1 {
		sex = Male
	} else {
		sex = Female
	}

	return true, sex, birthDate
}
