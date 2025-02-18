package config

import "text/template"

//holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}