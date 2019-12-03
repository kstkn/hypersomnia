package handler

import (
	"html/template"
	"net/http"

	"github.com/gietos/hypersomnia/micro"
	"github.com/gietos/hypersomnia/templates"
)

type IndexHandler struct {
	tmpl        *template.Template
	localClient micro.Client
	webClient   micro.Client
}

func NewIndexHandler(localClient micro.Client, webClient micro.Client) IndexHandler {
	tmpl := template.Must(template.New("index").Parse(templates.Index))
	return IndexHandler{tmpl, localClient, webClient}
}

func (h IndexHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h.tmpl.Execute(w, struct {
			Envs        []string
			JsTemplates template.HTML
		}{
			append(h.localClient.ListEnvs(), h.webClient.ListEnvs()...),
			templates.JsTemplates,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
