package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/TodoList/serializer"
	"golang/TodoList/service"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	// 1. 执行参数绑定和校验
	if err := c.ShouldBind(&userRegister); err != nil {
		// 校验失败：直接返回错误信息（使用统一的Response格式）
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    "参数校验失败：" + err.Error(), // 明确提示错误原因，如"user_name is required"
		})
		return
	}

	// 2. 校验通过，执行注册逻辑
	res := userRegister.Register()
	c.JSON(res.Status, res)
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	fmt.Println(userLogin)
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(400, serializer.Response{
			Status: 400,
			Msg:    "参数校验失败：" + err.Error(),
		})
		return
	}

	res := userLogin.Login()
	c.JSON(res.Status, res)
}
