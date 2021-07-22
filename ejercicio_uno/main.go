package main

import (
	"github.com/gin-gonic/gin"
)

type Data struct {
	Data string `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/myapi", func(c *gin.Context) {
		value := c.Query("data")
		data := Data {
			Data: value,
		}
		c.JSON(200, data)
	})
	r.Run(":8080")
}