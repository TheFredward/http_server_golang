package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) getServerCount(w http.ResponseWriter, r *http.Request) {
	// Display the count of metrics counter
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	webPage := fmt.Sprintf(`
	<html>
	<body>
	    <h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	</html>`, cfg.fileServerHits.Load())
	w.Write([]byte(webPage))
}

func (cfg *apiConfig) middlewareMetricsCounter(next http.Handler) http.Handler {
	// return handler function with updated visited times for specific page
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
