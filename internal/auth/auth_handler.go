package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
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
	authApi.POST("/users", auth.registerUser)
}

func (auth *AuthHandler) registerUser(c *gin.Context) {
	data := &dto.UserRegistrationRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[registerUser] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	err = auth.as.RegisterUser(c, data)
	if err != nil {
		log.Printf("[registerUser] failed to register a user: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", data)
	c.JSON(http.StatusCreated, response)
}
