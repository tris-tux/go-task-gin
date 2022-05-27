package schema

import "encoding/json"

type TaskAddRequest struct {
	Title         string      `json:"title" binding:"required"`
	ActionTime    json.Number `json:"action_time" binding:"required,number"`
	ObjectiveList []string    `json:"objective_list"`
}

type TaskUpdateRequest struct {
	Title         string           `json:"title" binding:"required"`
	ObjectiveList []DetailResponse `json:"objective_list"`
}
