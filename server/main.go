package main

import (
	"log"
	"server-ssm/db"
	"server-ssm/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func init() {
	LoadConfig()
	db.ConnectDB()
	db.SyncDB()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
	
	routes.PostRoute(r)
	routes.UserRoute(r)
	routes.AuthRoute(r)
	
	r.Run(":8080")
}
