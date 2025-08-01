package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Email                      string `gorm:"primaryKey"`
	Name                       string
	City                       string
	PhoneNumber                string
	Password                   string
	PostCode                   uint
	SchoolId                   string
	FieldOfStudyId             string
	DoesNotLiveInDenmark       bool
	NoLongerAttendingTheCourse bool
	TermsAccepted              bool
	Updated                    int64 `gorm:"autoUpdateTime" json:"updated_at"`
	Created                    int64 `gorm:"autoCreateTime" json:"created_at"`
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

func (db *UserDatabase) GetUserByEmail(email string) (user User, err error) {
	result := db.db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&user, "email = ?", email)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = errors.Join(ErrDatabaseGetUser, result.Error)
		return user, err
	}

	return user, err
}

func (db *UserDatabase) GetUsers() (users []User, err error) {
	result := db.db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Find(&users)

	if result.Error != nil {
		return users, errors.Join(ErrDatabaseGetUser, result.Error)
	}
	return users, err
}

func (db *UserDatabase) CreateUser(user User) (User, error) {
	result := db.db.Clauses(clause.Locking{
		Strength: "SHARE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Save(&user)
	if result.Error != nil {
		return User{}, errors.Join(ErrDatabaseCreateUser, result.Error)
	}
	return User{}, nil
}
