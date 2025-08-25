package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-crud-app/models"
	"net/http"
	"strconv"
	"time"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: cannot parse request": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error: cannot parse request": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func createEvent(c *gin.Context) {

	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.GetInt64("userId")

	fmt.Println(userId)
	event.UserID = userId

	fmt.Println(event)

	retId, err := event.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("something went wrong : %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": retId, "timestamp": time.Now().Format(time.RFC3339)})
}

func updateEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentEvent, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("could not fetch the event by id : %d", id))
		return
	}

	userId, ok := GetUserID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not logged in"})
		return
	}
	fmt.Println(userId)
	event.ID = currentEvent.ID

	if currentEvent.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot update an event that you did not create"})
		return
	}

	retId, err := event.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("something went wrong : %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": retId, "timestamp": time.Now().Format(time.RFC3339)})
}

func deleteById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentEvent, err := models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("could not fetch the event by id : %d", id))
		return
	}

	userId, ok := GetUserID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you are not logged in"})
		return
	}
	fmt.Println(userId)

	if currentEvent.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you cannot delete an event that you did not create"})
		return
	}

	err = models.DeleteById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("something went wrong : %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"timestamp": time.Now().Format(time.RFC3339)})
}

func GetUserID(c *gin.Context) (int64, bool) {
	val, exists := c.Get("userId")
	if !exists {
		return 0, false
	}

	userId, ok := val.(int64)
	return userId, ok
}
