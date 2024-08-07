package test

import (
	"encoding/json"
	"server-lu/models"
)

// SendHealth 发送健康信息
func SendHealth(user *models.User) {
	resp, err := json.Marshal(&models.Health{
		ID:       user.ID,
		IsHealth: true,
	})

	if err != nil {
		return
	}
	user.Send <- resp
}
