package dto

type UserLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserLoginReturnType struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
	AccessToken string `json:"accessToken"`
}