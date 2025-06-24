package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// getSecretKey reads JWT_SECRET from environment at runtime
func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token berlaku 24 jam
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(getSecretKey())
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(token string) (*Claims, error) {
	// Parse token dengan klaim khusus (Claims)
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Pastikan metode penandatanganan adalah HS256
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return getSecretKey(), nil
	})

	if err != nil {
		// Tangani error parsing token
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("token is malformed")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token is expired")
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token is not valid yet")
		} else {
			return nil, errors.New("failed to parse token")
		}
	}

	// Validasi klaim
	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
