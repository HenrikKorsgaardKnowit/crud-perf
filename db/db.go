package db

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrDatabaseConnection = errors.New("error connecting to the database")
	ErrDatabaseUpdateUser = errors.New("error updating user in database")
	ErrDatabaseGetUser    = errors.New("error get user(s) in database")
	ErrDatabaseCreateUser = errors.New("error creating user(s) in database")
	ErrDatabaseDeleteUser = errors.New("error deleting user in database")
)

type User struct {
	Email           string `gorm:"primaryKey"`
	Name            string
	City            string
	Phone           string
	Postal          uint
	SchoolID        uint
	FieldOfStudyID  uint
	DKResident      bool
	AttendingCourse bool
	Updated         int64 `gorm:"autoUpdateTime" json:"updated_at"`
	Created         int64 `gorm:"autoCreateTime" json:"created_at"`
}

type UserDatabase struct {
	db *gorm.DB
}

func New(host string) *UserDatabase {
	db, err := gorm.Open(sqlite.Open(host), &gorm.Config{Logger: logger.Default.LogMode((logger.Warn))})
	if err != nil {
		panic(errors.Join(ErrDatabaseConnection, err))
	}

	db.AutoMigrate(&User{})
	return &UserDatabase{
		db: db,
	}
}
