package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/routes"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/database"
)

func main() {
	config.Init()
	cfg := config.GetConfig()

	db := database.GetDatabase()
	r := gin.Default()

	routes.Init(r, db)
	port := fmt.Sprintf("%s:%s", "0.0.0.0", cfg.PORT)
	r.Run(port)
}
