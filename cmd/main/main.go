package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello world")
	r := gin.Default()
	r.Run()
}
