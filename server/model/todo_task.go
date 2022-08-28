package model

type ToDoTask struct {
	BaseModel
	Task_Description string `json:"task_description"`
	Is_Finished bool `json:"is_finished" gorm:"default:false"`
	Remind_Time int64 `json:"remind_time"`
	Is_Delay bool `json:"is_delay" gorm:"default:false"`
}

func (entity *ToDoTask) TableName () string {
	return "todo_tasks"
}