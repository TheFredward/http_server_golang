package main

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func (cfg *apiConfig) resetServerCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	cfg.fileServerHits.Store(0)
}

func main() {
	apicfg := apiConfig{}
	mux := http.NewServeMux()
	mux.Handle("/app/", apicfg.middlewareMetricsCounter(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/healthz" {
			http.NotFound(w, req)
			return
		}
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 OK"))
	})
	mux.HandleFunc("POST /api/validate_chirp", apicfg.handlerValidateChirp)
	mux.HandleFunc("GET /admin/metrics", apicfg.getServerCount)
	mux.HandleFunc("POST /admin/reset", apicfg.resetServerCount)
	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
