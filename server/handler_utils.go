package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WrongParams(context *gin.Context) {
	context.JSON(http.StatusOK, struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{Status: "Failed", Message: "Wrong Parameters"})
}

func (s *Server) CheckAuthentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		authentication := context.GetHeader("Authorization")
		var admin AdminTable
		if !s.DB.Where(&AdminTable{Token: authentication}).First(&admin).RecordNotFound() {
			context.JSON(http.StatusOK, struct {
				Auth bool
			}{Auth: false})
			context.Abort()
		}
	}
}
