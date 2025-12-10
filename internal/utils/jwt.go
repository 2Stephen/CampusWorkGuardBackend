package utils

import (
	"CampusWorkGuardBackend/internal/initialize"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func getJWTSecret() []byte {
	return []byte(initialize.AppConfig.JWTConfig.Secret)
}

func getExpireMinutes() int {
	return initialize.AppConfig.JWTConfig.Expires
}

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// 生成 Token
func GenerateJWTToken(userID int, email string, role string) (string, error) {
	JWTExpireMinutes := getExpireMinutes()
	JWTSecret := getJWTSecret()

	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(JWTExpireMinutes) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// 解析 Token
func ParseJWTToken(tokenString string) (*Claims, error) {
	JWTSecret := getJWTSecret()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func IsTokenExpired(err error) bool {
	return errors.Is(err, jwt.ErrTokenExpired)
}
