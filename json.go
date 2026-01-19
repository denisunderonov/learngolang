package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	resp := Response{
		Message: "User " + user.Name + " created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/user", handler)
	http.ListenAndServe(":8080", nil)
}
