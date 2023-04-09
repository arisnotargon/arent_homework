package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/arisnotargon/arent_homework/config"
	"github.com/davecgh/go-spew/spew"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(expireMinute int, UserId int) (string, error) {
	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(UserId),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration((expireMinute))).Unix(), // 無効になる時点
		IssuedAt:  time.Now().Unix(),                                                  // 作成時点

		NotBefore: time.Now().Unix(), // Token有効になる時点
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VarifyToken(tokenString string) (*jwt.StandardClaims, error) {
	spew.Dump("in VarifyToken")
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.JwtSecret), nil
	})

	// 检查Token是否有效
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	if err != nil {
		return nil, err
	}

	nowTs := time.Now().Unix()

	if claims.ExpiresAt < nowTs || claims.NotBefore > nowTs {
		return nil, errors.New("token not activity")
	}

	return claims, nil
}
