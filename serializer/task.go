package serializer

import "golang/TodoList/model"

type Task struct {
	ID      uint   `json:"id" example:"1"`
	Title   string `json:"title" example:"Title"`
	Content string `json:"content" example:"Content"`
	//View      uint   `json:"view" example:"1"` //浏览量
	Status    int   `json:"status" example:"0"`
	CreatedAt int64 `json:"created_at"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}
func BasicTask(item model.Task) Task {
	return Task{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
	}
}

func BuildTasks(items []model.Task) []Task {
	// 初始化一个空的 Task 切片，容量与输入切片一致
	tasks := make([]Task, 0, len(items))
	// 遍历所有任务，逐个转换并添加到结果切片
	for _, item := range items {
		tasks = append(tasks, BuildTask(item))
	}
	return tasks
}
