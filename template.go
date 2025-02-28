package vcf

import "text/template"

const tmpl = `BEGIN:VCARD
VERSION:3.0
FN:{{.FullName}}
{{if contains .FirstName .LastName .MiddleName .Prefix .Suffix }}N:{{.LastName}};{{.FirstName}};{{.MiddleName}};{{.Prefix}};{{.Suffix}}
{{end}}{{if .PhoneNumber}}TEL;TYPE={{.PhoneType}}:{{.PhoneNumber}}
{{end}}{{if .Email}}EMAIL:{{.Email}}
{{end}}{{if .Organization}}ORG:{{.Organization}}
{{end}}{{if .Title}}TITLE:{{.Title}}
{{end}}{{if contains .POBox .ExtendedAddress .StreetAddress .City .State .PostalCode .Country }}ADR:{{.POBox}};{{.ExtendedAddress}};{{.StreetAddress}};{{.City}};{{.State}};{{.PostalCode}};{{.Country}}
{{end}}{{if .URL}}URL:{{.URL}}
{{end}}{{if .Birthday}}BDAY:{{.Birthday}}
{{end}}{{if .PhotoURL}}PHOTO;TYPE={{.PhotoType}}:{{.PhotoURL}}
{{end}}{{if .Note}}NOTE:{{.Note}}
{{end}}END:VCARD`

var parsed *template.Template

func init() {
	var err error
	parsed = template.New("vcf")
	parsed.Funcs(template.FuncMap{
		"contains": containsNotEmpty,
	})

	if parsed, err = getParsed(tmpl); err != nil {
		panic(err)
	}
}
