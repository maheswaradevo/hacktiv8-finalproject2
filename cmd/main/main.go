package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/routes"
)

func main() {
	config.Init()
	cfg := config.GetConfig()

	r := gin.Default()
	routes.Init(r)
	r.Run(cfg.PORT)
}
