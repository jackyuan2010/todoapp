package model

type ToDoTask struct {
	BaseModel
	task_description string `json:"task_description"`
	is_finished bool `json:"is_finished"`
	remind_time int64 `json:"remind_time"`
	is_delay bool `json:"is_delay"`
}

func (entity *ToDoItem) TableName () string {
	return "todo_tasks"
}