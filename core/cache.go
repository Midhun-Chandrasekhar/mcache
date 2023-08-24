package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var cache = make(map[string]string)
var lock sync.RWMutex
var writeCacheCh = make(chan map[string]string, 100)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(
			"Method: ", r.Method, " | ",
			"Request IP: ", r.RemoteAddr, " | ",
			"Path: ", r.URL.Path, " | ",
			"Status: ", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Key is required"})
		return
	}

	lock.RLock()
	value, found := cache[key]
	lock.RUnlock()

	if !found {
		w.WriteHeader(http.StatusNotFound)
		log.Println(
			"Method: ", r.Method, " | ",
			"Request IP: ", r.RemoteAddr, " | ",
			"Path: ", r.URL.Path, " | ",
			"Status: ", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Key not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	log.Println(
		"Method: ", r.Method, " | ",
		"Request IP: ", r.RemoteAddr, " | ",
		"Path: ", r.URL.Path, " | ",
		"Status: ", http.StatusOK, " | ",
		"Key: ", key)

	json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
}

func SetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("key")
	value := r.FormValue("value")

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(
			"Method: ", r.Method, " | ",
			"Request IP: ", r.RemoteAddr, " | ",
			"Path: ", r.URL.Path, " | ",
			"Status: ", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Key is required"})
		return
	}

	log.Printf("ADD key %s", key)
	lock.Lock()
	if cache[key] != "" {
		log.Printf("Mcache: Key %s already exists, UPDATING", key)
	}
	cache[key] = value
	lock.Unlock()

	// Notify the persister that there are changes
	writeCacheCh <- cache

	w.WriteHeader(http.StatusOK)

	log.Println(
		"Method: ", r.Method, " | ",
		"Request IP: ", r.RemoteAddr, " | ",
		"Path: ", r.URL.Path, " | ",
		"Status: ", http.StatusOK, " | ",
		"Key: ", key)

	json.NewEncoder(w).Encode(map[string]string{"message": "Value stored"})
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	log.Printf("Mcache: DELETE key %s", key)

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(
			"Method: ", r.Method, " | ",
			"Request IP: ", r.RemoteAddr, " | ",
			"Path: ", r.URL.Path, " | ",
			"Status: ", http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Key is required"})
		return
	}

	lock.Lock()
	delete(cache, key)
	lock.Unlock()

	// Notify the persister that there are changes
	writeCacheCh <- cache

	w.WriteHeader(http.StatusOK)

	log.Println(
		"Method: ", r.Method, " | ",
		"Request IP: ", r.RemoteAddr, " | ",
		"Path: ", r.URL.Path, " | ",
		"Status: ", http.StatusOK, " | ",
		"Key: ", key)

	json.NewEncoder(w).Encode(map[string]string{"message": "Key deleted"})
}

func GetKeysHandler(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(cache))
	for k := range cache {
		keys = append(keys, k)
	}

	// Get the total size of the cache that store in memory (in bytes)
	cacheSize := 0
	for _, v := range cache {
		cacheSize += len(v)
	}

	log.Println(
		"Method: ", r.Method, " | ",
		"Request IP: ", r.RemoteAddr, " | ",
		"Path: ", r.URL.Path, " | ",
		"Status: ", http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{"keys": keys, "size": cacheSize})
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"version": "Mcache v1.0.0",
		"message": "Welcome to Mcache server",
		"status":  "Active"})
}

func PersistCacheWorker() {
	for {
		select {
		case c := <-writeCacheCh:
			persistCache(c)
		}
	}
}

func persistCache(c map[string]string) {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Println("Mcache: Error marshaling cache:", err)
		return
	}
	err = ioutil.WriteFile("cache.json", bytes, 0644)
	if err != nil {
		log.Println("Mcache: Error writing cache to disk:", err)
	}

	log.Println("Mcache: Cache persisted")
}

func LoadCache() error {
	bytes, err := ioutil.ReadFile("cache.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &cache)
}

type ExtendedRequest struct {
	*http.Request
}

type ExtendedResponseWriter struct {
	http.ResponseWriter
}

func (r *ExtendedRequest) Methods(methods ...string) bool {
	for _, method := range methods {
		if r.Method == method {
			return true
		}
	}
	return false
}

func (w *ExtendedResponseWriter) MethodNotAllowedResponse() {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
}

func (w *ExtendedResponseWriter) UnAutherizedResponse() {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"error": "UnAutherized request"})
}
