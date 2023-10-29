package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

const (
	projectID      = "your-project-id"
	topicID        = "your-topic"
	subscriptionID = "your-subscription"
)

func main() {
	ctx := context.Background()

	// Load Google Cloud credentials from a JSON file
	credentialFile := "path/to/your/credentials.json"
	opt := option.WithCredentialsFile(credentialFile)

	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, projectID, opt)
	if err != nil {
		log.Fatalf("Error creating Pub/Sub client: %v", err)
	}
	defer client.Close()

	// Create a topic if it doesn't exist
	topic := client.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking if topic exists: %v", err)
	}
	if !exists {
		if _, err = client.CreateTopic(ctx, topicID); err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	// Create a subscription if it doesn't exist
	sub := client.Subscription(subscriptionID)
	exists, err = sub.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking if subscription exists: %v", err)
	}
	if !exists {
		if _, err = client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic: topic,
		}); err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
		}
	}

	// Publish a message to the topic
	msg := &pubsub.Message{
		Data: []byte("Hello, Google Cloud Pub/Sub!"),
	}
	res := topic.Publish(ctx, msg)
	id, err := res.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("Published a message with ID: %s\n", id)

	// Receive messages from the subscription
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Error receiving message: %v", err)
	}
}
