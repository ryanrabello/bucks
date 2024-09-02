package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.HandleFunc("/code/{code}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		code := vars["code"]

		// TODO: Implement code validation and claiming logic
		// This is a placeholder implementation
		if code == "" {
			http.Error(w, "Invalid code", http.StatusBadRequest)
			return
		}

		// Simulate marking the code as claimed
		// In a real implementation, you would update a database or storage
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Code " + code + " has been claimed"))
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
