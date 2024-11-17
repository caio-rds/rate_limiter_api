package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/schemas"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Struct struct {
	*gorm.DB
}

func NewDb(db *gorm.DB) *Struct {
	value := Struct{
		db,
	}
	return &value
}

func getToken(id uint, username string) (string, error) {
	token, err := GenerateJwt(id, username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (db *Struct) TryLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(400, gin.H{"msg": "Invalid request"})
		return
	}

	var user *schemas.User
	db.Where("username = ?", username).First(&user)
	if ComparePassword(user.Password, password) {
		if token, err := getToken(user.ID, user.Username); err != nil {
			c.JSON(500, gin.H{"msg": "Internal server error"})
			return
		} else {
			c.JSON(200, gin.H{"token": token})
			return
		}
	}
	c.JSON(401, gin.H{"msg": "Invalid credentials"})
	return
}

func Password(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

func ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
