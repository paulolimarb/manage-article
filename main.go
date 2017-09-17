package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	router.LoadHTMLGlob("templates/*")

	initializeRoutes()
	router.Run(":3000")

}

func render(c *gin.Context, data gin.H, templateName string)  {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])

	case "application/xml":
		c.XML(http.StatusOK, data["payload"])

	default:
		c.HTML(http.StatusOK, templateName, data)

	}
}