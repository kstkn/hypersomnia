package main

import (
	"github.com/kstkn/hypersomnia/config"
	"github.com/kstkn/hypersomnia/handler"
	"github.com/kstkn/hypersomnia/micro"
	"github.com/micro/go-micro/client"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//go:generate go run templates.go

var conf config.Config

func init() {
	conf = config.NewConfig()

	log.SetLevel(conf.LogLevel)
	log.Debugf("configuration %+v", conf)

	localClient := micro.NewLocalClient(
		client.NewClient(client.Registry(conf.GetRegistry())),
		conf.GetRegistry(),
		conf.RpcRequestTimeout,
	)
	webClient := micro.NewMultiWebClient(conf.Environments)

	http.HandleFunc("/", handler.NewIndexHandler(localClient, webClient).Handle())
	http.HandleFunc("/service", handler.NewServiceHandler(localClient, webClient).Handle())
	http.HandleFunc("/services", handler.NewServicesHandler(localClient, webClient).Handle())
	http.HandleFunc("/call", handler.NewCallHandler(localClient, webClient).Handle())
}

func main() {
	log.Info("starting web server on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
