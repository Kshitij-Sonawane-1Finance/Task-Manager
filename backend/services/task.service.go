package services

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/models"
	"gorm.io/gorm"
)


type TaskService interface {
	CreateTask(ctx *fiber.Ctx, task models.Task) error
	FindTask(ctx *fiber.Ctx, id uint64, userId uint) (models.Task, error)
	FindAllTasks(ctx *fiber.Ctx, userId uint) ([]models.Task, error)
	DeleteTask(ctx *fiber.Ctx, id uint64, userId uint) ReturnType
	SoftDelete(ctx *fiber.Ctx, id uint64, userId uint) ReturnType
	UpdateTask(ctx *fiber.Ctx, id uint64, updateTask dto.UpdateTask, userId uint) ReturnType
	SearchTask(ctx *fiber.Ctx) error
	FindTasks(ctx *fiber.Ctx) error
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

func (s *taskService) CreateTask(ctx *fiber.Ctx, task models.Task) error {

	var db *gorm.DB = s.dbService.InitializeDB()

	result := db.Create(&task)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	fmt.Println("Inserted Row");
	fmt.Println(result);

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task Created",
	})

}


func (s *taskService) FindTask(ctx *fiber.Ctx, id uint64, userId uint) (models.Task, error) {

	var db *gorm.DB = s.dbService.InitializeDB()
	var task models.Task;

	err := db.Where("user_id = ?", userId).Where("active = ?", true).First(&task, id).Error;

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return task, err;
	}

	return task, nil;

}

func (s *taskService) FindAllTasks(ctx *fiber.Ctx, userId uint) ([]models.Task, error) {

	var db *gorm.DB = s.dbService.InitializeDB()
	// var tasks []dto.Task;
	var tasks []models.Task;

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

	now := time.Now();
	updateTask.UpdatedAt = now;

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


func (s *taskService) SearchTask(ctx *fiber.Ctx) error {

	userID := uint(ctx.Locals("user_id").(uint64));
	searchData := "%" + ctx.Params("searchData") + "%";

	var db *gorm.DB = s.dbService.InitializeDB();
	var tasks []models.Task

	result := db.Where("title ilike ? AND user_id = ?", searchData, userID).Find(&tasks)
	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(tasks);

}


// func (s *taskService) FindTaskCount(ctx *fiber.Ctx) error {

// 	userID := uint(ctx.Locals("user_id").(uint64));
	
// 	var count int64;

// 	db := s.dbService.InitializeDB();
// 	db.Model(&models.Task{}).Count(&count).Find()

// }


func (s *taskService) FindTasks(ctx *fiber.Ctx) error {

	searchData := ctx.Query("searchData");
	status := ctx.Query("status");
	priority := ctx.Query("priority");
	sortBy := ctx.Query("sortBy");
	orderBy := ctx.Query("orderBy");
	pageParam := ctx.Query("page", "1");
	limitParam := ctx.Query("limit", "10");

	
	userID := uint(ctx.Locals("user_id").(uint64));


	var tasks []models.Task;
	db := s.dbService.InitializeDB();

	query := db.Model(&models.Task{}).Where("user_id = ?", userID).Where("active = ?", true);

	if searchData != "" {
		query = query.Where("title ILIKE ?", "%"+searchData+"%");
	}

	if status != "" {
		query = query.Where("status = ?", status);
	}

	if priority != "" {
		query = query.Where("priority = ?", priority);
	}

	// sort by query start
	validSortFields := map[string]bool{
		"start_date": true,
		"end_date": true,
	}

	validOrder := map[string]bool{
		"asc": true,
		"desc": true,
	}

	if validSortFields[sortBy] && validOrder[orderBy] {
		query = query.Order(fmt.Sprintf("%s %s", sortBy, orderBy));
	}

	countQry := query;
	var count int64
	countQry.Count(&count)


	page, err := strconv.Atoi(pageParam);
	if err != nil || page < 1 {
		page = 1;
	}

	limit, err := strconv.Atoi(limitParam);
	if err != nil || limit < 1 {
		limit = 5;
	}

	offset := (page - 1) * limit;
	query = query.Limit(limit).Offset(offset)
	// sort by query end

	result := query.Find(&tasks)
	if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"count": count,
		"tasks": tasks,
	});

}