package comment

import (
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

type commentHandler struct {
	r  *gin.Engine
	cs CommentService
}

type CommentHandlerParams struct {
	R  *gin.Engine
	CS CommentService
}

func ProvideCommentHandler(params CommentHandlerParams) *commentHandler {
	return &commentHandler{
		r:  params.R,
		cs: params.CS,
	}
}

func (cmth *commentHandler) InitHandler() {
	protectedRoute := cmth.r.Group(constant.ROOT_API_PATH)
	protectedRoute.Use(middleware.AuthMiddleware())
	protectedRoute.POST("/comments", cmth.createComment)
	protectedRoute.GET("/comments", cmth.viewComment)
	protectedRoute.PUT("/comments/:commentId", cmth.updateComment)
	protectedRoute.DELETE("/comments/:commentId", cmth.deleteComment)
}

func (cmth *commentHandler) createComment(c *gin.Context) {
	var requestBody dto.CreateCommentRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))

	res, err := cmth.cs.CreateComment(c, &requestBody, userID)
	if err != nil {
		log.Printf("[createComment] failed to create user, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", res)
	c.JSON(http.StatusCreated, response)
}

func (cmth *commentHandler) viewComment(c *gin.Context) {
	res, err := cmth.cs.ViewComment(c)
	if err != nil {
		log.Printf("[viewComment] failed to view comment, err: %v", err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (cmth *commentHandler) updateComment(c *gin.Context) {
	data := dto.EditCommentRequest{}

	err := c.BindJSON(&data)
	if err != nil {
		errResponse := utils.NewErrorResponse(c.Writer, errors.ErrInvalidRequestBody)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	commentID := c.Param("commentId")
	commentIDConv, _ := strconv.ParseUint(commentID, 10, 64)

	res, err := cmth.cs.UpdateComment(c, commentIDConv, userID, &data)
	if err != nil {
		log.Printf("[updateComment] failed to update comment, id: %v, err: %v", commentIDConv, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}

func (cmth *commentHandler) deleteComment(c *gin.Context) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["user_id"].(float64))
	commentID := c.Param("commentId")
	commentIDConv, _ := strconv.ParseUint(commentID, 10, 64)

	res, err := cmth.cs.DeleteComment(c, commentIDConv, userID)
	if err != nil {
		log.Printf("[deleteComment] failed to delete comment, id: %v, err: %v", commentID, err)
		errResponse := utils.NewErrorResponse(c.Writer, err)
		c.JSON(errResponse.Error.Code, errResponse)
		return
	}
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusCreated, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}
