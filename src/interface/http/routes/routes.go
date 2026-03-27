package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

type Handler struct {
	eventService        services.EventService
	registrationService services.RegistrationService
	userService         services.UserService
}

func NewHandler(eventService services.EventService, registrationService services.RegistrationService, userService services.UserService) *Handler {
	return &Handler{
		eventService:        eventService,
		registrationService: registrationService,
		userService:         userService,
	}
}

func (h *Handler) Register(server *gin.Engine, authMiddleware gin.HandlerFunc) {
	server.POST("/signup", h.signup)
	server.POST("/login", h.login)

	server.GET("/events", h.getEvents)
	server.GET("/events/:id", h.getEventByID)

	authenticated := server.Group("/")

	authenticated.Use(authMiddleware)
	authenticated.POST("/events", h.createEvent)
	authenticated.PUT("/events/:id", h.updateEvent)
	authenticated.DELETE("/events/:id", h.deleteEvent)

	authenticated.POST("/events/:id/register", h.createRegistration)
	authenticated.DELETE("/events/:id/register", h.cancelRegistration)
}
