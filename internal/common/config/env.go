package config

type Database struct {
	Name     string
	User     string
	Password string
	Ssl_mode string
	Port     string
}

type RabbitMQ struct {
	User                     string
	Password                 string
	Host                     string
	Port                     string
	Vhost                    string
	Consumer_name            string
	Consumer_queue_name      string
	Notification_ex          string
	Notification_routing_key string
}

type Grpc struct {
	Port int
}

type Storage struct {
	Local_path string
	Local_file string
}
