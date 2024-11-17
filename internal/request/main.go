package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_limiter_rate/internal/schemas"
	"gorm.io/gorm"
)

type SQLite struct{ *gorm.DB }

func NewSQLite(db *gorm.DB) *SQLite {
	value := SQLite{
		db,
	}
	return &value
}

type Request struct {
	Endpoint string           `json:"endpoint"`
	Method   string           `json:"method"`
	Params   *json.RawMessage `json:"params"`
	Headers  *json.RawMessage `json:"headers"`
	Body     *json.RawMessage `json:"body"`
}

func (s *SQLite) Create(c *gin.Context) {
	var request *Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	apiKey := c.GetHeader("X-API-KEY")
	if apiKey == "" {
		c.JSON(401, gin.H{"error": "API Key is required"})
		return
	}

	var key *schemas.Key
	if err := s.DB.Where("key = ?", apiKey).First(&key).Error; err != nil {
		c.JSON(401, gin.H{"error": "API Key is invalid"})
		return
	}

	var packs []*schemas.Pack
	var total int
	if err := s.DB.Where("user_id = ?", key.UserID).Find(&packs).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(packs) == 0 {
		c.JSON(401, gin.H{"error": "no request packages"})
		return
	}

	for _, pack := range packs {
		total += pack.Amount
	}

	var CurrentRequests []*schemas.RequestApi
	if err := s.DB.Where("user_id = ?", key.UserID).Find(&CurrentRequests).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if total < len(CurrentRequests) {
		c.JSON(401, gin.H{"error": "no requests available"})
		return
	}

	// develop all logic to requests another apis

	var schemaReq *schemas.RequestApi
	schemaReq = &schemas.RequestApi{
		UserID:     key.UserID,
		ClientIP:   c.ClientIP(),
		Key:        key.Key,
		Endpoint:   request.Endpoint,
		Method:     request.Method,
		Params:     *request.Params,
		Headers:    *request.Headers,
		Body:       *request.Body,
		Content:    []byte(`{"message": "Hello World"}`),
		StatusCode: 200,
	}

	if err := s.DB.Create(schemaReq).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Request created successfully"})
	return

}

func (s *SQLite) Read(c *gin.Context) {
	apiKey := c.GetHeader("X-API-KEY")
	if apiKey == "" {
		c.JSON(401, gin.H{"error": "API Key is required"})
		return
	}

	var key *schemas.Key
	if err := s.DB.Where("key = ?", apiKey).First(&key).Error; err != nil {
		c.JSON(401, gin.H{"error": "API Key is invalid"})
		return
	}

	var requests []*schemas.RequestApi
	if err := s.DB.Where("user_id = ?", key.UserID).Find(&requests).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, requests)
	return
}
