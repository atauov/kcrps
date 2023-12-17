package handler

import (
	"errors"
	"github.com/atauov/kcrps/models/request"
	"github.com/atauov/kcrps/models/response"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create invoice
// @Security ApiKeyAuth
// @Tags invoices
// @Description create invoice
// @UUID invoice
// @Accept  json
// @Produce  json
// @Param input body request.Invoice true "invoice info"
// @Success 200 {object} idResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices [post]
func (h *Handler) createInvoice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input request.Invoice

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = validateInputInvoice(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userId

	uuId, err := h.services.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{
		Uuid: uuId,
	})

}

type getAllInvoicesResponse struct {
	Data []response.Invoice `json:"data"`
}

// @Summary Get All Invoices
// @Security ApiKeyAuth
// @Tags invoices
// @Description get all user invoices
// @UUID get-all-invoices
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllInvoicesResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/:pos-id [get]
func (h *Handler) getAllInvoices(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posId, err := uuid.Parse(c.Param("pos-id"))
	if err != nil {
		logrus.Printf("error parsing pos-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong pos-id parameter")
		return
	}

	input := request.Invoice{
		UserID: userId,
		PosID:  posId,
	}

	invoices, err := h.services.GetAll(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllInvoicesResponse{
		Data: invoices,
	})
}

// @Summary Get Invoice By UUID
// @Security ApiKeyAuth
// @Tags invoices
// @Description get invoice by id
// @UUID get-invoice-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Invoice
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/:pos-id/:id [get]
func (h *Handler) getInvoiceById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posId, err := uuid.Parse(c.Param("pos-id"))
	if err != nil {
		logrus.Printf("error parsing pos-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong pos-id parameter")
		return
	}

	invoiceId, err := strconv.Atoi(c.Param("invoice-id"))
	if err != nil {
		logrus.Printf("error parsing invoice-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong invoice-id parameter")
		return
	}

	input := request.Invoice{
		UUID:   invoiceId,
		UserID: userId,
		PosID:  posId,
	}

	invoice, err := h.services.GetById(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, invoice)
}

// @Summary Cancel Invoice By UUID
// @Security ApiKeyAuth
// @Tags invoices
// @Description cancel invoice by id
// @UUID cancel-invoice-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/cancel/:pos-id/:id [put]
func (h *Handler) cancelInvoice(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posId, err := uuid.Parse(c.Param("pos-id"))
	if err != nil {
		logrus.Printf("error parsing pos-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong pos-id parameter")
		return
	}

	invoiceId, err := strconv.Atoi(c.Param("invoice-id"))
	if err != nil {
		logrus.Printf("error parsing invoice-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong invoice-id parameter")
		return
	}

	invoice := request.Invoice{
		UserID: userId,
		PosID:  posId,
		UUID:   invoiceId,
	}

	if err = h.services.SetInvoiceForCancel(invoice); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Refund Invoice By UUID
// @Security ApiKeyAuth
// @Tags invoices
// @Description refund invoice by id
// @UUID refund-invoice-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/invoices/refund/:pos-id/:id [put]
func (h *Handler) cancelPayment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	posId, err := uuid.Parse(c.Param("pos-id"))
	if err != nil {
		logrus.Printf("error parsing pos-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong pos-id parameter")
		return
	}

	invoiceId, err := strconv.Atoi(c.Param("invoice-id"))
	if err != nil {
		logrus.Printf("error parsing invoice-id: %s", err)
		newErrorResponse(c, http.StatusBadRequest, "wrong invoice-id parameter")
		return
	}

	invoice := request.Invoice{
		UserID: userId,
		PosID:  posId,
		UUID:   invoiceId,
	}

	if err = h.services.SetInvoiceForRefund(invoice); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func validateInputInvoice(invoice *request.Invoice) error {
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
