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
// 	Password     string `mapstructure:"redis_password"`
// 	InstanceName string `mapstructure:"redis_instance_name"`
// }

// type Mongo struct {
// 	Connection string `mapstructure:"mongo_connection"`
// 	Options    string `mapstructure:"mongo_options"`
// 	DbName     string `mapstructure:"mongo_db_name"`
// 	IsDebug    bool   `mapstructure:"mongo_is_debug"`
// }

// type Postgres struct {
// 	Host     string `mapstructure:"postgres_host"`
// 	Port     int    `mapstructure:"postgres_port"`
// 	Database string `mapstructure:"postgres_database"`
// 	User     string `mapstructure:"postgres_user"`
// 	Password string `mapstructure:"postgres_password"`
// 	SSLMode  string `mapstructure:"postgres_ssl_mode"`
// 	IsDebug  bool   `mapstructure:"postgres_is_debug"`
// }

// type RabbitMQ struct {
// 	Host     string `mapstructure:"rabbitmq_host"`
// 	Port     int    `mapstructure:"rabbitmq_port"`
// 	User     string `mapstructure:"rabbitmq_user"`
// 	Password string `mapstructure:"rabbitmq_password"`
// 	VHost    string `mapstructure:"rabbitmq_vhost"`
// }
