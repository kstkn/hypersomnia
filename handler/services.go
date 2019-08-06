package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func (h ServicesHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
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

		var services []*registry.Service
		if req.Environment == micro.EnvLocal {
			services, err = h.localClient.ListServices(req.Environment)
		} else {
			services, err = h.dashboardClient.ListServices(req.Environment)
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

		bytes, _ := json.Marshal(services)
		fmt.Fprintln(w, string(bytes))
	}
}
