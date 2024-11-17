package user

import (
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/schemas"
	"gorm.io/gorm"
	"strconv"
)

type SQLite struct{ *gorm.DB }

func NewSQLite(db *gorm.DB) *SQLite {
	value := SQLite{
		db,
	}
	return &value
}

type PacksResponse struct {
	Total int            `json:"total"`
	Packs []PackResponse `json:"packs"`
}

type PackResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Amount    int    `json:"amount"`
	CreatedAt string `json:"created_at"`
}

func (s *SQLite) Create(c *gin.Context, userID uint) {
	amountStr := c.Param("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid amount"})
		return
	}

	if amount != 100 && amount != 500 && amount != 1000 {
		c.JSON(400, gin.H{"error": "invalid amount, must be 100, 500 or 1000"})
		return
	}

	var pack *schemas.Pack
	pack = &schemas.Pack{
		UserID: userID,
		Amount: amount,
	}

	if err := s.DB.Create(pack).Error; err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "pack created"})
}

func (s *SQLite) Read(c *gin.Context, UserId uint) {
	var packs []schemas.Pack
	if err := s.DB.Where("user_id = ?", UserId).Find(&packs).Error; err != nil {
		c.JSON(500, gin.H{"error": "internal server error"})
		return
	}

	var response PacksResponse
	var total int
	for _, pack := range packs {
		total += pack.Amount
		response.Packs = append(response.Packs, PackResponse{
			ID:        pack.ID,
			UserID:    pack.UserID,
			Amount:    pack.Amount,
			CreatedAt: pack.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	response.Total = total
	c.JSON(200, response)
}
