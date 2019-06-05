package config

import "html/template"

func TplManager() *template.Template {
	return template.Must(template.New("").ParseGlob("assets/templates/*"))
}
