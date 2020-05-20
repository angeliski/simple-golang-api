package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"url": "https://api.thecatapi.com/v1/images/search?size=full"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	log.Println("Server started! listening on port: %s", port)

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS()(loggedRouter)); err != nil {
		log.Fatalf("Error on http listen: %v", err)
	}
}