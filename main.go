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

	router.HandleFunc("/subjects", domains.GetSubjects).Methods("GET")
	router.HandleFunc("/subjects", domains.CreateSubject).Methods("POST")
	router.HandleFunc("/subjects/{id}", domains.DeleteSubject).Methods("DELETE")
	router.HandleFunc("/subjects/{id}", domains.UpdateSubject).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}
