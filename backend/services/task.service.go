package services

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/models"
	"gorm.io/gorm"
)

type ReturnType struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

type DeleteTaskStruct struct {
	ID uint `json:"id"`
}

func CreateTask(ctx *fiber.Ctx, task models.Task) ReturnType {

	var db *gorm.DB = conn.InitializeDB()

	result := db.Create(&task)
	if result.Error != nil {
		return ReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		};
	}

	fmt.Println("Inserted Row");
	fmt.Println(result);
	return ReturnType{
		StatusCode: http.StatusCreated,
		Message: "Task Created",
	};

}


func FindTask(ctx *fiber.Ctx, id uint64) (models.Task, error) {

	var db *gorm.DB = conn.InitializeDB();
	var task models.Task;

	// db.First(&task, id);
	err := db.First(&task, id).Error;

	if task.UserID == 0 {
		return task, err;
	}

	fmt.Println(task);

	return task, nil;

}


// func DeleteTask(ctx *fiber.Ctx, deleteTaskStruct DeleteTaskStruct) ReturnType {
// 	var db *gorm.DB = conn.InitializeDB();
// 	var task models.Task;

// 	db.First(&task, DeleteTaskStruct.ID);


// 	// result := db.Delete(&task, deleteTaskStruct.ID);
// 	// if result.Error != nil {
// 	// 	return ReturnType{
// 	// 		StatusCode: http.StatusInternalServerError,
// 	// 		Message: "Internal Server Error",
// 	// 	};
// 	// }
	
// 	// fmt.Println("Deleted Row");
// 	// fmt.Println(result);
// 	return ReturnType{
// 		StatusCode: http.StatusOK,
// 		Message: "Task Deleted",
// 	};
// }