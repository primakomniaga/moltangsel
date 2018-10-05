package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jemmycalak/mall-tangsel/api"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
)

var userService api.UserService

func Init(service api.UserService) {
	userService = service
}

func Register(c *gin.Context) {

	model := userService.NewUser()
	if err := c.BindJSON(&model); err != nil {
		u.ResponseError(c, http.StatusBadRequest, "Data not compelit")
		c.Abort()
		return
	}

	if err := userService.Register(c, model); err != nil {
		log.Println("error", err)
		u.ResponseError(c, http.StatusInternalServerError, "failed save data, please try again.")
		c.Abort()
		return
	}

	code := http.StatusOK

	payload := gin.H{
		"status":  "success",
		"message": "register berhasil",
		"code":    code,
	}
	u.ResponseSuccess(c, code, payload)
}

func Login(c *gin.Context) {

	token, err := u.GeneratorJWT(1)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "Something wrong, please try again")
		return
	}

	payload := gin.H{
		"status": "success",
		"code":   http.StatusOK,
		"token":  token,
	}
	u.ResponseSuccess(c, http.StatusOK, payload)
}

func GetUser(c *gin.Context) {
	a := userService.GetUser(c)

	fmt.Println(a)
	c.JSON(http.StatusOK, a)
}
