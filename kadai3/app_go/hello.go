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
	db, err := sql.Open("postgres", "host=postgres port=5432 user=root password=password dbname=Godb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

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
		var id int
		err = db.QueryRow("INSERT INTO users (name, email) VALUES($1,$2) RETURNING id", input.Name, input.Email).Scan(&id)
		if err != nil {
			fmt.Println(err)
		}

		var row string
		err = db.QueryRow("SELECT * FROM user where id = ?", id).Scan(&row)
		res, err := json.Marshal(row)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(res)

	}

	// if r.Method == "GET" {

	// }

	// if r.Method == "DELETE" {
	// }

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/users", usersHandler)
	http.ListenAndServe(":8080", nil)
}
