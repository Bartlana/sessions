package domains

import (
	"net/http"
	"encoding/json"
	"../utils"
)


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