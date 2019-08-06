package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gietos/hypersomnia/micro"

	"github.com/micro/go-micro/registry"
)

type ServiceHandler struct {
	localClient     micro.LocalClient
	dashboardClient micro.DashboardClient
}

func NewServiceHandler(localClient micro.LocalClient, dashboardClient micro.DashboardClient) ServiceHandler {
	return ServiceHandler{localClient, dashboardClient}
}

func (h ServiceHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			service, err = h.localClient.GetService(req.Environment, req.Name)
		} else {
			service, err = h.dashboardClient.GetService(req.Environment, req.Name)
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
	}
}
