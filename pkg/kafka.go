package pkg

// import (
// 	"log"

// 	"github.com/segmentio/kafka-go"
// 	"pt-xyz-multifinance/config"
// )

// func NewKafkaWriter(cfg *config.Config, topic string) *kafka.Writer {
// 	writer := kafka.NewWriter(kafka.WriterConfig{
// 		Brokers:  cfg.KafkaBrokers,
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	})
// 	// Optionally, test connection here with a dummy message or metadata fetch
// 	log.Println("Kafka writer initialized to topic:", topic)
// 	return writer
// }
