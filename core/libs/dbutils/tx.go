package dbutils

import (
	"myapp/core/services/database"

	"gorm.io/gorm"
)

// Transaction 在一个事务中执行函数
func Transaction(conn *gorm.DB, f func(tx *gorm.DB) error) error {
	if conn == nil {
		conn = database.Master()
	}
	tx := conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	err := f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
