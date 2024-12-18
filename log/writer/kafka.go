package writer

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/TiktokCommence/component/log/config"
	"io"
)

type KafkaWriter struct {
	producer sarama.SyncProducer
	topic    string
}

func newKafkaWriter(brokers []string, topic string) (*KafkaWriter, error) {
	// 配置 Kafka 生产者
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true            // 等待 Kafka 返回成功消息
	conf.Producer.RequiredAcks = sarama.WaitForLocal // 本地写入确认
	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %v", err)
	}
	return &KafkaWriter{producer: producer, topic: topic}, nil
}

func (k *KafkaWriter) Write(p []byte) (n int, err error) {
	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.StringEncoder(p),
	}
	_, _, err = k.producer.SendMessage(msg)
	return len(p), err
}

type KafkaBuilder struct {
	conf *config.KafkaConfig
}

func NewKafkaBuilder(conf *config.KafkaConfig) *KafkaBuilder {
	return &KafkaBuilder{conf: conf}
}

func (k *KafkaBuilder) Build() (io.Writer, error) {
	writer, err := newKafkaWriter(k.conf.BrokersAddr, k.conf.TopicName)
	if err != nil {
		return nil, fmt.Errorf("failed to create kafka writer: %v", err)
	}
	return writer, nil
}
