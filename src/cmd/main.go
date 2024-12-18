package main

import (
	_ "github.com/betterreads/docs"
	"github.com/betterreads/internal/application"
)

// @title BetterReads API
// @version 1.0
// @description This is a  server for Swagger with Gin.
// @host localhost:8080
// @BasePath /

func main() {
	r := application.NewRouter(":8080")
	r.Run()
}
