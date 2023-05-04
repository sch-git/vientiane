package models

import "time"

const contentTableName = "vientiane_content"

type Content struct {
	Id          int64     `json:"id"`
	ContentInfo string    `json:"content"`
	Ct          time.Time `json:"ct"`
	Ut          time.Time `json:"ut"`
}

func (m *Content) TableName() string {
	return contentTableName
}
