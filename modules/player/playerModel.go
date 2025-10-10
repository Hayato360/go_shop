package player

import "time"

type (
	PlayerProfile struct {
		Id       string `json:"_id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		CreateAt time.Time `json:"created_at"`
		UpdateAt time.Time `json:"updated_at"`
	}
)