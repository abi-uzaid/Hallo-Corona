package main

import (
	"fmt"
	"hallo-corona/database"
	"hallo-corona/pkg/mysql"
	"hallo-corona/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()

	database.RunMigration()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PATCH", "POST", "DELETE"},
		AllowHeaders: []string{"x-Requested-With", "Content-Type", "Authorization"},
	}))

	routes.RouteInit(r.Group("/api/v1"))

	r.Static("/uploads", "./uploads")

	r.Use(gin.Logger())

	var port = os.Getenv("PORT")

	fmt.Println("Server running localhost:" + port)

	r.Run(":" + port)

	// fmt.Println("Server Started")
	// http.ListenAndServe("localhost:5000", r)
}
