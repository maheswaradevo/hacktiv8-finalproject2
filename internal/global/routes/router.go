package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/ping"
)

func Init(router *gin.Engine) {
	pingService := ping.ProvidePingService()
	pingHandler := ping.ProvidePingHandler(pingService, router)
	pingHandler.InitHandler()
}
