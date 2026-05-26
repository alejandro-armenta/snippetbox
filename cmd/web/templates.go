package main

import (
	//"html/template"
	"path/filepath"

	"snippetbox.alexarmenta.net/internal/models"
)

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func ale() {
	//cache := map[string]*template.Template{}

	filepath.Glob("../../ui/html/pages/*.tmpl")
}
