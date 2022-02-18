package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Id       string `json:"id"`
}

func handleTest(c *gin.Context) {
	user := User{}

	errs := Validator{
		"username": "string",
		"age":      "int",
		"id":       "uuid",
	}.validateAndMarshalBody(c.Request.Body, &user)

	if len(errs) != 0 {
		c.JSON(400, gin.H{})
		return
	}

	fmt.Print(user)

	c.JSON(200, gin.H{})
}

func main() {
	r := gin.Default()

	r.POST("/testing", handleTest)
	r.Run()
}
