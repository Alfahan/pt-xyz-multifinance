package middleware

import (
	"log"
	"net/http"
	"pt-xyz-multifinance/pkg" // Pastikan path package benar
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari header Authorization
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is missing")
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing Authorization header"})
		}

		// Log token untuk debugging
		log.Println("Authorization header value:", authHeader)

		// Validasi format token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization header format")
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid Authorization header format"})
		}

		token := parts[1]
		log.Println("Extracted token:", token)

		// Validasi token
		claims, err := pkg.ValidateToken(token)
		if err != nil {
			log.Println("Token validation error:", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		}

		// Simpan user ID ke context
		c.Set("userID", claims.UserID)

		return next(c)
	}
}
