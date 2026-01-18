package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintln(w, "Hello, my name is Denis and I started learning Go in 2026")
	})
	http.ListenAndServe(":8080", nil)
}
