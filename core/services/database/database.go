package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	db *gorm.DB
}

// 连接指定数据库并返回句柄
func Connect(host, dbname, user, password string, port int) *Connection {
	var db *gorm.DB
	connectString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	logHandler := logger.New(log.New(os.Stdout, "\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Error,
		Colorful:      true,
	})

	var err error
	db, err = gorm.Open(mysql.Open(connectString), &gorm.Config{
		AllowGlobalUpdate: false, //禁止全表更新和删除
		Logger:            logHandler,
		CreateBatchSize:   100,
	})

	if err != nil {
		log.Fatalln(err, connectString)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetConnMaxLifetime(time.Second * 14400)
	sqlDB.SetMaxOpenConns(128)
	sqlDB.SetMaxIdleConns(10)
	// 强制屏蔽日志中的record not found
	db.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", func(d *gorm.DB) {
		d.Statement.RaiseErrorOnNotFound = false
	})
	return &Connection{
		db: db,
	}
}
