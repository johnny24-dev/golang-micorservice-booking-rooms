package mysql

import (
	"fmt"
	"gorm.io/gorm"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

var Conn *gorm.DB

func init() {
	mysqlDsn := "dbhuan:0866444202@tcp(103.173.254.82:3306)/doandb?charset=utf8mb4&parseTime=True&loc=Local"
	orm, err := gorm.Open(mysql.Open(mysqlDsn), &gorm.Config{})
	if nil != err {
		fmt.Println(err)
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	fmt.Println(fmt.Sprintf("MySQL connection established"))

	Conn = orm
}
