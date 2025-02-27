package main

import (
	"time"

	task "github.com/MinhVuongTran/chilley-test-intern-golang/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")
	{
		api.GET("/tasks", task.GetTasks)
		api.POST("/tasks", task.CreateTask)
		api.PUT("/tasks/:id", task.UpdateTask)
		api.DELETE("/tasks/:id", task.DeleteTask)
	}

	router.Run(":8080")
}
