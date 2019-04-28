package domains

import (
	"net/http"
	"encoding/json"
	"../utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)


type Group struct {
	Group_id int64	`json:"group_id" gorm:"primary_key"`
	Group_name string `json:"group_name"`
	Subgroup string `json:"subgroup"`
}

func GetGroups(w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var groups []Group
	db.Find(&groups)
	json.NewEncoder(w).Encode(&groups)
}

func CreateGroup (w http.ResponseWriter, r *http.Request){
	db := utils.GetDB()
	var group Group
	json.NewDecoder(r.Body).Decode(&group)
	db.Create(&group)
	json.NewEncoder(w).Encode(&group)
	PlugLog()
	var groups []Group
	db.Find(&groups)
	json.NewEncoder(w).Encode(&groups)
	log.WithFields(logrus.Fields{
		"group_id":     group.Group_id,
		"group_name":   group.Group_name,
		"subgroup":		group.Subgroup,
	}).Info("Created group")
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	db := utils.GetDB()
	params := mux.Vars(r)
	var group Group
	db.First(&group, params["id"])
	db.Delete(&group)

	var groups []Group
	db.Find(&groups)
	json.NewEncoder(w).Encode(&groups)
	PlugLog()
	log.WithFields(logrus.Fields{
		"group_id":    group.Group_id,
	}).Info("Deleted group")
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {

	db := utils.GetDB()

	params := mux.Vars(r)
	var updGroup Group
	db.First(&updGroup, params["id"])
	json.NewDecoder(r.Body).Decode(&updGroup)
	if err := db.Save(updGroup); err.Error != nil {
		return
	}
	var groups []Group
	db.Find(&groups)
	json.NewEncoder(w).Encode(&groups)
	PlugLog()
	log.WithFields(logrus.Fields{
		"group_id":    	updGroup.Group_id,
		"group_name":  	updGroup.Group_name,
		"subgroup":		updGroup.Subgroup,
	}).Info("Updated group")
}