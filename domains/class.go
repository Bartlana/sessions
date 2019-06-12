package domains

import (
	"encoding/json"
	"net/http"
	u "../utils"
	"github.com/gorilla/mux"
	"time"
)


type GetClass struct {
	Class_id int64	`json:"class_id" gorm:"primary_key"`
	Class_date string `json:"class_date"`
	Subject_name string `json:"subject"`
	Professor_name string `json:"professor"`
	Theme string `json:"theme"`
	Groups string `json:"groups" sql:"-"`
}
type Class struct {
	Class_id int64	`json:"class_id" gorm:"primary_key"`
	Class_date time.Time`json:"class_date"`
	Subject int64 `json:"subject"`
	Professor int64 `json:"professor"`
	Theme string `json:"theme"`
	}
type GroupInClass struct {
	Class_id int64 `json:"class_id"`
	Group_id int64 `json:"group_id"`
}

func GetClassesByProfessor(w http.ResponseWriter, r *http.Request){
	db := u.GetDB()
	var class []GetClass
	params := mux.Vars(r)

	db.Raw(" select classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups from classes " +
	"join subjects s2 on classes.subject = s2.subject_id " +
	"join professors p on classes.professor = p.professor_id " +
	"join group_in_classes on classes.class_id = group_in_classes.class_id " +
	"join groups g on group_in_classes.group_id = g.group_id " +
	"where professor_id = ?" +
	"group by classes.class_id, subject_name, professor_name", params["professor_id"]).Scan(&class)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	json.NewEncoder(w).Encode(&class)
}
func GetClassesBySubject(w http.ResponseWriter, r *http.Request){
	db := u.GetDB()
	var class []GetClass
	params := mux.Vars(r)
	db.Raw(" select classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups from classes " +
		"join subjects s2 on classes.subject = s2.subject_id " +
		"join professors p on classes.professor = p.professor_id " +
		"join group_in_classes on classes.class_id = group_in_classes.class_id " +
		"join groups g on group_in_classes.group_id = g.group_id " +
		"where subject_id = ?" +
		"group by classes.class_id, subject_name, professor_name", params["subject_id"]).Scan(&class)
	json.NewEncoder(w).Encode(&class)
}

func GetClassesByStudent(w http.ResponseWriter, r *http.Request){
	db := u.GetDB()
	var class []GetClass
	params := mux.Vars(r)
	db.Raw(" select classes.class_id, class_date, theme, subject_name, professor_name, string_agg(group_name, ', ') as groups from classes" +
		"	join subjects s2 on classes.subject = s2.subject_id" +
		"	join professors p on classes.professor = p.professor_id" +
		"	join group_in_classes on classes.class_id = group_in_classes.class_id" +
		"	join groups g on group_in_classes.group_id = g.group_id" +
		"	join students s3 on g.group_id = s3.group_id" +
		"	where student_id = ?" +
		"	group by classes.class_id, subject_name, professor_name", params["student_id"]).Scan(&class)
	json.NewEncoder(w).Encode(&class)
}

func CreateClass (w http.ResponseWriter, r *http.Request) {
	db := u.GetDB()
	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	db.Create(&class)
	db.First(&class)
	json.NewEncoder(w).Encode(&class)
}
func CreateGroupInClass (w http.ResponseWriter, r *http.Request) {
	db := u.GetDB()
	var groupInClass GroupInClass
	err := json.NewDecoder(r.Body).Decode(&groupInClass)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	db.Create(&groupInClass)
}