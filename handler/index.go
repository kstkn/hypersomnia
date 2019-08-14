package handler

import (
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
		if err := h.tmpl.Execute(w, struct {
			Envs        []string
			JsTemplates template.HTML
		}{
			append(h.localClient.ListEnvs(), h.dashboardClient.ListEnvs()...),
			templates.JsTemplates,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
