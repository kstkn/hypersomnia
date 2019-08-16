package main

import (
	"github.com/gietos/hypersomnia/config"
	"github.com/gietos/hypersomnia/handler"
	"github.com/gietos/hypersomnia/micro"
	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//go:generate go run templates.go

var conf config.Config

func init() {
	log.SetLevel(log.DebugLevel)
	conf = config.NewConfig()
	log.WithFields(log.Fields{
		"registry":          conf.Registry,
		"registryAddress":   conf.RegistryAddress,
		"environments":      conf.GetEnvironments(),
		"rpcRequestTimeout": conf.GetRpcRequestTimeout(),
	}).Debug("Configuration")

	localClient := micro.NewLocalClient(
		client.NewClient(client.Registry(conf.GetRegistry())),
		conf.GetRegistry(),
		conf.GetRpcRequestTimeout(),
	)
	dashboardClient := micro.NewDashboardClient(conf.GetEnvironments())

	http.HandleFunc("/", handler.NewIndexHandler(localClient, dashboardClient).Handle())
	http.HandleFunc("/service", handler.NewServiceHandler(localClient, dashboardClient).Handle())
	http.HandleFunc("/services", handler.NewServicesHandler(localClient, dashboardClient).Handle())
	http.HandleFunc("/call", handler.NewCallHandler(localClient, dashboardClient).Handle())
}

func main() {
	log.Info("Starting web server on " + conf.GetAddr())
	s := &http.Server{
		Addr: conf.GetAddr(),
	}
	log.Fatal(s.ListenAndServe())
}
