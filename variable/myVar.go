package variable

import "time"

type Updated struct {
	UpdatedAt time.Time `json:"updated_at"`
}

type Created struct {
	CreatedAt time.Time `json:"created_at"`
}

type Id struct {
	Id uint `json:"id"`
}

type UserId struct {
	UserId uint `json:"user_id"`
}
