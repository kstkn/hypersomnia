package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/gietos/hypersomnia/micro"

	"github.com/micro/go-micro/registry"
)

type ServicesHandler struct {
	localClient     micro.LocalClient
	dashboardClient micro.DashboardClient
}

func NewServicesHandler(localClient micro.LocalClient, dashboardClient micro.DashboardClient) ServicesHandler {
	return ServicesHandler{localClient, dashboardClient}
}

func (h ServicesHandler) getClient(env string) micro.Client {
	if env == micro.EnvLocal {
		return h.localClient
	}
	return h.dashboardClient
}

func (h ServicesHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
		}{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		services, err := h.getClient(req.Environment).ListServices(req.Environment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		messages := make(chan *registry.Service, len(services))
		var wg sync.WaitGroup
		wg.Add(len(services))
		for _, service := range services {
			go func(service *registry.Service) {
				defer wg.Done()
				var serviceInfo *registry.Service
				serviceInfo, err := h.getClient(req.Environment).GetService(req.Environment, service.Name)
				if err != nil {
					log.WithFields(log.Fields{
						"environment": req.Environment,
						"service":     service.Name,
					}).Error(err)
					messages <- nil
					return
				}
				if len(serviceInfo.Endpoints) == 0 {
					service = nil
				}
				messages <- service
			}(service)
		}
		wg.Wait()
		close(messages)

		var results []*registry.Service
		done := make(chan bool)
		go func() {
			for service := range messages {
				if service != nil {
					results = append(results, service)
				}
			}
			done <- true
		}()

		<-done
		bytes, _ := json.Marshal(results)
		w.Write(bytes)
	}
}
