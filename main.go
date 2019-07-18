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

type indexView struct {
	Services []*registry.Service
}

type request struct {
	Service  string
	Endpoint string
	Body     map[string]interface{}
}

type response struct {
	Body string
	Time string
}

func main() {
	var conf config
	var reg registry.Registry

	if err := envconfig.Process("hypersomnia", &conf); err != nil {
		log.Fatal(err.Error())
	}

	if conf.Registry == "mdns" {
		reg = mdns.NewRegistry()
	} else {
		reg = consul.NewRegistry()
	}

	rpcRequestTimeout, err := time.ParseDuration(conf.RpcRequestTimeout)
	if err != nil {
		log.Fatal(err.Error())
	}

	cl := client.NewClient(client.Registry(reg))

	tmpl, err := template.New("index").Funcs(template.FuncMap{
		"id": func(v string) string {
			return strings.ReplaceAll(snaker.CamelToSnake(v), ".", "-")
		},
		"formatRequestTemplate": func(v *registry.Value) string {
			return format.RequestTemplate(v, 0)
		},
	}).Parse(templates.Index)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var services []*registry.Service
		if services, err = reg.ListServices(); err != nil {
			log.Fatal(err.Error())
		}
		sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

		var servicesWithEndpoints []*registry.Service
		for _, service := range services {
			var serviceInfo []*registry.Service
			if serviceInfo, err = reg.GetService(service.Name); err != nil {
				log.Println(err.Error())
			}
			if len(serviceInfo) == 0 {
				continue
			}
			servicesWithEndpoints = append(servicesWithEndpoints, serviceInfo[0])
		}
		view := indexView{Services: servicesWithEndpoints}

		if err := tmpl.Execute(w, view); err != nil {
			fmt.Fprintln(w, err.Error())
		}
	})

	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		request := &request{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(request)
		if err != nil {
			resp := &response{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		var serviceResponse json.RawMessage

		start := time.Now()
		serviceRequest := cl.NewRequest(
			request.Service,
			request.Endpoint,
			request.Body,
			client.WithContentType("application/json"),
		)
		err = cl.Call(
			context.Background(),
			serviceRequest,
			&serviceResponse,
			client.WithRequestTimeout(rpcRequestTimeout),
		)

		response := response{
			Time: time.Since(start).Round(time.Millisecond).String(),
		}

		if err != nil {
			response.Body = err.Error()
		} else {
			response.Body = string(serviceResponse)
		}
		bytes, _ := json.Marshal(response)
		fmt.Fprintln(w, string(bytes))
	})

	log.Println("Starting webserver on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
