package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"slavamuravey/cors/pkg/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		tmpl, _ := template.ParseFiles(filepath.Join(getTemplatesDir(), "index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/api", handler.Handler)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":8181", nil))
}

func getCurrentFileDir() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	return path.Dir(filename)
}

func getTemplatesDir() string {
	return filepath.Join(getCurrentFileDir(), "../../templates")
}
