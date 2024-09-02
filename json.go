package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter , code int , msg string) {
	if code > 499 {
		log.Println("Error Responsed with 5XX error : " , msg)
	}

	type errResponse struct {
		Error string `json:"error"`

	}

	responseWithJSON(w , code , errResponse {
		Error : msg,
	})
}

func responseWithJSON(w http.ResponseWriter , code int , payload interface{}) {
	dat , err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal payload Response : %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}