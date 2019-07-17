package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gietos/hypersomnia/format"
	"github.com/gietos/hypersomnia/templates"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
	"github.com/serenize/snaker"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

//go:generate go run templates.go

type Config struct {
	Addr     string `default:":8083"`
	Registry string `default:"mdns"`
}

type ServicesIndexView struct {
	Services []*registry.Service
}

type Call struct {
	Service  string
	Endpoint string
	Body     map[string]interface{}
}

type Response struct {
	Body string
	Time string
}

func main() {
	var c Config
	var reg registry.Registry
	var tmpl *template.Template

	err := envconfig.Process("hypersomnia", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	if c.Registry == "mdns" {
		reg = mdns.NewRegistry()
	} else {
		reg = consul.NewRegistry()
	}

	tmpl, err = template.New("index").Funcs(template.FuncMap{
		"id": func(v string) string {
			return strings.ReplaceAll(snaker.CamelToSnake(v), ".", "-")
		},
		"formatEndpoint": func(v *registry.Value) string {
			return format.Endpoint(v, 0)
		},
	}).Parse(templates.Index)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		services, _ := reg.ListServices()
		sort.Slice(services, func(i, j int) bool { return services[i].Name < services[j].Name })

		var results []*registry.Service
		for _, service := range services {
			service, _ := registry.GetService(service.Name)
			results = append(results, service[0])
		}
		view := ServicesIndexView{Services: results}

		if err := tmpl.Execute(w, view); err != nil {
			fmt.Fprintln(w, err.Error())
		}
	})
	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		call := &Call{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(call)
		if err != nil {
			resp := &Response{
				Body: err.Error(),
			}
			bytes, _ := json.Marshal(resp)
			fmt.Fprintln(w, string(bytes))
			return
		}

		var serviceResponse json.RawMessage

		start := time.Now()
		serviceRequest := (*cmd.DefaultOptions().Client).NewRequest(call.Service, call.Endpoint, call.Body, client.WithContentType("application/json"))
		err = (*cmd.DefaultOptions().Client).Call(context.Background(), serviceRequest, &serviceResponse, client.WithRequestTimeout(time.Minute))

		response := Response{
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

	log.Println("Starting webserver on " + c.Addr)
	s := &http.Server{
		Addr: c.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
