package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	group sarama.ConsumerGroup
)

type exampleConsumerGroupHandler struct {
	signal *chan os.Signal
}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {

		select {
		case msg := <-claim.Messages():
			fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
			tweet := sess.Context().Value("msgChan").(chan<- []byte)
			tweet <- msg.Value
			sess.MarkMessage(msg, "")
		case <-*h.signal:
			fmt.Println("Stopping consuming message")
			return nil
		}
	}

}
func KafkaConsumer(msg chan<- []byte) {
	ctx := context.WithValue(context.Background(), "msgChan", msg)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	topics := []string{"tweets"}
	handler := exampleConsumerGroupHandler{&signals}

	err := group.Consume(ctx, topics, handler)
	if err != nil {
		glog.Errorf("error while consuming :%v", err)
		time.Sleep(5 * time.Second)
		KafkaConsumer(msg)
	}
}

func InitKafkaConsumer() {
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_0
	config.Consumer.Return.Errors = true
	var err error
	group, err = sarama.NewConsumerGroup([]string{"kafka:9092"}, "my-group", config)
	if err != nil {
		glog.Errorf("error creating kafka consumer:%v", err)
		time.Sleep(5 * time.Second)
		InitKafkaConsumer()
	}
	glog.Info("conected to kafka")
}

func CloseConsumner(){
	group.Close()
}