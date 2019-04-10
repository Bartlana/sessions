package main

import (
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"./domains"
	"./utils"
	"log"
)




func main() {
	router := mux.NewRouter()
	utils.InitDB()

	router.HandleFunc("/subjects", domains.GetSubjects).Methods("Get")

	log.Fatal(http.ListenAndServe(":8080", router))
}
