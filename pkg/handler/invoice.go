package handler

import (
	"dashboard"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createInvoice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input dashboard.Invoice
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllInvoicesResponse struct {
	Data []dashboard.Invoice `json:"data"`
}

func (h *Handler) getAllInvoices(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	invoices, err := h.services.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllInvoicesResponse{
		Data: invoices,
	})
}

func (h *Handler) getInvoiceById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	invoice, err := h.services.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) updateInvoice(c *gin.Context) {

}

func (h *Handler) deleteInvoice(c *gin.Context) {

}
