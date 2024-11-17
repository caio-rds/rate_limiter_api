package schemas

import (
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt `gorm:"index"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) ValidatePassword() (bool, error) {

	specialChars := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};:'",<>./?\\|]`)
	upperCase := regexp.MustCompile(`[A-Z]`)
	number := regexp.MustCompile(`[0-9]`)

	if len(u.Password) < 6 {
		return false, fmt.Errorf("password must be at least 6 characters")
	}

	if strings.Contains(u.Password, "password") {
		return false, fmt.Errorf("password cannot contain the word 'password'")
	}

	if !specialChars.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one special character")
	}

	if !number.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one number")
	}

	if !upperCase.MatchString(u.Password) {
		return false, fmt.Errorf("password must contain at least one uppercase letter")
	}

	return true, nil

}
