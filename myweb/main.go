package main

import (
	"github.com/eva81829/go-postgres/mypackage/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", GetUsers)                 //get users
	r.GET("/user/:user_id", GetUserById)     //get user by id
	r.POST("/user", AddUser)                 //add user
	r.PUT("/user/:user_id", ModUserById)     //modify user by id
	r.DELETE("/user/:user_id", DelUserById)  //delete user by id
	r.Run(":8080")
}