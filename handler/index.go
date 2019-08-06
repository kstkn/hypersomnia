package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gietos/hypersomnia/micro"
	"github.com/gietos/hypersomnia/templates"
)

type IndexHandler struct {
	tmpl            *template.Template
	localClient     micro.LocalClient
	dashboardClient micro.DashboardClient
}

func NewIndexHandler(localClient micro.LocalClient, dashboardClient micro.DashboardClient) IndexHandler {
	tmpl := template.Must(template.New("index").Parse(templates.Index))
	return IndexHandler{tmpl, localClient, dashboardClient}
}

func (h IndexHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.tmpl.Execute(w, struct {
			Envs        []string
			JsTemplates template.HTML
		}{
			append(h.localClient.ListEnvs(), h.dashboardClient.ListEnvs()...),
			templates.JsTemplates,
		})
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}
}
