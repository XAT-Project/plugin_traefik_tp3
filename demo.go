// Package plugintraefiktp3 a demo plugin.
package plugintraefiktp3

import (
	"context"
	"fmt"
	"net/http"
	"text/template"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Demo a Demo plugin.
type Demo struct {
	next     http.Handler
	headers  map[string]string
	name     string
	template *template.Template
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &Demo{
		headers:  config.Headers,
		next:     next,
		name:     name,
		template: template.New("demo").Delims("[[", "]]"),
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("authtoken")

	// Vérification qu'il n'y ai pas l'erreur ErrNoCoockie
	// Si il y en a un alors ca veut dire que le coockie n'existe pas
	if err != nil {
		// Le coockie existe donc retour 200
		rw.WriteHeader(http.StatusFound)
	} else {
		// Coockie n'existe pas donc retour 302
		rw.WriteHeader(http.StatusFound)
		// Redirection
		req.URL.Path = "http://monApp.localhost/monApp-api/Login"
	}

	// for key, value := range a.headers {
	// 	tmpl, err := a.template.Parse(value)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	writer := &bytes.Buffer{}

	// 	err = tmpl.Execute(writer, req)
	// 	if err != nil {
	// 		http.Error(rw, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	req.Header.Set(key, writer.String())
	// }

	a.next.ServeHTTP(rw, req)
}
