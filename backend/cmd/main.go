package main

import (
	"github.com/gin-gonic/gin"
	"password-manager/backend/api"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {})
	router.POST("/createUser", api.CreateUser)
	router.PUT("/updatePassword", api.UpdateUserPassword)
	router.DELETE("/deleteUser", api.DeleteUser)
}
