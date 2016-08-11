package controllers

import (
	"github.com/gin-gonic/gin"
	"enkhalifapro/trax/services"
	"enkhalifapro/trax/models"
	"enkhalifapro/trax/viewModels"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"enkhalifapro/knexpert-api/utilities"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(configUtil utilities.IConfigUtil) *UserController {
	controller := UserController{}
	controller.userService = services.NewUserService(configUtil)
	return &controller
}

func (controller UserController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "users_login.html", gin.H{
		"title": "Login" })
}

func (controller UserController) LoginPost(c *gin.Context) {
	var loginViewModel viewModels.LoginViewModel
	err := c.BindJSON(&loginViewModel)
	if err == nil {
		fmt.Println(loginViewModel.UserName)
		fmt.Println(loginViewModel.Password)
		var loginResult = controller.userService.Login(&loginViewModel)
		if loginResult == true {
			c.Header("Set-Cookie", "trToken=" + loginViewModel.Token)
			c.Redirect(http.StatusMovedPermanently, "/")
		}else {
			c.JSON(http.StatusInternalServerError, gin.H{"message":"invalid user or password"})
		}
	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
	}
}

func (controller UserController) Create(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err == nil {
		err = controller.userService.Insert(&user)
		if err == nil {
			c.JSON(http.StatusOK, &user)
		} else {
			c.JSON(http.StatusInternalServerError, err)
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": "unauthorized"})
	}

}

func (controller UserController) List(c *gin.Context) {
	userList := controller.userService.Find(&bson.M{})
	c.JSON(http.StatusOK, userList)
}
