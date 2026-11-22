package main

import (
	"encoding/json"
	"net/http"
)

type parameters struct {
	Body string `json:"body"`
}
type respondParams struct {
	Error string `json:"error,omitempty"`
	Valid bool   `json:"valid,omitempty"`
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	failedParms := respondParams{
		Error: msg,
	}
	data, err := json.Marshal(failedParms)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func (cfg *apiConfig) handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	// Logic for managing POSTS requests on validate_chirp.
	msg := []string{"Something went wrong", "Chirp is too long", "Empty body sent"}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, msg[0])
		return
	}
	if len(params.Body) > 140 {
		respondWithErr(w, 400, msg[1])
		return
	}
	if len(params.Body) == 0 {
		respondWithErr(w, 400, msg[2])
		return
	}
	resParams := respondParams{Valid: true}
	respondWithJson(w, 200, resParams)
}
