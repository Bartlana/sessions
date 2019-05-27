package domains

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
)


type Professors struct {
	Professor_id uint	`json:"professor_id" gorm:"primary_key"`
	Professor_name string `json:"professor_name"`
	Login string `json:"professor_login"`
	Password string `json:"professor_pass"`
	Token string `json:"token" sql:"-"`
}

var GetProfessors = func(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
}

var CreateProfessor = func (w http.ResponseWriter, r *http.Request){
	//db := utils.GetDB()
	professor := &Professors{}
	err := json.NewDecoder(r.Body).Decode(professor)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	//db.Create(&professor)
	resp := professor.Create()
	utils.Respond(w, resp)
	//json.NewEncoder(w).Encode(&professor)
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    professor.Professor_id,
		"professor_name":  professor.Professor_name,
		"professor_login": professor.Login,
	}).Info("Created professor")
}
func Authenticate (w http.ResponseWriter, r *http.Request) {

	account := &Professors{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := Login(account.Login, account.Password)
	utils.Respond(w, resp)
}


var DeleteProfessor = func(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	params := mux.Vars(r)
	var professor Professors
	db.First(&professor, params["id"])
	db.Delete(&professor)
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
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    updProfessor.Professor_id,
		"professor_name":  updProfessor.Professor_name,
	}).Info("Updated professor")
}