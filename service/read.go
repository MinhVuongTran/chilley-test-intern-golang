package task

import (
	"net/http"
	"sort"

	"github.com/MinhVuongTran/chilley-test-intern-golang/model"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks := model.GetTasks()
	taskList := make([]model.Task, 0, len(tasks))

	for _, task := range tasks {
		taskList = append(taskList, task)
	}

	sort.Slice(taskList, func(i, j int) bool {
		return taskList[i].ID < taskList[j].ID
	})

	c.JSON(http.StatusOK, taskList)
}
