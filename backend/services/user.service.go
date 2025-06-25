package services

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	conn "github.com/kshitij/taskManager/connection"
	"github.com/kshitij/taskManager/dto"
	"github.com/kshitij/taskManager/loadEnv"
	"github.com/kshitij/taskManager/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Interface for abstaraction, just display the signature of the Methods which are available in the services
type UserService interface {
	CreateUser(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	generateJWT(userID uint) (string, error)
	verifyPassword(dbPassword string, password string) bool
	hashPassword(password string) (string, error)
	FetchUser(ctx *fiber.Ctx, userID uint) error
}

// Struct which will consist of all the methods
// and additinal dependencies
type userService struct {
	dbService conn.DBService
	loadEnvService loadEnv.LoadEnvService
}


// constructor which returns all the methods
func NewUserService(dbService conn.DBService, loadEnvService loadEnv.LoadEnvService) UserService {
	return &userService{dbService, loadEnvService}
}


type fetchUserStruct struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


func (s *userService) hashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12);
	if err != nil {
		return string(hash), err;
	}

	return string(hash), nil;

}

func (s *userService) verifyPassword(dbPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password));

	return err == nil;

}


func (s *userService) generateJWT(userID uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	s.loadEnvService.LoadEnv();
	jwtSecret := os.Getenv("JWT_SECRET")
	
	tokenString, err := token.SignedString([]byte(jwtSecret));
	if err != nil {
		return "", err;
	}

	return tokenString, nil;

}

// this s *userService works like { this of javascript }
func (s *userService) CreateUser(ctx *fiber.Ctx) error {

	var user models.User;
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	var db *gorm.DB = s.dbService.InitializeDB();
	var dbUser models.User;

	db.First(&dbUser, "email = ?", user.Email);
	if dbUser.Email != "" {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User Already Exists",
		})
	}

	hashedPassword, err := s.hashPassword(user.Password);
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Hashing Password",
		})
	}

	user.Password = hashedPassword

	result := db.Create(&user);

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error Inserting into Database",
		})
	}

	fmt.Println("Inserted Row");
	fmt.Println(result);
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registered Successfully",
	})

}


func (s *userService) Login(ctx *fiber.Ctx) error {

	var userLogin dto.UserLogin;

	err := ctx.BodyParser(&userLogin);
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Request Body",
		})
	}

	var db *gorm.DB = s.dbService.InitializeDB();
	var dbUser models.User;

	db.First(&dbUser, "email = ?", userLogin.Email);
	if dbUser.Email == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Email",
		})
	}

	isPasswordMatch := s.verifyPassword(dbUser.Password, userLogin.Password)
	if !isPasswordMatch {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	accessToken, err := s.generateJWT(dbUser.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login Successful",
		"accessToken": accessToken,
	})

}


func (s *userService) FetchUser(ctx *fiber.Ctx, userID uint) error {

	var db *gorm.DB = s.dbService.InitializeDB();
	var user fetchUserStruct;

	res := db.Model(&models.User{}).Select("id", "name", "email").Where("id = ?", userID).Scan(&user)
	if res.Error != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": 404,
			"message": "User Not Found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user);
}