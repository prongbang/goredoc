package goredoc

import (
	_ "embed"
	"html/template"
	"net/http"
)

type Config struct {
	Title   string
	SpecURL string
}

// ReDocScript standalone javascript
//
//go:embed assets/redoc.standalone.js
var ReDocScript string

type ReDoc interface {
	Handler() http.HandlerFunc
}

type reDoc struct {
	Cfg Config
}

func (rd *reDoc) Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("assets/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			Title   string
			SpecURL string
			Script  template.JS
		}{
			Title:   rd.Cfg.Title,
			SpecURL: rd.Cfg.SpecURL,
			Script:  template.JS(ReDocScript),
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func New(c Config) ReDoc {
	return &reDoc{
		Cfg: c,
	}
}
