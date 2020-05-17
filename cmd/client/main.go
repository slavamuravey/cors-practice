package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(filepath.Join(getTemplatesDir(), "index.html"))
		tmpl.Execute(w, nil)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
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
