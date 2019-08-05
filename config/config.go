package config

import (
	"regexp"
	"strings"
)

type Config struct {
	Addr              string `default:"localhost:8083"`
	Registry          string `default:"mdns"`
	RpcRequestTimeout string `default:"1m"`
	Environments      string
}

func GetEnvMap(s string) map[string]string {
	envs := strings.Split(s, ";")
	r := regexp.MustCompile(`(?P<Name>[a-z]+?):(?P<BaseUrl>.+)`)
	m := map[string]string{}
	for _, env := range envs {
		m[r.FindStringSubmatch(env)[1]] = r.FindStringSubmatch(env)[2]
	}
	return m
}
