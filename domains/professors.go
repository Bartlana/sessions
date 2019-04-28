package domains

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
)


type Professors struct {
	Professor_id int64	`json:"professor_id" gorm:"primary_key"`
	Professor_name string `json:"professor_name"`
	Login string `json:"professor_login"`
	Password string `json:"professor_pass"`
}

func GetProfessors(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
}

func CreateProfessor (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var professor Professors
	json.NewDecoder(r.Body).Decode(&professor)
	db.Create(&professor)
	json.NewEncoder(w).Encode(&professor)
	PlugLog()
	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
	log.WithFields(logrus.Fields{
		"professor_id":    professor.Professor_id,
		"professor_name":  professor.Professor_name,
		"professor_login": professor.Login,
	}).Info("Created professor")
}

func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	params := mux.Vars(r)
	var professor Professors
	db.First(&professor, params["id"])
	db.Delete(&professor)

	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    professor.Professor_id,
	}).Info("Deleted professor")
}

func UpdateProfessor(w http.ResponseWriter, r *http.Request) {

	db := utils.GetDB()

	params := mux.Vars(r)
	var updProfessor Professors
	db.First(&updProfessor, params["id"])
	json.NewDecoder(r.Body).Decode(&updProfessor)
	if err := db.Save(updProfessor); err.Error != nil {
		return
	}
	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    updProfessor.Professor_id,
		"professor_name":  updProfessor.Professor_name,
	}).Info("Updated professor")
}