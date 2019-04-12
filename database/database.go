package database

import (
	"github.com/cocoagaurav/devops/models"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)
var db *gorm.DB


func Opendatabase() {
	var err error
	db, err = gorm.Open("mysql","root:password123@tcp(mysql:3306)/devop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		glog.Errorf("not able to connect to database err:%v", err)
		time.Sleep(5 * time.Second)
		Opendatabase()
	}else {
		db.DB().SetMaxIdleConns(10)
		db.DB().SetMaxOpenConns(100)
		glog.Infof("Connected to Database")
	}
}

func GetDB()*gorm.DB{
	return db
}


func InitDB(){
	db.AutoMigrate(&models.Registration{}, &models.Tweet{})

}

func CloseDb(){
	db.Close()
}