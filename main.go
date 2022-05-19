package main

import "github.com/gin-gonic/gin"

func defualtAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()

	r.GET("/ping", defualtAPI)
	r.Run(":9091") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
