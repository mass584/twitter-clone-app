package entity

import (
	"fmt"
	"unicode/utf8"
)

type User struct {
	Id          int64
	DisplayName string
}

func NewUser(
	Id int64,
	DisplayName string,
) (*User, error) {
	if utf8.RuneCountInString(DisplayName) > 20 {
		return nil, fmt.Errorf("DisplayName is too long.")
	}

	if utf8.RuneCountInString(DisplayName) == 0 {
		return nil, fmt.Errorf("DisplayName is empty.")
	}

	user := User{
		Id:          Id,
		DisplayName: DisplayName,
	}

	return &user, nil
}
