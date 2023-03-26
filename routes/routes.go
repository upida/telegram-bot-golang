package routes

import (
	"bot/controllers/api/task"
	"bot/controllers/telegram"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// API
	r.GET("/api/tasks", task.FindTasks)
	r.POST("/api/tasks", task.CreateTask)
	r.GET("/api/tasks/:id", task.FindTask)
	r.PATCH("/api/tasks/:id", task.UpdateTask)
	r.DELETE("api/tasks/:id", task.DeleteTask)

	// Telegram
	r.GET("/telegram", telegram.Telegram)
	return r
}
