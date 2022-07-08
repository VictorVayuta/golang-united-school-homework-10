package main

import (
	"fmt"
	"github.com/GolangUnited/helloweb/cmd/handlers"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/stretchr/testify/assert"
)

func AddHandlers(router *mux.Router) {
	router.HandleFunc("/name/{PARAM}", handlers.NameParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handlers.BadParam).Methods(http.MethodGet)
	router.HandleFunc("/data", handlers.BodyParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", handlers.HeadersParam).Methods(http.MethodPost)
}

func Start(host string, port int) {
	router := mux.NewRouter()

	AddHandlers(router)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 8081
	}

	Start(host, port)
}
