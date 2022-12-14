package photos

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

type photoHandler struct {
	r  *gin.Engine
	ps PhotoService
}

func ProvidePhotoHandler(r *gin.Engine, ps PhotoService) *photoHandler {
	return &photoHandler{r: r, ps: ps}
}

func (p *photoHandler) InitHandler() {
	photoRoute := p.r.Group(constant.ROOT_API_PATH)
	photoRoute.Use(middleware.AuthMiddleware())
	photoRoute.POST("/photos", p.postPhoto)
	photoRoute.GET("/photos", p.viewPhoto)
	photoRoute.PUT("/photos/:photoID", p.updatePhoto)
	photoRoute.DELETE("photos/:photoID", p.deletePhoto)
}

func (p *photoHandler) postPhoto(c *gin.Context) {
	data := &dto.PostPhotoRequest{}

	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[postPhoto] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	res, err := p.ps.PostPhoto(c, data, userID)
	if err != nil {
		log.Printf("[postPhoto] failed to post photo, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", res)
	c.JSON(http.StatusCreated, response)
}

func (p *photoHandler) viewPhoto(c *gin.Context) {
	res, err := p.ps.ViewPhoto(c)
	if err != nil {
		log.Printf("[viewPhoto] failed to view photo, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (p *photoHandler) updatePhoto(c *gin.Context) {
	data := &dto.EditPhotoRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(data)
	if err != nil {
		log.Printf("[updatePhoto] failed to parse json data: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	photoID := c.Param("photoID")
	photoIDConv, _ := strconv.ParseUint(photoID, 10, 64)

	res, err := p.ps.UpdatePhoto(c, data, photoIDConv, userID)
	if err != nil {
		log.Printf("[updatePhoto] failed to update photo, id: %v, err: %v", photoIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (p *photoHandler) deletePhoto(c *gin.Context) {
	photoID := c.Param("photoID")
	photoIDConv, _ := strconv.ParseUint(photoID, 10, 64)
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := p.ps.DeletePhoto(c, photoIDConv, userID)
	if err != nil {
		log.Printf("[deletePhoto] failed to delete photo, id: %v, err: %v", photoIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}
