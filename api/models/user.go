package models

type User struct {
	Id        uint   `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	LoginName string `json:"login_name" gorm:"column:login_name;unique_key"`
	Pwd       string `json:"pwd" gorm:"type:text"`
}

func (User) TableName() string {
	return "users"
}
