package micro

import (
	"context"
	"encoding/json"

	"github.com/micro/go-micro/registry"
)

type ClientWrapper interface {
	ListEnvs() []string
	ListServices(env string) ([]*registry.Service, error)
	GetService(env, name string) (*registry.Service, error)
	Call(ctx context.Context, env, service, endpoint string, body map[string]interface{}, resp *json.RawMessage) error
}

const EnvLocal = "local"
