package service

import (
	"golang/TodoList/model"
	"golang/TodoList/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" `
	Status  int    `json:"status" form:"status" ` //0 是未做，1 是已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title" binding:"required"`
	Content string `json:"content" form:"content" `
	Status  int    `json:"status" form:"status" `
}

type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

type DeleteTaskService struct {
}

// 创建备忘录服务
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Content:   service.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败" + err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
		Data:   serializer.BasicTask(task),
	}
}

// 展示备忘录服务
func (service *ShowTaskService) Show(uid uint, tid string) serializer.Response {
	var task model.Task
	err := model.DB.Where("id = ? AND uid = ?", tid, &uid).First(&task).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "备忘录不存在或无权访问",
			Error:  err.Error(),
		}
	}
	//查询成功
	return serializer.Response{
		Status: 200,
		Msg:    "查询成功",
		Data:   serializer.BuildTask(task),
	}
}

// 展示用户所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64 // 注意使用int64，与gorm的Count匹配
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	// 先查询总数（正确传入&count）
	model.DB.Model(&model.Task{}).Where("uid = ?", uid).Count(&count)
	if count == 0 {
		return serializer.Response{
			Status: 404,
			Msg:    "没有找到任何备忘录",
		}
	}
	// 再查询分页数据
	model.DB.Preload("User").Where("uid = ?", uid).
		Limit(service.PageSize).Offset(service.PageSize * (service.PageNum - 1)).
		Find(&tasks)
	// 返回正确的总数
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 更新备忘录
func (service *UpdateTaskService) Update(uid uint, tid string) serializer.Response {
	var task model.Task
	model.DB.Where("id = ? AND uid = ?", tid, &uid).First(&task)
	task.Title = service.Title
	task.Content = service.Content
	task.Status = service.Status
	model.DB.Save(&task)
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新成功",
	}
}
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	var count int64
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	// 正确的条件查询（同时过滤uid和搜索关键词）
	query := model.DB.Model(&model.Task{}).Where("uid = ?", uid)
	if service.Info != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%")
	}
	query.Count(&count) // 基于搜索条件统计总数
	// 空结果处理
	if count == 0 {
		return serializer.Response{
			Status: 404,
			Msg:    "没有找到匹配的备忘录",
		}
	}
	// 查询分页数据
	query.Preload("User").
		Limit(service.PageSize).Offset(service.PageSize * (service.PageNum - 1)).
		Find(&tasks)
	// 直接返回列表响应（避免嵌套）
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

// 删除备忘录
func (service *DeleteTaskService) Delete(uid uint, tid string) serializer.Response {
	var task model.Task
	model.DB.Where("id = ? AND uid = ?", tid, &uid).First(&task)
	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "删除成功",
		Data:   serializer.BasicTask(task),
	}
}
