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


type TaskService interface {
	CreateTask(ctx *fiber.Ctx, task models.Task) ReturnType
	FindTask(ctx *fiber.Ctx, id uint64, userId uint) (dto.Task, error)
	FindAllTasks(ctx *fiber.Ctx, userId uint) ([]dto.Task, error)
	DeleteTask(ctx *fiber.Ctx, id uint64, userId uint) ReturnType
	SoftDelete(ctx *fiber.Ctx, id uint64, userId uint) ReturnType
	UpdateTask(ctx *fiber.Ctx, id uint64, updateTask dto.UpdateTask, userId uint) ReturnType
}

type taskService struct {
	dbService conn.DBService
}

func NewTaskService(dbService conn.DBService) TaskService {
	return &taskService{dbService}
}


type ReturnType struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

func (s *taskService) CreateTask(ctx *fiber.Ctx, task models.Task) ReturnType {

	var db *gorm.DB = s.dbService.InitializeDB()

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


func (s *taskService) FindTask(ctx *fiber.Ctx, id uint64, userId uint) (dto.Task, error) {

	var db *gorm.DB = s.dbService.InitializeDB()
	var task dto.Task;

	err := db.Where("user_id = ?", userId).Where("active = ?", true).First(&task, id).Error;

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return task, err;
	}

	return task, nil;

}

func (s *taskService) FindAllTasks(ctx *fiber.Ctx, userId uint) ([]dto.Task, error) {

	var db *gorm.DB = s.dbService.InitializeDB()
	var tasks []dto.Task;

	err := db.Where("user_id = ?", userId).Where("active = ?", true).Find(&tasks).Error;

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks, err;
	}

	return tasks, nil;

}



func (s *taskService) DeleteTask(ctx *fiber.Ctx, id uint64, userId uint) ReturnType {
	var db *gorm.DB = s.dbService.InitializeDB()
	var task models.Task;

	_, err := s.FindTask(ctx, id, userId);
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


func (s *taskService) SoftDelete(ctx *fiber.Ctx, id uint64, userId uint) ReturnType {

	var db *gorm.DB = s.dbService.InitializeDB()
	var task models.Task;

	// err := db.First(&task, id).Error;
	_, err := s.FindTask(ctx, id, userId);
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


func (s *taskService) UpdateTask(ctx *fiber.Ctx, id uint64, updateTask dto.UpdateTask, userId uint) ReturnType {

	var db *gorm.DB = s.dbService.InitializeDB()
	var task models.Task;

	_, err := s.FindTask(ctx, id, userId);
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