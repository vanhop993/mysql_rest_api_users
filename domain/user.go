package domain

import "time"

type UserStruct struct {
	Id          string
	Username    string
	Email       string
	Phone       string
	DateOfBirth *time.Time
}

type UserRepository interface {
	GetAllUsersDB() ([]UserStruct, error)
	GetUserById(id string) (*UserStruct, error)
	InsertUserDB(newUser *UserStruct) (string, error)
	UpdateUserDB(user *UserStruct) (string, error)
	DeleteUserDB(id string) (string, error)
}
