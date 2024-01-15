package config

type database struct {
	name     string `mapstructure:"DATABASE_NAME"`
	user     string `mapstructure:"DATABASE_USER"`
	password string `mapstructure:"DATABASE_PASSWORD"`
	sslMode  string `mapstructure:"DATABASE_SSL_MODE"`
	port     string `mapstructure:"DATABASE_PORT"`
}

type jwt struct {
	secret    string `mapstructure:"JWT_SECRET"`
	expiredIn string `mapstructure:"JWT_EXPIREDIN"`
}

type rabbitMQ struct {
	user     string `mapstructure:"RABBITMQ_USER"`
	pass     string `mapstructure:"RABBITMQ_PASS"`
	host     string `mapstructure:"RABBITMQ_HOST"`
	port     string `mapstructure:"RABBITMQ_PORT"`
	vhost    string `mapstructure:"RABBITMQ_VHOST"`
	consumer string `mapstructure:"RABBITMQ_CONSUMER"`
	queue    string `mapstructure:"RABBITMQ_QUEUE"`
	exchange string `mapstructure:"RABBITMQ_EXCHANGE"`	
	dlx string `mapstructure:"RABBITMQ_DLX"`
	routingKey
}

type routingKey struct {
	userCreate string `mapstructure:"ROUTINGKEY_USERCREATE"`
}

type grpc struct {
	port int `mapstructure:"GRPC_PORT"`
}

type storage struct {
	localPath string `mapstructure:"LOCAL_STORAGE_PATH"`
	localFile string `mapstructure:"LOCAL_FILE"`
}
