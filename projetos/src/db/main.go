package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Email 	string 	`json:"email"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)
	http.ListenAndServe(":3000", mux)
}

// Conexão com o banco de dados, abrindo tabela, listando json
func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fechando a conexão
	defer db.Close() 

	rows, err := db.Query("select * from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fechando a conexão
	defer rows.Close() 

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}
	
	// Convertendo para json
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Criando usuário
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var u User
	// Convertendo para Struct
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := db.Exec("INSERT INTO users (id, name, email) VALUES(?, ?, ?)", u.ID, u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}