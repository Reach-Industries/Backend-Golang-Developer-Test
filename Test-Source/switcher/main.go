package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var switchWriter *kafka.Writer

func main() {
	fmt.Println("Starting Kafka Producer")
	switchWriter = getKafkaWriter("switch")
	writeLoop()
}

func getKafkaWriter(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(getBrokerURLs()...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func getBrokerURLs() []string {
	BrokerURLs := strings.Split(os.Getenv("MSK_BROKERS"), ",")
	for i := range BrokerURLs {
		BrokerURLs[i] = strings.TrimSpace(BrokerURLs[i])
	}
	fmt.Println("Brokers:", BrokerURLs)
	return BrokerURLs
}

func writeLoop() {
	fmt.Println("Starting Write Loop")
	messages := []string{"Counter", "Text", "Time", "Counter", "Text", "Time", "Counter", "Text", "Time", "Dump"}
	orgs := []string{"1", "2", "3", "4"}

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for {
		org := orgs[r.Intn(len(orgs))]
		msg := messages[r.Intn(len(messages))]
		writeMessage(org, msg)
		fmt.Println("Wrote message:", msg, "for org:", org)

		sleepTime := time.Duration(r.Intn(5-1+1)+1) * time.Second
		time.Sleep(sleepTime)
	}
}

func writeMessage(org string, msg string) {
	msgJson := fmt.Sprintf(`{"org": "%s", "msg": "%s"}`, org, msg)
	if err := switchWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(time.Now().String()),
			Value: []byte(msgJson),
		},
	); err != nil {
		fmt.Println("Error while writing to kafka", err)
	}
}
