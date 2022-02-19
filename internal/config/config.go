package config

import (
	"bytes"
	"log"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/go-playground/validator.v9"
)

const (
	cfgEnvPrefix = "smart-office"
)

type (
	Config struct {
		JWT           JWT           `mapstructure:"jwt"`
		LocalServer   LocalServer   `mapstructure:"localserver"`
		CentralServer CentralServer `mapstructure:"centralserver"`
		MQTT          MQTT          `maprstructure:"mqtt"`
	}

	JWT struct {
		Secret     string `mapstructure:"secret"`
		Expiration int    `mapstructure:"expiration"`
	}

	LocalServer struct {
		Address string `mapstructure:"address"`
	}

	CentralServer struct {
		Address string `mapstructure:"address"`
	}

	MQTT struct {
		Host   string
		Port   string
		Scheme string
	}
)

func (c Config) Validate() error {
	return validator.New().Struct(c)
}

func Init() Config {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewReader([]byte(defaultConfig))); err != nil {
		log.Fatalf("error loading default configs: %s", err.Error())
	}
	v.SetEnvPrefix(cfgEnvPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.UnmarshalExact(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config into struct: %s", err.Error())
	}
	return cfg
}

const defaultConfig = `
localserver:
  address: :9889
centralserver:
  address: :45554
jwt:
  secret: jdnfksdmfks
  expiration: 60
mqtt:
   port: 1883
   scheme: tcp
   host: localhost
  `
