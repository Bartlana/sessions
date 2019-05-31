package domains

import (
	"../utils"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Answer struct {
	Answer_id int64	`json:"answer_id" gorm:"primary_key"`
	Question int64 `json:"question_id"`
	Answer_text string `json:"answer_text"`
	Student int64 `json:"student_id"`
}

type Question struct {
	Question_id int64	`json:"question_id" gorm:"primary_key"`
	Question_text string `json:"question_text"`
}

type Presence struct {
	Presence_id int64	`json:"presence_id_id" gorm:"primary_key"`
	Student int64 		`json:"student_id"`
	Class int64			`json:"class_id"`
}

func CreateQuestion (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var question Question
	json.NewDecoder(r.Body).Decode(&question)
	db.Create(&question)
	db.First(&question)
	json.NewEncoder(w).Encode(&question)

	PlugLog()
	log.WithFields(logrus.Fields{
		"question_id":    question.Question_id,
		"question_text":  question.Question_text,
	}).Info("Created question")
}

func CreateAnswer (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var answer Answer
	json.NewDecoder(r.Body).Decode(&answer)
	db.Create(&answer)
	db.First(&answer)
	json.NewEncoder(w).Encode(&answer)
}

func CreatePresence (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var presence Presence
	json.NewDecoder(r.Body).Decode(&presence)
	db.Create(&presence)
}

