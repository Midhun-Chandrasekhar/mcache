package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Midhun-Chandrasekhar/mcache/core"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	persist := os.Getenv("PERSIST")

	if port == "" {
		port = "4567"
	}

	if host == "" {
		host = "0.0.0.0"
	}

	if persist == "true" {
		err := core.LoadCache()
		if err != nil && !os.IsNotExist(err) {
			log.Println("Error loading cache:", err)
			os.Exit(1)
		}
		go core.PersistCacheWorker()
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	router := core.NewRouter()
	log.Println("Booting Mcache server on", addr)
	err := http.ListenAndServe(addr, core.SecureRoutesMiddleware(router))
	if err != nil {
		log.Println("Error booting Mcache server:", err)
	}
}
