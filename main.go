package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	lvlFlag := flag.String("l", "info", "description of log level")

	flag.Parse()

	fmt.Printf("level is %s\n", *lvlFlag)
	fmt.Println("Gameworm API Blank Commit")
	fmt.Println("Hello, World!")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
