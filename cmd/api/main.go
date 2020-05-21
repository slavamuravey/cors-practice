package main

import (
	"github.com/slavamuravey/cors-practice/pkg/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api", handler.Handler)
	http.HandleFunc("/jsonp", jsonpHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func jsonpHandler(w http.ResponseWriter, r *http.Request) {
	callback := r.URL.Query().Get("callback")
	switch callback {
	case "gotWeather":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`gotWeather({
  temperature: 25,
  humidity: 78
});
`))
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
