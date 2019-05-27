package domains

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	u "../utils"
)

//type Professors struct {
//	Professor_id uint	`json:"professor_id" gorm:"primary_key"`
//	Professor_name string `json:"professor_name"`
//	Login string `json:"professor_login"`
//	Password string `json:"professor_pass"`
//	Token string `json:"token" sql:"-"`
//}
 //var Professors = domains.Professors{}


func (professor *Professors) Validate() (map[string] interface{}, bool) {
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

func (professor *Professors) Create() (map[string] interface{}) {
	db := u.GetDB()
	if resp, ok := professor.Validate(); !ok {
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

func Login(login, password string) (map[string]interface{}) {
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
		return u.Message(false, "Invalid login credentials. Please try again")
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

//func GetProfessor(u uint) *Professors {
//	db := u.GetDB()
//	acc := &Professors{}
//	db.Table("professors").Where("professor_id = ?", u).First(acc)
//	if acc.Login == "" { //User not found!
//		return nil
//	}
//
//	acc.Password = ""
//	return acc
//}