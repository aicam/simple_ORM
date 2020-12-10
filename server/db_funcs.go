package server

import (
	"github.com/jinzhu/gorm"
)

func CheckAdmin(db *gorm.DB, username string, password string) bool {
	var admin AdminTable
	recordNotFound := db.Where(&AdminTable{
		Username: username,
		Password: password,
	}).First(&admin).RecordNotFound()
	if recordNotFound {
		return false
	}
	return true
}

func SetAdminToken(db *gorm.DB, username string, token string) {
	admin := AdminTable{}
	db.Where(&AdminTable{Username: username}).First(&admin)
	admin.Token = token
	db.Save(&admin)
}
