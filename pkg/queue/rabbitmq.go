package queue

import (
	"github.com/Lucasvmarangoni/financial-file-manager/config"
	errors "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type IRabbitMQ interface {
	Connect() *amqp.Channel
	Consume(messageChannel chan amqp.Delivery, routingKey string)
	Publish(message string, contentType string, exchange string, routingKey string) error
}

type RabbitMQ struct {
	User              string
	Password          string
	Host              string
	Port              string
	Vhost             string
	ConsumerQueueName string
	ConsumerName      string
	AutoAck           bool
	Args              amqp.Table
	Channel           *amqp.Channel
}

func NewRabbitMQ() *RabbitMQ {

	rabbitMQArgs := amqp.Table{}
	rabbitMQArgs["x-dead-letter-exchange"] = config.GetEnvString("rabbitMQ", "dlx")

	rabbitMQ := RabbitMQ{
		User:              config.GetEnvString("rabbitMQ", "user"),
		Password:          config.GetEnvString("rabbitMQ", "pass"),
		Host:              config.GetEnvString("rabbitMQ", "host"),
		Port:              config.GetEnvString("rabbitMQ", "port"),
		Vhost:             config.GetEnvString("rabbitMQ", "vhost"),
		ConsumerQueueName: config.GetEnvString("rabbitMQ", "queue"),
		ConsumerName:      config.GetEnvString("rabbitMQ", "consumer"),
		AutoAck:           true,
		Args:              rabbitMQArgs,
	}
	return &rabbitMQ
}

func (r *RabbitMQ) Connect() *amqp.Channel {
	dsn := "amqp://" + r.User + ":" + r.Password + "@" + r.Host + ":" + r.Port + r.Vhost
	conn, err := amqp.Dial(dsn)
	errors.FailOnErrLog(err, "amqp.Dial", "Failed to connect to RabbitMQ")

	r.Channel, err = conn.Channel()
	errors.FailOnErrLog(err, "conn.Channel", "Failed to open a channel")

	return r.Channel
}

func (r *RabbitMQ) Consume(messageChannel chan amqp.Delivery, routingKey string) {
	q, err := r.Channel.QueueDeclare(
		r.ConsumerQueueName,
		true,
		false,
		false,
		false,
		r.Args,
	)
	errors.FailOnErrLog(err, "r.Channel.QueueDeclare", "Failed to declare a queue")

	err = r.Channel.QueueBind(
		q.Name,
		routingKey,
		config.GetEnvString("rabbitMQ", "exchange"),
		false,
		nil,
	)
	errors.FailOnErrLog(err, "r.Channel.QueueBind", "Failed to bind a queue")

	incomingMessage, err := r.Channel.Consume(
		q.Name,
		r.ConsumerName,
		r.AutoAck,
		false,
		false,
		false,
		r.Args,
	)
	errors.FailOnErrLog(err, "r.Channel.Consume", "Failed to register a consumer")

	go func() {
		for message := range incomingMessage {
			log.Info().Str("context", "RabbitMQ").Msg("New incoming message")
			messageChannel <- message
		}
		if err := r.Channel.Close(); err != nil {
			log.Error().Err(errors.ErrCtx(err, "Failed to close RabbitMQ channel")).Send()
		} else {
			log.Info().Str("context", "RabbitMQ").Msg("RabbitMQ channel closed gracefully")
		}
		close(messageChannel)
	}()
}

func (r *RabbitMQ) Publish(message string, contentType string, exchange string, routingKey string) error {
	err := r.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(message),
		},
	)
	if err != nil {
		return errors.ErrCtx(err, "Failed to publish message")
	}
	return nil
}
