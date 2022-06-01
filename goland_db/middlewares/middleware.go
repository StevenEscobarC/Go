package middlewares

import "net/http"

//recibe un handler(valida) y devuelve otro (si todo esta bien)
func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != "Autorizado" {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		next.ServeHTTP(w, r)
	}
}
