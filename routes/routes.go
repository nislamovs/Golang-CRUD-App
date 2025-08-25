package routes

import (
	"github.com/gin-gonic/gin"
	"golang-crud-app/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Auth)

	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteById)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	//server.POST("/events", middlewares.Auth, createEvent)
	//server.PUT("/events/:id", middlewares.Auth, updateEvent)
	//server.DELETE("/events/:id", middlewares.Auth, deleteById)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
