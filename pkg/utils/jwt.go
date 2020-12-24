package utils

import (
	"fmt"
	"gin_test/pkg/app"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type AuthUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Group    int    `json:"group"`
}
type Claims struct {
	*AuthUser
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(u *AuthUser) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		AuthUser: u,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*AuthUser, error) {
	claims := Claims{}
	if token == "" {
		return claims.AuthUser, jwt.NewValidationError(app.GetMsg(app.ERROR_AUTH_EMPTY), jwt.ValidationErrorEmpty)
	}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return claims.AuthUser, jwt.NewValidationError(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]), jwt.ValidationErrorUnverifiable)
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return claims.AuthUser, err
	}

	return claims.AuthUser, err
}
