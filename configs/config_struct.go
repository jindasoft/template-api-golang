package configs

import (
	"time"

	"go.uber.org/zap/zapcore"
)

type Configs struct {
	App     App     `mapstructure:"app"`
	Server  Server  `mapstructure:"server"`
	Service Service `mapstructure:"service"`
}

type App struct {
	Name   string `mapstructure:"name"`
	KvName string `mapstructure:"kv_name"`
	Env    string `mapstructure:"env"`
	Port   int    `mapstructure:"port"`
}

type Server struct {
	PrettyPrint  bool `mapstructure:"pretty_print"`
	LogLevel     zapcore.Level
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type Service struct {
	Debug    bool   `mapstructure:"debug"`
	VaultURL string `mapstructure:"vault_url"`
}
