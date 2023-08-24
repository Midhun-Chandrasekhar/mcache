package core

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/keys", GetKeysHandler)
	router.HandleFunc("/get", GetHandler)
	router.HandleFunc("/set", SetHandler)
	router.HandleFunc("/delete", DeleteHandler)
	router.HandleFunc("/", HealthHandler)

	PrintRoutes()
	return router
}
