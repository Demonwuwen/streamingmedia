package dbops

import (
	"database/sql"
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB
var err error

func init() {
	config, iniErr := ini.Load("api/conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Failed to read file :%v", iniErr)
		os.Exit(1)
	}

	ip := config.Section("Mysql").Key("ip").String()
	port := config.Section("Mysql").Key("port").String()
	user := config.Section("Mysql").Key("user").String()
	password := config.Section("Mysql").Key("password").String()
	database := config.Section("Mysql").Key("database").String()

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, password, ip, port, database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})
}

func openConn() *sql.DB {

	return nil
}

func AddUserCredential(loginName string, pwd string) error {
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	return "", nil
}
