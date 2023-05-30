package model

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"log"
)

var (
	DB *gorm.DB

	//username string = "root"
	//password string = "PetWorld2022"
	//dbName   string = "spiders"
	//dbHost   string = "119.91.152.112"
	//dbPort   int32  = 3306

	username string = "root"
	password string = "a123456"
	dbName   string = "spiders"
	dbHost   string = "localhost"
	dbPort   int32  = 3306
)

func init() {
	var err error
	dbArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, dbHost, dbPort, dbName)
	//DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName))
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf(" gorm.Open.err: %v", err)
	}
	//
	//DB.SingularTable(true)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return "sp_" + defaultTableName
	//}

}
