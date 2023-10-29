package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var PubSubClientInstance *GoogleProject

type GoogleProject struct {
	ProjectId string
	Client *pubsub.Client
}

func InitClient() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ctx := context.Background()

	// Load Google Cloud credentials from a JSON file
	credentialFile := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	opt := option.WithCredentialsFile(credentialFile)

	projectId := os.Getenv("GOOGLE_PROJECT_ID")

	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("Error creating Pub/Sub client: %v", err)
	}

	PubSubClientInstance = &GoogleProject{
		ProjectId: projectId,
		Client: client,
	}
}

func CloseClient() {
	PubSubClientInstance.Client.Close()
}