package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/TheFredward/http_server_golang/internal/database"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	db             *database.Queries
	environment    string
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}

func (cfg *apiConfig) resetServerCount(w http.ResponseWriter, r *http.Request) {
	if cfg.environment != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset only allowed in dev environment"))
		return
	}
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		w.Write([]byte(("Failed to delete users" + err.Error())))
		return
	}
	w.WriteHeader(http.StatusOK)
	cfg.fileServerHits.Store(0)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to import .env file:", err)
	}
	dbURL := os.Getenv("DB_URL")
	workEnv := os.Getenv("PLATFORM")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	dbQueries := database.New(db)
	apicfg := apiConfig{
		db:          dbQueries,
		environment: workEnv,
	}
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
	mux.HandleFunc("POST /api/chirps", apicfg.handlerChirps)
	mux.HandleFunc("GET /api/chirps", apicfg.handlerGetAllChirps)
	mux.HandleFunc("POST /api/users", apicfg.handlerCreateUser)
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
