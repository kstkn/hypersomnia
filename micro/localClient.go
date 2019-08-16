package micro

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
)

type LocalClient struct {
	microClient    client.Client
	microRegistry  registry.Registry
	requestTimeout time.Duration
}

func NewLocalClient(microClient client.Client, microRegistry registry.Registry, requestTimeout time.Duration) LocalClient {
	return LocalClient{
		microClient,
		microRegistry,
		requestTimeout,
	}
}

func (c LocalClient) ListEnvs() []string {
	return []string{EnvLocal}
}

func (c LocalClient) ListServices(env string) ([]*registry.Service, error) {
	if env != EnvLocal {
		return []*registry.Service{}, errors.New("local client can be used only with local environment")
	}

	services, err := c.microRegistry.ListServices()
	if err != nil {
		return []*registry.Service{}, err
	}

	return services, nil
}

func (c LocalClient) GetService(env, name string) (*registry.Service, error) {
	if env != EnvLocal {
		return nil, errors.New("local client can be used only with local environment")
	}

	services, err := c.microRegistry.GetService(name)
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

	serviceRequest := c.microClient.NewRequest(
		service,
		endpoint,
		body,
		client.WithContentType("application/json"),
	)
	err := c.microClient.Call(
		ctx,
		serviceRequest,
		response,
		client.WithRequestTimeout(c.requestTimeout),
	)

	if err != nil {
		return err
	}

	return nil
}
