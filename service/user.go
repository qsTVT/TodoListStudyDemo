package service

import (
	"github.com/jinzhu/gorm"
	"golang/TodoList/model"
	"golang/TodoList/pkg/utils"
	"golang/TodoList/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}

// 注册
func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有这个人了，不需要再注册",
		}
	}
	user.UserName = service.UserName
	//密码加密
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	result := model.DB.Create(&user)
	if result.Error != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "注册用户成功",
	}
}

// 登陆
func (service *UserService) Login() serializer.Response {
	var user model.User
	//先找数据库有没有这个人
	err := model.DB.Where("user_name = ?", service.UserName).First(&user).Error
	if err != nil {
		// 判断是否为"记录未找到"错误（用户不存在）
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		//是其他因素错误
		return serializer.Response{
			Status: 500,
			Msg:    "数据库查询错误" + err.Error(),
		}
	}
	//验证密码
	if user.CheckPassword(service.Password) == false {
		return serializer.Response{
			Status: 401,
			Msg:    "密码错误",
		}
	}
	//登陆成功,发一个token，为了其他功能需要身份验证所给前端存储的。
	//创建一个备忘录，就需要token,不然都不知道是谁创建的。
	token, err := utils.GenerateToken(user.ID, service.UserName)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登陆成功",
	}
}
