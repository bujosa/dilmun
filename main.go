package main

import (
	"context"
	"dilmun/shared"
	"dilmun/utils"
	"net/http"
)

const (
	topicID        = "test-topic"
	subscriptionID = "test-topic-sub"
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

	go func() {
		for {
			utils.ReceiveMessage(ctx, sub)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
