package domains

import (
	"net/http"
	"encoding/json"
	"../utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"os"
)
var log = logrus.New()

func PlugLog() {
	log.Out = os.Stdout
	file, err := os.OpenFile("session.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file")
	}
}


type Subject struct {
	Subject_id int64	`json:"subject_id" gorm:"primary_key"`
	Subject_name string `json:"subject_name"`
}

func GetSubjects(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var subjects []Subject
	db.Find(&subjects)
	json.NewEncoder(w).Encode(&subjects)
}

func CreateSubject (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var subject Subject
	json.NewDecoder(r.Body).Decode(&subject)
	db.Create(&subject)
	json.NewEncoder(w).Encode(&subject)
	PlugLog()
	log.WithFields(logrus.Fields{
		"subject_id":    subject.Subject_id,
		"subject_name":  subject.Subject_name,
	}).Info("Created subject")
}

func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	params := mux.Vars(r)
	var subject Subject
	db.First(&subject, params["id"])
	db.Delete(&subject)
	PlugLog()
	log.WithFields(logrus.Fields{
		"subject_id":    subject.Subject_id,
	}).Info("Deleted subject")
}

func UpdateSubject(w http.ResponseWriter, r *http.Request) {

	db := utils.GetDB()

	params := mux.Vars(r)
	var updSubject Subject
	db.First(&updSubject, params["id"])
	json.NewDecoder(r.Body).Decode(&updSubject)
	if err := db.Save(updSubject); err.Error != nil {
		return
	}
	PlugLog()
	log.WithFields(logrus.Fields{
		"subject_id":    updSubject.Subject_id,
		"subject_name":  updSubject.Subject_name,
	}).Info("Updated subject")
}