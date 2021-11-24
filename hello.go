package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World 3")
	})

	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	r.Run()
}

func test() {
    setupServer()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", pingEndpoint)

    return r
}

func pingEndpoint(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}
