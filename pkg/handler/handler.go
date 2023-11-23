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
	return &Handler{services: services}
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
