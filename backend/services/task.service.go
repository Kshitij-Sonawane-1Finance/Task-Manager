package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/models"
	"gorm.io/gorm"
)

type ReturnType struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
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


func FindTask(ctx *fiber.Ctx, id uint64, userId uint) (dto.Task, error) {

	var db *gorm.DB = conn.InitializeDB();
	var task dto.Task;

	err := db.Where("user_id = ?", userId).Where("active = ?", true).First(&task, id).Error;

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return task, err;
	}

	return task, nil;

}

func FindAllTasks(ctx *fiber.Ctx, userId uint) ([]dto.Task, error) {

	var db *gorm.DB = conn.InitializeDB();
	var tasks []dto.Task;

	err := db.Where("user_id = ?", userId).Where("active = ?", true).Find(&tasks).Error;

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks, err;
	}

	return tasks, nil;

}



func DeleteTask(ctx *fiber.Ctx, id uint64, userId uint) ReturnType {
	var db *gorm.DB = conn.InitializeDB();
	var task models.Task;

	_, err := FindTask(ctx, id, userId);
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusNotFound,
			Message: "Task Does not exist",
		}
	}

	result := db.Delete(&task, id);
	if result.Error != nil {
		return ReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		};
	}
	
	fmt.Println("Deleted Row");
	// fmt.Println(result);
	return ReturnType{
		StatusCode: http.StatusOK,
		Message: "Task Deleted",
	};
}


func SoftDelete(ctx *fiber.Ctx, id uint64, userId uint) ReturnType {

	var db *gorm.DB = conn.InitializeDB();
	var task models.Task;

	// err := db.First(&task, id).Error;
	_, err := FindTask(ctx, id, userId);
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusNotFound,
			Message: "Task Does Not Exist",
		}
	}
	
	// task.Active = false;

	// err = db.Save(&task).Error;

	err = db.Model(&task).Where("id = ?", id).Update("active", false).Error;
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return ReturnType{
		StatusCode: http.StatusOK,
		Message: "Task Deleted Successfully",
	}

}


func UpdateTask(ctx *fiber.Ctx, id uint64, updateTask dto.UpdateTask, userId uint) ReturnType {

	var db *gorm.DB = conn.InitializeDB();
	var task models.Task;

	_, err := FindTask(ctx, id, userId);
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusNotFound,
			Message: "Task Does not exist",
		}
	}

	err = db.Model(&task).Where("id = ?", id).Updates(updateTask).Error;
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return ReturnType{
		StatusCode: http.StatusOK,
		Message: "Task Updated Successfully",
	}

}