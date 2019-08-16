package micro

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
)

type DashboardClient struct {
	httpClient http.Client
	envs       map[string]string
}

func NewDashboardClient(envs map[string]string) DashboardClient {
	return DashboardClient{
		http.Client{},
		envs,
	}
}

func (c DashboardClient) ListEnvs() []string {
	v := make([]string, 0, len(c.envs))
	for name := range c.envs {
		v = append(v, name)
	}
	return v
}

func (c DashboardClient) ListServices(env string) ([]*registry.Service, error) {
	u, _ := url.Parse(c.envs[env] + "/registry")
	req, err := http.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer(nil))
	if err != nil {
		return []*registry.Service{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []*registry.Service{}, err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	services := struct {
		Services []*registry.Service
	}{}
	if err = json.Unmarshal(responseBody, &services); err != nil {
		return []*registry.Service{}, err
	}

	return services.Services, nil
}

func (c DashboardClient) GetService(env, name string) (*registry.Service, error) {
	u, _ := url.Parse(c.envs[env] + "/registry")
	q := u.Query()
	q.Set("service", name)
	u.RawQuery = q.Encode()

	service := &registry.Service{}
	req, err := http.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer(nil))
	if err != nil {
		return service, nil
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return service, err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	services := struct {
		Services []*registry.Service
	}{}

	json.Unmarshal(responseBody, &services)
	return services.Services[0], nil
}

func enrichFromContext(ctx context.Context, r *http.Request) {
	md, has := metadata.FromContext(ctx)
	if !has {
		return
	}

	for k, v := range md {
		r.Header.Add(k, v)
	}
}

func (c DashboardClient) Call(ctx context.Context, env, service, endpoint string, body map[string]interface{}, response *json.RawMessage) error {
	u, _ := url.Parse(c.envs[env] + "/rpc")

	payload := struct {
		Service  string                 `json:"service"`
		Endpoint string                 `json:"endpoint"`
		Request  map[string]interface{} `json:"request"`
	}{
		Service:  service,
		Endpoint: endpoint,
		Request:  body,
	}

	payloadString, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(payloadString))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	enrichFromContext(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(responseBody, response)
	return nil
}
