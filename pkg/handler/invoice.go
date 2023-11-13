package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createInvoice(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllInvoices(c *gin.Context) {

}

func (h *Handler) getInvoiceById(c *gin.Context) {

}

func (h *Handler) updateInvoice(c *gin.Context) {

}

func (h *Handler) deleteInvoice(c *gin.Context) {

}
