package core

import (
	"log"
	"net/http"
	"os"
)

func SecureRoutesMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path

		w.Header().Set("Content-Type", "application/json")

		auth_enabled := os.Getenv("AUTH_ENABLED")
		if auth_enabled == "true" {
			username, password, ok := r.BasicAuth()
			if ok {
				user := os.Getenv("AUTH_USER")
				pass := os.Getenv("AUTH_PASSWORD")
				if !(username == user && password == pass) {

					log.Println(
						"Method: ", r.Method, " | ",
						"Request IP: ", r.RemoteAddr, " | ",
						"Path: ", r.URL.Path, " | ",
						"Status: ", http.StatusUnauthorized)

					(&ExtendedResponseWriter{w}).UnAutherizedResponse()
					return
				}
			}
		}

		if r.Method == "GET" {
			next.ServeHTTP(w, r)
			return
		}

		switch path {
		case "/get":
			if !(&ExtendedRequest{r}).Methods("GET") {
				(&ExtendedResponseWriter{w}).MethodNotAllowedResponse()
				return
			}
		case "/set":
			if !(&ExtendedRequest{r}).Methods("POST") {
				(&ExtendedResponseWriter{w}).MethodNotAllowedResponse()
				return
			}
		case "/delete":
			if !(&ExtendedRequest{r}).Methods("DELETE") {
				(&ExtendedResponseWriter{w}).MethodNotAllowedResponse()
				return
			}
		default:
			if !(&ExtendedRequest{r}).Methods("GET") {
				(&ExtendedResponseWriter{w}).MethodNotAllowedResponse()
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
