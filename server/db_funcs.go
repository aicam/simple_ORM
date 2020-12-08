package server

import "github.com/jinzhu/gorm"

func CheckAdmin(db *gorm.DB, username string, password string) bool {
	recordNotFound := db.Where(&Admin{
		Username: username,
		Password: password,
	}).RecordNotFound()
	if recordNotFound {
		return false
	}
	return true
}

func SetAdminToken(db *gorm.DB, username string, token string) {
	admin := Admin{}
	db.Where(&Admin{Username:username}).First(&admin)
	admin.Token = token
	db.Save(&admin)
}
