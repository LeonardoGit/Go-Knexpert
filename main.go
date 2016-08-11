package main

import (
	"enkhalifapro/trax/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"enkhalifapro/knexpert-api/utilities"
)

type person struct {
	name string
	age  int
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// get enviroment configuration
	configUtil := utilities.NewConfigUtil()

	usersController := controllers.NewUserController(configUtil)
	router.GET("/login", usersController.Login)
	router.POST("/login", usersController.LoginPost)
	router.GET("/user", usersController.List)
	router.POST("/user", usersController.Create)

	// Listen and server
	port := GetPort()
	log.Println("Port is " + port)
	router.Run(port)
}
