package server

import "github.com/jinzhu/gorm"

type AdminTable struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type CustomersTable struct {
	gorm.Model
	Name           string `json:"name"`
	Mobile         string `json:"mobile"`
	WhatsappNumber string `json:"whatsapp_number"`
	Address        string `json:"address"`
	Sites          string `json:"sites"`
}

type PaymentTable struct {
	gorm.Model
	CustomerName string `json:"customer_name"`
	Amount       int    `json:"amount"`
	Method       string `json:"method"`
	Comment      string `json:"comment"`
}
