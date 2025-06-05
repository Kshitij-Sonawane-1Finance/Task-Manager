package services

import (
	"fmt"
	"net/http"
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


func hashPassword(password string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12);
	if err != nil {
		return string(hash), err;
	}

	return string(hash), nil;

}

func verifyPassword(dbPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password));

	return err == nil;

}


func generateJWT(userID uint) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	loadEnv.LoadEnv();
	jwtSecret := os.Getenv("JWT_SECRET")
	
	tokenString, err := token.SignedString([]byte(jwtSecret));
	if err != nil {
		return "", err;
	}

	return tokenString, nil;

}


func CreateUser(ctx *fiber.Ctx, user models.User) ReturnType {

	var db *gorm.DB = conn.InitializeDB();
	var dbUser models.User;

	db.First(&dbUser, "email = ?", user.Email);
	if dbUser.Email != "" {
		return ReturnType{
			StatusCode: http.StatusConflict,
			Message: "User Already Exists",
		}
	}

	hashedPassword, err := hashPassword(user.Password);
	if err != nil {
		return ReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	user.Password = hashedPassword

	result := db.Create(&user);

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
		Message: "Registered Successfully",
	};

}


func Login(ctx *fiber.Ctx, userLogin dto.UserLogin) dto.UserLoginReturnType {

	var db *gorm.DB = conn.InitializeDB();
	var dbUser models.User;

	db.First(&dbUser, "email = ?", userLogin.Email);
	if dbUser.Email == "" {
		return dto.UserLoginReturnType{
			StatusCode: http.StatusUnauthorized,
			Message: "Invalid Email",
		}
	}

	isPasswordMatch := verifyPassword(dbUser.Password, userLogin.Password)
	if !isPasswordMatch {
		return dto.UserLoginReturnType{
			StatusCode: http.StatusUnauthorized,
			Message: "Invalid Password",
		}
	}

	accessToken, err := generateJWT(dbUser.ID)
	if err != nil {
		return dto.UserLoginReturnType{
			StatusCode: http.StatusInternalServerError,
			Message: "Internal Server Error",
		}
	}

	return dto.UserLoginReturnType{
		StatusCode: http.StatusOK,
		Message: "Login Successful",
		AccessToken: accessToken,
	}

}