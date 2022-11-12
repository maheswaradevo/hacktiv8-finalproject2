package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/comment"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/photos"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/ping"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/socialmedia"
)

func Init(router *gin.Engine, db *sql.DB) {
	pingService := ping.ProvidePingService()
	pingHandler := ping.ProvidePingHandler(pingService, router)
	pingHandler.InitHandler()

	authService := auth.ProvideAuthService(db)
	authHandler := auth.ProvideAuthHandler(router, authService)
	authHandler.InitHandler()

	scmdService := socialmedia.ProvideSocialMediaService(db)
	scmdHandler := socialmedia.ProvideSocialMediaHandler(router, scmdService)
	scmdHandler.InitHandler()

	photoService := photos.ProvidePhotoService(db)
	photoHandler := photos.ProvidePhotoHandler(router, photoService)
	photoHandler.InitHandler()

	cmtService := comment.ProvideCommentService(db)
	cmtHandler := comment.ProvideCommentHandler(comment.CommentHandlerParams{
		R:  router,
		CS: cmtService,
	})
	cmtHandler.InitHandler()
}
