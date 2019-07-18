package format

import (
	"encoding/json"
	"github.com/micro/go-micro/registry"
)

func Pack(v *registry.Value) map[string]interface{} {
	m := map[string]interface{}{}
	for _, i := range v.Values {
		if len(i.Values) == 0 {
			if i.Type == "string" {
				m[i.Name] = ""
			} else {
				m[i.Name] = 0
			}
		} else {
			m[i.Name] = Pack(i)
		}
	}
	return m
}

func RequestTemplateAsString(v *registry.Value) string {
	js, err := json.Marshal(Pack(v))
	if err != nil {
		return ""
	}
	return string(js)
}
