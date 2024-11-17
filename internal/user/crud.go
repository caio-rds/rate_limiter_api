package user

import (
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/login"
	"go_limiter_rate/internal/schemas"
	"gorm.io/gorm"
)

func (s *SQLite) Create(c *gin.Context) {
	var user *schemas.User
	var createUser CreateUser
	if err := c.BindJSON(&createUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user = &schemas.User{
		Username: createUser.Username,
		Email:    createUser.Email,
		Password: createUser.Password,
		Name:     createUser.Name,
	}

	if _, err := user.ValidatePassword(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Password = login.Password(user.Password)

	if err := s.DB.Create(&user).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			c.JSON(400, gin.H{"error": "Username already exists"})
			return
		}
		if err.Error() == "UNIQUE constraint failed: users.email" {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user_id": user.ID})
}

func (s *SQLite) Read(c *gin.Context, UserId uint) {

	var user *schemas.User
	if err := s.DB.First(&user, UserId).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var response *ResponseUser
	response = &ResponseUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.Format("02/01/2006 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("02/01/2006 15:04:05"),
	}

	c.JSON(200, &response)
}

func (s *SQLite) Update(c *gin.Context, UserId uint) {
	var user *schemas.User
	var updateUser RequestUpdateUser
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := s.DB.First(&user, UserId).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if updateUser.Username != nil {
		user.Username = *updateUser.Username
	}
	if updateUser.Email != nil {
		user.Email = *updateUser.Email
	}
	if updateUser.Name != nil {
		user.Name = *updateUser.Name
	}

	if err := s.DB.Save(&user).Error; err != nil {
		if err.Error() == "UNIQUE constraint failed: users.username" {
			c.JSON(400, gin.H{"error": "Username already exists"})
			return
		}
		if err.Error() == "UNIQUE constraint failed: users.email" {
			c.JSON(400, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User updated successfully"})
}

func (s *SQLite) Delete(c *gin.Context, UserId uint) {

	var user *schemas.User
	if err := s.DB.First(&user, UserId).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	user.DeletedAt = &gorm.DeletedAt{
		Time:  user.DeletedAt.Time,
		Valid: true,
	}

	if err := s.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (s *SQLite) Restore(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(400, gin.H{"error": "User ID is required"})
		return
	}

	var user *schemas.User
	if err := s.DB.First(&user, userId).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if err := s.DB.Unscoped().Model(&user).Update("deleted_at", nil).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User restored successfully"})
}
