package models

import "time"

type Comment struct {
	Id       string `json:"id" gorm:"primaryKey;not null"`
	VideoId  string `json:"video_id" gorm:"column:video_id"`
	AuthorId uint
	Content  string
	Time     time.Time `gorm:"autoCreateTime"`
}

func (Comment) TableName() string {
	return "comments"
}
