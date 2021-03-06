package micro

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/micro/go-micro/registry"
)

type MultiWebClient struct {
	httpClient http.Client
	webClients map[string]WebClient
}

func NewMultiWebClient(envs map[string]string) MultiWebClient {
	c := MultiWebClient{
		http.Client{},
		map[string]WebClient{},
	}
	for env, uri := range envs {
		c.webClients[env] = NewWebClient(uri)
	}
	return c
}

func (c MultiWebClient) ListEnvs() []string {
	v := make([]string, 0, len(c.webClients))
	for name := range c.webClients {
		v = append(v, name)
	}
	return v
}

func (c MultiWebClient) ListServices(env string) ([]*registry.Service, error) {
	return c.webClients[env].ListServices()
}

func (c MultiWebClient) GetService(env, name string) (*registry.Service, error) {
	return c.webClients[env].GetService(name)
}

func (c MultiWebClient) Call(ctx context.Context, env, service, endpoint string, body map[string]interface{}, response *json.RawMessage) error {
	req := c.webClients[env].NewRequest(service, endpoint, body)

	return c.webClients[env].Call(ctx, req, response)
}
