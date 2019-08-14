// +build ignore

package main

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"os"
	"regexp"
	"text/template"
)

func main() {
	src := `// Code generated by go generate; DO NOT EDIT.
package templates; const Index = ` + "`{{ .Contents }}`\nconst JsTemplates = " + "`{{ .JsTemplates }}`"
	h, _ := os.Open("templates/index.html")
	defer h.Close()

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)

	indexHtmlBuf := new(bytes.Buffer)
	if err := m.Minify("text/html", indexHtmlBuf, h); err != nil {
		panic(err)
	}

	hj, _ := os.Open("templates/js.html")
	defer hj.Close()

	jsHtmlBuf := new(bytes.Buffer)
	if err := m.Minify("text/html", jsHtmlBuf, hj); err != nil {
		panic(err)
	}

	f, _ := os.Create("templates/index.go")
	defer f.Close()
	tmpl, _ := template.New("").Parse(src)
	tmpl.Execute(f, struct {
		Contents    string
		JsTemplates string
	}{
		indexHtmlBuf.String(),
		jsHtmlBuf.String(),
	})
}