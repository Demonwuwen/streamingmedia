package models

import "time"

type VideoInfo struct {
	Id           string    `json:"id" gorm:"type:VARCHAR(64);primaryKey;not null"`
	AuthorID     uint      `json:"author_id" gorm:"uint"`
	Name         string    `json:"name" gorm:"type:text"`
	DisplayCtime string    `json:"display_ctime" gorm:"column:display_ctime;type:text"`
	CreateTime   time.Time `gorm:"autoCreateTime"`
}

func (VideoInfo) TableName() string {
	return "video_info"
}
