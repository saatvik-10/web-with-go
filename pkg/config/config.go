package config

import "text/template"

//holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
