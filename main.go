package main

import (
	"github.com/gietos/hypersomnia/config"
	"github.com/gietos/hypersomnia/handler"
	"github.com/gietos/hypersomnia/micro"
	"log"
	"net/http"

	"github.com/micro/go-micro/client"
)

//go:generate go run templates.go

func main() {
	conf := config.NewConfig()
	log.Println("Using [" + conf.Registry + "] registry")

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

	log.Println("Starting web server on " + conf.GetAddr())
	s := &http.Server{
		Addr: conf.GetAddr(),
	}
	log.Fatal(s.ListenAndServe())
}
