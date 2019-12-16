package main

import (
	"emeli/snippetbox/pkg/forms"
	"html/template"
	//"html/template"
	"path/filepath"

	"emeli/snippetbox/pkg/models"
)


type templateData struct {
	AuthenticatedUser *models.User
	CSRFToken string
	CurrentYear int
	Flash string
	Form *forms.Form
	Snippet *models.Snippet
	Snippets []*models.Snippet
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
