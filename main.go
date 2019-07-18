package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/gietos/hypersomnia/format"
	"github.com/gietos/hypersomnia/templates"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/serenize/snaker"
)

//go:generate go run templates.go

type config struct {
	Addr              string `default:"localhost:8083"`
	Registry          string `default:"mdns"`
	RpcRequestTimeout string `default:"1m"`
}

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getServiceView(service *registry.Service) interface{} {
	serviceView := &struct {
		Name      string
		Endpoints []interface{}
	}{
		Name: service.Name,
	}
	for _, e := range service.Endpoints {
		endpoint := &struct {
			Name            string
			RequestTemplate string
		}{
			Name:            e.Name,
			RequestTemplate: format.RequestTemplateAsString(e.Request),
		}
		serviceView.Endpoints = append(serviceView.Endpoints, endpoint)
	}
	return serviceView
}

func main() {
	var conf config
	var reg registry.Registry

	err := envconfig.Process("hypersomnia", &conf)
	check(err)

	if conf.Registry == "mdns" {
		reg = mdns.NewRegistry()
	} else {
		reg = consul.NewRegistry()
	}

	rpcRequestTimeout, err := time.ParseDuration(conf.RpcRequestTimeout)
	check(err)

	cl := client.NewClient(client.Registry(reg))

	tmpl, err := template.New("index").Funcs(template.FuncMap{
		"id": func(v string) string {
			return strings.ReplaceAll(snaker.CamelToSnake(v), ".", "-")
		},
	}).Parse(templates.Index)
	check(err)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		services, err := reg.ListServices()
		check(err)
		sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

		servicesView := struct {
			Services []interface{}
		}{}

		for _, service := range services {
			var serviceInfo []*registry.Service
			if serviceInfo, err = reg.GetService(service.Name); err != nil {
				log.Println(err.Error())
			}
			if len(serviceInfo) == 0 {
				continue
			}
			servicesView.Services = append(servicesView.Services, getServiceView(serviceInfo[0]))
		}

		if err := tmpl.Execute(w, servicesView); err != nil {
			fmt.Fprintln(w, err.Error())
		}
	})

	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Service  string
			Endpoint string
			Body     map[string]interface{}
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

		var serviceResponse json.RawMessage

		start := time.Now()
		serviceRequest := cl.NewRequest(
			req.Service,
			req.Endpoint,
			req.Body,
			client.WithContentType("application/json"),
		)
		err = cl.Call(
			context.Background(),
			serviceRequest,
			&serviceResponse,
			client.WithRequestTimeout(rpcRequestTimeout),
		)

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

	log.Println("Starting webserver on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
