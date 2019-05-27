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
	router.Use(utils.JwtAuthentication)

	// Subjects
	router.HandleFunc("/subjects", domains.GetSubjects).Methods("GET")
	router.HandleFunc("/subjects", domains.CreateSubject).Methods("POST")
	router.HandleFunc("/subjects/{id}", domains.DeleteSubject).Methods("DELETE")
	router.HandleFunc("/subjects/{id}", domains.UpdateSubject).Methods("PUT")

	//Groups
	router.HandleFunc("/groups", domains.GetGroups).Methods("GET")
	router.HandleFunc("/groups", domains.CreateGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", domains.DeleteGroup).Methods("DELETE")
	router.HandleFunc("/groups/{id}", domains.UpdateGroup).Methods("PUT")

	//Students
	router.HandleFunc("/students/{group_id}", domains.GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", domains.UpdateStudent).Methods("PUT")

	//Professors
	router.HandleFunc("/professors", domains.GetProfessors).Methods("GET")
	router.HandleFunc("/professor/new", domains.CreateProfessor).Methods("POST")
	router.HandleFunc("/professors/{id}", domains.DeleteProfessor).Methods("DELETE")
	router.HandleFunc("/professors/{id}", domains.UpdateProfessor).Methods("PUT")
	router.HandleFunc("/professor/login", domains.Authenticate).Methods("POST")

	//Class
	router.HandleFunc("/class/professor/{professor_id}", domains.GetClassesByProfessor).Methods("GET")
	router.HandleFunc("/class/subject/{subject_id}", domains.GetClassesBySubject).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
