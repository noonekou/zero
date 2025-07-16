package auth

import (
	errs "bookstore/common/error"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"userId"`
}

func GenerateToken(accessSecret string, accessExpire int64, userId int64) (string, error) {
	now := time.Now().Unix()

	claims := CustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "bookstore-api",
			ExpiresAt: jwt.NewNumericDate(time.Unix(now+accessExpire, 0)),
			IssuedAt:  jwt.NewNumericDate(time.Unix(now, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(accessSecret))
	if err != nil {
		logx.Errorf("generate token error: %v", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(accessSecret string, token string) (int64, error) {
	var rawToken string = token
	if strings.HasPrefix(token, "Bearer ") {
		rawToken = strings.TrimPrefix(token, "Bearer ")
	}

	tk, err := jwt.ParseWithClaims(rawToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil {
		return 0, errs.ErrTokenInvalid
	}

	claims, ok := tk.Claims.(*CustomClaims)
	if !ok || !tk.Valid {
		return 0, errs.ErrTokenInvalid
	}

	return claims.UserId, nil
}
