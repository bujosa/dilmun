package utils

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

// Publish a message to a topic
func PublishMessage(ctx context.Context, topic *pubsub.Topic, message string) {
	msg := &pubsub.Message{
		Data: []byte(message),
	}
	res := topic.Publish(ctx, msg)
	id, err := res.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("Published a message with ID: %s\n", id)
}

// Receive messages from a subscription
func ReceiveMessage(ctx context.Context, sub *pubsub.Subscription) {
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}
}