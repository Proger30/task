package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PeopleInfoRequest struct {
	Name  string `json:"name"`
	Iin   string `json:"iin"`
	Phone string `json:"phone"`
}

type People struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PeopleInfoRequest
}

func (p *People) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}

type PeopleInfoGetResponse struct {
	PeopleInfoRequest
}
