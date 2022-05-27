package schema

type DetailResponse struct {
	ObjectName string `json:"object_name"`
	IsFinished bool   `json:"is_finished"`
}

type TaskResponse struct {
	ID         int              `json:"id"`
	Title      string           `json:"title"`
	ActionTime int              `json:"action_time"`
	CreateTime int              `json:"create_time"`
	UpdateTime int              `json:"update_time"`
	IsFinished bool             `json:"is_finished"`
	ObjectList []DetailResponse `json:"obejct_list"`
}
