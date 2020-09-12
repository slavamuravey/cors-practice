package main

import (
	"github.com/slavamuravey/cors"
	"github.com/slavamuravey/cors-practice/pkg/handler"
	"log"
	"net/http"
)

func main() {
	config := &cors.Config{
		AllowAllOrigin: true,
	}

	corsMiddleware := cors.CreateMiddleware(config)

	http.HandleFunc("/api-no-cors", handler.Handler)
	http.HandleFunc("/api-with-cors", corsMiddleware(http.HandlerFunc(handler.Handler)))

	config1 := &cors.Config{
		AllowCredentials:           false,
		AllowOriginPattern:         "http://127.0.0.1:8181",
	}

	corsMiddleware1 := cors.CreateMiddleware(config1)
	http.HandleFunc("/api-with-cors-with-credentials-error", corsMiddleware1(http.HandlerFunc(handler.Handler)))

	config2 := &cors.Config{
		AllowCredentials:           true,
		AllowOriginPattern:         "http://127.0.0.1:8181",
	}

	corsMiddleware2 := cors.CreateMiddleware(config2)
	http.HandleFunc("/api-with-cors-with-credentials-success", corsMiddleware2(http.HandlerFunc(handler.Handler)))

	config3 := &cors.Config{
		AllowCredentials:           true,
		AllowAllOrigin: true,
	}

	corsMiddleware3 := cors.CreateMiddleware(config3)
	http.HandleFunc("/api-with-cors-with-credentials-wildcard-error", corsMiddleware3(http.HandlerFunc(handler.Handler)))

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
