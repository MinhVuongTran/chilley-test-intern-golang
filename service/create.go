package task

import (
	"net/http"

	"github.com/MinhVuongTran/chilley-test-intern-golang/model"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var task model.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"id": nil, "error": "Invalid request body"})
		return
	}

	task.ID = model.GetNextID()
	model.GetTasks()[task.ID] = task
	model.IncrementNextID()
	c.JSON(http.StatusCreated, gin.H{"id": task.ID, "error": nil})
}
