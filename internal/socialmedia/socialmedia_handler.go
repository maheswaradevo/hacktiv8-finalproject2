package socialmedia

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

type SocialMediaHandler struct {
	r  *gin.Engine
	as SocialMediaService
}

func ProvideSocialMediaHandler(r *gin.Engine, as SocialMediaService) *SocialMediaHandler {
	return &SocialMediaHandler{
		r:  r,
		as: as,
	}
}

func (scmd *SocialMediaHandler) InitHandler() {
	protectedRoute := scmd.r.Group(constant.ROOT_API_PATH)
	protectedRoute.Use(middleware.AuthMiddleware())
	protectedRoute.POST("/social-media", scmd.createSocialMedia)
}

func (scmd *SocialMediaHandler) createSocialMedia(c *gin.Context) {
	data := &dto.CreateSocialMediaRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[createSocialMedia] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	err = scmd.as.CreateSocialMedia(c, data, userID)
	if err != nil {
		log.Printf("[createSocialMedia] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", data)
	c.JSON(http.StatusCreated, response)
}
