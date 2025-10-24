package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang/TodoList/pkg/utils"
	"golang/TodoList/service"
	"net/http"
)

// 创建备忘录
func CreateTask(c *gin.Context) {
	var creatTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&creatTask); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := creatTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

// 展示单个备忘录详细信息
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showTask); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := showTask.Show(claim.Id, c.Param("id"))
		c.JSON(200, res)
	}
}

// 展示全部备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		log.Error(err)
	}
}

// 更新备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&updateTask); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := updateTask.Update(claim.Id, c.Param("id"))
		c.JSON(200, res)
	}
}

// 查询
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	}
}

// 删除备忘录
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTask); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		res := deleteTask.Delete(claim.Id, c.Param("id"))
		c.JSON(200, res)
	}
}
