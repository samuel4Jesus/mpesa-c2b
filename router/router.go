package router

import (
	"log"
	"mpesa/api"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func InitServer() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Starting M-Pesa C2B API Server")

	//create a new router
	router := mux.NewRouter()
	log.Println("Creating routes Completed")

	//Endpoints
	router.HandleFunc("/mpesa/collection", api.MPesaC2BStagging).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))

	//start and listen to requests
	go http.ListenAndServe(":8443", router)

	select {}
}
