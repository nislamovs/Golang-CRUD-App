package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-crud-app/models"
	"golang-crud-app/utils"
	"net/http"
	"time"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("something went wrong : %s", err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "user registered", "timestamp": time.Now().Format(time.RFC3339)})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Could not authenticate user : error : %s", err.Error()))
		return
	}
	fmt.Printf("userid is : %d", user.ID)
	fmt.Print(user)

	token, err := utils.GenerateJWT(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Could not generate token : error : %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "user logged in", "timestamp": time.Now().Format(time.RFC3339), "token": token})
}
