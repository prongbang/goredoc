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

// RedocScript standalone javascript
//
//go:embed assets/redoc.standalone.js
var RedocScript string

type Redoc interface {
	Handler() http.HandlerFunc
}

type redoc struct {
	Cfg Config
}

func (rd *redoc) Handler() http.HandlerFunc {
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
			Script:  template.JS(RedocScript),
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func New(c Config) Redoc {
	return &redoc{
		Cfg: c,
	}
}
