package http

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/mall-tangsel/api"
	"github.com/jemmycalak/mall-tangsel/api/http/product"
	"github.com/jemmycalak/mall-tangsel/api/http/user"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
)

type Server struct {
	server         *http.Server
	UserService    api.UserService
	ProductService api.ProductService
}

func (s *Server) Serve(lis net.Listener) error {
	s.server = &http.Server{}

	// init all handler
	user.Init(s.UserService)
	product.Init(s.ProductService)

	// import all route into server handler
	s.server.Handler = Handler()

	return s.server.Serve(lis)
}

func Handler() *gin.Engine {

	c := gin.Default()
	c.MaxMultipartMemory = 1 << 2
	// if c.MaxMultipartMemory <= 1 {
	// 	log.Println("max values")
	// 	return nil
	// }
	v1 := c.Group("/api/v1")
	{
		//user
		v1.POST("/account", user.Register)
		v1.GET("/account", user.GetUser)

		//product
		v1.POST("/product", product.NewProduct)
		v1.GET("/product/:id", product.Product)
		v1.GET("/products", product.Products)
		v1.PUT("/product", product.EditProduct)
		v1.DELETE("/product/:id", product.DeleteProduct)
		v1.PUT("/product/image/:id", product.UpdateImage)
		v1.GET("/products/limit", product.ProductLimit)
	}
	c.NoRoute(func(c *gin.Context) {
		u.ResponseError(c, http.StatusBadGateway, "I dont know what are you looking for !")
		c.Abort()
		return
	})
	return c
}
