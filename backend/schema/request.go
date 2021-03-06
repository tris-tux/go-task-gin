package schema

type TaskAddRequest struct {
	Title         string   `json:"title" binding:"required"`
	ActionTime    int      `json:"action_time" binding:"required"`
	ObjectiveList []string `json:"objective_list" binding:"required"`
}

type TaskUpdateRequest struct {
	Title         string           `json:"title" binding:"required"`
	ObjectiveList []DetailResponse `json:"objective_list" binding:"required"`
}
