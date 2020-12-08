package server

import "github.com/jinzhu/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token string `json:"token"`
}

type Customer struct {
	gorm.Model
	Name string `json:"name"`
	Mobile string `json:"mobile"`
	WhatsappNumber string `json:"whatsapp_number"`
	Address string `json:"address"`
	Sites []string `json:"sites"`
}

