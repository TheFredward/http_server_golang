package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type parameters struct {
	Body string `json:"body"`
}
type respondParams struct {
	Cleaned_body string `json:"cleaned_body,omitempty"`
	Error        string `json:"error,omitempty"`
	Valid        bool   `json:"valid,omitempty"`
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	// Returns the respondparams struct with specific error message
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

func checkMsg(msg string) (string, error) {
	// Filter user message.
	// Usage of map[struct]string for 0(n) sspeed
	var finalMsg string
	profaneWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	sliceMsg := strings.Split(msg, " ")
	for i, word := range sliceMsg {
		if _, ok := profaneWords[strings.ToLower(word)]; ok {
			sliceMsg[i] = "****"
		}
	}
	finalMsg = strings.Join(sliceMsg, " ")
	return finalMsg, nil
}
