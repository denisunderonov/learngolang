package main 

import (
	"fmt"
	"net/http"
)

func timeOutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Создание контекста с тайм-аутом в 2 секунды
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()
		// Обновление запроса с новым контекстом
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}

func main() {
	http.Handle("/", timeOutMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(3 * time.Second):
			fmt.Fprintln(w, "Request processed")
		case <-r.Context().Done():
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		}
	})))

	http.ListenAndServe(":8080", nil)
}