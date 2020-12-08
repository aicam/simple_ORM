package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)




func WrongParams(context *gin.Context) {
	context.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		Message string `json:"message"`
	}{Status: "Failed", Message: "Wrong Parameters"})
}
