package format

import (
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/serenize/snaker"
	"strings"
)

func Value(v *registry.Value, r int) string {
	// default format is tabbed plus the value plus new line
	fparts := []string{"", "\"%s\": %s", "\n"}
	for i := 0; i < r; i++ {
		fparts[0] += "  "
	}
	// its just a primitive of sorts so return
	if len(v.Values) == 0 {
		def := "0,"
		if v.Type == "string" {
			def = "\"\","
		}
		return fmt.Sprintf(strings.Join(fparts, ""), snaker.CamelToSnake(v.Name), def)
	}

	// this thing has more things, it's complex
	var vals []interface{}
	if r == 0 {
		fparts[1] = "{"
	} else {
		fparts[1] += "{"
		vals = []interface{}{snaker.CamelToSnake(v.Name), ""}
	}

	for _, val := range v.Values {
		fparts = append(fparts, "%s")
		vals = append(vals, Value(val, r+1))
	}

	// at the end
	l := len(fparts) - 1
	vl := len(vals) - 1
	vals[vl] = strings.TrimRight(fmt.Sprintf("%v", vals[vl]), ",\n") + "\n"

	for i := 0; i < r; i++ {
		fparts[l] += "  "
	}
	fparts = append(fparts, "},\n")

	return fmt.Sprintf(strings.Join(fparts, ""), vals...)
}

func RequestTemplate(v *registry.Value, r int) string {
	return strings.TrimRight(Value(v, r), ",\n") + "\n"
}
