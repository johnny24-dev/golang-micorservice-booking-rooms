package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var ConWrite *gorm.DB
var ConRead *gorm.DB

func DbWrite() *gorm.DB {
	mysqlDsn := "root:minhvipoi2000@tcp(127.0.0.1:3306)/doan?charset=utf8mb4&parseTime=True&loc=Local"
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

	return orm
}

func DbRead() *gorm.DB {
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

	return orm
}

func syncAccount(dbr, dbw *gorm.DB) {
	var data []map[string]interface{}
	query := dbr.Table("payments").Select("*").Find(&data)
	if err := query.Error; err != nil {
		fmt.Println(err)
	}

	for _, val := range data {
		queryI := dbw.Table("payments").Create(val)
		if err := queryI.Error; err != nil {
			fmt.Println(err, val)
			return
		}
	}
	//
	//queryI := dbw.Table("users").Create(data)
	//if err := queryI.Error; err != nil {
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println("OK")
}
