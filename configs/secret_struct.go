package configs

type Secrets struct {
	// Redis Redis `mapstructure:",squash"`
	// Mongo Mongo `mapstructure:",squash"`
	// Postgres Postgres `mapstructure:",squash"`
	// RabbitMQ RabbitMQ `mapstructure:",squash"`
}

// type Redis struct {
// 	Host         string `mapstructure:"redis_host"`
// 	Port         int    `mapstructure:"redis_port"`
// 	InstanceName string `mapstructure:"redis_instance_name"`
// 	Password     string `mapstructure:"redis_password"`
// }

// type Mongo struct {
// 	Host       string `mapstructure:"mongo_host"`
// 	Port       int    `mapstructure:"mongo_port"`
// 	Database   string `mapstructure:"mongo_database"`
// 	Username   string `mapstructure:"mongo_user"`
// 	Password   string `mapstructure:"mongo_password"`
// 	IsDebug    bool   `mapstructure:"mongo_is_debug"`
// 	Options    string `mapstructure:"mongo_options"`
// }

// type Postgres struct {
// 	Host     string `mapstructure:"postgres_host"`
// 	Port     int    `mapstructure:"postgres_port"`
// 	Database string `mapstructure:"postgres_database"`
// 	Username string `mapstructure:"postgres_user"`
// 	Password string `mapstructure:"postgres_password"`
// 	IsDebug  bool   `mapstructure:"postgres_is_debug"`
// 	SSLMode  string `mapstructure:"postgres_ssl_mode"`
// }

// type RabbitMQ struct {
// 	Host     string `mapstructure:"rabbitmq_host"`
// 	Port     int    `mapstructure:"rabbitmq_port"`
// 	Username string `mapstructure:"rabbitmq_user"`
// 	Password string `mapstructure:"rabbitmq_password"`
// 	VHost    string `mapstructure:"rabbitmq_vhost"`
// }
