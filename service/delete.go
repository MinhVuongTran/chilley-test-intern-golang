package task

import (
	"net/http"
	"strconv"

	"github.com/MinhVuongTran/chilley-test-intern-golang/model"
	"github.com/gin-gonic/gin"
)

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	tasks := model.GetTasks()
	_, exists := tasks[id]

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	delete(tasks, id)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
