package key

import (
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/schemas"
	"gorm.io/gorm"
	"time"
)

func (db *Struct) Create(c *gin.Context, uid uint) {
	key, err := generateKey()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate key"})
		return
	}

	newKey := schemas.Key{
		UserID: uid,
		Key:    *key,
	}
	result := db.DB.Create(&newKey)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create key"})
		return
	}
	c.JSON(200, gin.H{"key": newKey.Key})
}

func (db *Struct) Read(c *gin.Context, uid uint) {
	var key []*schemas.Key
	var response []*Key

	result := db.Find(&key, "user_id = ?", uid)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to read key"})
		return
	}

	for _, k := range key {
		response = append(response, &Key{
			ID:        k.ID,
			UserID:    k.UserID,
			Key:       k.Key,
			CreatedAt: k.CreatedAt,
		})
	}

	c.JSON(200, gin.H{"keys": response})
}

func (db *Struct) Delete(c *gin.Context, uid uint, kid uint) {
	var user *schemas.User
	if err := db.DB.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(500, gin.H{"error": "User not found"})
		return
	}

	var key *schemas.Key
	if err := db.DB.First(&key, "id = ? AND user_id = ?", kid, uid).Error; err != nil {
		c.JSON(500, gin.H{"error": "Key not found"})
		return
	}

	key.DeletedAt = &gorm.DeletedAt{
		Time:  time.Now(),
		Valid: true,
	}

	if err := db.DB.Save(&key).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete key"})
		return
	}

	c.JSON(200, gin.H{"message": "Key deleted"})
}
