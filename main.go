package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imavdhoot/GO-AVD-project1/constant"
	"github.com/imavdhoot/GO-AVD-project1/model"
	"io"
	"os"
)

import "github.com/imavdhoot/GO-AVD-project1/mod"

func defualtAPI(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
func main() {
	r := gin.Default()

	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("AVD-project1.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(gin.Recovery())

	r.GET("/ping", defualtAPI)
	r.POST("/merchant/add", mod.AddMerchant)
	fmt.Println(constant.SeverStarted)
	model.InitDB()
	r.Run(":9091") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
