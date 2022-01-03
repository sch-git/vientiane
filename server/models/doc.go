package models

import "time"

const docTableName = "vientiane_account"

type Doc struct {
	Id         int64     `json:"id"`
	Content    string    `json:"content"`
	CategoryId int64     `json:"category_id"`
	Author     string    `json:"author"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (m *Doc) TableName() string {
	return docTableName
}
