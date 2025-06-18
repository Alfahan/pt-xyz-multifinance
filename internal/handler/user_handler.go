package handler

import (
	"net/http"
	"pt-xyz-multifinance/internal/handler/dto"
	"pt-xyz-multifinance/internal/usecase"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user endpoints
type UserHandler struct {
	uc usecase.UserUsecase
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

// Register godoc
// @Summary      Register new user
// @Description  Create/register a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      dto.RegisterRequest  true  "User Data"
// @Success      201  {object}  dto.RegisterResponse
// @Failure      400  {object}  map[string]string "invalid request or validation error"
// @Router       /api/v1/register [post]
func (h *UserHandler) Register(c echo.Context) error {
	var req dto.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	res, err := h.uc.Register(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, res)
}

// Login godoc
// @Summary      Login user
// @Description  Authenticate user and generate a token
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      dto.LoginRequest  true  "Login Credentials"
// @Success      200  {object}  dto.LoginResponse
// @Failure      400  {object}  map[string]string "invalid request or validation error"
// @Failure      401  {object}  map[string]string "invalid credentials"
// @Router       /api/v1/login [post]
func (h *UserHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	res, err := h.uc.Login(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}
