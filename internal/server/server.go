package server

import (
	"github.com/Tristan-Muggridge/custom-oauth-provider/internal/auth"
	"github.com/gin-gonic/gin"
)

var apiPrefix string = "/api/v1/"
var routes = []struct {
	path    string
	handler func(c *gin.Context)
}{
	{apiPrefix + "hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	}},
	{apiPrefix + "auth/" + "authorize", auth.AuthController.Authorize},
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	for _, route := range routes {
		r.GET(route.path, route.handler)
	}

	return r
}

func StartServer() {
	r := SetupRouter()
	r.Run((":8080"))
}
