package main

import (
	"encoding/json"
	"net/http"
)

type Index struct {
	Message string `json:"message"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index := Index{"Hello World!!"}

	res, err := json.Marshal(index)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
