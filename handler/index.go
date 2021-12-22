package handler

import (
	"html/template"
	"net/http"

	"github.com/kstkn/hypersomnia/micro"
	"github.com/kstkn/hypersomnia/templates"
)

type IndexHandler struct {
	tmpl        *template.Template
	localClient micro.ClientWrapper
	webClient   micro.ClientWrapper
}

func NewIndexHandler(localClient micro.ClientWrapper, webClient micro.ClientWrapper) IndexHandler {
	tmpl := template.Must(template.New("index").Parse(templates.Index))
	return IndexHandler{tmpl, localClient, webClient}
}

func (h IndexHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if err := h.tmpl.Execute(w, struct {
		Envs        []string
		JsTemplates template.HTML
	}{
		append(h.localClient.ListEnvs(), h.webClient.ListEnvs()...),
		template.HTML(templates.JsTemplates),
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
