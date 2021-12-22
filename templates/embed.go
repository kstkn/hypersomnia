package templates

import _ "embed"

//go:embed index.html
var Index string

//go:embed js.html
var JsTemplates string
