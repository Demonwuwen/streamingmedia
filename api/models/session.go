package models

type Sessions struct {
	SessionId int    `gorm:"type:tinytext;primaryKey;not null"`
	TTL       string `gorm:"type:tinyText"`
	LoginName string `gorm:"type:varchar(64)"`
}
