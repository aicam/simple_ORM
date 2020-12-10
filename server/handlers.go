package server

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Server) CheckToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var status bool
		var ad AdminTable
		if s.DB.Where(&AdminTable{Token: context.Param("token")}).First(&ad).RecordNotFound() {
			status = false
		} else {
			status = true
		}
		context.JSON(http.StatusOK, struct {
			Status bool `json:"status"`
		}{Status: status})
	}
}

func (s *Server) AdminLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var JSONStruct struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		err := context.BindJSON(&JSONStruct)
		if err != nil {
			WrongParams(context)
			return
		}
		if !CheckAdmin(s.DB, JSONStruct.Username, JSONStruct.Password) || JSONStruct.Username == "" {
			context.JSON(http.StatusOK, struct {
				Status   bool   `json:"status"`
				Messsage string `json:"messsage"`
			}{Messsage: "Wrong username or password!", Status: false})
			return
		}

		token := sha256.Sum256([]byte(JSONStruct.Username + JSONStruct.Password + time.Now().String()))
		SetAdminToken(s.DB, JSONStruct.Username, hex.EncodeToString(token[:]))
		context.JSON(http.StatusOK, struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
			Token   string `json:"token"`
		}{Message: "Login successful", Token: hex.EncodeToString(token[:]), Status: true})
		return
	}
}

func (s *Server) AddCustomer() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js CustomersTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Save(&js)
		context.JSON(http.StatusOK, struct {
			Status bool
		}{Status: true})
	}
}

func (s *Server) GetCustomers() gin.HandlerFunc {
	return func(context *gin.Context) {
		var customers []CustomersTable
		s.DB.Find(&customers)
		context.JSON(http.StatusOK, customers)
	}
}

func (s *Server) AddPayment() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js PaymentTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Save(&js)
		context.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Added Successfully"})
	}
}

func (s *Server) GetPayments() gin.HandlerFunc {
	return func(context *gin.Context) {
		var payments []PaymentTable
		s.DB.Find(&payments)
		context.JSON(http.StatusOK, payments)
	}
}

func (s *Server) AddProduct() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js ProductsTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Save(&js)
		context.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Added Successfully"})
	}
}

func (s *Server) GetProducts() gin.HandlerFunc {
	return func(context *gin.Context) {
		var products []ProductsTable
		s.DB.Find(&products)
		context.JSON(http.StatusOK, products)
	}
}

func (s *Server) AddOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js OrdersTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Save(&js)
		context.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Added Successfully"})
	}
}

func (s *Server) GetOrders() gin.HandlerFunc {
	return func(context *gin.Context) {
		var products []OrdersTable
		s.DB.Find(&products)
		context.JSON(http.StatusOK, products)
	}
}

func (s *Server) UpdateOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js OrdersTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Save(&js)
		context.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Updated Successfully"})
	}
}

func (s *Server) RemoveOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		var js OrdersTable
		err := context.BindJSON(&js)
		if err != nil {
			WrongParams(context)
			return
		}
		s.DB.Delete(&js)
		context.JSON(http.StatusOK, struct {
			Status  bool
			Message string
		}{Status: true, Message: "Deleted Successfully"})
	}
}

func (s *Server) AddAdmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		s.DB.Save(&AdminTable{
			Username: context.Param("user"),
			Password: context.Param("pass"),
			Token:    "",
		})
		context.String(200, "ok")
	}
}
