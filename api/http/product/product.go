package product

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/mall-tangsel/api"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
	m "github.com/jemmycalak/mall-tangsel/service/product"
)

var productService api.ProductService
var path string = "www.image.com/mall-tangsel/"

func Init(service api.ProductService) {
	productService = service
}

//new product
func NewProduct(c *gin.Context) {

	model := productService.NewProduct()
	if err := c.ShouldBind(model); err != nil {
		fmt.Println(err)
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "data not complit")
		return
	}

	file, _ := c.MultipartForm()

	var mImage []string
	images := file.File["image[]"]
	if len(images) == 0 {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "image must be sumbitted")
		return
	}

	//path image
	for _, element := range images {
		nPath := path + NameImage(element.Filename)
		mImage = append(mImage, nPath)
	}

	/*
	*	do send to server image
	 */

	bind := m.Product{
		Name:        model.Name,
		Descript:    model.Descript,
		Price:       model.Price,
		Image:       mImage,
		Stock:       model.Stock,
		Condition:   model.Condition,
		Url:         CustemUrl(model.Name),
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

//show all product
func Products(c *gin.Context) {
	products, err := productService.Products()
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "Failed get data product, please try again")
		return
	}

	payload := gin.H{
		"code":   http.StatusOK,
		"status": "success",
		"data":   products,
	}

	u.ResponseSuccess(c, http.StatusOK, payload)
}

//product by id
func Product(c *gin.Context) {
	idProduct := c.Params.ByName("id")
	mid, err := strconv.Atoi(idProduct)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "id must be integer")
		return
	}

	product, err := productService.Product(mid)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "product not found")
		return
	}
	payload := gin.H{
		"status": "success",
		"code":   http.StatusOK,
		"data":   product,
	}
	u.ResponseSuccess(c, http.StatusOK, payload)
}

//edit product
func EditProduct(c *gin.Context) {
	var model m.Product
	if err := c.BindJSON(&model); err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "data not complit")
		return
	}

	bind := m.Product{
		ProductId:  model.ProductId,
		Name:       model.Name,
		Descript:   model.Descript,
		Price:      model.Price,
		Stock:      model.Stock,
		Condition:  model.Condition,
		Url:        CustemUrl(model.Name),
		Preorder:   model.Preorder,
		CategoryId: model.CategoryId,
		MerchantId: model.MerchantId,
		UpdateAt:   time.Now(),
	}

	if err := productService.EditProduct(&bind); err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "failed to edit data, please try again")
		return
	}

	payload := gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "berhasil mengedit data",
	}
	u.ResponseSuccess(c, http.StatusOK, payload)

}

//delet product
func DeleteProduct(c *gin.Context) {
	idproduct := c.Params.ByName("id")
	mid, err := strconv.Atoi(idproduct)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "id must be integer")
		return
	}

	if err := productService.DeleteProduct(mid); err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "failed")
		return
	}

	payload := gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "berhasil menghapus product",
	}
	u.ResponseSuccess(c, http.StatusOK, payload)
}

//update image product
func UpdateImage(c *gin.Context) {
	idproduct := c.Params.ByName("id")
	mid, err := strconv.Atoi(idproduct)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "id must be integer")
		return
	}
	file, _ := c.MultipartForm()
	var mImage []string
	images := file.File["image[]"]
	if len(images) == 0 {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "image must be submited")
		return
	}
	for _, element := range images {
		npath := path + NameImage(element.Filename)
		mImage = append(mImage, npath)
	}

	if err := productService.UpdateImage(mid, mImage); err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusInternalServerError, "failed update image, please try again")
		return
	}

	payload := gin.H{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "berhasil update image",
	}
	u.ResponseSuccess(c, http.StatusOK, payload)

}

func ProductLimit(c *gin.Context) {
	limit := c.DefaultQuery("limit", "50")
	page := c.DefaultQuery("page", "1")

	mlimit, err := strconv.Atoi(limit)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "limit must be integer")
		return
	}
	mpage, err := strconv.Atoi(page)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "page must be integer")
		return
	}

	newPage := (mlimit * mpage) - mlimit
	fmt.Println(newPage)
	model, err := productService.ProductLimit(newPage, mlimit)
	if err != nil {
		c.Abort()
		u.ResponseError(c, http.StatusBadRequest, "data not found")
		return
	}
	payload := gin.H{
		"status": "success",
		"code":   http.StatusOK,
		"data":   model,
	}
	u.ResponseSuccess(c, http.StatusOK, payload)
}

func NameImage(s string) string {
	//set Custem image name
	t := time.Now()
	format := "02012006030405"
	sname := strings.Split(s, " ")
	var nname string
	for _, element := range sname {
		nname += "_" + element
	}
	return "" + t.Format(format) + nname
}

func CustemUrl(s string) string {
	//set custem url
	t := time.Now().Local()
	format := "02012006030405"

	sname := strings.Split(s, " ")
	mname := "/" + t.Format(format)
	for _, element := range sname {
		mname += "_" + element
	}
	return mname
}
