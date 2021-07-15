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

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "my-topic",
	})
	fmt.Println("Producer name: ", producer.Name())
	defer producer.Close()

	// pb data
	st := &common.Status{
		Code:   200,
		Reason: "Billing_ok",
	}

	msg := &common.Message{
		Type:   common.Message_BILLING,
		Status: st,
	}

	out, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Failed to encode billing message:", err)
		return
	}

	ID, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		// Payload:      []byte("hello"),
		Payload:      out,
		DeliverAfter: 3 * time.Second,
	})
	fmt.Println(ID)

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}
