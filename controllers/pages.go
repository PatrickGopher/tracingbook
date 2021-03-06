package controllers

import (
	"errors"
	"net/http"
	"tracingbook/dtos"
	"tracingbook/services"

	"github.com/gin-gonic/gin"
)

func RegisterPageRoutes(router *gin.RouterGroup) {
	router.GET("", Home)
	router.GET("/home", Home)

}

func Home(c *gin.Context) {

	books, err := services.FetchBooks()
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.CreateDetailedErrorDto("comments", errors.New("Somethign went wrong")))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateHomeResponse(c.Request, books))
}
