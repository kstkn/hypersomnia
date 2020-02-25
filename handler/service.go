package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kstkn/hypersomnia/micro"
)

type ServiceHandler struct {
	localClient micro.ClientWrapper
	webClient   micro.ClientWrapper
}

func NewServiceHandler(localClient micro.ClientWrapper, webClient micro.ClientWrapper) ServiceHandler {
	return ServiceHandler{localClient, webClient}
}

func (h ServiceHandler) getClient(env string) micro.ClientWrapper {
	if env == micro.EnvLocal {
		return h.localClient
	}
	return h.webClient
}

func (h ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
