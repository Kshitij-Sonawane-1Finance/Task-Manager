package conn

import (
	"log"
	"os"

	"github.com/kshitij/taskManager/loadEnv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBService interface {
	InitializeDB() *gorm.DB
}

type dbService struct {
	loadEnvService loadEnv.LoadEnvService
}

func NewDBService(loadEnvService loadEnv.LoadEnvService) DBService {
	return &dbService{loadEnvService};
}

func (c *dbService) InitializeDB() *gorm.DB {

	c.loadEnvService.LoadEnv();

	dsn := os.Getenv("DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	if err != nil {
		log.Fatal(err);
	}

	return DB;

}