package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Index struct {
	Message string `json:"message"`
}

type Input struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Users []User

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
		user := User{}
		err = db.QueryRow("INSERT INTO users (name, email) VALUES($1,$2) RETURNING id, name, email, created_at, updated_at", input.Name, input.Email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}

		res, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write(res)
	}

	if r.Method == "GET" {
		rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users Users
		user := User{}

		for rows.Next() {
			if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
				log.Fatal(err)
			}
			users = append(users, user)
		}
		res, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}

}

func usersidHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	db, err := sql.Open("postgres", "host=postgres port=5432 user=root password=password dbname=Godb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	if r.Method == "DELETE" {
		_, err := db.Exec("delete from users WHERE id = '" + id + "'")
		if err != nil {
			fmt.Println(err)
		}
		w.WriteHeader(204)
	}

	if r.Method == "GET" {
		user := User{}
		err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = '"+id+"' LIMIT 1").Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}

		res, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}

	if r.Method == "PUT" {
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

		_, err = db.Exec("UPDATE users SET name = '" + input.Name + "', email = '" + input.Email + "' WHERE id = " + id)
		if err != nil {
			fmt.Println(err)
		}
		user := User{}
		err = db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users WHERE id = '"+id+"' LIMIT 1").Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println(err)
		}

		res, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(res)
	}

}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexHandler)
	myRouter.HandleFunc("/users", usersHandler)
	myRouter.HandleFunc("/users/{id}", usersidHandler)
	http.ListenAndServe(":8080", myRouter)
}
