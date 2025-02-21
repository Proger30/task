package config

import (
	"encoding/json"
	"os"
)

type Configuraiton struct {
	Port    string        `json:"port"`
	DbCreds DbCredentials `json:"dbCredentials"`
}

type DbCredentials struct {
	UserDB     string `json:"userDB"`
	PasswordDB string `json:"passwordDB"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	DBName     string `json:"dbName"`
}

func (c *Configuraiton) ReadFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}
}

func NewConfiguration(path string) *Configuraiton {
	var configuraiton Configuraiton
	configuraiton.ReadFile(path)
	return &configuraiton
}
