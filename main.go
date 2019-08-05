package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gietos/hypersomnia/config"
	"github.com/gietos/hypersomnia/handler"
	"github.com/gietos/hypersomnia/micro"
	"github.com/gietos/hypersomnia/templates"

	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/registry/mdns"
)

//go:generate go run templates.go

var conf config.Config
var reg registry.Registry
var localClient micro.Client
var dashboardClient micro.Client
var tmpl *template.Template

func check(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func registerHandlers() {
	http.HandleFunc("/", handler.IndexHandler{
		Tmpl:            tmpl,
		LocalClient:     localClient,
		DashboardClient: dashboardClient,
	}.Handle())

	http.HandleFunc("/services", handler.ServicesHandler{
		LocalClient:     localClient,
		DashboardClient: dashboardClient,
	}.Handle())

	http.HandleFunc("/service", handler.ServiceHandler{
		LocalClient:     localClient,
		DashboardClient: dashboardClient,
	}.Handle())

	http.HandleFunc("/call", handler.CallHandler{
		LocalClient:     localClient,
		DashboardClient: dashboardClient,
	}.Handle())
}

func main() {
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
		envs = config.GetEnvMap(conf.Environments)
	}
	dashboardClient = micro.DashboardClient{
		Envs: envs,
	}

	tmpl, err = template.New("index").Parse(templates.Index)
	check(err)

	registerHandlers()

	log.Println("Starting web server on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
