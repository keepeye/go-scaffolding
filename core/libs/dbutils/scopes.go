package dbutils

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page > 0 && pageSize > 0 {
			offset := (page - 1) * pageSize
			return db.Offset(offset).Limit(pageSize)
		} else {
			return db
		}
	}
}

func PreloadAssociations() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(clause.Associations)
	}
}

func Preload(names string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload(names)
	}
}

func Limit(num int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(num)
	}
}
