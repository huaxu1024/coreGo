package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)
var MysqlDb *sql.DB

var MysqlDbErr error

const (
	USERNAME = "root"
	PASSWORD = "root"
	HOST	 = "dev.docker.db01"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8"
)

// 初始化连接
func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)

	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	if MysqlDbErr != nil {
		log.Println("dbDNS: " + dbDSN)
		panic("数据源配置不正确：" + MysqlDbErr.Error())
	}
	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库连接失败： " + MysqlDbErr.Error())
	}
}