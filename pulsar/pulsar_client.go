package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://101.64.234.16:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}
	defer client.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	fmt.Println("Producer name: ", producer.Name())
	defer producer.Close()

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "subName",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	ID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:      []byte("hello"),
		DeliverAfter: 3 * time.Second,
	})
	fmt.Println(ID)

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")

	ctx, canc := context.WithTimeout(context.Background(), 1*time.Second)
	msg, err := consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Payload())
	canc()

	ctx, canc = context.WithTimeout(context.Background(), 5*time.Second)
	msg, err = consumer.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Payload())
	canc()

	consumer.Ack(msg)
	if err := consumer.Unsubscribe(); err != nil {
		log.Fatal()
	} else {
		fmt.Println("Succeed to unsubscribe consumer")
	}
}
