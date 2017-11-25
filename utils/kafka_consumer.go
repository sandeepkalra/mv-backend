package utils

import (
	"fmt"

	"github.com/Shopify/sarama"

	"os"
	"os/signal"
)

// Consumer is a kafka consumer
func Consumer(topic, myIP, redisIP string) (sarama.Consumer, *RedisDb, error) {
	dial := myIP + ":9092"
	redisStartIndexTag := topic + "_MsgStartIndex"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// Specify brokers address. This is default one
	brokers := []string{dial}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create consumer")
	}
	b, r := FastMemInit(redisIP)
	if b != true {
		master.Close()
		return nil, nil, fmt.Errorf("Error connecting to Redis")
	}

	startOffsetIter64Bit := sarama.OffsetOldest
	if r.R.Get(redisStartIndexTag).Val() != "" {
		startOffsetIter64Bit, _ = r.R.Get(redisStartIndexTag).Int64()
	}
	// How to decide partition, is it fixed value...?
	consumer, err := master.ConsumePartition(topic, 0, startOffsetIter64Bit)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Get signnal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
				//SendEmail(mySql, &cfg, msg)
				startOffsetIter64Bit++
				r.R.Set(redisStartIndexTag, startOffsetIter64Bit, 0)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	return master, r, nil
}
