package main

import (
	"github.com/eva81829/go-postgres/mypackage/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user", controller.GetUsers)                 //get users
	r.GET("/user/:user_id", controller.GetUserById)     //get user by id
	r.POST("/user", controller.AddUser)                 //add user
	r.PUT("/user/:user_id", controller.ModUserById)     //modify user by id
	r.DELETE("/user/:user_id", controller.DelUserById)  //delete user by id
	r.Run(":8080")
}