package migration

import (
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/models"
	"gorm.io/gorm"
)

func Migrate() {

	var db *gorm.DB = conn.InitializeDB();

	db.Migrator().CreateTable(&models.User{});
	db.Migrator().CreateTable(&models.Admin{});
	db.Migrator().CreateTable(&models.Task{});

}