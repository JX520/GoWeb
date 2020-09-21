package initDB

import (
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

//var Db *sql.DB
var Db *gorm.DB
func init()  {
	var err error
	//Db, err = sql.Open("mysql", "test:2100M4105av@tcp(127.0.0.1:3306)/go")
	Db, err = gorm.Open("mysql", "test:2100M4105av@tcp(127.0.0.1:3306)/go")
	if err != nil{
		log.Println("err：", err.Error())
	} else {
		log.Println("数据库连接成功！")
	}
	//Db.SetMaxOpenConns(10) //最大连接数
	//Db.SetMaxIdleConns(10) //最大空闲连接数
}

