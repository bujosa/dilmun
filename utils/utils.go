package utils

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

// Create a topic if it doesn't exist
func CreateTopic(ctx context.Context, topicID string) (*pubsub.Topic) {
	client := PubSubClientInstance.Client
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

	return topic
}

// Create a subscription if it doesn't exist
func CreateSubscription(ctx context.Context, subscriptionID string, topic *pubsub.Topic) (*pubsub.Subscription) {
	client := PubSubClientInstance.Client
	sub := client.Subscription(subscriptionID)
	exists, err := sub.Exists(ctx)
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

	return sub
}