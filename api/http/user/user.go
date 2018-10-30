package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jemmycalak/mall-tangsel/api"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
	m "github.com/jemmycalak/mall-tangsel/service/user"
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

	//validation data
	if userService.ValidEmailPhone(model) {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "Email or Phone is already exist")
		return
	}

	pass, err := u.Encryptor(model.Password)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "Something wrong, please try again")
		return
	}
	//new model
	nmodel := m.User{
		Name:     model.Name,
		Password: pass,
		Email:    model.Email,
		Phone:    model.Phone,
		CreateAt: model.CreateAt,
		Level:    model.Level,
	}

	if err := userService.Register(c, &nmodel); err != nil {
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

	var model m.Login
	if err := c.BindJSON(&model); err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusOK, "email and password required")
		return
	}

	nmodel, err := userService.Login(&model)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "Email not found")
		return
	}

	pass, err := u.Decryptor(nmodel.Password)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "Something wrong, please try again")
		return
	}

	if u.NewPassword(model.Password) != pass {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "Woops wrong password")
		return
	}

	token, err := u.GeneratorJWT(nmodel.ID)
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

	token := c.Request.Header.Get("Authorization")
	_, idUser := u.ValidatorJWT(token, c)

	model, err := userService.GetUser(idUser)
	if err != nil {
		fmt.Println("error ", err)
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "Data not found, please try again")
		return
	}

	payload := gin.H{
		"status": "success",
		"code":   http.StatusOK,
		"data":   model,
	}

	u.ResponseSuccess(c, http.StatusOK, payload)
}
