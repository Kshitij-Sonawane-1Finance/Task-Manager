package conn

import (
	"log"
	"os"

	"github.com/kshitij/taskManager/loadEnv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB() *gorm.DB {

	loadEnv.LoadEnv();

	dsn := os.Getenv("DSN")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});
	if err != nil {
		log.Fatal(err);
	}

	return DB;

}