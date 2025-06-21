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

var dbInstance *gorm.DB;

func (c *dbService) InitializeDB() *gorm.DB {

	// For avoiding creating multiple open connections with the db
	if dbInstance != nil {
		return dbInstance
	}

	c.loadEnvService.LoadEnv();

	dsn := os.Getenv("DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	if err != nil {
		log.Fatal(err);
	}

	dbInstance = DB;
	return DB;

}