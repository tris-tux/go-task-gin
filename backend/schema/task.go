package schema

type Task struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	ActionTime int    `json:"action_time"`
	CreateTime int    `json:"create_time"`
	UpdateTime int    `json:"update_time"`
	IsFinished bool   `json:"is_finished"`
}

type Detail struct {
	ID           int    `json:"id"`
	ObjectTaskFK int    `json:"object_task_fk"`
	ObjectName   string `json:"object_name" binding:"required"`
	IsFinished   bool   `json:"is_finished"`
}
