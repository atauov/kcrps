package handler

import (
	_ "github.com/atauov/kcrps/docs"
	"github.com/atauov/kcrps/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
			// invoices.POST("/multi", h.createMultiInvoiceFromFile)
			invoices.PUT("/cancel/:pos-id/:invoice-id", h.cancelInvoice)
			invoices.PUT("/refund/:pos-id/:invoice-id", h.cancelPayment)
			invoices.GET("/:pos-id", h.getAllInvoices)
			invoices.GET("/:pos-id/:invoice-id", h.getInvoiceById)
		}
	}

	router.Static("/.well-known", "./.well-known")

	return router
}
