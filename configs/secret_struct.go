package configs

type Secrets struct {
	Redis Redis `mapstructure:",squash"`
}

type Redis struct {
	Host         string `mapstructure:"redis_host"`
	Port         int    `mapstructure:"redis_port"`
	Password     string `mapstructure:"redis_password"`
	InstanceName string `mapstructure:"redis_instance_name"`
}
