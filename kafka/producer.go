package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"log"
	"time"
)

var KafkaProducer sarama.SyncProducer

func NewKafkaProducer() {
	var err error
	configs :=  sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Return.Errors = true

	KafkaProducer, err = sarama.NewSyncProducer([]string{"localhost:9092"}, configs)
	if err != nil {
		glog.Errorf("error creating kafka producer err:%v", err)
		time.Sleep(5 * time.Second)
		NewKafkaProducer()
	}
	glog.Info("kafkaproducer connected")

}

func SentMessg(b []byte) {
	msg := &sarama.ProducerMessage{Topic: "tweets", Value: sarama.ByteEncoder(b)}
	partition, offset, err := KafkaProducer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

}

func CloseProducer() {
	if err := KafkaProducer.Close(); err != nil {
		glog.Errorf("error while closing producer")
	}
}
