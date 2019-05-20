package domains

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
)


type Student struct {
	Student_id int64	`json:"student_id" gorm:"primary_key"`
	Student_name string `json:"student_name"`
	Login string `json:"student_login"`
	Group_id int64 `json:"group" gorm:"foreignkey:Group_id"`
	Password string `json:"student_pass"`
}

func GetStudents(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var students []Student
	params := mux.Vars(r)
	db.Joins("join groups on students.group_id = groups.group_id").Where("groups.group_id = ?", params["group_id"]).Find(&students)
	json.NewEncoder(w).Encode(&students)
}

func CreateStudent (w http.ResponseWriter, r *http.Request){
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

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	params := mux.Vars(r)
	var student Student
	db.First(&student, params["id"])
	db.Delete(&student)

	var students []Student
	db.Find(&students)
	json.NewEncoder(w).Encode(&students)
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    student.Student_id,
	}).Info("Deleted student")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

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