package main

import (
	"context"
	"dilmun/shared"
	"dilmun/utils"
	"net/http"
)

const (
	topicID        = "your-topic"
	subscriptionID = "your-subscription"
)

func main() {
	shared.LoadEnv()
	ctx := context.Background()

	utils.InitClient()
	defer utils.CloseClient()

	// Create topic
	topic := utils.CreateTopic(ctx, topicID)

	// Create subscription
	sub := utils.CreateSubscription(ctx, subscriptionID, topic)

	// Expose an HTTP endpoint to publish messages	
	http.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		message := r.FormValue("message")

		utils.PublishMessage(ctx, topic, message)
	})


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", nil)

	// Publish a message to the topic
	utils.PublishMessage(ctx, topic, "Hello, World! From Dilmun!")

	// Receive messages from the subscription
	for {
		utils.ReceiveMessage(ctx, sub)
	}

}
