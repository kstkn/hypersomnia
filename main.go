package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gietos/hypersomnia/micro"
	"github.com/gietos/hypersomnia/templates"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
)

//go:generate go run templates.go

type config struct {
	Addr              string `default:"localhost:8083"`
	Registry          string `default:"mdns"`
	RpcRequestTimeout string `default:"1m"`
	Environments      string
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getEnvMap(s string) map[string]string {
	envs := strings.Split(s, ";")
	r := regexp.MustCompile(`(?P<Name>[a-z]+?):(?P<BaseUrl>.+)`)
	m := map[string]string{}
	for _, env := range envs {
		m[r.FindStringSubmatch(env)[1]] = r.FindStringSubmatch(env)[2]
	}
	return m
}

func main() {
	var conf config
	var reg registry.Registry
	var localClient micro.Client
	var dashboardClient micro.Client

	err := envconfig.Process("hypersomnia", &conf)
	check(err)

	fmt.Println()
	if conf.Registry == "mdns" {
		reg = mdns.NewRegistry()
	} else {
		reg = consul.NewRegistry()
	}

	rpcRequestTimeout, err := time.ParseDuration(conf.RpcRequestTimeout)
	check(err)

	localClient = micro.LocalClient{
		MicroClient:    client.NewClient(client.Registry(reg)),
		MicroRegistry:  reg,
		RequestTimeout: rpcRequestTimeout,
	}

	envs := map[string]string{}
	if conf.Environments != "" {
		envs = getEnvMap(conf.Environments)
	}
	dashboardClient = micro.DashboardClient{
		Envs: envs,
	}

	tmpl, err := template.New("index").Parse(templates.Index)
	check(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, struct {
			Envs []string
		}{
			Envs: append(localClient.ListEnvs(), dashboardClient.ListEnvs()...),
		})
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)
		if err != nil {
			resp := &struct {
				Body string
			}{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		var services []*registry.Service
		if req.Environment == micro.EnvLocal {
			services, err = localClient.ListServices(req.Environment)
		} else {
			services, err = dashboardClient.ListServices(req.Environment)
		}

		if err != nil {
			resp := &struct {
				Body string
			}{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		bytes, _ := json.Marshal(services)
		fmt.Fprintln(w, string(bytes))
	})

	http.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
			Name        string
		}{}

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)
		if err != nil {
			resp := &struct {
				Body string
			}{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		var service *registry.Service
		if req.Environment == micro.EnvLocal {
			service, err = localClient.GetService(req.Environment, req.Name)
		} else {
			service, err = dashboardClient.GetService(req.Environment, req.Name)
		}
		if err != nil {
			resp := &struct {
				Body string
			}{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		bytes, _ := json.Marshal(service)
		fmt.Fprintln(w, string(bytes))
	})

	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
			Service     string
			Endpoint    string
			Body        map[string]interface{}
		}{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(req)
		if err != nil {
			bytes, _ := json.Marshal(&struct{ Body string }{Body: err.Error()})
			fmt.Fprintln(w, string(bytes))
			return
		}

		var serviceResponse json.RawMessage
		start := time.Now()
		if req.Environment == micro.EnvLocal {
			err = localClient.Call(
				req.Environment,
				req.Service,
				req.Endpoint,
				req.Body,
				&serviceResponse,
			)
		} else {
			err = dashboardClient.Call(
				req.Environment,
				req.Service,
				req.Endpoint,
				req.Body,
				&serviceResponse,
			)
		}

		resp := struct {
			Body string
			Time string
		}{
			Time: time.Since(start).Round(time.Millisecond).String(),
		}
		if err != nil {
			resp.Body = err.Error()
		} else {
			resp.Body = string(serviceResponse)
		}
		bytes, _ := json.Marshal(resp)
		fmt.Fprintln(w, string(bytes))
	})

	log.Println("Starting web server on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
