package configs

import (
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

func GetConfig() *Configs {
	var config Configs

	// viper
	viper.AddConfigPath(strings.Join([]string{"./configs"}, "/"))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("TGRENV")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %s", err)
	}

	config.Server.IdleTimeout = parseDurationWithDefault(viper.GetString("server.idle_timeout"), 1*time.Minute)
	config.Server.ReadTimeout = parseDurationWithDefault(viper.GetString("server.read_timeout"), 10*time.Second)
	config.Server.WriteTimeout = parseDurationWithDefault(viper.GetString("server.write_timeout"), 30*time.Second)

	logLevel := viper.GetString("server.log_level")
	if err := config.Server.LogLevel.UnmarshalText([]byte(logLevel)); err != nil {
		config.Server.LogLevel = zapcore.InfoLevel
	}

	return &config
}

func parseDurationWithDefault(duration string, defaultDuration time.Duration) time.Duration {
	if duration == "" {
		return defaultDuration
	}

	d, err := time.ParseDuration(duration)
	if err != nil {
		return defaultDuration
	}

	return d
}
