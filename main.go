package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("Fatal error occurred in server setup: ", err)
	}
}

//go:embed index.html
var indexTemplate string

func run() error {
	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	addr := "0.0.0.0:9090"
	srv := &http.Server{
		Addr: addr,
		Handler: reqHandler(
			tmpl,
			os.Getenv("POD_NAME"),
			os.Getenv("GITHUB_ACTIONS_RUN_URL"),
		),
	}

	log.Println("Listening on: ", addr)
	return srv.ListenAndServe()
}

const (
	colorBlue  = "SkyBlue"
	colorGreen = "LightGreen"
	colorPink  = "LightPink"

	// Change me ! Commit, push, and see the magic happen!
	configuredColor = colorPink
)

func reqHandler(
	tmpl *template.Template,
	podName string,
	runURL string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, map[string]any{
			"color":               configuredColor,
			"podName":             podName,
			"githubActionsRunURL": runURL,
		})
		if err != nil {
			log.Println("Failed to write response: ", err)
		}
	}
}
