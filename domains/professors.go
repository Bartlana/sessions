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


type Professors struct {
	Professor_id uint	`json:"professor_id" gorm:"primary_key"`
	Professor_name string `json:"professor_name"`
	Login string `json:"professor_login"`
	Password string `json:"professor_pass"`
	Token string `json:"token" sql:"-"`
}
func (professor *Professors) ProfessorValidate() (map[string] interface{}, bool) {
	db := u.GetDB()
	if !strings.Contains(professor.Login, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(professor.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Professors{}

	//check for errors and duplicate emails
	err := db.Table("professors").Where("login = ?", professor.Login).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Login != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (professor *Professors) ProfessorCreate() (map[string] interface{}) {
	db := u.GetDB()
	if resp, ok := professor.ProfessorValidate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(professor.Password), bcrypt.DefaultCost)
	professor.Password = string(hashedPassword)

	db.Create(professor)

	if professor.Professor_id <= 0 {
		return u.Message(false, "Failed to create account, connection error.")
	}

	//Create new JWT token for the newly registered account
	tk := &u.Token{UserId: professor.Professor_id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	professor.Token = tokenString

	professor.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = professor
	return response
}

func ProfessorLogin(login, password string) (map[string]interface{}) {
	db := u.GetDB()
	professor := &Professors{}
	err := db.Table("professors").Where("login = ?", login).First(professor).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(professor.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Your password is invalid. Please try again")
	}
	//Worked! Logged In
	professor.Password = ""

	//Create JWT token
	tk := &u.Token{UserId: professor.Professor_id}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	professor.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["professor"] = professor
	return resp
}

var GetProfessors = func(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var professors []Professors
	db.Find(&professors)
	json.NewEncoder(w).Encode(&professors)
}

func CreateProfessor(w http.ResponseWriter, r *http.Request){
	professor := &Professors{}
	err := json.NewDecoder(r.Body).Decode(professor)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	resp := professor.ProfessorCreate()
	utils.Respond(w, resp)
	PlugLog()
	log.WithFields(logrus.Fields{
		"professor_id":    professor.Professor_id,
		"professor_name":  professor.Professor_name,
		"professor_login": professor.Login,
	}).Info("Created professor")
}
func ProfessorAuthenticate (w http.ResponseWriter, r *http.Request) {

	account := &Professors{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	resp := ProfessorLogin(account.Login, account.Password)
	utils.Respond(w, resp)
}


func DeleteProfessor(w http.ResponseWriter, r *http.Request) {
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