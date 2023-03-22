package main

import (
	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.String(200, "Hello this is a simple gin server")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", helloWorld)
	return r
}

func main() {
	r := setupRouter()
	if err := r.Run(":8080"); err != nil {
		return
	}

}
