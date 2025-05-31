package models

import (
	"fmt"

	"github.com/Jonaires777/image-uploader/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *User) Validate() error {
	if u.Firstname == "" {
		return fmt.Errorf("firstname is required")
	}
	if u.Lastname == "" {
		return fmt.Errorf("lastname is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Password == "" {
		return fmt.Errorf("password is required")
	}
	if len(u.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	return nil
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Save() error {
	const query = `
		INSERT INTO users (firstname, lastname, email, password) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Firstname, u.Lastname, u.Email, u.Password).Scan(&u.ID)
	return err
}
