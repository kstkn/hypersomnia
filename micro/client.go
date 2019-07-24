package micro

import (
	"encoding/json"
	"github.com/micro/go-micro/registry"
)

type Client interface {
	ListEnvs() []string
	ListServices(env string) ([]*registry.Service, error)
	GetService(env, name string) (*registry.Service, error)
	Call(env, service, endpoint string, body map[string]interface{}, resp *json.RawMessage) error
}

const EnvLocal = "local"
