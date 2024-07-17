package fs

import (
	"embed"
	"html/template"
)

//go:embed templates/*
var templatesFS embed.FS

// GetTemplates returns the embedded templates
func GetTemplates() *template.Template {
	// Create a template set from the embedded files
	tmpl := template.Must(template.New("").ParseFS(templatesFS, "templates/*"))
	return tmpl
}
