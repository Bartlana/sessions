package domains

import (
	"encoding/json"
	"net/http"
	"../utils"
	"time"
	"github.com/gorilla/mux"
)


type Class struct {
	Class_id int64	`json:"class_id" gorm:"primary_key"`
	Class_date time.Time `json:"class_date"`
	Subject_name string `json:"subject"`
	Professor_name string `json:"professor"`
	Theme string `json:"theme"`
	Groups string `json:"groups"`
}

func GetClassesByProfessor(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var class []Class
	params := mux.Vars(r)

	//var results []Result
	//db.Table("classes").Select(
	//	"classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups").Joins(
	//		"join subjects s2 on classes.subject = s2.subject_id").Joins(
	//			"join professors p on classes.professor = p.professor_id").Joins(
	//				"join group_in_class on classes.class_id = group_in_class.class_id").Joins(
	//					"join groups g on group_in_class.group_id = g.group_id").Group("classes.class_id, subject_name, professor_name").Find(&results)
	db.Raw(" select classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups from classes " +
	"join subjects s2 on classes.subject = s2.subject_id " +
	"join professors p on classes.professor = p.professor_id " +
	"join group_in_class on classes.class_id = group_in_class.class_id " +
	"join groups g on group_in_class.group_id = g.group_id " +
	"where professor_id = ?" +
	"group by classes.class_id, subject_name, professor_name", params["professor_id"]).Scan(&class)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	json.NewEncoder(w).Encode(&class)
}
func GetClassesBySubject(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var class []Class
	params := mux.Vars(r)
	db.Raw(" select classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups from classes " +
		"join subjects s2 on classes.subject = s2.subject_id " +
		"join professors p on classes.professor = p.professor_id " +
		"join group_in_class on classes.class_id = group_in_class.class_id " +
		"join groups g on group_in_class.group_id = g.group_id " +
		"where subject_id = ?" +
		"group by classes.class_id, subject_name, professor_name", params["subject_id"]).Scan(&class)
	json.NewEncoder(w).Encode(&class)
}
