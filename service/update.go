package task

import (
	"net/http"
	"strconv"

	"github.com/MinhVuongTran/chilley-test-intern-golang/model"
	"github.com/gin-gonic/gin"
)

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var updateData struct {
		Completed bool `json:"completed"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tasks := model.GetTasks()
	task, exists := tasks[id]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	task.Completed = updateData.Completed
	tasks[id] = task

	c.JSON(http.StatusOK, gin.H{"message": "Updated task successfully"})
}
