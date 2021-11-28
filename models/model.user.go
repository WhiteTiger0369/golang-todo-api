package models

import "time"

type ModelUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreateDate time.Time
	UpdateDate time.Time
}
