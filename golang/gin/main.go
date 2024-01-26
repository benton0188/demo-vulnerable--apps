package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// The Run() call below is commented out since this is the vulnerable method.
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Hello, World!")
}
