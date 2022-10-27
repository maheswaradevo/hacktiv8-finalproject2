package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/auth"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/ping"
)

func Init(router *gin.Engine, db *sql.DB) {
	pingService := ping.ProvidePingService()
	pingHandler := ping.ProvidePingHandler(pingService, router)
	pingHandler.InitHandler()

	authService := auth.ProvideAuthService(db)
	authHandler := auth.ProvideAuthHandler(router, authService)
	authHandler.InitHandler()
}
