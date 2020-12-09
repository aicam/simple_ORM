package server

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
		if !CheckAdmin(s.DB, JSONStruct.Username, JSONStruct.Password) {
			context.JSON(http.StatusOK, struct {
				Messsage string `json:"messsage"`
			}{Messsage: "Wrong username or password!"})
			return
		}

		token := sha256.Sum256([]byte(JSONStruct.Username + JSONStruct.Password + time.Now().String()))
		SetAdminToken(s.DB, JSONStruct.Username, hex.EncodeToString(token[:]))
		context.JSON(http.StatusOK, struct {
			Message string
			Token   string
		}{Message: "Login successful", Token: hex.EncodeToString(token[:])})
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
