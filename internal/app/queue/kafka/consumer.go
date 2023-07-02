package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/mailgun/mailgun-go/v4"
)

type Consumer interface {
	StartConsumerGroup(ready chan<- bool, topic string) error
	CloseConsumerGroup()
}

type consumer struct {
	consumer sarama.ConsumerGroup
}

func NewKafkaConsumer(kafkaBroker string, groupConsumer string) (Consumer, error) {
	// Set up configuration for the consumer group
	config := sarama.NewConfig()
	//config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	//config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// Create a new consumer group
	consumerGroup, err := sarama.NewConsumerGroup([]string{kafkaBroker}, groupConsumer, config)
	if err != nil {
		return nil, err
	}

	return consumer{consumerGroup}, nil
}

// Helper function to start the consumer group and handle messages
func (csmr consumer) StartConsumerGroup(ready chan<- bool, topic string) error {
	for {
		log.Println("start worker kafka")
		err := csmr.consumer.Consume(context.Background(), []string{topic}, &MyConsumerGroupHandler{})
		if err != nil {
			return err
		}

		// Check if the consumer group is still running
		select {
		case <-csmr.consumer.Errors():
			err := <-csmr.consumer.Errors()
			// Handle the error here
			fmt.Println("Error occurred:", err)
			// You can add your custom error handling logic
		default:
			// The consumer group has stopped, signal that it's ready to restart
			ready <- false
		}
	}
}

func (csmr consumer) CloseConsumerGroup() {
	csmr.consumer.Close()
}

// Custom consumer group handler
type MyConsumerGroupHandler struct{}

func (h MyConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	// Called when the consumer group session is set up
	return nil
}

func (h MyConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	// Called when the consumer group session is closed
	return nil
}

func (h MyConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Process each message in a separate goroutine
	for message := range claim.Messages() {
		go func(msg *sarama.ConsumerMessage) {
			// Retry logic
			retryCount := 0
			maxRetries := 3
			retryInterval := 5 * time.Second
			var processingErr error
			err := processMessage(msg)
			if err == nil {
				log.Println("Message processed successfully")
				session.MarkMessage(msg, "")
				session.Commit()
			}
			if err != nil {
				for retryCount < maxRetries {

					log.Printf("Error processing message: %s", err)
					log.Printf("Retrying message processing in %s...", retryInterval)
					time.Sleep(retryInterval)
					err := processMessage(msg)
					if err == nil {
						log.Println("Message processed successfully after retry.")
						session.MarkMessage(msg, "")
						session.Commit()
						return
					}
					processingErr = err
					retryCount++
				}

				log.Printf("Max retries reached. Message processing failed: %s", processingErr)
				// Perform any necessary error handling, such as logging or sending to an error topic
			}
		}(message)
	}

	return nil
}

func processMessage(msg *sarama.ConsumerMessage) error {
	fmt.Printf("Received message: %s\n", string(msg.Value))

	mg := mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_KEY"))
	message := mg.NewMessage(
		"noreply@gmail.com",
		"info pesanan",
		"Pesanan anda sudah kami terima",
		string(msg.Value),
	)

	_, _, err := mg.Send(context.Background(), message)
	if err != nil {
		return err
	}

	log.Println("Email sent successfully with kafka!")

	return nil
}
