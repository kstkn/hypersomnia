package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gietos/hypersomnia/micro"
)

type ServiceHandler struct {
	localClient     micro.LocalClient
	dashboardClient micro.DashboardClient
}

func NewServiceHandler(localClient micro.LocalClient, dashboardClient micro.DashboardClient) ServiceHandler {
	return ServiceHandler{localClient, dashboardClient}
}

func (h ServiceHandler) getClient(env string) micro.Client {
	if env == micro.EnvLocal {
		return h.localClient
	}
	return h.dashboardClient
}

func (h ServiceHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
			Name        string
		}{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		service, err := h.getClient(req.Environment).GetService(req.Environment, req.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bytes, _ := json.Marshal(service)
		w.Write(bytes)
	}
}
