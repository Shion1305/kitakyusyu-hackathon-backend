package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	"kitakyusyu-hackathon/pkg/openai"
	"kitakyusyu-hackathon/svc/pkg/schema"
	"log"
	"net/http"
)

type ChatHandler struct {
	openai *openai.OpenAI
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		openai: openai.NewOpenAI(),
	}
}

func (h *ChatHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
		req := schema.ChatRequest{}
		err = json.Unmarshal(data, &req)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
		stream, err := h.openai.GetStreamResponse(c, req.Messages)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		w, r := c.Writer, c.Request
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}

		defer conn.Close()
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				c.AbortWithStatus(200)
				fmt.Println("Stream finished")
				return
			}

			if err != nil {
				fmt.Printf("Stream error: %v\n", err)
				c.AbortWithError(500, err)
				return
			}

			fmt.Printf("Stream response: %v\n", response)
			if err := conn.WriteMessage(
				websocket.TextMessage, []byte(response.Choices[0].Delta.Content)); err != nil {
				log.Printf("failed to write message: %v", err)
				c.AbortWithError(500, err)
				return
			}
		}
	}
}
