// +build ignore

package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	src := `// Code generated by go generate; DO NOT EDIT.
package templates; const Index = ` + "`{{ .Contents }}`\nconst JsTemplates = " + "`{{ .JsTemplates }}`"
	indexHtml, err := ioutil.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}

	jsHtml, err := ioutil.ReadFile("templates/js.html")
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("templates/index.go")
	defer f.Close()
	tmpl, _ := template.New("").Parse(src)
	tmpl.Execute(f, struct {
		Contents    string
		JsTemplates string
	}{
		string(indexHtml),
		string(jsHtml),
	})
}
