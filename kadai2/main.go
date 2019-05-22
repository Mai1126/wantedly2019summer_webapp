package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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
	port := os.Getenv("PORT")
	fmt.Printf("Starting server at Port %s", port)
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
