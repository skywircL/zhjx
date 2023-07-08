/**
 * @Author: lrc
 * @Date: 2023/5/16-22:09
 * @Desc:
 **/

package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"videoStream/model"
)

var (
	DB       *gorm.DB
	user     = "root"
	password = "123456"
	//password = "sjk123456"
	//host     = "redrock-mariadb-primary"
	host   = "127.0.0.1"
	port   = "3306"
	dbname = "videostream"
)

func MySQLInit() {
	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	DB = db

	AutoMigrateAll(model.DeviceIp{}, model.FfmpegStatus{}, model.SuperResolution{}, model.Video{}, model.AbImgRecord{}, model.User{})
}

func AutoMigrateAll(is ...interface{}) {
	var err error
	for _, i := range is {
		err = DB.AutoMigrate(&i)
		if err != nil {
			panic(err)
		}
	}
}
