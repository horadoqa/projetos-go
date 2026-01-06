package main

import (

	"net/http"

)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", listUsersHandler)
	http.ListenAndServe(":3000", mux)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WORKING..."))
}