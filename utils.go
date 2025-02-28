package vcf

import (
	"bytes"
	"fmt"
	"text/template"
)

func containsNotEmpty(vals ...string) bool {
	for _, v := range vals {
		if len(v) > 0 {
			return true
		}
	}

	return false
}

func getOutput(v any) (out string, err error) {
	buf := bytes.NewBuffer(nil)
	if err = parsed.Execute(buf, v); err != nil {
		return
	}

	out = buf.String()
	return
}

func getParsed(tmpl string) (out *template.Template, err error) {
	if out, err = parsed.Parse(tmpl); err != nil {
		err = fmt.Errorf("error parsing template: %v", err)
		return
	}

	return
}
