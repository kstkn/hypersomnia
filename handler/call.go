package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/metadata"

	"github.com/gietos/hypersomnia/micro"
)

type CallHandler struct {
	localClient     micro.LocalClient
	dashboardClient micro.DashboardClient
}

func NewCallHandler(localClient micro.LocalClient, dashboardClient micro.DashboardClient) CallHandler {
	return CallHandler{localClient, dashboardClient}
}

func createContext(correlationId string) context.Context {
	ctx := context.Background()
	md := metadata.Metadata{}
	ctx = metadata.NewContext(ctx, md)
	md["X-Correlation-Id"] = correlationId
	return ctx
}

func (h CallHandler) getClient(env string) micro.Client {
	if env == micro.EnvLocal {
		return h.localClient
	}
	return h.dashboardClient
}

func (h CallHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &struct {
			Environment string
			Service     string
			Endpoint    string
			Body        map[string]interface{}
		}{}

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var serviceResponse json.RawMessage
		start := time.Now()
		correlationId, _ := uuid.NewRandom()
		ctx := createContext(correlationId.String())

		resp := struct {
			CorrelationId string
			Time          string
			Body          string
		}{correlationId.String(), "", ""}

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
}
