package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"strings"
	"time"
)

var (
	Topic         string
	Address       string //{"http://192.168.1.211:9092", "http://192.168.1.212:9092", "http://192.168.1.213:9092"}
	MessageCount  int
	ConsumerGroup string
	operationType string
)

type OperationType string

const (
	OptSyncProduceMessage OperationType = "SyncProduceMessage"
	OptConsumeMessage     OperationType = "ConsumeMessage"
	OptDisplayClusterInfo OperationType = "DisplayClusterInfo"
	OptDisplayTopicInfo   OperationType = "DisplayTopicInfo"
)

func main() {
	flag.StringVar(&Topic, "Topic", "kafka-test", "Kafka Test Topic Name")
	flag.StringVar(&ConsumerGroup, "ConsumerGroup", "kafka-test", "Kafka Test Consumer Group")
	flag.StringVar(&Address, "Address", "127.0.0.1:9092", "Kafka Brokers Address")
	flag.IntVar(&MessageCount, "MessageCount", 1, "Message Count")
	flag.StringVar(&operationType, "OperationType", "", "Operation Type")
	flag.Parse()

	addresses := strings.Split(Address, ",")

	// validate parameter
	if operationType == "" {
		fmt.Println("Must Pass Param Operator Type.")
		return
	}

	switch OperationType(operationType) {
	case OptSyncProduceMessage:
		// construct producer
		producer, err := NewDefaultProducer(addresses)
		if err != nil {
			panic(err)
		}
		defer producer.Close()
		fmt.Println("construct kafka sync producer success.")

		SyncProduceMessage(producer, Topic, MessageCount)
	case OptConsumeMessage:
		// construct consumer
		consumer, err := NewDefaultConsumer(addresses, ConsumerGroup, []string{Topic})
		if err != nil {
			panic(err)
		}
		defer consumer.Close()
		fmt.Println("construct kafka consumer success.")

		ConsumeMessage(consumer, MessageCount)
	case OptDisplayClusterInfo:
		// construct client
		client, err := NewClusterClient(addresses)
		if err != nil {
			panic(err)
		}
		defer client.Close()
		fmt.Println("construct kafka client success.")

		DisplayClusterInfo(client)
	case OptDisplayTopicInfo:
		// construct client
		client, err := NewClusterClient(addresses)
		if err != nil {
			panic(err)
		}
		defer client.Close()
		fmt.Println("construct kafka client success.")

		DisplayTopicInfo(client, Topic)
	}

}

func NewDefaultProducer(address []string) (sarama.SyncProducer, error) {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Version, err = sarama.ParseKafkaVersion("0.10.2.1")

	if err != nil {
		fmt.Println("parse kafka version failed. Error:", err)
		return nil, err
	}

	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("construct kafka producer failed. Error:", err)
		return nil, err
	}

	return producer, nil
}

func NewClusterClient(address []string) (*cluster.Client, error) {
	config := cluster.NewConfig()
	version, err := sarama.ParseKafkaVersion("0.10.2.1")
	if err != nil {
		fmt.Println("parse kafka version failed. Error:", err)
		return nil, err
	}
	config.Version = version

	return cluster.NewClient(address, config)
}

func NewDefaultConsumer(address []string, groupID string, topics []string) (*cluster.Consumer, error) {
	var err error
	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version, err = sarama.ParseKafkaVersion("0.10.2.1")
	if err != nil {
		fmt.Println("parse kafka version failed. Error:", err)
		return nil, err
	}

	consumer, err := cluster.NewConsumer(address, groupID, topics, config)
	if err != nil {
		fmt.Println("construct kafka consumer failed. Error:", err)
		return nil, err
	}
	return consumer, nil
}

func SyncProduceMessage(producer sarama.SyncProducer, topic string, count int) {
	messageFormat := "this is health. index=%d"
	for index := 0; index < count; index++ {
		message := fmt.Sprintf(messageFormat, index)
		partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(message),
		})
		if err != nil {
			fmt.Printf("send message(%s) err=%s \n", message, err)
		} else {
			fmt.Printf("%s send success, partition=%d, offset=%d \n", message, partition, offset)
		}
	}
}

func ConsumeMessage(consumer *cluster.Consumer, count int) {
	index := 0
	for {
		select {
		case message, ok := <-consumer.Messages():
			if !ok {
				fmt.Printf("can't get message from consumer, msg = %+v, ok = %v, consumer = %+v", message, ok, consumer)
				return
			}
			fmt.Printf("[%d] [Topic:%s, Partition: %d, Offset: %d, TimeStamp: %v] Key: %s, Value: %s\n",
				index, message.Topic, message.Partition, message.Offset, message.Timestamp, string(message.Key), string(message.Value))
			consumer.MarkOffset(message, "")
			if index >= count {
				return
			}

			index++
		}
	}
}

func DisplayClusterInfo(client *cluster.Client) error {
	config := client.ClusterConfig()
	fmt.Println(config)

	brokers := client.Brokers()
	for index, broker := range brokers {
		fmt.Printf("[index: %d] ID: %d, Addr: %s\n", index, broker.ID(), broker.Addr())
	}

	topics, err := client.Topics()
	if err != nil {
		fmt.Println("Get Topic Failed, Error:", err)
		return err
	}
	fmt.Println("Topic List:", topics)

	return nil
}

func DisplayTopicInfo(client *cluster.Client, topic string) error {
	partitions, err := client.Partitions(topic)
	if err != nil {
		fmt.Printf("Get Partitions For Topic [%s] Failed, Error:%s\n", topic, err.Error())
		return err
	}
	for index, partition := range partitions {
		offsetOldest, err := client.GetOffset(topic, partition, sarama.OffsetOldest)
		offsetNewest, err := client.GetOffset(topic, partition, sarama.OffsetNewest)
		broker, err := client.Leader(topic, partition)
		inSyncRep, err := client.InSyncReplicas(topic, partition)

		fmt.Printf("[index: %d] ID: %d, Offset: [%d, %d), Leader: %s, InSyncReplicas: %v\n",
			index, partition, offsetOldest, offsetNewest, broker.Addr(), inSyncRep)

		replicas, err := client.Replicas(topic, partition)
		if err != nil {
			fmt.Printf("Get Replica For [Topic: %s, Partition: %d] Failed, Error: %s\n", topic, partition, err.Error())
		}
		for _, replica := range replicas {
			fmt.Printf("\tReplica ID: %d\n", replica)
		}
	}

	return nil
}
