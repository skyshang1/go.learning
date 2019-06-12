package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var Address1 = []string{"jstack-kafka-jvessel01.stack-suqian-1a.jdcloud.local:9092"}

func main() {
	// fmt.Printf("tags:")
	// fmt.Printf(strings.Replace(strings.Split(Address1[0], ":")[0], "-", "_", -1) + ":status" + "\n")
	syncProducer(Address1)
}

//同步消息模式
func syncProducer(address []string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// SASL enable
	config.Net.SASL.Enable = true
	config.Net.SASL.Handshake = true
	config.Net.SASL.User = "jvesselRW"
	config.Net.SASL.Password = "KVw0vvXnWpQ7AzIL"

	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Println("construct kafka producer failed. Error:", err)
		return
	}
	fmt.Println("construct kafka producer success.")

	defer p.Close()
	topic := "test_midop1"

	srcValue := "this is hearth. index=%d"
	for i := 0; i < 1; i++ {
		value := fmt.Sprintf(srcValue, i)
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		part, offset, err := p.SendMessage(msg)
		if err != nil {
			fmt.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Printf("%s send success, partition=%d, offset=%d \n", value, part, offset)
		}
		time.Sleep(2 * time.Second)
	}
}
