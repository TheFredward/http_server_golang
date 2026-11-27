package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TheFredward/http_server_golang/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {

	type params struct {
		Body    string    `json:"body"`
		User_id uuid.UUID `json:"user_id"`
	}
	// Logic for managing POSTS requests on validate_chirp.
	msg := []string{"Something went wrong", "Chirp is too long", "Empty body sent"}
	var cleanedMsg string
	var err error

	de := json.NewDecoder(r.Body)
	reqData := params{}
	err = de.Decode(&reqData)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Failed to decode data" + err.Error()))
		return
	}
	if len(reqData.Body) > 140 {
		respondWithErr(w, 400, msg[1])
		return
	}
	if len(reqData.Body) == 0 {
		respondWithErr(w, 400, msg[2])
		return
	}
	if len(reqData.User_id) == 0 {
		respondWithErr(w, 400, "user_id required")
		return
	}
	cleanedMsg, err = checkMsg(reqData.Body)
	if err != nil {
		respondWithErr(w, 400, err.Error())
		return
	}
	cleanedData := database.CreateChirpParams{
		Body:   cleanedMsg,
		UserID: reqData.User_id,
	}
	createdChirp, err := cfg.db.CreateChirp(r.Context(), cleanedData)
	if err != nil {
		respondWithErr(w, 400, "Unable to save to database"+err.Error())
		return
	}
	localChirp := Chirp{
		ID:        createdChirp.ID,
		CreatedAt: createdChirp.CreatedAt,
		UpdatedAt: createdChirp.UpdatedAt,
		Body:      createdChirp.Body,
		UserID:    createdChirp.UserID,
	}
	respondWithJson(w, 201, localChirp)
}

func (cfg *apiConfig) handlerGetAllChirps(w http.ResponseWriter, r *http.Request) {
	var listChirps []Chirp
	chirpsFound, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithErr(w, 400, "Unable to fetch all chirps!")
	}
	for _, chirp := range chirpsFound {
		localChirp := Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
		listChirps = append(listChirps, localChirp)
	}
	respondWithJson(w, 200, listChirps)
}
