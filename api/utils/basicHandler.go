package utils

import (
	"errors"
	"fmt"
	"net/http"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTModel struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

var tokenPetter string = "secret*#key#*for*#AES&encryption"

func ApikeyValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.Request.Header.Get("apikey")

		if apikey == "" {
			c.Abort()
			c.Error(errors.New("No auth APIKEY"))
			ResponseError(c, http.StatusBadRequest, "APIKEY required")
			return
		}
		c.Next()
	}
}

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.Abort()
			c.Error(errors.New("No auth Authorization"))
			ResponseError(c, http.StatusBadRequest, "Authorization required")
			return
		}

		if !ValidatorJWT(token, c) {
			c.Abort()
			c.Error(errors.New("Auht token is invalid"))
			ResponseError(c, http.StatusBadRequest, "Auth token is invalid")
			return
		}
		c.Set("Bearer", token)
		c.Next()
	}
}

func ValidatorJWT(token string, c *gin.Context) bool {
	nToken, err := jwt.ParseWithClaims(token, &JWTModel{}, func(nToken *jwt.Token) (interface{}, error) {
		return []byte(tokenPetter), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := nToken.Claims.(*JWTModel); ok && nToken.Valid {
		c.Set("userid", claims.UserId)
		return true
	}
	return false
}

func GeneratorJWT(userId int) (string, error) {
	claims := JWTModel{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 7200,
			Issuer:    "CRUD",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingString, err := token.SignedString([]byte(tokenPetter))
	if err != nil {
		fmt.Println("error generet jwt ", err)
		return "", err
	}
	return signingString, nil

}
