package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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

const (
	barear_schema = "Bearer "
)

func TokenValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		fmt.Println(token)

		if token == "" || !strings.HasPrefix(token, barear_schema) {
			c.Abort()
			c.Error(errors.New("No auth Authorization"))
			ResponseError(c, http.StatusBadRequest, "Authorization required")
			return
		}
		val, _ := ValidatorJWT(token, c)
		if !val {
			c.Abort()
			c.Error(errors.New("Auht token is invalid"))
			ResponseError(c, http.StatusBadRequest, "Auth token is invalid")
			return
		}
		splitToken := strings.Split(token, "Bearer ")
		c.Set("Bearer", splitToken[1])
		c.Next()
	}
}

func ValidatorJWT(token string, c *gin.Context) (bool, int) {

	splitToken := strings.Split(token, "Bearer ")

	nToken, err := jwt.ParseWithClaims(splitToken[1], &JWTModel{}, func(nToken *jwt.Token) (interface{}, error) {
		return []byte(tokenPetter), nil
	})
	if err != nil {
		return false, 0
	}
	if claims, ok := nToken.Claims.(*JWTModel); ok && nToken.Valid {
		c.Set("userid", claims.UserId)
		return true, claims.UserId
	}
	return false, 0
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
	return "Bearer " + signingString, nil

}
