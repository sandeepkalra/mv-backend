package utils

import (
	"fmt"

	"github.com/Shopify/sarama"

	"os"
	"os/signal"
)

func Consumer(topic, mq_ip, redis_ip string) (sarama.Consumer, *RedisDb, error) {
	dial := mq_ip + ":9092"
	redis_start_index_tag := topic + "_MsgStartIndex"

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// Specify brokers address. This is default one
	brokers := []string{dial}

	// Create new consumer
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to create consumer")
	}
	b, r := FastMemInit(redis_ip)
	if b != true {
		master.Close()
		return nil, nil, fmt.Errorf("Error connecting to Redis")
	}

	var start_offset_iter64 int64 = sarama.OffsetOldest
	if r.R.Get(redis_start_index_tag).Val() != "" {
		start_offset_iter64, _ = r.R.Get(redis_start_index_tag).Int64()
	}
	// How to decide partition, is it fixed value...?
	consumer, err := master.ConsumePartition(topic, 0, start_offset_iter64)
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
				start_offset_iter64++
				r.R.Set(redis_start_index_tag, start_offset_iter64, 0)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	return master, r, nil
}
