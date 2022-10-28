package app

import (
	"activity/internal/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGormDB() (*gorm.DB, func(), error) {

	db, err := NewGormDB()
	if err != nil {
		return nil, nil, err
	}
	cleanFunc := func() {

	}
	return db, cleanFunc, nil
}

func NewGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.C.MySql), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
