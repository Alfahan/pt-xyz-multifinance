package handler

import (
	"net/http"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/usecase"

	"github.com/labstack/echo/v4"
)

// CustomerHandler handles customer endpoints
type CustomerHandler struct {
	uc usecase.CustomerUsecase
}

func NewCustomerHandler(uc usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{uc: uc}
}

// Create godoc
// @Summary      Register new customer
// @Description  Create/register a new customer
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        customer  body      dto.CreateCustomerRequest  true  "Customer Data"
// @Success      201  {object}  dto.CustomerResponse
// @Failure      400  {object}  map[string]string "invalid request or validation error"
// @Router       /api/v1/customers [post]
// @Security     BearerAuth
func (h *CustomerHandler) Create(c echo.Context) error {
	var req dto.CreateCustomerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	res, err := h.uc.Create(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, res)
}

// GetByID godoc
// @Summary      Get customer by ID
// @Description  Get detail of customer by ID
// @Tags         customer
// @Produce      json
// @Param        id   path      string  true  "Customer ID"
// @Success      200  {object}  dto.CustomerResponse
// @Failure      404  {object}  map[string]string "customer not found"
// @Router       /api/v1/customers/{id} [get]
// @Security     BearerAuth
func (h *CustomerHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.uc.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
