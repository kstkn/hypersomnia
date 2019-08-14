package config

import (
	"regexp"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
)

type Config struct {
	Addr              string `default:"localhost:8083"`
	Registry          string `default:"mdns"`
	RpcRequestTimeout string `default:"1m"`
	Environments      string
}

func NewConfig() Config {
	c := Config{}
	if err := envconfig.Process("hypersomnia", &c); err != nil {
		panic(err.Error())
	}
	return c
}

func (c Config) GetAddr() string {
	return c.Addr
}

// Parses micro web environment urls into map.
func (c Config) GetEnvironments() map[string]string {
	if c.Environments == "" {
		return map[string]string{}
	}
	envs := strings.Split(c.Environments, ";")
	r := regexp.MustCompile(`(?P<Name>[a-z]+?):(?P<Url>.+)`)
	m := map[string]string{}
	for _, env := range envs {
		m[r.FindStringSubmatch(env)[1]] = r.FindStringSubmatch(env)[2]
	}
	return m
}

func (c Config) GetRpcRequestTimeout() time.Duration {
	t, err := time.ParseDuration(c.RpcRequestTimeout)
	if err != nil {
		panic(err.Error())
	}
	return t
}

func (c Config) GetRegistry() registry.Registry {
	if c.Registry == "consul" {
		return consul.NewRegistry()
	}
	return mdns.NewRegistry()
}
