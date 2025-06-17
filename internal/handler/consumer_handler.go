package handler

import (
	"net/http"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/usecase"

	"github.com/labstack/echo/v4"
)

// ConsumerHandler handles consumer endpoints
type ConsumerHandler struct {
	uc usecase.ConsumerUsecase
}

func NewConsumerHandler(uc usecase.ConsumerUsecase) *ConsumerHandler {
	return &ConsumerHandler{uc: uc}
}

// Create godoc
// @Summary      Register new consumer
// @Description  Create/register a new consumer
// @Tags         consumer
// @Accept       json
// @Produce      json
// @Param        consumer  body      dto.CreateConsumerRequest  true  "Consumer Data"
// @Success      201  {object}  dto.ConsumerResponse
// @Failure      400  {object}  map[string]string "invalid request or validation error"
// @Router       /api/v1/consumers [post]
// @Security     BearerAuth
func (h *ConsumerHandler) Create(c echo.Context) error {
	var req dto.CreateConsumerRequest
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
// @Summary      Get consumer by ID
// @Description  Get detail of consumer by ID
// @Tags         consumer
// @Produce      json
// @Param        id   path      string  true  "Consumer ID"
// @Success      200  {object}  dto.ConsumerResponse
// @Failure      404  {object}  map[string]string "consumer not found"
// @Router       /api/v1/consumers/{id} [get]
// @Security     BearerAuth
func (h *ConsumerHandler) GetByID(c echo.Context) error {
	id := c.Param("id")
	res, err := h.uc.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, res)
}
