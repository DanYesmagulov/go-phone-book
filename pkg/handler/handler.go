package handler

import (
	"github.com/DanYesmagulov/go-phone-book/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		contacts := api.Group("/contacts")
		{
			contacts.POST("/", h.createContact)
			contacts.GET("/", h.getAllContacts)
			contacts.GET("/:id", h.getContactById)
			contacts.PUT("/:id", h.updateContactById)
			contacts.DELETE("/:id", h.deleteContactById)
		}
	}

	return router
}
