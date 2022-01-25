package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kelseyhightower/envconfig"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushMessageToQueue(message []byte) error {
	var cfg configKafka
	err := envconfig.Process("CAR", &cfg)
	brokersUrl := []string{cfg.BrokerUrl}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: cfg.Topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", cfg.Topic, partition, offset)

	return nil
}

type configKafka struct {
	//Kafka configuration
	BrokerUrl string `default:"localhost:9092"`
	Topic     string `default:"car_events_report"`
}
