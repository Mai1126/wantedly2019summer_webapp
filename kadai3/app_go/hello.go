package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

type Index struct {
	Message string `json:"message"`
}

type Input struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Output struct {
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
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

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// var id := r.URL.RawQuery

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		input := Input{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	// if r.Method == "GET" {

	// }

	// if r.Method == "DELETE" {
	// }

}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5555 user=root password=password dbname=Godb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
