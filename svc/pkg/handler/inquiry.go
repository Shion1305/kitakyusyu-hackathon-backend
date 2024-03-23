package handler

import (
	"github.com/gin-gonic/gin"
	"kitakyusyu-hackathon/svc/pkg/schema"
)

type InquiryHandler struct {
}

func NewInquiryHandler() *InquiryHandler {
	return &InquiryHandler{}
}

func (h *InquiryHandler) HandleInquiry() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := schema.InquiryData{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ok",
		})
	}
}
