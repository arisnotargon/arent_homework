package utils

import (
	"strconv"
	"time"

	"github.com/arisnotargon/arent_homework/config"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(expireMinute int, UserId int) (string, error) {
	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(UserId),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 無効になる時点
		IssuedAt:  time.Now().Unix(),                     // 作成時点
		NotBefore: time.Now().Unix(),                     // Token有効になる時点
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
