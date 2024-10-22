package repository

import (
	"github.com/RomanshkVolkov/test-api/internal/core/domain"
	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func SigninJWT(user domain.User) (string, error) {
	sign := []byte(GetEnv("JWT_SECRET"))
	claims := &CustomClaims{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			Issuer: "test-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(sign)

	if err != nil {
		panic(err)
	}
	return ss, nil
}

func ExtractDataByToken(tokenString string) (CustomClaims, error) {
	claims := &CustomClaims{}
	sign := []byte(GetEnv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return sign, nil
	})

	if err != nil {
		return CustomClaims{}, err
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return CustomClaims{}, err
	}

	return *claims, nil
}
