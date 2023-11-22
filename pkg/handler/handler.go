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
		invoices := api.Group("/invoices")
		{
			invoices.POST("/", h.createInvoice)
			invoices.POST("/:id", h.cancelInvoice)
			invoices.GET("/", h.getAllInvoices)
			invoices.GET("/:id", h.getInvoiceById)
		}
	}

	router.LoadHTMLGlob("templates/*")

	router.GET("/panel", h.getPanelPage)
	router.GET("/login", h.getLoginPage)
	router.GET("/register", h.getRegisterPage)

	return router
}
