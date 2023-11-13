package handler

import (
	"dashboard/pkg/service"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/", h.createInvoice)
		api.GET("/", h.getAllInvoices)
		api.GET("/:id", h.getInvoiceById)
		api.PUT("/:id", h.updateInvoice)
		api.DELETE("/:id", h.deleteInvoice)
	}
	return router
}
