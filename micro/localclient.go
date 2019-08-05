package micro

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"sort"
	"time"
)

type LocalClient struct {
	MicroClient    client.Client
	MicroRegistry  registry.Registry
	RequestTimeout time.Duration
}

func (c LocalClient) ListEnvs() []string {
	return []string{EnvLocal}
}

func (c LocalClient) ListServices(env string) ([]*registry.Service, error) {
	if env != EnvLocal {
		return []*registry.Service{}, errors.New("local client can be used only with local environment")
	}

	services, err := c.MicroRegistry.ListServices()
	if err != nil {
		return []*registry.Service{}, err
	}

	// We want them sorted, so each time order is the same
	sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

	return services, nil
}

func (c LocalClient) GetService(env, name string) (*registry.Service, error) {
	if env != EnvLocal {
		return nil, errors.New("local client can be used only with local environment")
	}

	services, err := c.MicroRegistry.GetService(name)
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, errors.New("failed to get service")
	}

	return services[0], nil
}

func (c LocalClient) Call(ctx context.Context, env, service, endpoint string, body map[string]interface{}, response *json.RawMessage) error {
	if env != EnvLocal {
		return errors.New("local client can be used only with local environment")
	}

	serviceRequest := c.MicroClient.NewRequest(
		service,
		endpoint,
		body,
		client.WithContentType("application/json"),
	)
	err := c.MicroClient.Call(
		context.Background(),
		serviceRequest,
		response,
		client.WithRequestTimeout(c.RequestTimeout),
	)

	if err != nil {
		return err
	}

	return nil
}
