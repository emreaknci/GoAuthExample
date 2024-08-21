package app

import (
	"fmt"
	"os"
	"path/filepath"

	route "github.com/emreaknci/goauthexample/api/router"
	"github.com/emreaknci/goauthexample/internal/ioc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	loadEnv()

	container, err := ioc.BuildContainer()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("http://localhost:%v", os.Getenv("WEB_APP_PORT"))},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	
	route.RegisterRoutes(container, router)

	router.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		panic("Error loading .env file")
	}
}
