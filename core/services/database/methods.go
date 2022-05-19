package database

import (
	"strings"

	"gorm.io/gorm"
)

type Where map[string]interface{}
type SqlOrder map[string]string

func (conn *Connection) queryBuilder(wheres *Where, orders *SqlOrder) *gorm.DB {
	db := conn.db
	if wheres != nil {
		for k, v := range *wheres {
			if strings.Count(k, "?") > 1 {
				if arrayOfInterfaces, ok := v.([]interface{}); ok {
					db = db.Where(k, arrayOfInterfaces...)
				}
				continue
			}
			db = db.Where(k, v)
		}
	}
	if orders != nil {
		for k, v := range *orders {
			db = db.Order(k + " " + v)
		}
	}
	return db
}

// GetDB 获取gorm.DB对象
func (conn *Connection) GetDB() *gorm.DB {
	return conn.db
}

// Count 查询
func (conn *Connection) Count(model interface{}, wheres *Where) int64 {
	var c int64
	conn.queryBuilder(wheres, nil).Model(model).Count(&c)
	return c
}

// Take 查询
func (conn *Connection) Take(dest interface{}, wheres *Where, scopes ...func(*gorm.DB) *gorm.DB) error {
	q := conn.queryBuilder(wheres, nil)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	return q.Take(dest).Error
}

// TakeByID 查询指定ID的一条记录
func (conn *Connection) TakeByID(dest interface{}, id interface{}, scopes ...func(*gorm.DB) *gorm.DB) error {
	q := conn.queryBuilder(&Where{"id=?": id}, nil)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	return q.Take(dest).Error
}

// Find 查询列表
func (conn *Connection) Find(dest interface{}, wheres *Where, orderBy *SqlOrder, scopes ...func(*gorm.DB) *gorm.DB) error {
	q := conn.queryBuilder(wheres, orderBy)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	return q.Find(dest).Error
}

// Exists 查询指定条件的行是否存在
func (conn *Connection) Exists(table string, wheres *Where) bool {
	q := conn.queryBuilder(wheres, nil).Table(table).Limit(1)
	var ret struct {
		N uint8 `json:"n"`
	}
	conn.db.Raw("select EXISTS(?) as n", q).Scan(&ret)
	return ret.N > 0
}

// Transaction 在一个事务中执行函数
func (conn *Connection) Transaction(f func(tx *gorm.DB) error) error {
	tx := conn.db.Begin()
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
