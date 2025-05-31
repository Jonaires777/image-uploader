package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jonaires777/image-uploader/db"
	"github.com/Jonaires777/image-uploader/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
