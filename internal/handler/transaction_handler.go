package handler

import (
	"net/http"

	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionUC usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUC usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{transactionUC: transactionUC}
}

// CreateTransaction godoc
// @Summary Create a new transaction
// @Param X-CSRF-Token header string true "CSRF Token"
// @Description Endpoint to create a new transaction. Validates consumer limit and processes the transaction.
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body dto.TransactionRequest true "Transaction data"
// @Success 201 {object} map[string]string "Transaction created successfully"
// @Failure 400 {object} map[string]string "Invalid request or limit exceeded"
// @Router /transactions [post]
func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	var req dto.TransactionRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	err := h.transactionUC.CreateTransaction(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "transaction created successfully"})
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Param X-CSRF-Token header string true "CSRF Token"
// @Description Fetch details of a transaction by its ID.
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} dto.TransactionResponse "Transaction details"
// @Failure 404 {object} map[string]string "Transaction not found"
// @Router /transactions/{id} [get]
func (h *TransactionHandler) GetTransactionByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.transactionUC.GetTransactionByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "transaction not found"})
	}

	return c.JSON(http.StatusOK, res)
}
