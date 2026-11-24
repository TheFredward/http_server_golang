package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type usrParams struct {
	Email string `json:"email"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	usrParms := usrParams{}
	err := decode.Decode(&usrParms)
	if err != nil {
		log.Fatal("Unable to decode req body from user: ", err)
	}
	if len(usrParms.Email) == 0 {
		respondWithErr(w, 400, "Email required!")
		return
	}
	dbUsr, err := cfg.db.CreateUser(r.Context(), usrParms.Email)
	if err != nil {
		log.Fatal("Unable to create user", err)
		return
	}
	createdUsr := User{
		ID:        dbUsr.ID,
		CreatedAt: dbUsr.CreatedAt,
		UpdatedAt: dbUsr.UpdatedAt,
		Email:     dbUsr.Email,
	}
	respondWithJson(w, 201, createdUsr)
}
