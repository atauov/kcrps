package handler

import (
	"dashboard"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create invoice
// @Security ApiKeyAuth
// @Tags invoices
// @Description create invoice
// @ID invoice
// @Accept  json
// @Produce  json
// @Param input body dashboard.Invoice true "invoice info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices [post]
func (h *Handler) createInvoice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input dashboard.Invoice

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = validateInputInvoice(&input); err != nil {
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

	// if err = sendInvoice(&input); err != nil {
	// 	fmt.Println(err)
	// }
}

type getAllInvoicesResponse struct {
	Data []dashboard.Invoice `json:"data"`
}

// @Summary Get All Invoices
// @Security ApiKeyAuth
// @Tags invoices
// @Description get all user invoices
// @ID get-all-invoices
// @Accept  json
// @Produce  json
// @Success 200 {integer} getAllInvoicesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices [get]
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

// @Summary Get Invoice By Id
// @Security ApiKeyAuth
// @Tags invoices
// @Description get invoice by id
// @ID get-invoice-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} dashboard.Invoice
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/:id [get]
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

// @Summary Cancel Invoice By Id
// @Security ApiKeyAuth
// @Tags invoices
// @Description cancel invoice by id
// @ID cancel-invoice-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} dashboard.Invoice
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/:id [delete]
func (h *Handler) cancelInvoice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.services.Cancel(userId, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func validateInputInvoice(invoice *dashboard.Invoice) error {
	if invoice.Amount < 1 || invoice.Amount > 999999999 {
		return errors.New("incorrect amount")
	}
	if len(invoice.Account) != 11 {
		return errors.New("incorrect account")
	}
	if invoice.Account[:2] != "87" {
		return errors.New("incorrect account")
	}
	if _, err := strconv.Atoi(invoice.Account); err != nil {
		return err
	}
	if len([]rune(invoice.Message)) > 40 {
		invoice.Message = string([]rune(invoice.Message)[:40])
	}
	return nil
}
