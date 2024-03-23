package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)

var (
	client *firestore.Client
)

func New() *firestore.Client {
	if client == nil {
		ctx := context.Background()
		c, err := firestore.NewClient(ctx, "kitakyusyu-hackathon")
		if err != nil {
			log.Fatalf("firebase.NewClient err: %v", err)
		}
		client = c
	}
	return client
}
