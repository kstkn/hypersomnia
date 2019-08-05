package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gietos/hypersomnia/micro"

	"github.com/google/uuid"
	"github.com/micro/go-micro/metadata"
)

type CallHandler struct {
	LocalClient     micro.Client
	DashboardClient micro.Client
}

func createContext(correlationId string) context.Context {
	ctx := context.Background()
	md := metadata.Metadata{}
	ctx = metadata.NewContext(ctx, md)
	md["X-Correlation-Id"] = correlationId
	return ctx
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
		err := decoder.Decode(req)
		if err != nil {
			bytes, _ := json.Marshal(&struct{ Body string }{Body: err.Error()})
			fmt.Fprintln(w, string(bytes))
			return
		}

		var serviceResponse json.RawMessage
		start := time.Now()
		correlationId, _ := uuid.NewRandom()
		ctx := createContext(correlationId.String())

		if req.Environment == micro.EnvLocal {
			err = h.LocalClient.Call(
				ctx,
				req.Environment,
				req.Service,
				req.Endpoint,
				req.Body,
				&serviceResponse,
			)
		} else {
			err = h.DashboardClient.Call(
				ctx,
				req.Environment,
				req.Service,
				req.Endpoint,
				req.Body,
				&serviceResponse,
			)
		}

		resp := struct {
			CorrelationId string
			Body          string
			Time          string
		}{
			CorrelationId: correlationId.String(),
			Time:          time.Since(start).Round(time.Millisecond).String(),
		}
		if err != nil {
			resp.Body = err.Error()
		} else {
			resp.Body = string(serviceResponse)
		}
		bytes, _ := json.Marshal(resp)
		fmt.Fprintln(w, string(bytes))
	}
}
