package db

import (
	"errors"

	"github.com/shirinebadi/smart-office/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("smartoffice.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Activity{}, &model.Admin{}, &model.CentralUser{}, &model.LocalUser{}, &model.Office{})
	return err
}

func Init() (*gorm.DB, error) {
	db, err := NewDB()
	if err != nil {
		return nil, errors.New("Error in DB Creation")
	}

	if err = migrate(db); err != nil {
		return nil, errors.New("Error in DB Creation" + err.Error())
	}
	return db, nil
}
