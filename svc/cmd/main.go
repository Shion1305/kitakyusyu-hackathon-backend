package main

import (
	"github.com/gin-gonic/gin"
	"kitakyusyu-hackathon/pkg/firestore"
	"kitakyusyu-hackathon/svc/pkg/handler"
	"kitakyusyu-hackathon/svc/pkg/middleware"
	"log"
)

func main() {
	e := gin.Default()
	fs := firestore.New()
	inquiryHandler := handler.NewInquiryHandler(fs)
	chatHandler := handler.NewChatHandler()
	cors := middleware.NewCORSMiddleware()
	e.Use(cors.Handle)

	rg := e.Group("/api/v1").
		Use(cors.Handle)
	rg.POST("/inquiry", inquiryHandler.HandleInquiry())
	rg.GET("/chat", chatHandler.Handle())
	rg.POST("/chat", chatHandler.Handle())
	e.NoRoute(func(c *gin.Context) {
		log.Println("NoRoute")
		c.JSON(200, gin.H{
			"message": "no api found",
		})
	})
	if err := e.Run(":8080"); err != nil {
		panic(err)
	}
}
