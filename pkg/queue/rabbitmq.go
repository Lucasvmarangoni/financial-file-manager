package queue

import (
	"github.com/Lucasvmarangoni/financial-file-manager/config"
	logella "github.com/Lucasvmarangoni/logella/err"
	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

type IRabbitMQ interface {
	Connect() *amqp.Channel
	Consume(messageChannel chan amqp.Delivery, routingKey string)
	Publish(message string, contentType string, exchange string, routingKey string)
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
	rabbitMQArgs["x-dead-letter-exchange"] = config.GetEnv("rabbitMQ_dlx").(string)

	rabbitMQ := RabbitMQ{
		User:              config.GetEnv("rabbitMQ_user").(string),
		Password:          config.GetEnv("rabbitMQ_pass").(string),
		Host:              "localhost", //config.GetEnv("rabbitMQ_host").(string),
		Port:              config.GetEnv("rabbitMQ_port").(string),
		Vhost:             config.GetEnv("rabbitMQ_vhost").(string),
		ConsumerQueueName: config.GetEnv("rabbitMQ_queue").(string),
		ConsumerName:      config.GetEnv("rabbitMQ_consumer").(string),
		AutoAck:           true,
		Args:              rabbitMQArgs,
	}
	return &rabbitMQ
}

func (r *RabbitMQ) Connect() *amqp.Channel {
	dsn := "amqp://" + r.User + ":" + r.Password + "@" + r.Host + ":" + r.Port + r.Vhost
	conn, err := amqp.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")

	r.Channel, err = conn.Channel()
	failOnError(err, "Failed to open a channel")

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

	failOnError(err, "Failed to declare a queue")

	err = r.Channel.QueueBind(
		q.Name,
		routingKey,
		config.GetEnv("rabbitMQ_exchange").(string),
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	incomingMessage, err := r.Channel.Consume(
		q.Name,
		r.ConsumerName,
		r.AutoAck,
		false,
		false,
		false,
		r.Args,
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for message := range incomingMessage {
			log.Info().Str("context", "RabbitMQ").Msg("New incoming message")
			messageChannel <- message
		}
		if err := r.Channel.Close(); err != nil {
			logella.ErrCtx(err, "Failed to close RabbitMQ channel")
		} else {
			log.Info().Str("context", "RabbitMQ").Msg("RabbitMQ channel closed gracefully")
		}
		close(messageChannel)
	}()
}

func (r *RabbitMQ) Publish(message string, contentType string, exchange string, routingKey string) {
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
	failOnError(err, "Failed to publish message:")
	log.Info().Str("context", "RabbitMQ").Msgf("Message published to exchange '%s' with routing key '%s'", exchange, routingKey)
}

func failOnError(err error, msg string) {
	if err != nil {
		logella.ErrCtx(err, msg)
	}
}
