package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Post struct {
	Description string `json:"description"`
}

func GetJSON(w http.ResponseWriter, r *http.Request) {
	var post Post
	w.Header().Set("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"error": true}`)
	}

	json.Unmarshal(reqBody, &post)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(post)
}
