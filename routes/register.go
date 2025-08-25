package routes

import (
	"github.com/gin-gonic/gin"
	"golang-crud-app/models"
	"net/http"
	"strconv"
	"time"
)

func registerForEvent(c *gin.Context) {
	userId, ok := GetUserID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not logged in"})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not register user for event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "registered!!!!", "timestamp": time.Now().Format(time.RFC3339)})
}

func cancelRegistration(c *gin.Context) {
	userId, ok := GetUserID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not logged in"})
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not cancel user registration for event"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"msg": "Registration canceled!!", "timestamp": time.Now().Format(time.RFC3339)})
}
