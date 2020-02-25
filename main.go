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

	http.Handle("/", handler.NewIndexHandler(localClient, webClient))
	http.Handle("/service", handler.NewServiceHandler(localClient, webClient))
	http.Handle("/services", handler.NewServicesHandler(localClient, webClient))
	http.Handle("/call", handler.NewCallHandler(localClient, webClient))
}

func main() {
	log.Info("starting web server on " + conf.Addr)
	s := &http.Server{
		Addr: conf.Addr,
	}
	log.Fatal(s.ListenAndServe())
}
