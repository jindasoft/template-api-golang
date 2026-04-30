package configs

import (
	"log"
	"strings"

	"github.com/jindasoft/jinda-platform/xlogger"
	"github.com/jindasoft/jinda-platform/xvault"
	"github.com/spf13/viper"
)

func GetSecret() *Secrets {
	var secret Secrets

	conf := GetConfig()
	env := conf.App.Env

	if env == "local" {
		// local env (secret.json)
		viper.AddConfigPath(strings.Join([]string{"./configs"}, "/"))
		viper.SetConfigName("secret")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Failed to read secret file: %s", err)
		}

		err = viper.Unmarshal(&secret)
		if err != nil {
			log.Fatalf("Failed to unmarshal secret file: %s", err)
		}

		xlogger.SysInfof("Fetching secrets from env")
	} else {
		// vault
		vaultUrl := conf.Service.VaultUrl
		appName := conf.App.Name
		kvName := conf.App.KvName

		vaultSecret, err := xvault.LoadVaultSecret(vaultUrl, appName, kvName, secret)
		if err != nil {
			log.Fatalf("Failed to load vault secret: %s", err)
		}

		secret = *vaultSecret

		xlogger.SysInfof("Fetching secrets from vault")
	}

	return &secret
}
