package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Cant't marshal JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(dat))
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	type Error struct {
		Error string `json:"error"`
	}
	if code > 499 {
		fmt.Println("Server get 5XX code", code)
	}
	Err := Error{
		Error: msg,
	}
	RespondWithJSON(w, code, Err)
}
