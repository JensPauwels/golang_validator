package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string
	Age      int
	Id       string
}

func handleTest(c *gin.Context) {
	errs, body := Validator{
		"username": "string",
		"age":      "int",
		"id":       "uuid",
	}.validateReqBody(c.Request.Body)

	fmt.Print(body)

	if len(errs) != 0 {
		c.JSON(400, gin.H{})
		return
	}

	c.JSON(200, gin.H{})
}

func main() {
	r := gin.Default()

	r.POST("/testing", handleTest)
	r.Run()
}
