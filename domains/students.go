package domains

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"github.com/gorilla/mux"
	"../utils"
	"strings"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
	u "../utils"
)


type Student struct {
	Student_id uint	`json:"student_id" gorm:"primary_key"`
	Student_name string `json:"student_name"`
	Login string `json:"student_login"`
	Group_id int64 `json:"group" gorm:"foreignkey:Group_id"`
	Password string `json:"student_pass"`
	Token string `json:"token" sql:"-"`
}
func (student *Student) StudentValidate() (map[string] interface{}, bool) {
	db := u.GetDB()
	if !strings.Contains(student.Login, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(student.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Student{}

	//check for errors and duplicate emails
	err := db.Table("students").Where("login = ?", student.Login).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Login != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (student *Student) StudentCreate() (map[string] interface{}) {
	db := u.GetDB()
	if resp, ok := student.StudentValidate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	student.Password = string(hashedPassword)

	db.Create(student)

	if student.Student_id <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &u.Token{UserId: student.Student_id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	student.Token = tokenString

	student.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = student
	return response
}

func StudentLogin(login, password string) (map[string]interface{}) {
	db := u.GetDB()
	student := &Student{}
	err := db.Table("students").Where("login = ?", login).First(student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Your password is invalid. Please try again")
	}
	//Worked! Logged In
	student.Password = ""

	//Create JWT token
	tk := &u.Token{UserId: student.Student_id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	student.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["student"] = student
	return resp
}
func StudentAuthenticate (w http.ResponseWriter, r *http.Request) {

	account := &Student{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := StudentLogin(account.Login, account.Password)
	utils.Respond(w, resp)
}
func GetStudents(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var students []Student
	params := mux.Vars(r)
	db.Joins("join groups on students.group_id = groups.group_id").Where("groups.group_id = ?", params["group_id"]).Find(&students)
	json.NewEncoder(w).Encode(&students)
}

func CreateStudent (w http.ResponseWriter, r *http.Request){
	student := &Student{}
	err := json.NewDecoder(r.Body).Decode(student)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := student.StudentCreate()
	u.Respond(w, resp)
	//db.Create(&student)
	//json.NewEncoder(w).Encode(&student)
	//var students []Student
	//db.Find(&students)
	//json.NewEncoder(w).Encode(&students)
	PlugLog()
	log.WithFields(logrus.Fields{
		"student_id":    student.Student_id,
		"student_name":  student.Student_name,
		"student_login": student.Login,
	}).Info("Created student")
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
		"student_id":    student.Student_id,
	}).Info("Deleted student")
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {

	db := utils.GetDB()
	params := mux.Vars(r)
	var updStudent Student
	db.First(&updStudent, params["id"])
	json.NewDecoder(r.Body).Decode(&updStudent)
	if err := db.Save(updStudent); err.Error != nil {
		return
	}
	//var students []Student
	//db.Find(&students)
	//json.NewEncoder(w).Encode(&students)
	PlugLog()
	log.WithFields(logrus.Fields{
		"student_id":    updStudent.Student_id,
		"student_name":  updStudent.Student_name,
	}).Info("Updated student")
}
