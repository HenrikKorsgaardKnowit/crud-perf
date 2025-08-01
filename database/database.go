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
	Email                      string `gorm:"primaryKey" json:"email"`
	Name                       string `json:"name"`
	City                       string `json:"city"`
	PhoneNumber                string `json:"phonenumber"`
	Password                   string `json:"password"`
	PostCode                   uint   `json:"posctode"`
	SchoolId                   string `json:"schoolid"`
	FieldOfStudyId             string `json:"fieldofstudyid"`
	DoesNotLiveInDenmark       bool   `json:"doesnotliveindenmark"`
	NoLongerAttendingTheCourse bool   `json:"nolongerattendingthecourse"`
	TermsAccepted              bool   `json:"termsaccepted"`
	Updated                    int64  `gorm:"autoUpdateTime" json:"updated_at"`
	Created                    int64  `gorm:"autoCreateTime" json:"created_at"`
}

type UserDatabase struct {
	db *gorm.DB
}

func New(host string) *UserDatabase {

	db, err := gorm.Open(sqlite.Open(host), &gorm.Config{Logger: logger.Default.LogMode((logger.Error))})
	if err != nil {
		panic(errors.Join(ErrDatabaseConnection, err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(errors.Join(ErrDatabaseConnection, err))
	}

	sqlDB.SetMaxOpenConns(1000)

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
