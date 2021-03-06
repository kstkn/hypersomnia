package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"

	"github.com/kstkn/hypersomnia/micro"
)

type CallHandler struct {
	localClient micro.ClientWrapper
	webClient   micro.ClientWrapper
}

func NewCallHandler(localClient micro.ClientWrapper, webClient micro.ClientWrapper) CallHandler {
	return CallHandler{localClient, webClient}
}

func createContext(values map[string]string) context.Context {
	ctx := context.Background()
	md := metadata.Metadata{}
	ctx = metadata.NewContext(ctx, md)
	for k, v := range values {
		md[k] = v
	}
	return ctx
}

func (h CallHandler) getClient(env string) micro.ClientWrapper {
	if env == micro.EnvLocal {
		return h.localClient
	}
	return h.webClient
}

func (h CallHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &struct {
		Environment string
		Service     string
		Endpoint    string
		Body        map[string]interface{}
		Context     map[string]string
	}{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for k, v := range req.Context {
		fmt.Println(k, v)
	}

	var serviceResponse json.RawMessage
	start := time.Now()
	correlationId, _ := uuid.NewRandom()
	req.Context["X-Correlation-Id"] = correlationId.String()
	ctx := createContext(req.Context)

	resp := struct {
		CorrelationId string
		Time          string
		Body          string
	}{correlationId.String(), "", ""}

	log.
		With("environment", req.Environment).
		With("service", req.Service).
		With("endpoint", req.Endpoint).
		With("correlationId", correlationId.String()).
		Info("sending RPC request")

	if err := h.getClient(req.Environment).Call(
		ctx,
		req.Environment,
		req.Service,
		req.Endpoint,
		req.Body,
		&serviceResponse,
	); err != nil {
		resp.Body = err.Error()
	} else {
		resp.Body = string(serviceResponse)
	}
	resp.Time = time.Since(start).Round(time.Millisecond).String()

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
