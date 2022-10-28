package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/middleware"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/utils"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
)

type AuthHandler struct {
	r  *gin.Engine
	as AuthService
}

func ProvideAuthHandler(r *gin.Engine, as AuthService) *AuthHandler {
	return &AuthHandler{
		r:  r,
		as: as,
	}
}

func (auth *AuthHandler) InitHandler() {
	authApi := auth.r.Group(constant.ROOT_API_PATH)
	authApi.POST("/users/register", auth.registerUser)
	authApi.POST("/users/login", auth.loginUser)

	protectedRoute := auth.r.Group(constant.ROOT_API_PATH)
	protectedRoute.Use(middleware.AuthMiddleware())
	protectedRoute.PUT("/users", auth.updateUser)
	protectedRoute.DELETE("/users", auth.deleteUser)
}

func (auth *AuthHandler) registerUser(c *gin.Context) {
	data := &dto.UserRegistrationRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[registerUser] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	err = auth.as.RegisterUser(c, data)
	if err != nil {
		log.Printf("[registerUser] failed to register a user: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", data)
	c.JSON(http.StatusCreated, response)
}

func (auth *AuthHandler) loginUser(c *gin.Context) {
	data := &dto.UserSignInRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[loginUser] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	token, err := auth.as.LoginUser(c, data)
	if err != nil {
		log.Printf("[loginUser] user failed to login, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", token)
	c.JSON(http.StatusOK, response)
}

func (auth *AuthHandler) updateUser(c *gin.Context) {
	data := &dto.UserEditProfileRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[loginUser] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := auth.as.UpdateUser(c, data, userID)
	if err != nil {
		log.Printf("[updateUser] failed to update user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", res)
	c.JSON(http.StatusCreated, response)
}

func (auth *AuthHandler) deleteUser(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := auth.as.DeleteUser(c, userID)
	if err != nil {
		log.Printf("[deleteUser] failed to delete user, id: %v, err: %v", userID, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
	}

	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}
