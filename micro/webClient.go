package micro

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/micro/go-micro/registry"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
)

type WebClient struct {
	client     client.Client
	httpClient http.Client
	baseUri    string
}

func NewWebClient(baseUri string) WebClient {
	return WebClient{
		client.DefaultClient,
		http.Client{},
		baseUri,
	}
}

func (c WebClient) NewRequest(service, endpoint string, req interface{}, reqOpts ...client.RequestOption) client.Request {
	return c.client.NewRequest(service, endpoint, req, reqOpts...)
}

func (c WebClient) Call(ctx context.Context, req client.Request, rsp interface{}, _ ...client.CallOption) error {
	u, _ := url.Parse(c.baseUri + "/rpc")

	payload := struct {
		Service  string      `json:"service"`
		Endpoint string      `json:"endpoint"`
		Request  interface{} `json:"request"`
	}{
		Service:  req.Service(),
		Endpoint: req.Endpoint(),
		Request:  req.Body(),
	}

	payloadString, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(payloadString))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/json")
	c.enrichFromContext(ctx, r)

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(responseBody, rsp); err != nil {
		return err
	}
	return nil
}

func (c WebClient) enrichFromContext(ctx context.Context, r *http.Request) {
	md, has := metadata.FromContext(ctx)
	if !has {
		return
	}

	for k, v := range md {
		r.Header.Add(k, v)
	}
}

func (c WebClient) ListServices() ([]*registry.Service, error) {
	u, _ := url.Parse(c.baseUri + "/registry")
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

func (c WebClient) GetService(name string) (*registry.Service, error) {
	u, _ := url.Parse(c.baseUri + "/registry")
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

func (c WebClient) Init(...client.Option) error {
	return nil
}

func (c WebClient) Options() client.Options {
	return client.Options{}
}

func (c WebClient) NewMessage(topic string, msg interface{}, opts ...client.MessageOption) client.Message {
	return c.client.NewMessage(topic, msg, opts...)
}

func (c WebClient) Stream(_ context.Context, _ client.Request, _ ...client.CallOption) (client.Stream, error) {
	return nil, nil
}

func (c WebClient) Publish(_ context.Context, _ client.Message, _ ...client.PublishOption) error {
	return nil
}

func (c WebClient) String() string {
	return "micro-web"
}
