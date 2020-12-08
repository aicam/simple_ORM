package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



func (s *Server) AdminLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var JSONStruct struct{
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


	}
}
