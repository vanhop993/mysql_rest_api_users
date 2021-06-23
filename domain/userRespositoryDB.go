package domain

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepositoryDb(DB *sql.DB) UserRepositoryDB {
	return UserRepositoryDB{db: DB}
}

func (d UserRepositoryDB) GetAllUsersDB() ([]UserStruct, error) {
	query := "Select id, username, email, phone, date_of_birth from users"
	rows, err := d.db.Query(query)
	if err != nil {
		log.Println("Error while querying users table" + err.Error())
		return nil, err
	}
	users := make([]UserStruct, 0)
	for rows.Next() {
		var c UserStruct
		err := rows.Scan(&c.Id, &c.Username, &c.Email, &c.Phone, &c.DateOfBirth)
		if err != nil {
			log.Println("Error while scanning users" + err.Error())
			return nil, err
		}
		users = append(users, c)
	}
	return users, nil
}

func (d UserRepositoryDB) GetUserById(id string) (*UserStruct, error) {
	query := "select id, username, email, phone, date_of_birth from users where id = ?"
	var user UserStruct
	err := d.db.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
	if err != nil {
		if err != nil {
			errMsg := err.Error()
			if strings.Compare(fmt.Sprintf(errMsg), "0 row(s) returned") == 0 {
				return nil, nil
			} else {
				return nil, err
			}
		}
	}
	return &user, nil
}

func (d UserRepositoryDB) InsertUserDB(newUser *UserStruct) (string, error) {
	query := "insert into users (id, username, email, phone, date_of_birth) values (?, ?, ?, ?, ?)"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	_, er1 := stmt.Exec(newUser.Id, newUser.Username, newUser.Email, newUser.Phone, newUser.DateOfBirth)
	if er1 != nil {
		panic(err.Error())
	}
	resultString := "Create new user success"
	return resultString, nil
}

func (d UserRepositoryDB) UpdateUserDB(user *UserStruct) (string, error) {
	query := "update users set username = ?, email = ?, phone = ?, date_of_birth = ? where id = ?"
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(user.Username, user.Email, user.Phone, user.DateOfBirth, user.Id)
	if err != nil {
		return "", err
	}
	resultString := "Update success"
	return resultString, nil
}

func (d UserRepositoryDB) DeleteUserDB(id string) (string, error) {
	queryFind := "select id, username, email, phone, date_of_birth from users where id = ?"
	var user UserStruct
	err := d.db.QueryRow(queryFind, id).Scan(&user.Id, &user.Username, &user.Email, &user.Phone, &user.DateOfBirth)
	if err != nil {
		if err != nil {
			errMsg := err.Error()
			if strings.Compare(fmt.Sprintf(errMsg), "0 row(s) returned") == 0 {
				return "", nil
			} else {
				return "", err
			}
		}
	}
	query := "delete from users where id = ?"
	stmt, er0 := d.db.Prepare(query)
	if er0 != nil {
		return "", nil
	}
	_, er1 := stmt.Exec(id)
	if er1 != nil {
		return "", er1
	}
	resultString := "Delete success"
	return resultString, nil
}
