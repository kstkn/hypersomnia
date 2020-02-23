package config

import (
	"time"

	"github.com/kstkn/envconfig"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel          log.Level     `default:"info" split_words:"true"`
	Addr              string        `default:"localhost:8083"`
	Registry          string        `default:"consul"`
	RegistryAddr      string        `default:"localhost:8500" split_words:"true"`
	RpcRequestTimeout time.Duration `default:"1m" split_words:"true" `
	Environments      map[string]string
}

func NewConfig() Config {
	c := Config{}
	if err := envconfig.Process("hypersomnia", &c); err != nil {
		panic(err.Error())
	}
	return c
}

func (c Config) GetRegistry() registry.Registry {
	if c.Registry == "consul" {
		return consul.NewRegistry(registry.Addrs(c.RegistryAddr))
	}
	return mdns.NewRegistry()
}
