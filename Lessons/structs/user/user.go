package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "---",
			CreatedAt: time.Now(),
		},
	}
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, Last name and birthDate are required.")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		CreatedAt: time.Now(),
	}, nil
}
