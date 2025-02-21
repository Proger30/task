package model

type IinCheckResponce struct {
	Correct     bool   `json:"correct"`
	Sex         string `json:"sex"` // ToDo: change tipe
	DateOfBirth string `json:"date_of_birth"`
}
