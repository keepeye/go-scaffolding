package dbutils

import (
	"myapp/core/services/database"
	"strings"

	"gorm.io/gorm"
)

type Where map[string]interface{}
type SqlOrder map[string]string

func queryBuilder(query *gorm.DB, wheres *Where, orders *SqlOrder) *gorm.DB {
	if query == nil {
		query = database.Master()
	}
	if wheres != nil {
		for k, v := range *wheres {
			if strings.Count(k, "?") > 1 {
				if arrayOfInterfaces, ok := v.([]interface{}); ok {
					query = query.Where(k, arrayOfInterfaces...)
				}
				continue
			}
			query = query.Where(k, v)
		}
	}
	if orders != nil {
		for k, v := range *orders {
			query = query.Order(k + " " + v)
		}
	}
	return query
}

// Count 查询
func Count(conn *gorm.DB, model interface{}, wheres *Where) int64 {
	var c int64
	queryBuilder(conn, wheres, nil).Model(model).Count(&c)
	return c
}

// Take 查询
func Take[T any](conn *gorm.DB, wheres *Where, dest *T, scopes ...func(*gorm.DB) *gorm.DB) *T {
	q := queryBuilder(conn, wheres, nil)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	q.Take(dest)
	return dest
}

// TakeByID 查询指定ID的一条记录
func TakeByID[T any](conn *gorm.DB, id interface{}, dest *T, scopes ...func(*gorm.DB) *gorm.DB) *T {
	q := queryBuilder(conn, &Where{"id=?": id}, nil)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	q.Take(dest)
	return dest
}

// Find 查询列表
func Find[T any](conn *gorm.DB, dest []*T, wheres *Where, orderBy *SqlOrder, scopes ...func(*gorm.DB) *gorm.DB) []*T {
	q := queryBuilder(conn, wheres, orderBy)
	if len(scopes) > 0 {
		q = q.Scopes(scopes...)
	}
	q.Find(&dest)
	return dest
}

// Exists 查询指定条件的行是否存在
func Exists(conn *gorm.DB, table string, wheres *Where) bool {
	q := queryBuilder(conn, wheres, nil).Table(table).Limit(1)
	var ret struct {
		N uint8 `json:"n"`
	}
	conn.Raw("select EXISTS(?) as n", q).Scan(&ret)
	return ret.N > 0
}
