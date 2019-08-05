package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gietos/hypersomnia/micro"
	"github.com/gietos/hypersomnia/templates"
)

type IndexHandler struct {
	Tmpl            *template.Template
	LocalClient     micro.Client
	DashboardClient micro.Client
}

func (h IndexHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h.Tmpl.Execute(w, struct {
			Envs        []string
			JsTemplates template.HTML
		}{
			append(h.LocalClient.ListEnvs(), h.DashboardClient.ListEnvs()...),
			template.HTML(templates.JsTemplates),
		})
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}
}
