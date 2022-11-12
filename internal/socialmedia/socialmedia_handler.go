package socialmedia

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	protectedRoute.POST("/social-medias", scmd.createSocialMedia)
	protectedRoute.GET("/social-medias", scmd.viewSocialMedia)
	protectedRoute.PUT("/social-medias/:socialMediaID", scmd.updateSocialMedia)
	protectedRoute.DELETE("social-medias/:socialMediaID", scmd.deleteSocialMedia)
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

	res, err := scmd.as.CreateSocialMedia(c, data, userID)
	if err != nil {
		log.Printf("[createSocialMedia] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", res)
	c.JSON(http.StatusCreated, response)
}

func (scmd *SocialMediaHandler) viewSocialMedia(c *gin.Context) {
	res, err := scmd.as.ViewSocialMedia(c)
	if err != nil {
		log.Printf("[viewSocialMedia] failed to view social media, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (scmd *SocialMediaHandler) updateSocialMedia(c *gin.Context) {
	data := &dto.EditSocialMediaRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[updateSocialMedia] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	socialMediaID := c.Param("socialMediaID")
	socialMediaIDConv, _ := strconv.ParseUint(socialMediaID, 10, 64)

	res, err := scmd.as.UpdateSocialMedia(c, data, socialMediaIDConv, userID)
	if err != nil {
		log.Printf("[UpdateSocialMedia] failed to update social media, id: %v, err: %v", socialMediaIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (scmd *SocialMediaHandler) deleteSocialMedia(c *gin.Context) {
	socialMediaID := c.Param("socialMediaID")
	socialMediaIDConv, _ := strconv.ParseUint(socialMediaID, 10, 64)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := scmd.as.DeleteSocialMedia(c, socialMediaIDConv, userID)
	if err != nil {
		log.Printf("[deleteSocialMedia] failed to delete social media, id: %v, err: %v", socialMediaIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}