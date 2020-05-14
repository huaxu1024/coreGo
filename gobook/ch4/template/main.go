package template

import (
	"html/template"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}---------------------
Number: {{ .Number}}
User: 	{{ .User.Login}}
Title: 	{{ .Title | printf "%.64s"}}
Age:	{{ .CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
}