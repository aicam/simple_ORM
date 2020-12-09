package server

import (
	"github.com/jinzhu/gorm"
	"time"
)

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
	CustomerName string    `json:"customer_name"`
	Amount       int       `json:"amount"`
	Method       string    `json:"method"`
	Comment      string    `json:"comment"`
	TimeAdded    time.Time `json:"time_added"`
}

type ProductsTable struct {
	gorm.Model
	PName string `json:"p_name"`
	PUnit string `json:"p_unit"`
}

type OrdersTable struct {
	gorm.Model
	ProductName string `json:"product_name"`
	ProductUnit string `json:"product_unit"`
	Quantity    int    `json:"quantity"`
	RatePerUnit int    `json:"rate_per_unit"`
	Comments    string `json:"comments"`
}
