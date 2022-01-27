package main

import (
	"book_api/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	apiPrefixV1 string = "/api/v1"
)

var (
	port               string
	bookResourcePrefix string = apiPrefixV1 + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Can't load .env file: ", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	router := mux.NewRouter()
	log.Println("Start on port " + port)
	utils.BuildBookResource(router, bookResourcePrefix)
	log.Println("Pre listen")
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
