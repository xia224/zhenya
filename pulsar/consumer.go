package main

import (
	"context"
	"fmt"
	"log"
	"pulsar/common"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/golang/protobuf/proto"
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

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "my-topic",
		SubscriptionName: "subName",
		Type:             pulsar.Shared,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()
	// panic("aaa")

	msg, err := consumer.Receive(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg.Payload())
	bill_msg := &common.Message{}
	if err := proto.Unmarshal(msg.Payload(), bill_msg); err != nil {
		log.Fatalln("Failed to decode message:", err)
	}
	fmt.Println("Received data: %+v", bill_msg)

	consumer.Ack(msg)
	if err := consumer.Unsubscribe(); err != nil {
		log.Fatalln("Failed to unsubscribe:", err)
	} else {
		fmt.Println("Succeed to unsubscribe consumer")
	}
}
