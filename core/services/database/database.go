package database

import (
	"fmt"
	"log"
	"myapp/core/config"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var connections map[string]*gorm.DB

func init() {
	// 在初始化的时候把所有数据库连接上，并将句柄保存到connections，使用的时候可以直接按identify获取
	for _, identify := range []string{"master"} {
		host := config.GetString("databases.%s.host", identify)
		dbname := config.GetString("databases.%s.dbname", identify)
		user := config.GetString("databases.%s.user", identify)
		password := config.GetString("databases.%s.password", identify)
		port := config.GetInt("databases.%s.port", identify)
		connections[identify] = connect(host, dbname, user, password, port)
	}
}

// 返回默认主数据库的连接
func Master() *gorm.DB {
	return Conn("master")
}

// 返回指定的数据库连接
func Conn(identify string) *gorm.DB {
	return connections[identify]
}

// 连接指定数据库并返回句柄
func connect(host, dbname, user, password string, port int) *gorm.DB {
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
	return db
}
