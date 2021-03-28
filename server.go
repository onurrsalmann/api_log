package main

import (
	"github.com/onurrsalmann/api_log/config"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("/")
	//{
	//	authRoutes.POST("/login", authController.Login)
	//	authRoutes.POST("/register", authController.Register)
	//}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}