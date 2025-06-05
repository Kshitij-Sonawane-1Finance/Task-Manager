package loadEnv

import (
	"log"

	"github.com/joho/godotenv"
)

type LoadEnvService interface {
	LoadEnv()
}

type loadEnvService struct {}

func NewLoadEnvService() LoadEnvService {
	return &loadEnvService{};
}

func (l *loadEnvService) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}