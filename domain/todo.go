package domain

import "time"

type TodoReq struct {
	Task string `json:"task"`
}

type TodoRes struct {
	Id          uint      `json:"id"`
	Task        string    `json:"task"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
