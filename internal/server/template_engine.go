package server

import (
	"fmt"
	"html/template"
	"io"
	"sync"
)

type TemplateEngine struct {
	mu        sync.RWMutex
	templates *template.Template
}

func NewTemplateEngine() *TemplateEngine {
	return &TemplateEngine{
		templates: template.New("root"),
	}
}

func (e *TemplateEngine) Load(pattern string) error {
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		return err
	}

	e.mu.Lock()
	e.templates = tmpl
	e.mu.Unlock()

	return nil
}

func (e *TemplateEngine) Render(w io.Writer, name string, data any) error {
	e.mu.RLock()
	tmpl := e.templates
	e.mu.RUnlock()

	if tmpl.Lookup(name) == nil {
		return fmt.Errorf("template %q not found", name)
	}

	return tmpl.ExecuteTemplate(w, name, data)
}
