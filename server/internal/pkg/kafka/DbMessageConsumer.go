package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/config"
	"github.com/ahmed3hamdan/kafka-chat/server/internal/pkg/model"
	"github.com/sirupsen/logrus"
)

func NewDbMessageConsumer() (*DbConsumer, error) {
	var (
		dbConsumer DbConsumer
		err        error
	)
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V1_0_0_0
	saramaConfig.Consumer.Return.Errors = true
	dbConsumer.consumerGroup, err = sarama.NewConsumerGroup([]string{config.KafkaAddress}, config.KafkaDbMessageConsumer, saramaConfig)
	if err != nil {
		return nil, err
	}
	return &dbConsumer, nil
}

type DbConsumer struct {
	consumerGroup sarama.ConsumerGroup
}

func (c *DbConsumer) Start() {
	ctx := context.Background()
	for {
		err := c.consumerGroup.Consume(ctx, []string{config.KafkaMessageTopic}, dbConsumerGroupHandler{})
		if err != nil {
			logrus.Fatalln(err)
		}
	}
}

type dbConsumerGroupHandler struct{}

func (dbConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (dbConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h dbConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		var body MessageBody
		if err := json.Unmarshal(msg.Value, &body); err != nil {
			return err
		}

		messagesMap := make(map[int64]model.Message)
		messagesMap[body.FromUserID] = model.Message{
			OwnerUserID: body.FromUserID,
			FromUserID:  body.FromUserID,
			ToUserID:    body.ToUserID,
			Key:         body.Key,
			Content:     body.Content,
		}
		messagesMap[body.ToUserID] = model.Message{
			OwnerUserID: body.ToUserID,
			FromUserID:  body.FromUserID,
			ToUserID:    body.ToUserID,
			Key:         body.Key,
			Content:     body.Content,
		}

		messages := make([]model.Message, 0, len(messagesMap))
		for _, message := range messagesMap {
			messages = append(messages, message)
		}

		if err := model.InsertMessages(sess.Context(), messages); err != nil {
			return err
		}

		logrus.WithField("message", body).Debugln("message consumed to database successfully")
		sess.MarkMessage(msg, "")
	}
	return nil
}
