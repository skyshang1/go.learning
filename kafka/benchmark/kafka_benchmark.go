package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

var (
	Topic         string
	TopicList     string
	Address       string //{"http://192.168.1.211:9092", "http://192.168.1.212:9092", "http://192.168.1.213:9092"}
	MessageCount  int
	MessageSize   int
	ConsumerGroup string
	ConsumeMode   string
	ConsumerCount int
	FetchSize     int
	operationType string
	UserName      string
	PassWord      string
)

type OperationType string

const (
	OptSyncProduceMessage  OperationType = "SyncProduceMessage"
	OptAsyncProduceMessage OperationType = "AsyncProduceMessage"
	OptConsumeMessage      OperationType = "ConsumeMessage"

	ConsumeModeNormal          = "Normal"
	ConsumeModeConnectionReuse = "ConnectionReuse"
)

func main() {
	flag.StringVar(&Topic, "Topic", "kafka-test", "Kafka Test Topic Name")
	flag.StringVar(&TopicList, "TopicList", "", "Kafka Test Topic Name List")
	flag.StringVar(&ConsumerGroup, "ConsumerGroup", "kafka-test", "Kafka Test Consumer Group")
	flag.StringVar(&ConsumeMode, "ConsumeMode", "ConnectionReuse/Normal", "Consume Mode")
	flag.StringVar(&Address, "Address", "127.0.0.1:9092", "Kafka Brokers Address")
	flag.IntVar(&MessageSize, "MessageSize", 1, "Message Size")
	flag.IntVar(&MessageCount, "MessageCount", 1, "Message Count")
	flag.StringVar(&operationType, "OperationType", "", "Operation Type")
	flag.IntVar(&ConsumerCount, "ConsumerCount", 1, "Consumer Count")
	flag.IntVar(&FetchSize, "FetchSize", 1024, "Fitch Size")
	flag.StringVar(&UserName, "UserName", "", "UserName")
	flag.StringVar(&PassWord, "PassWord", "", "PassWord")

	flag.Parse()

	addresses := strings.Split(Address, ",")
	topicList := strings.Split(TopicList, ",")

	// validate parameter
	if operationType == "" {
		fmt.Println("Must Pass Param Operator Type.")
		return
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	switch OperationType(operationType) {
	case OptSyncProduceMessage:
		// construct producer
		producer, err := NewDefaultProducer(addresses)
		if err != nil {
			panic(err)
		}
		defer producer.Close()
		fmt.Println("construct kafka sync producer success.")

		SyncProduceMessage(producer, Topic, MessageSize, MessageCount)
	case OptAsyncProduceMessage:
		producer, err := NewAsyncProducer(addresses)
		if err != nil {
			panic(err)
		}
		defer producer.AsyncClose()

		AsyncProduceMessage(producer, Topic, MessageSize, MessageCount, signals)
		fmt.Println("produce message success")
	case OptConsumeMessage:
		if len(ConsumeMode) == 0 {
			fmt.Println("must pass consume mode")
			return
		}

		switch ConsumeMode {
		case ConsumeModeNormal:
			ConsumeMessage(addresses, ConsumerGroup, topicList, ConsumerCount, MessageCount)
		case ConsumeModeConnectionReuse:
			ConsumeMessageWithConnectionReuse(addresses, ConsumerGroup, topicList, ConsumerCount, MessageCount)
		}
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

	if UserName != "" && PassWord != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = UserName
		config.Net.SASL.Password = PassWord
	}

	producer, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("construct kafka producer failed. Error:", err)
		return nil, err
	}

	return producer, nil
}

func NewAsyncProducer(address []string) (sarama.AsyncProducer, error) {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Version, err = sarama.ParseKafkaVersion("0.10.2.1")
	if err != nil {
		fmt.Println("parse kafka version failed. Error:", err)
		return nil, err
	}

	if UserName != "" && PassWord != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = UserName
		config.Net.SASL.Password = PassWord
	}

	producer, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		fmt.Println("construct kafka async producer failed. Error:", err)
		return nil, err
	}

	return producer, nil
}

func NewDefaultConsumer(address []string, groupID string) (sarama.ConsumerGroup, error) {
	var err error
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Version, err = sarama.ParseKafkaVersion("0.10.2.1")
	if FetchSize != 0 {
		config.Consumer.Fetch.Default = int32(FetchSize)
	}

	if UserName != "" && PassWord != "" {
		config.Net.SASL.Enable = true
		config.Net.SASL.Handshake = true
		config.Net.SASL.User = UserName
		config.Net.SASL.Password = PassWord
	}

	if err != nil {
		fmt.Println("parse kafka version failed. Error:", err)
		return nil, err
	}

	consumer, err := sarama.NewConsumerGroup(address, groupID, config)
	if err != nil {
		fmt.Println("construct kafka consumer group failed. Error:", err)
		return nil, err
	}
	return consumer, nil
}

func SyncProduceMessage(producer sarama.SyncProducer, topic string, size, count int) {
	var message string
	for i := 0; i < size; i++ {
		message += string('a' + rand.Int()%26)
	}

	startTime := time.Now()
	for index := 0; index < count; index++ {
		_, _, err := producer.SendMessage(&sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(message),
		})
		if err != nil {
			fmt.Printf("send message(%s) err=%s \n", message, err)
		}
	}
	fmt.Printf("Produce %v Messages Elapsed Time: %vs\n", count, time.Now().Sub(startTime).Seconds())
}

func AsyncProduceMessage(producer sarama.AsyncProducer, topic string, size, count int, stopC <-chan os.Signal) {
	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func(stopC <-chan os.Signal) {
		defer wg.Done()
		ticker := time.NewTicker(10 * time.Second)

		select {
		case <-producer.Successes():
			successes++
		case <-producer.Errors():
			errors++
		case <-ticker.C:
			fmt.Printf("[Topic: %s] successfully produce %d messages in 10s, %d messages has occured error\n", topic, successes, errors)
			successes, errors = 0, 0

		case <-stopC:
			return
		}
	}(stopC)

	wg.Add(1)
	messageChan := make(chan *sarama.ProducerMessage, 1024)

	go func(stopC <-chan os.Signal, messageChan chan<- *sarama.ProducerMessage) {
		var message string
		for i := 0; i < size; i++ {
			message += string('a' + rand.Int()%26)
		}

		for index := 0; index < count; index++ {
			messageChan <- &sarama.ProducerMessage{
				Topic: topic,
				Value: sarama.ByteEncoder(message),
			}
		}
	}(stopC, messageChan)

	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case message := <-messageChan:
			producer.Input() <- message
			enqueued++
		case <-ticker.C:
			fmt.Printf("[Topic: %s] asyn produce %s message in ten seconds\n", topic, enqueued)
			enqueued = 0
		case <-stopC:
			break
		}
	}

	wg.Wait()
}

type ConsumerGroupHandler struct{}

func (c ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Printf("can't get message from consumer, msg = %+v, ok = %v, consumer = %+v", message, ok, claim.Partition())
				return nil
			}
			session.MarkMessage(message, "")
			count++
			size += float32(len(message.Value))
		case <-ticker.C:
			fmt.Printf("[Topic: %s, ConsumerID: %v, Times: %v, Mode: Normal] consume %v messages minute, %v MB data per minute\n", claim.Topic(), claim.Partition(), times, count, size/(1024*1024))
			times++
			count = 0
			size = 0
		}
	}
}

var (
	times, count, size = 0, 0, float32(0)
	ticker             = time.NewTicker(time.Minute)
)

func ConsumeMessage(addresses []string, group string, topicList []string, consumerCount, messageCount int) {
	var waitGroup sync.WaitGroup

	for _, topic := range topicList {
		for index := 0; index < consumerCount; index++ {
			consumer, err := NewDefaultConsumer(addresses, group)
			if err != nil {
				panic(err)
			}
			fmt.Printf("construct kafka consumer %v success.\n", index)
			waitGroup.Add(1)

			go func(waitGroup *sync.WaitGroup, consumer sarama.ConsumerGroup, topic string, id int) {
				defer func() {
					consumer.Close()
					waitGroup.Done()
					fmt.Printf("[Topic: %s, ConsumerID: %v] closed.\n", topic, id)
				}()

				//
				ctx := context.Background()
				for {
					topics := []string{topic}
					handler := ConsumerGroupHandler{}

					err := consumer.Consume(ctx, topics, handler)
					if err != nil {
						panic(err)
					}
				}
			}(&waitGroup, consumer, topic, index)
		}
	}

	waitGroup.Wait()
}

type statistics struct {
	count int
	size  float32
}

type ReuseConsumerGroupHandler struct {
	times          int
	mutex          sync.Mutex
	StatisticsInfo map[string]*statistics
}

func NewReuseConsumerGroupHandler() *ReuseConsumerGroupHandler {
	return &ReuseConsumerGroupHandler{
		StatisticsInfo: make(map[string]*statistics),
	}
}

func (c *ReuseConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	go func() {
		ticker := time.NewTicker(time.Minute)
		for {
			select {
			case <-ticker.C:
				c.mutex.Lock()
				for topic, info := range c.StatisticsInfo {
					fmt.Printf("[Topic: %s, Times: %v, Mode: ConnectionReuse] consume %v messages per minute, %v MB data per minute\n", topic, c.times, info.count, info.size/(1024*1024))
					c.StatisticsInfo[topic] = &statistics{}
				}
				c.mutex.Unlock()
				c.times++
			}
		}
	}()
	return nil
}

func (c *ReuseConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *ReuseConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				fmt.Printf("can't get message from consumer, msg = %+v, ok = %v, consumer = %+v", message, ok, claim.Partition())
				return nil
			}
			session.MarkMessage(message, "")
			// statistics
			c.mutex.Lock()
			info, exist := c.StatisticsInfo[message.Topic]
			if !exist {
				info = &statistics{}
				c.StatisticsInfo[message.Topic] = info
			}
			info.count++
			info.size += float32(len(message.Value))
			c.mutex.Unlock()
		}
	}
}

func ConsumeMessageWithConnectionReuse(addresses []string, group string, topicList []string, consumerCount, messageCount int) {
	var waitGroup sync.WaitGroup

	for _, topic := range topicList {
		for index := 0; index < consumerCount; index++ {
			consumer, err := NewDefaultConsumer(addresses, group)
			if err != nil {
				panic(err)
			}
			fmt.Printf("construct kafka consumer %v success.\n", index)
			waitGroup.Add(1)

			go func(waitGroup *sync.WaitGroup, consumer sarama.ConsumerGroup, id int) {
				defer func() {
					consumer.Close()
					waitGroup.Done()
					fmt.Printf("[consumer_%v] closed.\n", id)
				}()

				for {
					err := consumer.Consume(context.Background(), []string{topic}, NewReuseConsumerGroupHandler())
					if err != nil {
						fmt.Println("Consume error:", err)
					}
				}

			}(&waitGroup, consumer, index)
		}
	}

	waitGroup.Wait()
}
