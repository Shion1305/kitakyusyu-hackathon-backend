package middleware

import "github.com/gin-gonic/gin"

type CORSMiddleware struct {
}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (m *CORSMiddleware) Handle(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Max-Age", "600")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Content-Type", "application/json")
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}
