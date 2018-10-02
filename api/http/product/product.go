package product

import (
	"fmt"
	"net/http"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/mall-tangsel/api"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
	m "github.com/jemmycalak/mall-tangsel/service/product"
)

var productService api.ProductService

func Init(service api.ProductService) {
	productService = service
}

func NewProduct2(c *gin.Context) {

	model := productService.NewProduct()
	if err := c.ShouldBind(model); err != nil {
		fmt.Println(err)
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "data not complit")
		return
	}

	path := "www.image.com/mall-tangsel/"
	file, _ := c.MultipartForm()

	var mImage []string
	images := file.File["image[]"]
	if len(images) == 0 {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "image must be sumbitted")
		return
	}

	t := time.Now().Local()
	format := "020120060304"
	for _, image := range images {
		nPath := path + t.Format(format) + "_" + image.Filename
		mImage = append(mImage, nPath)
	}

	/*
	*	do send to server image
	 */

	sname := strings.Split(model.Name, " ")
	mname := "/"
	for _, element := range sname {
		mname += "_" + element
	}
	bind := m.Product{
		Name:        model.Name,
		Descript:    model.Descript,
		Price:       model.Price,
		Image:       mImage,
		Stock:       model.Stock,
		Condition:   model.Condition,
		Url:         mname,
		CountReview: model.CountReview,
		Preorder:    model.Preorder,
		Rating:      model.Rating,
		CategoryId:  model.CategoryId,
		MerchantId:  model.MerchantId,
		InsertBy:    model.InsertBy,
		CreateAt:    model.CreateAt,
		IsDelete:    model.IsDelete,
	}

	if err := productService.CreateProduct(&bind); err != nil {
		u.ResponseError(c, http.StatusInternalServerError, "failed add product, please try again")
		c.Abort()
		return
	}

	payload := gin.H{
		"status":  "success",
		"message": "berhasil menambah product",
		"code":    http.StatusOK,
	}
	u.ResponseSuccess(c, http.StatusOK, payload)
}
